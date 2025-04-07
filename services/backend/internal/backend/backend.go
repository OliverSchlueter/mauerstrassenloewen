package backend

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/chatbot"
	ch "github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/chatbot/handler"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/frontend"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/openapi"
	"github.com/nats-io/nats.go"
	"net/http"
)

const apiPrefix = "/api/v1"

type Configuration struct {
	Mux  *http.ServeMux
	Nats *nats.Conn
}

func Start(cfg Configuration) {
	frontendHandler := frontend.NewHandler(frontend.Configuration{
		Files: frontend.Files,
	})
	frontendHandler.Register(cfg.Mux, "")

	openApiHandler := openapi.NewHandler()
	openApiHandler.Register(cfg.Mux, apiPrefix)

	chatbotService := chatbot.NewService(chatbot.Configuration{
		Nats: cfg.Nats,
	})
	chatbotHandler := ch.NewHandler(ch.Configuration{
		Service: chatbotService,
	})
	chatbotHandler.Register(cfg.Mux, apiPrefix)
}
