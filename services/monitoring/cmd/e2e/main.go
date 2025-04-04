package main

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/monitoring/internal/backend"
	"github.com/OliverSchlueter/sloki/sloki"
	"github.com/nats-io/nats.go"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Setup logging
	lokiService := sloki.NewService(sloki.Configuration{
		URL:          "http://localhost:3100/loki/api/v1/push",
		Service:      "monitoring",
		ConsoleLevel: slog.LevelDebug,
		LokiLevel:    slog.LevelInfo,
	})
	slog.SetDefault(slog.New(lokiService))

	// Setup NATS
	natsClient, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		slog.Error("Could not connect to NATS", slog.Any("err", err.Error()))
		os.Exit(1)
	}

	backend.Start(backend.Configuration{
		NatsClient: natsClient,
	})

	slog.Info("NATS logging handler started")

	// Wait for a signal to exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
