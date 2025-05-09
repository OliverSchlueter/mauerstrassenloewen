package authhandler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/authentication"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/usermanagement"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"log/slog"
	"net/http"
)

type Store interface {
	CreateAuthToken(ctx context.Context, userID string) (string, error)
	GetAuthToken(ctx context.Context, tokenID string) (*authentication.Token, error)
	GetAuthTokensByUserID(ctx context.Context, userID string) ([]authentication.Token, error)
	DeleteAuthToken(ctx context.Context, tokenID string) error
}

type Handler struct {
	store   Store
	getUser func(ctx context.Context) *usermanagement.User
}

type Configuration struct {
	Store   Store
	GetUser func(ctx context.Context) *usermanagement.User
}

func NewHandler(cfg Configuration) *Handler {
	return &Handler{
		store:   cfg.Store,
		getUser: cfg.GetUser,
	}
}

func (h *Handler) Register(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(prefix+"/auth-token", h.handleAuthTokens)
	mux.HandleFunc(prefix+"/auth-token/{tokenid}", h.handleAuthToken)
}

func (h *Handler) handleAuthTokens(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getUserTokens(w, r)
	case http.MethodPost:
		h.createAuthToken(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handleAuthToken(w http.ResponseWriter, r *http.Request) {
	tid := r.PathValue("tokenid")

	switch r.Method {
	case http.MethodDelete:
		h.deleteAuthToken(tid, w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getUserTokens(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") != "application/json" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	usr := h.getUser(r.Context())

	tokens, err := h.store.GetAuthTokensByUserID(r.Context(), usr.ID)
	if err != nil {
		slog.Error("Could not get user tokens", sloki.WrapError(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(tokens)
	if err != nil {
		slog.Error("Could not marshal tokens", sloki.WrapError(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) createAuthToken(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") != "text/plain" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	usr := h.getUser(r.Context())

	token, err := h.store.CreateAuthToken(r.Context(), usr.ID)
	if err != nil {
		slog.Error("Could not create auth token", sloki.WrapError(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(token))
}

func (h *Handler) deleteAuthToken(tokenid string, w http.ResponseWriter, r *http.Request) {
	usr := h.getUser(r.Context())

	token, err := h.store.GetAuthToken(r.Context(), tokenid)
	if err != nil {
		if errors.Is(err, authentication.ErrTokenNotFound) {
			return
		}
		slog.Error("Could not get auth token", sloki.WrapError(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if token.UserID != usr.ID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if err := h.store.DeleteAuthToken(r.Context(), tokenid); err != nil {
		slog.Error("Could not delete auth token", sloki.WrapError(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
