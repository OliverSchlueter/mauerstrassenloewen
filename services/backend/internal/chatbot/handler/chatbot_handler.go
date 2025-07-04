package handler

import (
	"encoding/json"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/authentication"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/chatbot"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/natsdto"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
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

	mux.HandleFunc(prefix+"/chatbot/chat", h.handleStartChat)
	mux.HandleFunc(prefix+"/chatbot/chat/{id}", h.handleGetChat)
	mux.HandleFunc(prefix+"/chatbot/chat/{id}/new-message", h.handleSendMessage)
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
		slog.Error("Could not unmarshal request body", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if req.Message == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	job, err := h.service.NewPromptRequest(req.Message, chatbot.FinancialAdvisor)
	if err != nil {
		slog.Error("Could not create new prompt request", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	jobData, err := json.Marshal(job)
	if err != nil {
		slog.Error("Could not marshal job response", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jobData)
}

func (h *Handler) handleStartChat(w http.ResponseWriter, r *http.Request) {
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

	var req natsdto.StartChatRequest
	if err := json.Unmarshal(body, &req); err != nil {
		slog.Error("Could not unmarshal request body", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if req.SystemMsg == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	u := authentication.UserFromCtx(r.Context())
	profileData, err := json.Marshal(u.Profile)
	if err != nil {
		slog.Error("Could not marshal user data", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	req.UserMsg += "\n\nUser Data:\n"
	req.UserMsg += fmt.Sprintf("I am %s (%s) and my email is %s.\n", u.Name, u.ID, u.Email)
	req.UserMsg += fmt.Sprintf("My profile data is: %s\n", string(profileData))

	chat, err := h.service.StartChat(req)
	if err != nil {
		slog.Error("Could not create new chat request", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	chatData, err := json.Marshal(chat)
	if err != nil {
		slog.Error("Could not marshal chat response", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(chatData)
}

func (h *Handler) handleGetChat(w http.ResponseWriter, r *http.Request) {
	cid := r.PathValue("id")
	if cid == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	req := natsdto.GetChatRequest{ChatID: cid}
	chat, err := h.service.GetChat(req)
	if err != nil {
		slog.Error("Could not get chat", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	chatData, err := json.Marshal(chat)
	if err != nil {
		slog.Error("Could not marshal chat response", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(chatData)
}

func (h *Handler) handleSendMessage(w http.ResponseWriter, r *http.Request) {
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

	var req natsdto.SendChatMessageRequest
	if err := json.Unmarshal(body, &req); err != nil {
		slog.Error("Could not unmarshal request body", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	cid := r.PathValue("id")
	req.ChatID = cid

	if req.UserMsg == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	chat, err := h.service.SendMessage(req)
	if err != nil {
		slog.Error("Could not create new chat request", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	chatData, err := json.Marshal(chat)
	if err != nil {
		slog.Error("Could not marshal chat response", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(chatData)
}
