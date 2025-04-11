package handler

import (
	"bytes"
	"context"
	"github.com/OliverSchlueter/mauerstrassenloewen/simulation/internal/simulation"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetSimulation(t *testing.T) {
	tests := []struct {
		name           string
		simID          string
		mockStore      *mockStore
		expectedStatus int
	}{
		{
			name:  "Success",
			simID: "123",
			mockStore: &mockStore{
				GetSimulationFunc: func(ctx context.Context, id string) (*simulation.Simulation, error) {
					return &simulation.Simulation{ID: "123", Name: "Test Simulation"}, nil
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:  "Not Found",
			simID: "123",
			mockStore: &mockStore{
				GetSimulationFunc: func(ctx context.Context, id string) (*simulation.Simulation, error) {
					return nil, simulation.ErrSimulationNotFound
				},
			},
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewHandler(Configuration{Store: tt.mockStore})
			req := httptest.NewRequest(http.MethodGet, "/simulation/"+tt.simID, nil)
			req.Header.Set("Accept", "application/json")
			rr := httptest.NewRecorder()

			mux := http.NewServeMux()
			handler.Register(mux, "")
			mux.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
		})
	}
}

func TestHandleCreateSimulation(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		mockStore      *mockStore
		expectedStatus int
	}{
		{
			name: "Success",
			body: `{"name": "New Simulation"}`,
			mockStore: &mockStore{
				CreateSimulationFunc: func(ctx context.Context, req *simulation.CreateOrUpdateSimulationRequest) error {
					return nil
				},
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Conflict",
			body: `{"name": "Existing Simulation"}`,
			mockStore: &mockStore{
				CreateSimulationFunc: func(ctx context.Context, req *simulation.CreateOrUpdateSimulationRequest) error {
					return simulation.ErrSimulationAlreadyExists
				},
			},
			expectedStatus: http.StatusConflict,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewHandler(Configuration{Store: tt.mockStore})
			req := httptest.NewRequest(http.MethodPost, "/simulation", bytes.NewReader([]byte(tt.body)))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept", "application/json")
			rr := httptest.NewRecorder()

			mux := http.NewServeMux()
			handler.Register(mux, "")
			mux.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
		})
	}
}

func TestHandleUpdateSimulation(t *testing.T) {
	tests := []struct {
		name           string
		simID          string
		body           string
		mockStore      *mockStore
		expectedStatus int
	}{
		{
			name:  "Success",
			simID: "123",
			body:  `{"name": "Updated Simulation"}`,
			mockStore: &mockStore{
				UpdateSimulationFunc: func(ctx context.Context, req *simulation.CreateOrUpdateSimulationRequest) error {
					return nil
				},
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:  "Not Found",
			simID: "123",
			body:  `{"name": "Nonexistent Simulation"}`,
			mockStore: &mockStore{
				UpdateSimulationFunc: func(ctx context.Context, req *simulation.CreateOrUpdateSimulationRequest) error {
					return simulation.ErrSimulationNotFound
				},
			},
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewHandler(Configuration{Store: tt.mockStore})
			req := httptest.NewRequest(http.MethodPut, "/simulation/"+tt.simID, bytes.NewReader([]byte(tt.body)))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept", "application/json")
			rr := httptest.NewRecorder()

			mux := http.NewServeMux()
			handler.Register(mux, "")
			mux.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
		})
	}
}

func TestHandleDeleteSimulation(t *testing.T) {
	tests := []struct {
		name           string
		simID          string
		mockStore      *mockStore
		expectedStatus int
	}{
		{
			name:  "Success",
			simID: "123",
			mockStore: &mockStore{
				DeleteSimulationFunc: func(ctx context.Context, id string) error {
					return nil
				},
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:  "Not Found",
			simID: "123",
			mockStore: &mockStore{
				DeleteSimulationFunc: func(ctx context.Context, id string) error {
					return simulation.ErrSimulationNotFound
				},
			},
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewHandler(Configuration{Store: tt.mockStore})
			req := httptest.NewRequest(http.MethodDelete, "/simulation/"+tt.simID, nil)
			req.Header.Set("Accept", "application/json")
			rr := httptest.NewRecorder()

			mux := http.NewServeMux()
			handler.Register(mux, "")
			mux.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
		})
	}
}

type mockStore struct {
	GetSimulationFunc    func(ctx context.Context, id string) (*simulation.Simulation, error)
	CreateSimulationFunc func(ctx context.Context, req *simulation.CreateOrUpdateSimulationRequest) error
	UpdateSimulationFunc func(ctx context.Context, req *simulation.CreateOrUpdateSimulationRequest) error
	DeleteSimulationFunc func(ctx context.Context, id string) error
}

func (m *mockStore) GetSimulation(ctx context.Context, id string) (*simulation.Simulation, error) {
	return m.GetSimulationFunc(ctx, id)
}

func (m *mockStore) CreateSimulation(ctx context.Context, req *simulation.CreateOrUpdateSimulationRequest) error {
	return m.CreateSimulationFunc(ctx, req)
}

func (m *mockStore) UpdateSimulation(ctx context.Context, req *simulation.CreateOrUpdateSimulationRequest) error {
	return m.UpdateSimulationFunc(ctx, req)
}

func (m *mockStore) DeleteSimulation(ctx context.Context, id string) error {
	return m.DeleteSimulationFunc(ctx, id)
}
