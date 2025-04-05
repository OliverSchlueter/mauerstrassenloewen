package backend

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/frontend"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/frontend/handler"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/openapi"
	"net/http"
)

const apiPrefix = "/api/v1"

type Configuration struct {
	Mux *http.ServeMux
}

func Start(cfg Configuration) {
	frontendHandler := handler.NewHandler(handler.Configuration{
		Files: frontend.Files,
	})
	frontendHandler.Register(cfg.Mux, "")

	openApiHandler := openapi.NewHandler()
	openApiHandler.Register(cfg.Mux, apiPrefix)
}
