package simulation

type Simulation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateOrUpdateSimulationRequest struct {
	Name string `json:"name"`
}
