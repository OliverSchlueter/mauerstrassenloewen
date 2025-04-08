package chatbot

import (
	"common/natsdto"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

type Service struct {
	nats *nats.Conn
}

type Configuration struct {
	Nats *nats.Conn
}

func NewService(config Configuration) *Service {
	return &Service{
		nats: config.Nats,
	}
}

// NewPromptRequest creates a new prompt request for the chatbot and returns the job.
func (s *Service) NewPromptRequest(userMsg string, systemMsg natsdto.SystemMessage) (*natsdto.SimplePromptJob, error) {
	req := natsdto.SimplePromptRequest{
		UserMsg:   userMsg,
		SystemMsg: systemMsg,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("could not marshal request: %w", err)
	}

	cmd, err := s.nats.Request("msl.chatbot.simple_prompt", data, time.Second*50)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %w", err)
	}

	var resp natsdto.SimplePromptJob
	if err := json.Unmarshal(cmd.Data, &resp); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %w", err)
	}

	if resp.JobID == "" {
		return nil, fmt.Errorf("empty job ID in response")
	}

	return &resp, nil
}
