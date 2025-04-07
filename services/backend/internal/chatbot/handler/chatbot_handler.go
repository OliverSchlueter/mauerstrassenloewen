package handler

import (
	"encoding/json"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/chatbot"
	"io"
	"log/slog"
	"net/http"
)

type Handler struct {
	service *chatbot.Service
}

type Configuration struct {
	Service *chatbot.Service
}

func NewHandler(config Configuration) *Handler {
	return &Handler{
		service: config.Service,
	}
}

func (h *Handler) Register(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(prefix+"/chatbot/simple-prompt", h.handleSimplePrompt)
}

func (h *Handler) handleSimplePrompt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var req SimplePromptRequest
	if err := json.Unmarshal(body, &req); err != nil {
		slog.Error("Could not unmarshal request body", slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if req.Message == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	job, err := h.service.NewPromptRequest(req.Message, chatbot.FinancialAdvisor)
	if err != nil {
		slog.Error("Could not create new prompt request", slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	jobData, err := json.Marshal(job)
	if err != nil {
		slog.Error("Could not marshal job response", slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jobData)
}
