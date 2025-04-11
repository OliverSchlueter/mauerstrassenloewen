package openapi

import (
	"common/sloki"
	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/fflags"
	"log/slog"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Register(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(prefix+"/openapi.json", h.handleOpenAPI)
	mux.HandleFunc(prefix+"/openapi", h.handleScalar)
}

func (h *Handler) handleOpenAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !fflags.EndToEndEnvironment.IsEnabled() {
		w.Header().Set("Cache-Control", "max-age=86400") // 24h
	}
	w.WriteHeader(http.StatusOK)
	w.Write(SpecContent)
}

func (h *Handler) handleScalar(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecContent: string(SpecContent),
		CustomOptions: scalar.CustomOptions{
			PageTitle: "Mauerstrassenloewen API",
		},
		DarkMode: true,
	})
	if err != nil {
		slog.Error("Could not generate openapi html", slog.Any("err", err.Error()), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if !fflags.EndToEndEnvironment.IsEnabled() {
		w.Header().Set("Cache-Control", "max-age=86400") // 24h
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(htmlContent))
}
