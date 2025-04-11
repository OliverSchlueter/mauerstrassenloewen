package backend

import (
	"net/http"
)

type Configuration struct {
	Mux *http.ServeMux
}

func Start(cfg Configuration) {

}
