package simulation

type Database interface {
	GetSimulationByID(id string) (*Simulation, error)
	CreateSimulation(simulation *Simulation) error
	UpdateSimulation(simulation *Simulation) error
	DeleteSimulation(id string) error
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

func (s *Store) GetSimulation(id string) (*Simulation, error) {
	return s.db.GetSimulationByID(id)
}

func (s *Store) CreateSimulation(simulation *Simulation) error {
	return s.db.CreateSimulation(simulation)
}

func (s *Store) UpdateSimulation(simulation *Simulation) error {
	return s.db.UpdateSimulation(simulation)
}

func (s *Store) DeleteSimulation(id string) error {
	return s.db.DeleteSimulation(id)
}
