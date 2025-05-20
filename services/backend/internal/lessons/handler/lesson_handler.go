package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/authentication"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/lessons"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/usermanagement"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"log/slog"
	"net/http"
)

type Store interface {
	GetLesson(ctx context.Context, userID string) (*lessons.Lesson, error)
	UpsertLesson(ctx context.Context, userID, lessonID string) error
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
	mux.HandleFunc(prefix+"/lesson", h.handleGetLesson)
	mux.HandleFunc(prefix+"/lesson/{lesson_id}", h.handleUpsertLesson)
}

func (h *Handler) handleGetLesson(w http.ResponseWriter, r *http.Request) {
	user := authentication.UserFromCtx(r.Context())
	if r.Header.Get("Accept") != "application/json" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	lesson, err := h.store.GetLesson(r.Context(), user.ID)
	if err != nil {
		if errors.Is(err, usermanagement.ErrUserNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		slog.Error("Could not get lesson for user: "+user.ID, sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	lessonData, err := json.Marshal(lesson)
	if err != nil {
		slog.Error("Could not marshal lesson response", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(lessonData)
}

func (h *Handler) handleUpsertLesson(w http.ResponseWriter, r *http.Request) {
	user := authentication.UserFromCtx(r.Context())
	if r.Method != http.MethodPut {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	lessonID := r.PathValue("lesson_id")
	if lessonID == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := h.store.UpsertLesson(r.Context(), user.ID, lessonID)
	if err != nil {
		if errors.Is(err, usermanagement.ErrUserNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		slog.Error("Could not upsert lesson with id: "+lessonID, sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
