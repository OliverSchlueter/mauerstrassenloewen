package chatgpt

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

type Client struct {
	client *openai.Client
	model  string
}

type Configuration struct {
	AuthToken string
	Model     string
}

func NewClient(cfg Configuration) *Client {
	client := openai.NewClient(cfg.AuthToken)

	return &Client{
		client: client,
		model:  cfg.Model,
	}
}

func (c *Client) Chat(ctx context.Context, message string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: c.model,
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
