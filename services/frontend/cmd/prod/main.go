package main

import (
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/middleware"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"github.com/OliverSchlueter/mauerstrassenloewen/frontend/internal/backend"
	"github.com/OliverSchlueter/mauerstrassenloewen/frontend/internal/fflags"
	"github.com/justinas/alice"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	lokiUrlEnv = "LOKI_URL"
)

func main() {
	// Setup feature flags
	fflags.EndToEndEnvironment.Disable()

	lokiService := sloki.NewService(sloki.Configuration{
		URL:          mustGetLokiURL(),
		Service:      "frontend",
		ConsoleLevel: slog.LevelDebug,
		LokiLevel:    slog.LevelInfo,
		EnableLoki:   fflags.SendLogsToLoki.IsEnabled(),
	})
	slog.SetDefault(slog.New(lokiService))

	// Start the web server
	mux := &http.ServeMux{}
	port := "8081"

	appCfg := backend.Configuration{
		Mux: mux,
	}
	backend.Start(appCfg)

	go func() {
		chain := alice.New(
			middleware.Logging,
			middleware.RecoveryMiddleware,
		).Then(mux)

		err := http.ListenAndServe(":"+port, chain)
		if err != nil {
			slog.Error("Could not start server on port "+port, sloki.WrapError(err))
			os.Exit(1)
		}
	}()

	slog.Info(fmt.Sprintf("Started frontend service on http://localhost:%s\n", port))

	// Wait for a signal to exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	switch <-sig {
	case os.Interrupt:
		slog.Info("Received interrupt signal, shutting down...")

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
