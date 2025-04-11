package main

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"github.com/OliverSchlueter/mauerstrassenloewen/monitoring/internal/backend"
	"github.com/OliverSchlueter/mauerstrassenloewen/monitoring/internal/fflags"
	"github.com/nats-io/nats.go"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fflags.EndToEndEnvironment.Enable()
	//fflags.SendLogsToLoki.Enable()

	// Setup logging
	lokiService := sloki.NewService(sloki.Configuration{
		URL:          "http://localhost:3100/loki/api/v1/push",
		Service:      "monitoring",
		ConsoleLevel: slog.LevelDebug,
		LokiLevel:    slog.LevelInfo,
		EnableLoki:   fflags.SendLogsToLoki.IsEnabled(),
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
