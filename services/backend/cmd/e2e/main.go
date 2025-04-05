package main

import (
	"context"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/backend"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/containers"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/featureflags"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/middleware"
	"github.com/nats-io/nats.go"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	featureflags.EndToEndEnvironment.Enable()
	slog.SetLogLoggerLevel(slog.LevelDebug)

	ctx := context.Background()

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
	natsClient, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		slog.Error("Could not connect to NATS", slog.Any("err", err.Error()))
		os.Exit(1)
	}
	_ = natsClient.Publish("foo", []byte("bar"))

	// Start the web server
	mux := &http.ServeMux{}
	port := "8080"

	appCfg := backend.Configuration{
		Mux: mux,
	}
	backend.Start(appCfg)

	go func() {
		err := http.ListenAndServe(":"+port, middleware.Logging(middleware.RecoveryMiddleware(mux)))
		if err != nil {
			slog.Error("Could not start server", slog.Any("port", port), slog.Any("err", err.Error()))
			os.Exit(1)
		}
	}()

	slog.Info(fmt.Sprintf("Started server on http://localhost:%s\n", port))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	switch <-c {
	case os.Interrupt:
		slog.Info("Received interrupt signal, shutting down...")
		if err := containers.StopAllContainers(ctx); err != nil {
			slog.Error("Could not stop containers", slog.Any("err", err.Error()))
		}

		time.Sleep(5 * time.Second)
		slog.Info("All test containers stopped")
	}
}
