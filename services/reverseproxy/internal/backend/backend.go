package backend

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/common/healthcheck"
	"github.com/OliverSchlueter/mauerstrassenloewen/reverseproxy/internal/endpoint"
	"github.com/OliverSchlueter/mauerstrassenloewen/reverseproxy/internal/proxy"
	"net/http"
	"net/url"
)

type Configuration struct {
	Mux       *http.ServeMux
	Endpoints []endpoint.Endpoint
}

func Start(cfg Configuration) {
	for _, e := range cfg.Endpoints {
		dest := urlMustParse(e.Destination)
		p := proxy.NewProxy(dest)
		handler := proxy.RequestHandler(p, dest, e)
		cfg.Mux.HandleFunc(e.Endpoint+"/", handler)
	}

	// Healthcheck
	healthcheck.Register(cfg.Mux)
}

func urlMustParse(raw string) *url.URL {
	u, _ := url.Parse(raw)
	return u
}
