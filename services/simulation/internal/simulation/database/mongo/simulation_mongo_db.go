package mongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/simulation/internal/simulation"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DB struct {
	coll *mongo.Collection
}

type Configuration struct {
	Mongo *mongo.Database
}

func NewDB(config *Configuration) *DB {
	coll := config.Mongo.Collection("simulation")

	return &DB{
		coll: coll,
	}
}

func (db *DB) GetSimulationByID(ctx context.Context, id string) (*simulation.Simulation, error) {
	res := db.coll.FindOne(ctx, simulation.Simulation{
		ID: id,
	})

	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, simulation.ErrSimulationNotFound
		}
		return nil, fmt.Errorf("could not find simulation: %w", res.Err())
	}

	var sim simulation.Simulation
	if err := res.Decode(&sim); err != nil {
		return nil, fmt.Errorf("could not decode simulation: %w", err)
	}

	return &sim, nil
}

func (db *DB) CreateSimulation(ctx context.Context, sim *simulation.Simulation) error {
	_, err := db.coll.InsertOne(ctx, sim)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return simulation.ErrSimulationAlreadyExists
		}

		return fmt.Errorf("could not create simulation: %w", err)
	}

	return nil
}

func (db *DB) UpdateSimulation(ctx context.Context, sim *simulation.Simulation) error {
	_, err := db.coll.UpdateOne(ctx, simulation.Simulation{ID: sim.ID}, bson.M{"$set": sim})
	if err != nil {
		return fmt.Errorf("could not update simulation: %w", err)
	}

	return nil
}

func (db *DB) DeleteSimulation(ctx context.Context, id string) error {
	_, err := db.coll.DeleteOne(ctx, simulation.Simulation{ID: id})
	if err != nil {
		return fmt.Errorf("could not delete simulation: %w", err)
	}

	return nil
}
