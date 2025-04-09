package handler

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/simulation/internal/simulation"
	"net/http"
)

type Store interface {
	GetSimulation(id string) (*simulation.Simulation, error)
	CreateSimulation(sim *simulation.Simulation) error
	UpdateSimulation(sim *simulation.Simulation) error
	DeleteSimulation(id string) error
}

type Handler struct {
	store *Store
}

type Configuration struct {
	Store *Store
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
	// TODO: Implement handleGetSimulation
}

func (h *Handler) handleCreateSimulation(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handleCreateSimulation
}

func (h *Handler) handleUpdateSimulation(w http.ResponseWriter, r *http.Request, simID string) {
	// TODO: Implement handleUpdateSimulation
}

func (h *Handler) handleDeleteSimulation(w http.ResponseWriter, r *http.Request, simID string) {
	// TODO: Implement handleDeleteSimulation
}
