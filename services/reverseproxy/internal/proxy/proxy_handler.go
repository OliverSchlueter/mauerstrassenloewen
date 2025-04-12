package proxy

import (
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"github.com/OliverSchlueter/mauerstrassenloewen/reverseproxy/internal/endpoint"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func NewProxy(target *url.URL) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		slog.Warn("Error while proxying request", slog.String("error", err.Error()))
		http.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)
	}

	return proxy
}

func RequestHandler(proxy *httputil.ReverseProxy, url *url.URL, e endpoint.Endpoint) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Request received", sloki.WrapRequest(r))

		r.Host = url.Host
		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

		path := r.URL.Path
		r.URL.Path = strings.TrimPrefix(path, e.Endpoint)

		slog.Info(fmt.Sprintf("Redirecting"), slog.String("endpoint", e.Name), slog.String("to", r.URL.String()))
		proxy.ServeHTTP(w, r)
	}
}
