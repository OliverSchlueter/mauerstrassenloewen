package handler

import (
	"embed"
	"log/slog"
	"net/http"
)

type Handler struct {
	files embed.FS
}

type Configuration struct {
	Files embed.FS
}

func NewHandler(cfg Configuration) *Handler {
	return &Handler{
		files: cfg.Files,
	}
}

func (h *Handler) Register(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(prefix+"/{$}", h.handleIndex)
	mux.HandleFunc(prefix+"/", h.handleAssets)
}

func (h *Handler) handleIndex(w http.ResponseWriter, r *http.Request) {
	file, err := h.files.ReadFile("assets/index.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		slog.Error("Could not read index.html", slog.Any("err", err.Error()))
		return
	}

	w.Write(file)
}

func (h *Handler) handleAssets(w http.ResponseWriter, r *http.Request) {
	path := "assets" + r.URL.Path

	file, err := h.files.ReadFile(path)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(file)
}
