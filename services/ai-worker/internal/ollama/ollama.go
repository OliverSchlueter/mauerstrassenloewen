package ollama

import (
	"context"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/tools"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/natsdto"
	"github.com/google/uuid"
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
	ollama         *api.Client
	model          string
	embeddingModel string
	tools          tools.Service
}

type Configuration struct {
	BaseURL        string
	Model          string
	EmbeddingModel string
	Tools          tools.Service
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
		ollama:         ollama,
		model:          cfg.Model,
		embeddingModel: cfg.EmbeddingModel,
		tools:          cfg.Tools,
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

func (c *Client) StartChat(ctx context.Context, req natsdto.StartChatRequest) (*natsdto.Chat, error) {
	chat := &natsdto.Chat{
		ID: uuid.New().String(),
		Messages: []natsdto.Message{
			{
				Role:    "system",
				Content: string(req.SystemMsg),
				SentAt:  time.Now(),
			},
		},
	}

	err := c.Chat(ctx, chat, req.UserMsg)
	if err != nil {

	}

	return chat, nil
}

func (c *Client) Chat(ctx context.Context, chat *natsdto.Chat, next string) error {
	ragResp, err := c.executeRAG(next)
	if err != nil {
		return fmt.Errorf("failed to execute RAG: %w", err)
	}

	chat.AppendMsg(natsdto.Message{
		Role:    "system",
		Content: ragResp,
		SentAt:  time.Now(),
	})

	chat.AppendMsg(natsdto.Message{
		Role:    "user",
		Content: next,
		SentAt:  time.Now(),
	})
	initialResp, err := c.nextMsg(ctx, toOllamaMsgs(chat))
	if err != nil {
		return fmt.Errorf("failed to get initial chat response: %w", err)
	}

	if initialResp.Message.ToolCalls == nil || len(initialResp.Message.ToolCalls) == 0 {
		chat.AppendMsg(natsdto.Message{
			Role:    "assistant",
			Content: initialResp.Message.Content,
			SentAt:  time.Now(),
		})
		return nil
	}

	toolResp, err := c.executeToolCalls(ctx, initialResp.Message.ToolCalls)
	if err != nil {
		return fmt.Errorf("failed to execute tool calls: %w", err)
	}

	chat.AppendMsg(natsdto.Message{
		Role:    "assistant",
		Content: toolResp,
		SentAt:  time.Now(),
	})

	finalResp, err := c.nextMsg(ctx, toOllamaMsgs(chat))
	if err != nil {
		return fmt.Errorf("failed to get final chat response: %w", err)
	}

	chat.AppendMsg(natsdto.Message{
		Role:    "assistant",
		Content: finalResp.Message.Content,
		SentAt:  time.Now(),
	})

	return nil
}

func (c *Client) nextMsg(ctx context.Context, msgs []api.Message) (*api.ChatResponse, error) {
	req := api.ChatRequest{
		Model:    c.model,
		Stream:   &_false,
		Messages: msgs,
		Tools:    tools.ToOllama(c.tools.GetTools()),
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

func (c *Client) CreateEmbed(ctx context.Context, input string) ([][]float32, error) {
	resp, err := c.ollama.Embed(ctx, &api.EmbedRequest{
		Model: c.embeddingModel,
		Input: input,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get embedding: %w", err)
	}

	return resp.Embeddings, nil
}

func (c *Client) CreateEmbedding(ctx context.Context, prompt string) ([]float64, error) {
	resp, err := c.ollama.Embeddings(ctx, &api.EmbeddingRequest{
		Model:  c.embeddingModel,
		Prompt: prompt,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get embedding: %w", err)
	}

	return resp.Embedding, nil
}

func (c *Client) executeToolCalls(ctx context.Context, calls []api.ToolCall) (string, error) {
	if len(calls) == 0 {
		return "", nil
	}

	msg := "Executed the following tool calls:\n"

	for _, tc := range calls {
		switch tc.Function.Name {
		case "get_user_info":
			resp, err := tools.GetUserInfo(ctx)
			if err != nil {
				return "", fmt.Errorf("failed to execute tool call: %w", err)
			}
			msg += fmt.Sprintf("Tool: %s\nResponse: %s\n\n", tc.Function.Name, resp)
		}
	}

	return msg, nil
}

func (c *Client) executeRAG(query string) (string, error) {
	// TODO: implement RAG execution
	return "Found relevant documents:", nil
}
