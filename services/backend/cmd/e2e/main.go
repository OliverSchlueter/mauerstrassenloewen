package main

import (
	"common/middleware"
	"common/middleware/metricshandler"
	"context"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/backend"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/containers"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/fflags"
	"github.com/OliverSchlueter/sloki/sloki"
	"github.com/justinas/alice"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
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
	fflags.EndToEndEnvironment.Enable()
	//fflags.StartTestContainers.Enable()
	//fflags.SendLogsToLoki.Enable()

	// Setup logging
	lokiLevel := slog.LevelInfo
	if !fflags.SendLogsToLoki.IsEnabled() {
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
	if fflags.StartTestContainers.IsEnabled() {
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
	}

	// Setup NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		slog.Error("Could not connect to NATS", slog.Any("err", err.Error()))
		os.Exit(1)
	}

	// Setup MongoDB
	mc, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		slog.Error("Could not connect to MongoDB", slog.Any("err", err.Error()))
		os.Exit(1)
	}
	err = mc.Ping(ctx, readpref.Primary())
	if err != nil {
		slog.Error("Could not ping MongoDB", slog.Any("err", err.Error()))
		os.Exit(1)
	}
	mdb := mc.Database("msl_backend")

	// Start the web server
	mux := &http.ServeMux{}
	port := "8080"

	appCfg := backend.Configuration{
		Mux:     mux,
		Nats:    nc,
		MongoDB: mdb,
	}
	authMiddleware := backend.Start(appCfg)

	metricshandler.Register(mux, "")
	middleware.RegisterPrometheusHttpLogging()

	go func() {
		chain := alice.New(
			middleware.Logging,
			authMiddleware,
			middleware.RecoveryMiddleware,
		).Then(mux)

		err := http.ListenAndServe(":"+port, chain)
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
		if fflags.StartTestContainers.IsEnabled() {
			if err := containers.StopAllContainers(ctx); err != nil {
				slog.Error("Could not stop containers", slog.Any("err", err.Error()))
			}

			time.Sleep(5 * time.Second)
			slog.Info("All test containers stopped")
		}
	}
}
