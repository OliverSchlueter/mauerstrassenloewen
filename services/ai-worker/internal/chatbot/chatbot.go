package chatbot

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/ollama"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/natsdto"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"time"
)

type Store interface {
	GetChatByID(id string) (*natsdto.Chat, error)
	UpsertChat(chat *natsdto.Chat) error
}

type Service struct {
	store  Store
	nats   *nats.Conn
	ollama *ollama.Client
}

type Configuration struct {
	Store  Store
	Nats   *nats.Conn
	Ollama *ollama.Client
}

func NewService(cfg Configuration) *Service {
	return &Service{
		store:  cfg.Store,
		nats:   cfg.Nats,
		ollama: cfg.Ollama,
	}
}

func (s *Service) Register() error {
	if _, err := s.nats.Subscribe("msl.chatbot.simple_prompt", s.handleSimplePrompt); err != nil {
		return fmt.Errorf("could not subscribe to nats subject: %w", err)
	}

	if _, err := s.nats.Subscribe("msl.chatbot.start_chat", s.handleStartChat); err != nil {
		return fmt.Errorf("could not subscribe to nats subject: %w", err)
	}

	if _, err := s.nats.Subscribe("msl.chatbot.send_chat_message", s.handleSendChatMessage); err != nil {
		return fmt.Errorf("could not subscribe to nats subject: %w", err)
	}

	return nil
}

func (s *Service) handleSimplePrompt(msg *nats.Msg) {
	receivedAt := time.Now()

	var req natsdto.StartChatRequest
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to unmarshal request: %v", err)))
		return
	}

	output, err := s.ollama.Generate(context.Background(), req.UserMsg)
	if err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to get response from ollama: %v", err)))
		return
	}

	resp := natsdto.Chat{
		ChatID: uuid.New().String(),
		Messages: []natsdto.Message{
			{
				Role:    "system",
				Content: string(req.SystemMsg),
				SentAt:  receivedAt,
			},
			{
				Role:    "user",
				Content: req.UserMsg,
				SentAt:  receivedAt,
			},
			{
				Role:    "assistant",
				Content: output,
				SentAt:  time.Now(),
			},
		},
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

func (s *Service) handleStartChat(msg *nats.Msg) {
	var req natsdto.StartChatRequest
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to unmarshal request: %v", err)))
		return
	}

	chat, err := s.ollama.StartChat(context.Background(), req)
	if err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to get response from ollama: %v", err)))
		return
	}

	err = s.store.UpsertChat(chat)
	if err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to upsert chat: %v", err)))
		return
	}

	data, err := json.Marshal(chat)
	if err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to marshal response: %v", err)))
		return
	}

	if err := msg.Respond(data); err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to respond to request: %v", err)))
		return
	}
}

func (s *Service) handleSendChatMessage(msg *nats.Msg) {
	var req natsdto.SendChatMessageRequest
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to unmarshal request: %v", err)))
		return
	}

	chat, err := s.store.GetChatByID(req.ChatID)
	if err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to get chat by ID: %v", err)))
		return
	}

	err = s.ollama.Chat(context.Background(), chat, req.UserMsg)
	if err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to get response from ollama: %v", err)))
		return
	}

	err = s.store.UpsertChat(chat)
	if err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to upsert chat: %v", err)))
		return
	}

	data, err := json.Marshal(chat)
	if err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to marshal response: %v", err)))
		return
	}

	if err := msg.Respond(data); err != nil {
		s.nats.Publish(msg.Reply, []byte(fmt.Sprintf("failed to respond to request: %v", err)))
		return
	}
}
