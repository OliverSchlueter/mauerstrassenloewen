package backend

import (
	_ "embed"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/authentication"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/chatbot"
	ch "github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/chatbot/handler"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/docs"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/frontend"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/openapi"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
)

const apiPrefix = "/api/v1"
const frontendPrefix = ""

//go:embed openapi.yml
var openapiSpecContent []byte

type Configuration struct {
	Mux     *http.ServeMux
	Nats    *nats.Conn
	MongoDB *mongo.Database
}

func Start(cfg Configuration) (authMiddleware func(next http.Handler) http.Handler) {
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

	chatbotService := chatbot.NewService(chatbot.Configuration{
		Nats: cfg.Nats,
	})
	chatbotHandler := ch.NewHandler(ch.Configuration{
		Service: chatbotService,
	})
	chatbotHandler.Register(cfg.Mux, apiPrefix)

	authStore := authentication.NewStore(authentication.StoreConfiguration{
		GlobalToken: "GlobalToken",
	})

	return authStore.Middleware
}
