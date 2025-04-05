package backend

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/frontend"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/openapi"
	"net/http"
)

const apiPrefix = "/api/v1"

type Configuration struct {
	Mux *http.ServeMux
}

func Start(cfg Configuration) {
	frontendHandler := frontend.NewHandler(frontend.Configuration{
		Files: frontend.Files,
	})
	frontendHandler.Register(cfg.Mux, "")

	openApiHandler := openapi.NewHandler()
	openApiHandler.Register(cfg.Mux, apiPrefix)
}
