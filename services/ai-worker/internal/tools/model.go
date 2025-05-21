package tools

import "github.com/ollama/ollama/api"

type Tool struct {
	Type     string   `json:"type"`
	Function Function `json:"function"`
}

type Function struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Parameters  Parameters `json:"parameters"`
}

type Parameters struct {
	Type       string   `json:"type"`
	Required   []string `json:"required"`
	Properties map[string]struct {
		Type        string   `json:"type"`
		Description string   `json:"description"`
		Enum        []string `json:"enum,omitempty"`
	} `json:"properties"`
}

func (t *Tool) toOllama() api.Tool {
	return api.Tool{
		Type: t.Type,
		Function: api.ToolFunction{
			Name:        t.Function.Name,
			Description: t.Function.Description,
			Parameters: struct {
				Type       string   `json:"type"`
				Required   []string `json:"required"`
				Properties map[string]struct {
					Type        string   `json:"type"`
					Description string   `json:"description"`
					Enum        []string `json:"enum,omitempty"`
				} `json:"properties"`
			}{
				Type:       t.Function.Parameters.Type,
				Required:   t.Function.Parameters.Required,
				Properties: t.Function.Parameters.Properties,
			},
		},
	}
}
