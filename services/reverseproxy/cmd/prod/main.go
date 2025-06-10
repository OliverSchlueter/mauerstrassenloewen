package main

import (
	"encoding/json"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"github.com/OliverSchlueter/mauerstrassenloewen/reverseproxy/internal/backend"
	"github.com/OliverSchlueter/mauerstrassenloewen/reverseproxy/internal/endpoint"
	"github.com/OliverSchlueter/mauerstrassenloewen/reverseproxy/internal/fflags"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	lokiUrlEnv    = "LOKI_URL"
	configPathEnv = "CONFIG_PATH"
)

func main() {
	fflags.EndToEndEnvironment.Disable()

	// Setup logging
	lokiService := sloki.NewService(sloki.Configuration{
		URL:          mustGetLokiURL(),
		Service:      "reverseproxy",
		ConsoleLevel: slog.LevelDebug,
		LokiLevel:    slog.LevelInfo,
		EnableLoki:   fflags.SendLogsToLoki.IsEnabled(),
	})
	slog.SetDefault(slog.New(lokiService))

	mux := http.NewServeMux()
	port := "8080"

	backend.Start(backend.Configuration{
		Mux:       mux,
		Endpoints: mustLoadConfig(),
	})

	go func() {
		err := http.ListenAndServe(":"+port, mux)
		if err != nil {
			slog.Error("Could not start server on port "+port, sloki.WrapError(err))
			os.Exit(1)
		}
	}()

	slog.Info(fmt.Sprintf("Reverse proxy service started on http://localhost:%s", port))

	// Wait for a signal to exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	switch <-sig {
	case os.Interrupt:
		slog.Info("Received interrupt signal, shutting down...")

		slog.Info("Shutdown complete")
	}
}

func mustLoadConfig() []endpoint.Endpoint {
	path := mustGetConfigPath()
	data, err := os.ReadFile(path)
	if err != nil {
		slog.Error("Could not read config file", sloki.WrapError(err))
		os.Exit(1)
	}

	var endpoints []endpoint.Endpoint
	err = json.Unmarshal(data, &endpoints)
	if err != nil {
		slog.Error("Could not unmarshal config file", sloki.WrapError(err))
		os.Exit(1)
	}

	return endpoints
}

func mustGetLokiURL() string {
	lokiURL := os.Getenv(lokiUrlEnv)
	if lokiURL == "" && fflags.SendLogsToLoki.IsEnabled() {
		panic(lokiUrlEnv + " environment variable not set")
	}
	return lokiURL
}

func mustGetConfigPath() string {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		slog.Error(configPathEnv + " environment variable not set")
		os.Exit(1)
	}
	return path
}
