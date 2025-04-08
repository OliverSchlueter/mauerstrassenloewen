package backend

import (
	"github.com/nats-io/nats.go"
	"net/http"
)

type Configuration struct {
	Mux  *http.ServeMux
	Nats *nats.Conn
}

func Start(cfg Configuration) {

}
