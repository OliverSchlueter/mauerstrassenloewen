package main

import (
	"context"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/backend"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/containers"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/featureflags"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/middleware"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/middleware/metricshandler"
	"github.com/OliverSchlueter/sloki/sloki"
	"github.com/nats-io/nats.go"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()

	// Setup feature flags
	featureflags.EndToEndEnvironment.Enable()
	featureflags.SendLogsToLoki.Enable()

	// Setup logging
	lokiLevel := slog.LevelInfo
	if !featureflags.SendLogsToLoki.IsEnabled() {
		lokiLevel = 100_0000
	}
	lokiService := sloki.NewService(sloki.Configuration{
		URL:          "http://localhost:3100/loki/api/v1/push",
		Service:      "backend",
		ConsoleLevel: slog.LevelDebug,
		LokiLevel:    lokiLevel,
	})
	slog.SetDefault(slog.New(lokiService))

	// Start test containers
	_, err := containers.StartMongoDB(ctx)
	if err != nil {
		slog.Error("Could not start MongoDB", slog.Any("err", err.Error()))
		os.Exit(1)
	}

	_, err = containers.StartNATS(ctx)
	if err != nil {
		slog.Error("Could not start NATS", slog.Any("err", err.Error()))
		os.Exit(1)
	}

	// Setup NATS
	_, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		slog.Error("Could not connect to NATS", slog.Any("err", err.Error()))
		os.Exit(1)
	}

	// Start the web server
	mux := &http.ServeMux{}
	port := "8080"

	appCfg := backend.Configuration{
		Mux: mux,
	}
	backend.Start(appCfg)

	metricshandler.Register(mux, "")
	middleware.RegisterPrometheusHttpLogging()

	go func() {
		err := http.ListenAndServe(":"+port, middleware.Logging(middleware.RecoveryMiddleware(mux)))
		if err != nil {
			slog.Error("Could not start server on port "+port, slog.Any("err", err.Error()))
			os.Exit(1)
		}
	}()

	slog.Info(fmt.Sprintf("Started server on http://localhost:%s\n", port))

	// Wait for a signal to exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	switch <-sig {
	case os.Interrupt:
		slog.Info("Received interrupt signal, shutting down...")
		if err := containers.StopAllContainers(ctx); err != nil {
			slog.Error("Could not stop containers", slog.Any("err", err.Error()))
		}

		time.Sleep(5 * time.Second)
		slog.Info("All test containers stopped")
	}
}
