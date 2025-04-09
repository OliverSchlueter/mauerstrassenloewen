package simulation

import "errors"

var (
	ErrSimulationNotFound      = errors.New("could not find simulation")
	ErrSimulationAlreadyExists = errors.New("simulation already exists")
)
