package backend

import (
	_ "embed"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/openapi"
	"github.com/go-pg/pg/v10"
	"github.com/nats-io/nats.go"
	"github.com/questdb/go-questdb-client/v3"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
)

const apiPrefix = "/api/v1"
const frontendPrefix = ""

//go:embed openapi.yml
var openapiSpecContent []byte

type Configuration struct {
	Mux              *http.ServeMux
	Nats             *nats.Conn
	Mongo            *mongo.Database
	QuestDBIngestion *questdb.LineSender
	QuestDBQuery     *pg.DB
}

func Start(cfg Configuration) {
	openApiHandler := openapi.NewHandler(openapi.Configuration{
		Specification: openapiSpecContent,
	})
	openApiHandler.Register(cfg.Mux, frontendPrefix)
}
