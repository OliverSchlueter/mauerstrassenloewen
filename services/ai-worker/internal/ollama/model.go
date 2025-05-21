package ollama

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/common/natsdto"
	"github.com/ollama/ollama/api"
)

func toOllamaMsgs(c *natsdto.Chat) []api.Message {
	var messages []api.Message
	for _, msg := range c.Messages {
		messages = append(messages, api.Message{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}
	return messages
}
