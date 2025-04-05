package openapi

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/featureflags"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Register(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(prefix+"/openapi.json", h.handleOpenAPI)
}

func (h *Handler) handleOpenAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !featureflags.EndToEndEnvironment.IsEnabled() {
		w.Header().Set("Cache-Control", "max-age=86400") // 24h
	}
	w.WriteHeader(http.StatusOK)
	w.Write(OpenApiSpec)
}
