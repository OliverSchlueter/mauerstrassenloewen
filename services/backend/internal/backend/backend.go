package backend

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/frontend"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/frontend/handler"
	"net/http"
)

type Configuration struct {
	Mux *http.ServeMux
}

func Start(cfg Configuration) {
	frontendHandler := handler.NewHandler(handler.Configuration{
		Files: frontend.Files,
	})
	frontendHandler.Register(cfg.Mux, "")
}
