package metricshandler

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func Register(mux *http.ServeMux, prefix string) {
	mux.Handle(prefix+"/metrics", promhttp.Handler())
}
