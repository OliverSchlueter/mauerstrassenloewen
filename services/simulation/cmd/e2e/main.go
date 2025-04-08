package e2e

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/simulation/internal/backend"
	"github.com/OliverSchlueter/mauerstrassenloewen/simulation/internal/fflags"
	"github.com/OliverSchlueter/sloki/sloki"
	"github.com/nats-io/nats.go"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fflags.EndToEndEnvironment.Enable()
	//fflags.SendLogsToLoki.Enable()

	// Setup logging
	lokiLevel := slog.LevelInfo
	if !fflags.SendLogsToLoki.IsEnabled() {
		lokiLevel = 100_0000
	}
	lokiService := sloki.NewService(sloki.Configuration{
		URL:          "http://localhost:3100/loki/api/v1/push",
		Service:      "simulation",
		ConsoleLevel: slog.LevelDebug,
		LokiLevel:    lokiLevel,
	})
	slog.SetDefault(slog.New(lokiService))

	// Setup NATS
	natsClient, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		slog.Error("Could not connect to NATS", slog.Any("err", err.Error()))
		os.Exit(1)
	}

	mux := http.NewServeMux()

	backend.Start(backend.Configuration{
		Mux:  mux,
		Nats: natsClient,
	})

	slog.Info("Simulation service started")

	// Wait for a signal to exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
