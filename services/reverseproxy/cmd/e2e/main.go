package main

import (
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

func main() {
	fflags.EndToEndEnvironment.Enable()
	//fflags.SendLogsToLoki.Enable()

	// Setup logging
	lokiService := sloki.NewService(sloki.Configuration{
		URL:          "http://localhost:3100/loki/api/v1/push",
		Service:      "reverseproxy",
		ConsoleLevel: slog.LevelDebug,
		LokiLevel:    slog.LevelInfo,
		EnableLoki:   fflags.SendLogsToLoki.IsEnabled(),
	})
	slog.SetDefault(slog.New(lokiService))

	mux := http.NewServeMux()
	port := "8080"

	backend.Start(backend.Configuration{
		Mux: mux,
		Endpoints: []endpoint.Endpoint{
			{
				Name:        "Frontend",
				Endpoint:    "",
				Destination: "http://localhost:8081",
			},
			{
				Name:        "Backend",
				Endpoint:    "/msl",
				Destination: "http://localhost:8082",
			},
			{
				Name:        "Simulation",
				Endpoint:    "/simulation",
				Destination: "http://localhost:8083",
			},
		},
	})

	go func() {
		err := http.ListenAndServe(":"+port, mux)
		if err != nil {
			slog.Error("Could not start server on port "+port, slog.Any("err", err.Error()))
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
