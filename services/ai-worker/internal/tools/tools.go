package tools

type Service struct {
	tools []Tool
}

func NewService() *Service {
	return &Service{
		tools: []Tool{
			{
				Type: "function",
				Function: Function{
					Name:        "get_user_info",
					Description: "Gets user information",
					Parameters: Parameters{
						Type:     "object",
						Required: []string{},
						Properties: map[string]struct {
							Type        string   `json:"type"`
							Description string   `json:"description"`
							Enum        []string `json:"enum,omitempty"`
						}{},
					},
				},
			},
		},
	}
}

func (s *Service) GetTools() []Tool {
	return s.tools
}
