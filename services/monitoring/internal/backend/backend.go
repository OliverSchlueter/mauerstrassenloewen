package backend

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/monitoring/internal/natslogging"
	"github.com/nats-io/nats.go"
	"log/slog"
	"os"
)

type Configuration struct {
	NatsClient *nats.Conn
}

func Start(cfg Configuration) {
	natsLoggingHandler := natslogging.NewHandler(natslogging.Configuration{
		Nats: cfg.NatsClient,
	})
	if err := natsLoggingHandler.Register(); err != nil {
		slog.Error("Could not register NATS logging handler", slog.Any("err", err.Error()))
		os.Exit(1)
	}
}
