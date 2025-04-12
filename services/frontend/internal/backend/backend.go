package backend

import (
	_ "embed"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/openapi"
	"github.com/OliverSchlueter/mauerstrassenloewen/frontend/internal/docs"
	"github.com/OliverSchlueter/mauerstrassenloewen/frontend/internal/frontend"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
)

const frontendPrefix = ""

//go:embed openapi.yml
var openapiSpecContent []byte

type Configuration struct {
	Mux     *http.ServeMux
	Nats    *nats.Conn
	MongoDB *mongo.Database
}

func Start(cfg Configuration) {
	frontendHandler := frontend.NewHandler(frontend.Configuration{
		Files: frontend.Files,
	})
	frontendHandler.Register(cfg.Mux, frontendPrefix)

	openApiHandler := openapi.NewHandler(openapi.Configuration{
		Specification: openapiSpecContent,
	})
	openApiHandler.Register(cfg.Mux, frontendPrefix)

	docsHandler := docs.NewHandler(docs.Configuration{
		Files: docs.Files,
	})
	docsHandler.Register(cfg.Mux)
}
