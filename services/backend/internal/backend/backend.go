package backend

import (
	_ "embed"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/authentication"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/authentication/authhandler"
	amongo "github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/authentication/database/mongo"
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
	// User Management
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

	// Authentication
	adb := amongo.NewDB(amongo.Configuration{
		DB: cfg.MongoDB,
	})
	as := authentication.NewStore(authentication.StoreConfiguration{
		DB:          adb,
		UM:          ums,
		GlobalToken: "GlobalToken",
	})
	ah := authhandler.NewHandler(authhandler.Configuration{
		Store:   as,
		GetUser: authentication.UserFromCtx,
	})
	ah.Register(cfg.Mux, apiPrefix)

	// Chatbot
	chatbotService := chatbot.NewService(chatbot.Configuration{
		Nats: cfg.Nats,
	})
	chatbotHandler := ch.NewHandler(ch.Configuration{
		Service: chatbotService,
	})
	chatbotHandler.Register(cfg.Mux, apiPrefix)

	return as.Middleware
}
