package simulation

type Database interface {
	GetSimulationByID(id string) (*Simulation, error)
	CreateSimulation(sim *Simulation) error
	UpdateSimulation(sim *Simulation) error
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

func (s *Store) CreateSimulation(sim *Simulation) error {
	return s.db.CreateSimulation(sim)
}

func (s *Store) UpdateSimulation(sim *Simulation) error {
	return s.db.UpdateSimulation(sim)
}

func (s *Store) DeleteSimulation(id string) error {
	return s.db.DeleteSimulation(id)
}
