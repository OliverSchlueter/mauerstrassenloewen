package endpoint

type Endpoint struct {
	Name        string `json:"name"`
	Endpoint    string `json:"endpoint"`
	Destination string `json:"destination"`
}
