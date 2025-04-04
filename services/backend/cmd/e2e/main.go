package main

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/frontend"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/frontend/handler"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	slog.Info("Hallo Welt")

	mux := &http.ServeMux{}

	frontendHandler := handler.NewHandler(handler.Configuration{
		Files: frontend.Files,
	})
	frontendHandler.Register(mux, "")

	port := "8080"

	go func() {
		err := http.ListenAndServe(":"+port, mux)
		if err != nil {
			slog.Error("Could not start server", slog.Any("port", port), slog.Any("err", err.Error()))
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	<-c
}
