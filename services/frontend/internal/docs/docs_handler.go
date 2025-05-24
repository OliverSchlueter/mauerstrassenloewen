package docs

import (
	"embed"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"github.com/OliverSchlueter/mauerstrassenloewen/frontend/internal/fflags"
	"log/slog"
	"net/http"
	"strings"
)

var contentTypes = map[string]string{
	".css":   "text/css",
	".js":    "application/javascript",
	".html":  "text/html",
	".json":  "application/json",
	".png":   "image/png",
	".jpg":   "image/jpeg",
	".gif":   "image/gif",
	".svg":   "image/svg+xml",
	".woff":  "font/woff",
	".woff2": "font/woff2",
	".ttf":   "font/ttf",
	".ico":   "image/x-icon",
	".webp":  "image/webp",
	".mp4":   "video/mp4",
	".mp3":   "audio/mpeg",
	".ogg":   "audio/ogg",
	".wav":   "audio/wav",
	".pdf":   "application/pdf",
	".xml":   "application/xml",
	".zip":   "application/zip",
	".tar":   "application/x-tar",
	".gz":    "application/gzip",
	".xz":    "application/x-xz",
	".rar":   "application/x-rar-compressed",
	".csv":   "text/csv",
	".txt":   "text/plain",
}

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

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/docs/{$}", h.handleIndex)
	mux.HandleFunc("/docs/{path...}", h.handleAssets)
}

func (h *Handler) handleIndex(w http.ResponseWriter, r *http.Request) {
	file, err := h.files.ReadFile("assets/index.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		slog.Error("Could not read index.html", sloki.WrapError(err), sloki.WrapRequest(r))
		return
	}

	if !fflags.EndToEndEnvironment.IsEnabled() {
		w.Header().Set("Cache-Control", "max-age=3600")
	}
	w.Write(file)
}

func (h *Handler) handleAssets(w http.ResponseWriter, r *http.Request) {
	path := "assets/" + strings.TrimPrefix(r.URL.Path, "/docs/")

	if r.URL.Path != "/" && strings.HasSuffix(path, "/") {
		path += "index.html"
	}

	file, err := h.files.ReadFile(path)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	for ext, ct := range contentTypes {
		if strings.HasSuffix(path, ext) {
			w.Header().Set("Content-Type", ct)
			break
		}
	}

	if !fflags.EndToEndEnvironment.IsEnabled() {
		w.Header().Set("Cache-Control", "max-age=3600")
	}
	w.Write(file)
}
