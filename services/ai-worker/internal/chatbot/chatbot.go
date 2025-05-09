package chatbot

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/ollama"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/natsdto"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type Service struct {
	nats   *nats.Conn
	ollama *ollama.Client
}

type Configuration struct {
	Nats   *nats.Conn
	Ollama *ollama.Client
}

func NewService(cfg Configuration) *Service {
	return &Service{
		nats:   cfg.Nats,
		ollama: cfg.Ollama,
	}
}

func (s *Service) Register() error {
	if _, err := s.nats.Subscribe("msl.chatbot.simple_prompt", s.handleSimplePromptHandler); err != nil {
		return fmt.Errorf("could not subscribe to nats subject: %w", err)
	}

	return nil
}

func (s *Service) handleSimplePromptHandler(msg *nats.Msg) {
	var req natsdto.SimplePromptRequest
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to unmarshal request: %v", err)))
		return
	}

	output, err := s.ollama.Generate(context.Background(), req.UserMsg)
	if err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to get response from ollama: %v", err)))
		return
	}

	resp := natsdto.SimplePromptJob{
		JobID:  uuid.New().String(),
		Result: output,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to marshal response: %v", err)))
		return
	}

	if err := msg.Respond(data); err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to respond to request: %v", err)))
		return
	}
}
