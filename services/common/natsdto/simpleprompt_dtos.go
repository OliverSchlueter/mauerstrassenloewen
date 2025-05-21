package natsdto

import "time"

type SystemMessage string

type StartChatRequest struct {
	UserMsg   string        `json:"user_msg"`
	SystemMsg SystemMessage `json:"system_msg"`
}

type Chat struct {
	ChatID   string    `json:"chat_id"`
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
