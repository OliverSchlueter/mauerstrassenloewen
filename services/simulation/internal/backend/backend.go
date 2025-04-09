package backend

import (
	"github.com/go-pg/pg/v10"
	"github.com/nats-io/nats.go"
	"github.com/questdb/go-questdb-client/v3"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
)

type Configuration struct {
	Mux              *http.ServeMux
	Nats             *nats.Conn
	Mongo            *mongo.Database
	QuestDBIngestion *questdb.LineSender
	QuestDBQuery     *pg.DB
}

func Start(cfg Configuration) {

}
