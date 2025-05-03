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
	stream := true
	req := api.GenerateRequest{
		Model:  c.model,
		Prompt: message,
		Stream: &stream,
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
