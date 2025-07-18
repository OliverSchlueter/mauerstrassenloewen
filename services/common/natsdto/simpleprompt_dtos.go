package natsdto

import "time"

type SystemMessage string

type StartChatRequest struct {
	UserMsg   string `json:"user_msg"`
	SystemMsg string `json:"system_msg"`
}

type SendChatMessageRequest struct {
	ChatID  string `json:"chat_id"`
	UserMsg string `json:"user_msg"`
}

type GetChatRequest struct {
	ChatID string `json:"chat_id"`
}

type Chat struct {
	ID       string    `json:"id"`
	Messages []Message `json:"messages"`
}

func (c *Chat) AppendMsg(m Message) {
	c.Messages = append(c.Messages, m)
}

type Message struct {
	Role    string    `json:"role"`
	Content string    `json:"content"`
	SentAt  time.Time `json:"sent_at"`
}
