package openapi

import (
	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"log/slog"
	"net/http"
)

type Handler struct {
	specification []byte
}

type Configuration struct {
	Specification []byte
}

func NewHandler(config Configuration) *Handler {
	return &Handler{
		specification: config.Specification,
	}
}

func (h *Handler) Register(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(prefix+"/openapi.yml", h.handleOpenAPI)
	mux.HandleFunc(prefix+"/openapi", h.handleScalar)
}

func (h *Handler) handleOpenAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/yaml")
	//if !fflags.EndToEndEnvironment.IsEnabled() {
	w.Header().Set("Cache-Control", "max-age=86400") // 24h
	//}
	w.WriteHeader(http.StatusOK)
	w.Write(h.specification)
}

func (h *Handler) handleScalar(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecContent: string(h.specification),
		CustomOptions: scalar.CustomOptions{
			PageTitle: "Mauerstrassenloewen API",
		},
		DarkMode: true,
	})
	if err != nil {
		slog.Error("Could not generate openapi html", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	//if !fflags.EndToEndEnvironment.IsEnabled() {
	w.Header().Set("Cache-Control", "max-age=86400") // 24h
	//}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(htmlContent))
}
