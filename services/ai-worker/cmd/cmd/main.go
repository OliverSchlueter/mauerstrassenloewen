package main

import (
	"context"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/chatgpt"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/ollama"
	"log/slog"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	ctx := context.Background()
	askOpenAI(ctx)
}

func askOpenAI(ctx context.Context) {
	client := chatgpt.NewClient(chatgpt.Configuration{
		AuthToken: "",
	})

	resp, err := client.Chat(ctx, "Schreibe ein Gedicht in 100 Wörtern über Softwareentwicklung.")
	if err != nil {
		slog.Error("failed to get response", "error", err)
		return
	}

	fmt.Println(resp)
}

func askOllama(ctx context.Context) {
	client, err := ollama.NewClient(ollama.Configuration{
		BaseURL: "http://localhost:11434",
		Model:   "deepseek-r1:14b",
	})
	if err != nil {
		slog.Error("failed to create ollama client", "error", err)
		return
	}

	resp, err := client.Chat(ctx, "Schreibe ein Gedicht in 100 Wörtern über Softwareentwicklung.")
	if err != nil {
		slog.Error("failed to get response", "error", err)
		return
	}
	fmt.Println(resp)
}
