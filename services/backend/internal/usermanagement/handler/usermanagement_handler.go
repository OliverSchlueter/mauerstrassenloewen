package handler

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
	GetUserByID(ctx context.Context, id string) (*usermanagement.User, error)
	CreateUser(ctx context.Context, user *usermanagement.User) error
	UpdateUser(ctx context.Context, user *usermanagement.User) error
	DeleteUser(ctx context.Context, id string) error
}

type Handler struct {
	store Store
}

type Configuration struct {
	Store Store
}

func NewHandler(config Configuration) *Handler {
	return &Handler{
		store: config.Store,
	}
}

func (h *Handler) Register(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(prefix+"/user/register", h.handleRegisterUser)
	mux.HandleFunc(prefix+"/user/me", h.handleMe)
	mux.HandleFunc(prefix+"/user/{id}", h.handleUser)
}

func (h *Handler) handleUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")
	if userID == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.handleGetUser(w, r, userID)
	case http.MethodPut:
		h.handleUpdateUser(w, r, userID)
	case http.MethodDelete:
		h.handleDeleteUser(w, r, userID)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}

	if r.Header.Get("Accept") != "application/json" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var req usermanagement.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Could not unmarshal request body", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := h.store.CreateUser(r.Context(), &req)
	if err != nil {
		if errors.Is(err, usermanagement.ErrUserAlreadyExists) {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		}

		slog.Error("Could not create user", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request, userID string) {
	if r.Header.Get("Accept") != "application/json" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := h.store.GetUserByID(r.Context(), userID)
	if err != nil {
		if errors.Is(err, usermanagement.ErrUserNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		slog.Error("Could not get user with id: "+userID, sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	userData, err := json.Marshal(user)
	if err != nil {
		slog.Error("Could not marshal user response", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userData)
}

func (h *Handler) handleMe(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") != "application/json" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user := authentication.UserFromCtx(r.Context())

	userData, err := json.Marshal(user)
	if err != nil {
		slog.Error("Could not marshal user response", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userData)
}

func (h *Handler) handleUpdateUser(w http.ResponseWriter, r *http.Request, userID string) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}

	if r.Header.Get("Accept") != "application/json" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var req usermanagement.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Could not unmarshal request body", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	req.ID = userID

	err := h.store.UpdateUser(r.Context(), &req)
	if err != nil {
		if errors.Is(err, usermanagement.ErrUserNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		slog.Error("Could not update user", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) handleDeleteUser(w http.ResponseWriter, r *http.Request, userID string) {
	if r.Header.Get("Accept") != "application/json" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := h.store.DeleteUser(r.Context(), userID)
	if err != nil {
		if errors.Is(err, usermanagement.ErrUserNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		slog.Error("Could not delete user", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
