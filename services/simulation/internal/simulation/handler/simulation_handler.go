package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"github.com/OliverSchlueter/mauerstrassenloewen/simulation/internal/simulation"
	"log/slog"
	"net/http"
)

type Store interface {
	GetSimulation(ctx context.Context, id string) (*simulation.Simulation, error)
	CreateSimulation(ctx context.Context, req *simulation.CreateOrUpdateSimulationRequest) error
	UpdateSimulation(ctx context.Context, sim *simulation.CreateOrUpdateSimulationRequest) error
	DeleteSimulation(ctx context.Context, id string) error
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
	mux.HandleFunc(prefix+"/simulation", h.handleSimulationCollection)
	mux.HandleFunc(prefix+"/simulation/{simulationID}", h.handleSimulation)
}

func (h *Handler) handleSimulationCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handleCreateSimulation(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) handleSimulation(w http.ResponseWriter, r *http.Request) {
	simID := r.PathValue("simulationID")
	if simID == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.handleGetSimulation(w, r, simID)
	case http.MethodPut:
		h.handleUpdateSimulation(w, r, simID)
	case http.MethodDelete:
		h.handleDeleteSimulation(w, r, simID)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) handleGetSimulation(w http.ResponseWriter, r *http.Request, simID string) {
	if r.Header.Get("Accept") != "application/json" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	sim, err := h.store.GetSimulation(r.Context(), simID)
	if err != nil {
		if errors.Is(err, simulation.ErrSimulationNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		slog.Error("Could not get simulation with id: "+simID, sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	simData, err := json.Marshal(sim)
	if err != nil {
		slog.Error("Could not marshal simulation response", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(simData)
}

func (h *Handler) handleCreateSimulation(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}

	if r.Header.Get("Accept") != "application/json" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var req simulation.CreateOrUpdateSimulationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Could not unmarshal request body", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := h.store.CreateSimulation(r.Context(), &req)
	if err != nil {
		if errors.Is(err, simulation.ErrSimulationAlreadyExists) {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		}

		slog.Error("Could not create simulation", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) handleUpdateSimulation(w http.ResponseWriter, r *http.Request, simID string) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}

	if r.Header.Get("Accept") != "application/json" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var req simulation.CreateOrUpdateSimulationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Could not unmarshal request body", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := h.store.UpdateSimulation(r.Context(), &req)
	if err != nil {
		if errors.Is(err, simulation.ErrSimulationNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		slog.Error("Could not update simulation", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) handleDeleteSimulation(w http.ResponseWriter, r *http.Request, simID string) {
	if r.Header.Get("Accept") != "application/json" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := h.store.DeleteSimulation(r.Context(), simID)
	if err != nil {
		if errors.Is(err, simulation.ErrSimulationNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		slog.Error("Could not delete simulation", sloki.WrapError(err), sloki.WrapRequest(r))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
