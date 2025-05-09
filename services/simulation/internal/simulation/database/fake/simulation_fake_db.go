package fake

import (
	"context"
	"github.com/OliverSchlueter/mauerstrassenloewen/simulation/internal/simulation"
	"sync"
)

type DB struct {
	Data map[string]simulation.Simulation
	mu   sync.Mutex
}

func NewDB() *DB {
	return &DB{
		Data: make(map[string]simulation.Simulation),
		mu:   sync.Mutex{},
	}
}

func (db *DB) GetSimulationByID(ctx context.Context, id string) (*simulation.Simulation, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	s, exists := db.Data[id]
	if !exists {
		return nil, simulation.ErrSimulationNotFound
	}

	return &s, nil
}

func (db *DB) CreateSimulation(ctx context.Context, sim *simulation.Simulation) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, exists := db.Data[sim.ID]
	if exists {
		return simulation.ErrSimulationAlreadyExists
	}

	db.Data[sim.ID] = *sim

	return nil
}

func (db *DB) UpdateSimulation(ctx context.Context, sim *simulation.Simulation) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, exists := db.Data[sim.ID]
	if !exists {
		return simulation.ErrSimulationNotFound
	}

	db.Data[sim.ID] = *sim

	return nil
}

func (db *DB) DeleteSimulation(ctx context.Context, id string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, exists := db.Data[id]
	if !exists {
		return simulation.ErrSimulationNotFound
	}

	delete(db.Data, id)

	return nil
}
