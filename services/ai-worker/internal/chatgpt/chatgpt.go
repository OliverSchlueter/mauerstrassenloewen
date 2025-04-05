package chatgpt

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

type Client struct {
	client *openai.Client
}

type Configuration struct {
	AuthToken string
}

func NewClient(cfg Configuration) *Client {
	client := openai.NewClient(cfg.AuthToken)

	return &Client{
		client: client,
	}
}

func (c *Client) Chat(ctx context.Context, message string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: message,
			},
		},
	}

	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to create chat completion: %w", err)
	}

	return resp.Choices[0].Message.Content, nil
}
