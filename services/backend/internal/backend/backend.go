package backend

import (
	_ "embed"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/authentication"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/chatbot"
	ch "github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/chatbot/handler"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/usermanagement"
	ummongo "github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/usermanagement/database/mongo"
	umhandler "github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/usermanagement/handler"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
)

const apiPrefix = "/api/v1"

type Configuration struct {
	Mux     *http.ServeMux
	Nats    *nats.Conn
	MongoDB *mongo.Database
}

func Start(cfg Configuration) (authMiddleware func(next http.Handler) http.Handler) {
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

	umdb := ummongo.NewDB(&ummongo.Configuration{
		Mongo: cfg.MongoDB,
	})
	ums := usermanagement.NewStore(&usermanagement.Configuration{
		DB: umdb,
	})
	umh := umhandler.NewHandler(umhandler.Configuration{
		Store: ums,
	})
	umh.Register(cfg.Mux, apiPrefix)

	return authStore.Middleware
}
