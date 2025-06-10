package main

import (
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/middleware"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"github.com/OliverSchlueter/mauerstrassenloewen/monitoring/internal/backend"
	"github.com/OliverSchlueter/mauerstrassenloewen/monitoring/internal/fflags"
	"github.com/justinas/alice"
	"github.com/nats-io/nats.go"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	lokiUrlEnv       = "LOKI_URL"
	natsUrlEnv       = "NATS_URL"
	natsAuthTokenEnv = "NATS_AUTH_TOKEN"
)

func main() {
	fflags.EndToEndEnvironment.Disable()

	// Setup logging
	lokiService := sloki.NewService(sloki.Configuration{
		URL:          mustGetLokiURL(),
		Service:      "monitoring",
		ConsoleLevel: slog.LevelDebug,
		LokiLevel:    slog.LevelInfo,
		EnableLoki:   fflags.SendLogsToLoki.IsEnabled(),
	})
	slog.SetDefault(slog.New(lokiService))

	// Setup NATS
	natsClient, err := nats.Connect(mustGetEnvStr(natsUrlEnv), nats.Token(mustGetEnvStr(natsAuthTokenEnv)))
	if err != nil {
		slog.Error("Could not connect to NATS", sloki.WrapError(err))
		os.Exit(1)
	}

	mux := http.NewServeMux()
	port := "8084"

	backend.Start(backend.Configuration{
		Mux:        mux,
		NatsClient: natsClient,
	})

	go func() {
		chain := alice.New(
			middleware.CORS,
			middleware.Logging,
			middleware.RecoveryMiddleware,
		).Then(mux)

		err := http.ListenAndServe(":"+port, chain)
		if err != nil {
			slog.Error("Could not start server on port "+port, sloki.WrapError(err))
			os.Exit(1)
		}
	}()

	slog.Info(fmt.Sprintf("Started server on http://localhost:%s\n", port))

	// Wait for a signal to exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	switch <-sig {
	case os.Interrupt:
		slog.Info("Received interrupt signal, shutting down...")

		natsClient.Close()

		slog.Info("Shutdown complete")
	}
}

func mustGetLokiURL() string {
	lokiURL := os.Getenv(lokiUrlEnv)
	if lokiURL == "" && fflags.SendLogsToLoki.IsEnabled() {
		panic(lokiUrlEnv + " environment variable not set")
	}
	return lokiURL
}

func mustGetEnvStr(env string) string {
	val := os.Getenv(env)
	if val == "" {
		panic(env + " environment variable not set")
	}
	return val
}
