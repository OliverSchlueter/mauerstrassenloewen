package simulation

import (
	"context"
	"github.com/google/uuid"
)

type Database interface {
	GetSimulationByID(ctx context.Context, id string) (*Simulation, error)
	CreateSimulation(ctx context.Context, sim *Simulation) error
	UpdateSimulation(ctx context.Context, sim *Simulation) error
	DeleteSimulation(ctx context.Context, id string) error
}

type Store struct {
	db Database
}

type Configuration struct {
	DB Database
}

func NewStore(config *Configuration) *Store {
	return &Store{
		db: config.DB,
	}
}

func (s *Store) GetSimulation(ctx context.Context, id string) (*Simulation, error) {
	return s.db.GetSimulationByID(ctx, id)
}

func (s *Store) CreateSimulation(ctx context.Context, sim *Simulation) error {
	newSim := &Simulation{
		ID:   uuid.NewString(),
		Name: sim.Name,
	}

	return s.db.CreateSimulation(ctx, newSim)
}

func (s *Store) UpdateSimulation(ctx context.Context, sim *Simulation) error {
	return s.db.UpdateSimulation(ctx, sim)
}

func (s *Store) DeleteSimulation(ctx context.Context, id string) error {
	return s.db.DeleteSimulation(ctx, id)
}
