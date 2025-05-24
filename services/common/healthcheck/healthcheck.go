package healthcheck

import (
	"fmt"
	"net/http"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/health", handleHealthCheck)
}

func handleHealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, http.StatusText(http.StatusOK))
}
