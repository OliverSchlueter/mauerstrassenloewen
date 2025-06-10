package backend

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/chatbot"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/chatbot/store"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/ollama"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/healthcheck"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"github.com/nats-io/nats.go"
	"log/slog"
	"net/http"
)

type Configuration struct {
	Mux          *http.ServeMux
	Nats         *nats.Conn
	OllamaClient *ollama.Client
}

func Start(cfg Configuration) {
	chatbotStore := store.NewStore(store.Configuration{})
	chatbotService := chatbot.NewService(chatbot.Configuration{
		Store:  chatbotStore,
		Nats:   cfg.Nats,
		Ollama: cfg.OllamaClient,
	})
	if err := chatbotService.Register(); err != nil {
		slog.Error("failed to register chatbot service", sloki.WrapError(err))
		return
	}

	healthcheck.Register(cfg.Mux)
}
