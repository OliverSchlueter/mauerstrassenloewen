package main

import (
	"context"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/frontend"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/frontend/handler"
	"github.com/OliverSchlueter/mauerstrassenloewen/libs/containers"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	ctx := context.Background()

	// Start test containers
	_, err := containers.StartMongoDB(ctx)
	if err != nil {
		slog.Error("Could not start MongoDB", slog.Any("err", err.Error()))
		os.Exit(1)
	}

	// Start the web server
	mux := &http.ServeMux{}
	port := mustGetPort()

	frontendHandler := handler.NewHandler(handler.Configuration{
		Files: frontend.Files,
	})
	frontendHandler.Register(mux, "")

	go func() {
		err := http.ListenAndServe(":"+port, mux)
		if err != nil {
			slog.Error("Could not start server", slog.Any("port", port), slog.Any("err", err.Error()))
			os.Exit(1)
		}
	}()

	slog.Info(fmt.Sprintf("Started server on http://localhost:%s\n", port))

	c := make(chan os.Signal, 1)
	<-c
}

func mustGetPort() string {
	if port := os.Getenv("MSL_BACKEND_PORT"); port != "" {
		return port
	}

	return "8080"
}
