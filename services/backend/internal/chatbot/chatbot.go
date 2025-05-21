package chatbot

import (
	"encoding/json"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/natsdto"
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
func (s *Service) NewPromptRequest(userMsg string, systemMsg natsdto.SystemMessage) (*natsdto.Chat, error) {
	req := natsdto.StartChatRequest{
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

	var resp natsdto.Chat
	if err := json.Unmarshal(cmd.Data, &resp); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %w", err)
	}

	if resp.ID == "" {
		return nil, fmt.Errorf("empty job ID in response")
	}

	return &resp, nil
}

// StartChat starts a new chat with the given user message and system message.
func (s *Service) StartChat(req natsdto.StartChatRequest) (*natsdto.Chat, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("could not marshal request: %w", err)
	}

	cmd, err := s.nats.Request("msl.chatbot.start_chat", data, time.Second*50)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %w", err)
	}

	var resp natsdto.Chat
	if err := json.Unmarshal(cmd.Data, &resp); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %w", err)
	}

	if resp.ID == "" {
		return nil, fmt.Errorf("empty job ID in response")
	}

	return &resp, nil
}

// SendMessage sends a message to the chatbot and returns the response.
func (s *Service) SendMessage(req natsdto.SendChatMessageRequest) (*natsdto.Chat, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("could not marshal request: %w", err)
	}

	cmd, err := s.nats.Request("msl.chatbot.send_message", data, time.Second*50)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %w", err)
	}

	var resp natsdto.Chat
	if err := json.Unmarshal(cmd.Data, &resp); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %w", err)
	}

	if resp.ID == "" {
		return nil, fmt.Errorf("empty job ID in response")
	}

	return &resp, nil
}
