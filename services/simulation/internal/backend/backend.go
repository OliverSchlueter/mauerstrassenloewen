package backend

import (
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
)

type Configuration struct {
	Mux   *http.ServeMux
	Nats  *nats.Conn
	Mongo *mongo.Database
}

func Start(cfg Configuration) {

}
