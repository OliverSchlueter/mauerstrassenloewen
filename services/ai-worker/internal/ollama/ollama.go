package ollama

import (
	"context"
	"fmt"
	"github.com/ollama/ollama/api"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var _false = false
var _true = true

type Client struct {
	ollama *api.Client
	model  string
}

type Configuration struct {
	BaseURL string
	Model   string
}

func NewClient(cfg Configuration) (*Client, error) {
	baseURL, err := url.Parse(cfg.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	httpClient := &http.Client{
		Timeout: 10 * time.Minute,
	}

	ollama := api.NewClient(baseURL, httpClient)

	return &Client{
		ollama: ollama,
		model:  cfg.Model,
	}, nil
}

func (c *Client) Generate(ctx context.Context, message string) (string, error) {
	req := api.GenerateRequest{
		Model:  c.model,
		Prompt: message,
		Stream: &_true,
		Raw:    false,
	}

	var resp strings.Builder
	respFunc := func(gr api.GenerateResponse) error {
		_, err := resp.WriteString(gr.Response)
		if err != nil {
			return fmt.Errorf("failed to write response: %w", err)
		}

		return nil
	}

	slog.Info("Generating response", slog.String("model", c.model))

	err := c.ollama.Generate(ctx, &req, respFunc)
	if err != nil {
		return "", fmt.Errorf("failed to get response: %w", err)
	}

	return resp.String(), nil
}

func (c *Client) StartChat(ctx context.Context, systemMsg, msg string) (string, error) {
	msgs := []api.Message{
		{
			Role:    "system",
			Content: systemMsg,
		},
	}

	return c.Chat(ctx, msgs, msg)
}

func (c *Client) Chat(ctx context.Context, msgs []api.Message, next string) (string, error) {
	ragRespMsg := c.executeRAG(next)
	if ragRespMsg != nil {
		msgs = append(msgs, *ragRespMsg)
	}

	initialResp, err := c.nextMsg(ctx, msgs, api.Message{
		Role:    "user",
		Content: next,
	})
	if err != nil {
		return "", fmt.Errorf("failed to get initial chat response: %w", err)
	}

	if initialResp.Message.ToolCalls == nil || len(initialResp.Message.ToolCalls) == 0 {
		return initialResp.Message.Content, nil
	}

	toolRespMsg := c.executeToolCalls(initialResp.Message.ToolCalls)
	if toolRespMsg == nil {
		return initialResp.Message.Content, nil
	}

	finalResp, err := c.nextMsg(ctx, msgs, *toolRespMsg)
	if err != nil {
		return "", fmt.Errorf("failed to get final chat response: %w", err)
	}

	return finalResp.Message.Content, nil
}

func (c *Client) nextMsg(ctx context.Context, msgs []api.Message, next api.Message) (*api.ChatResponse, error) {
	// TODO: register tool calls

	req := api.ChatRequest{
		Model:    c.model,
		Stream:   &_false,
		Messages: append(msgs, next),
	}

	var resp api.ChatResponse
	respFunc := func(cr api.ChatResponse) error {
		resp = cr
		return nil
	}

	slog.Info("Generating next chat response", slog.String("model", c.model))

	err := c.ollama.Chat(ctx, &req, respFunc)
	if err != nil {
		return nil, fmt.Errorf("failed to get response: %w", err)
	}

	return &resp, nil
}

func (c *Client) executeToolCalls(calls []api.ToolCall) *api.Message {
	// TODO: implement tool call execution

	return &api.Message{
		Role:    "assistant",
		Content: "Executed tool calls:",
	}
}

func (c *Client) executeRAG(query string) *api.Message {
	// TODO: implement RAG execution

	return &api.Message{
		Role:    "system",
		Content: "Documents found:",
	}
}
