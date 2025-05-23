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

func (t *Tool) ToOllama() api.Tool {
	return api.Tool{
		Type: t.Type,
		Function: api.ToolFunction{
			Name:        t.Function.Name,
			Description: t.Function.Description,
			Parameters: struct {
				Type       string   `json:"type"`
				Defs       any      `json:"$defs,omitempty"`
				Items      any      `json:"items,omitempty"`
				Required   []string `json:"required"`
				Properties map[string]struct {
					Type        api.PropertyType `json:"type"`
					Items       any              `json:"items,omitempty"`
					Description string           `json:"description"`
					Enum        []any            `json:"enum,omitempty"`
				} `json:"properties"`
			}{
				Type:     "object",
				Defs:     nil,
				Items:    nil,
				Required: t.Function.Parameters.Required,
				Properties: map[string]struct {
					Type        api.PropertyType `json:"type"`
					Items       any              `json:"items,omitempty"`
					Description string           `json:"description"`
					Enum        []any            `json:"enum,omitempty"`
				}{},
			},
		},
	}
}

func ToOllama(tools []Tool) []api.Tool {
	ollamaTools := make([]api.Tool, len(tools))
	for i, tool := range tools {
		ollamaTools[i] = tool.ToOllama()
	}

	return ollamaTools
}
