package ollama

import (
	"context"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/telemetry"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/tools"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/natsdto"
	"github.com/google/uuid"
	"github.com/ollama/ollama/api"
	"github.com/qdrant/go-client/qdrant"
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
	tools          *tools.Service
	rag            *RAGStore
	telemetry      *telemetry.Service
}

type Configuration struct {
	BaseURL        string
	Model          string
	EmbeddingModel string
	Tools          *tools.Service
	QC             *qdrant.Client
	Telemetry      *telemetry.Service
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

	rag := NewRAGStore(RAGConfiguration{
		QC:             cfg.QC,
		Ollama:         ollama,
		EmbeddingModel: cfg.EmbeddingModel,
	})

	return &Client{
		ollama:         ollama,
		model:          cfg.Model,
		embeddingModel: cfg.EmbeddingModel,
		tools:          cfg.Tools,
		rag:            rag,
		telemetry:      cfg.Telemetry,
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
	slog.Debug("Starting new chat", slog.String("userMsg", req.UserMsg), slog.String("systemMsg", string(req.SystemMsg)))
	c.telemetry.TrackNewChat()

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

	newChat, err := c.Chat(ctx, chat, req.UserMsg)
	if err != nil {
		return nil, fmt.Errorf("failed to start chat: %w", err)
	}

	return newChat, nil
}

func (c *Client) Chat(ctx context.Context, chat *natsdto.Chat, next string) (*natsdto.Chat, error) {
	slog.Debug("Continuing chat", slog.String("chatID", chat.ID), slog.String("nextMessage", next))

	slog.Debug("Executing RAG for next message")
	ragResp, err := c.executeRAG(next)
	if err != nil {
		return nil, fmt.Errorf("failed to execute RAG: %w", err)
	}
	slog.Debug("RAG response", slog.String("ragResp", ragResp))

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

	slog.Debug("Chat response")
	initialResp, err := c.nextMsg(ctx, toOllamaMsgs(chat))
	if err != nil {
		return nil, fmt.Errorf("failed to get initial chat response: %w", err)
	}
	c.telemetry.TrackOllamaResponse(initialResp)
	slog.Debug("Initial chat response", slog.String("content", initialResp.Message.Content))

	if initialResp.Message.ToolCalls == nil || len(initialResp.Message.ToolCalls) == 0 {
		chat.AppendMsg(natsdto.Message{
			Role:    "assistant",
			Content: initialResp.Message.Content,
			SentAt:  time.Now(),
		})
		return chat, nil
	}
	slog.Debug("Chat response", slog.String("content", initialResp.Message.Content))

	slog.Debug("Executing tool calls", slog.Any("toolCalls", initialResp.Message.ToolCalls))
	toolResp, err := c.executeToolCalls(ctx, initialResp.Message.ToolCalls)
	if err != nil {
		return nil, fmt.Errorf("failed to execute tool calls: %w", err)
	}
	slog.Debug("Tool response", slog.String("toolResp", toolResp))

	chat.AppendMsg(natsdto.Message{
		Role:    "assistant",
		Content: toolResp,
		SentAt:  time.Now(),
	})

	slog.Debug("Getting final chat response after tool execution")
	finalResp, err := c.nextMsg(ctx, toOllamaMsgs(chat))
	if err != nil {
		return nil, fmt.Errorf("failed to get final chat response: %w", err)
	}
	c.telemetry.TrackOllamaResponse(finalResp)
	slog.Debug("Final chat response", slog.String("content", finalResp.Message.Content))

	chat.AppendMsg(natsdto.Message{
		Role:    "assistant",
		Content: finalResp.Message.Content,
		SentAt:  time.Now(),
	})

	slog.Debug("Finished chat message", slog.String("chatID", chat.ID))
	return chat, nil
}

func (c *Client) nextMsg(ctx context.Context, msgs []api.Message) (*api.ChatResponse, error) {
	req := api.ChatRequest{
		Model:    c.model,
		Stream:   &_false,
		Messages: msgs,
		//Tools:    tools.ToOllama(c.tools.GetTools()),
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

func (c *Client) executeToolCalls(ctx context.Context, calls []api.ToolCall) (string, error) {
	return "", nil

	//if len(calls) == 0 {
	//	return "", nil
	//}
	//
	//msg := "Executed the following tool calls:\n"
	//
	//for _, tc := range calls {
	//	switch tc.Function.Name {
	//	case "get_user_info":
	//		resp, err := tools.GetUserInfo(ctx)
	//		if err != nil {
	//			return "", fmt.Errorf("failed to execute tool call: %w", err)
	//		}
	//		msg += fmt.Sprintf("Tool: %s\nResponse: %s\n\n", tc.Function.Name, resp)
	//	}
	//}
	//
	//return msg, nil
}

func (c *Client) executeRAG(query string) (string, error) {
	return "", nil

	//res, err := c.rag.Search(context.Background(), query, 5)
	//if err != nil {
	//	return "", fmt.Errorf("failed to search RAG: %w", err)
	//}
	//
	//var resp string
	//if len(res) == 0 {
	//	resp = "No relevant documents found."
	//	return resp, nil
	//} else {
	//	resp = "Found relevant documents:\n"
	//}
	//
	//for _, r := range res {
	//	resp += fmt.Sprintf("- %s\n", r)
	//}
	//
	//return resp, nil
}
