package store

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/chatbot"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/natsdto"
)

type Store struct {
	Chats map[string]*natsdto.Chat
}

type Configuration struct {
}

func NewStore(cfg Configuration) *Store {
	return &Store{
		Chats: make(map[string]*natsdto.Chat),
	}
}

func (s *Store) GetChatByID(id string) (*natsdto.Chat, error) {
	chat, ok := s.Chats[id]
	if !ok {
		return nil, chatbot.ErrChatNotFound
	}

	return chat, nil
}

func (s *Store) UpsertChat(chat *natsdto.Chat) error {
	s.Chats[chat.ID] = chat
	return nil
}
