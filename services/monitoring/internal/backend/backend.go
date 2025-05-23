package backend

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/common/healthcheck"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"github.com/OliverSchlueter/mauerstrassenloewen/monitoring/internal/natslogging"
	"github.com/nats-io/nats.go"
	"log/slog"
	"net/http"
	"os"
)

type Configuration struct {
	Mux        *http.ServeMux
	NatsClient *nats.Conn
}

func Start(cfg Configuration) {
	natsLoggingHandler := natslogging.NewHandler(natslogging.Configuration{
		Nats: cfg.NatsClient,
	})
	if err := natsLoggingHandler.Register(); err != nil {
		slog.Error("Could not register NATS logging handler", sloki.WrapError(err))
		os.Exit(1)
	}

	healthcheck.Register(cfg.Mux)
}
