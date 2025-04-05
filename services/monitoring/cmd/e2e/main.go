package main

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/monitoring/internal/backend"
	"github.com/nats-io/nats.go"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

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
