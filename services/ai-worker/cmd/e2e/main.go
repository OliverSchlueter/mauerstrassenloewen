package main

import (
	"context"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/chatgpt"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/ollama"
	"github.com/OliverSchlueter/sloki/sloki"
	"github.com/nats-io/nats.go"
	"log/slog"
	"os"
)

func main() {
	// Setup logging
	lokiService := sloki.NewService(sloki.Configuration{
		URL:          "http://localhost:3100/loki/api/v1/push",
		Service:      "ai-worker",
		ConsoleLevel: slog.LevelDebug,
		//LokiLevel:    slog.LevelInfo,
		LokiLevel: 100_0000,
	})
	slog.SetDefault(slog.New(lokiService))

	ctx := context.Background()

	// Setup NATS
	natsClient, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		slog.Error("Could not connect to NATS", slog.Any("err", err.Error()))
		os.Exit(1)
	}
	_ = natsClient.Publish("foo", []byte("bar"))

	askOllama(ctx)
}

func askChatGPT(ctx context.Context) {
	client := chatgpt.NewClient(chatgpt.Configuration{
		AuthToken: "",
	})

	resp, err := client.Chat(ctx, "Schreibe ein Gedicht in 100 Wörtern über Softwareentwicklung.")
	if err != nil {
		slog.Error("failed to get response", slog.Any("err", err.Error()))
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
		slog.Error("failed to create ollama client", slog.Any("err", err.Error()))
		return
	}

	resp, err := client.Chat(ctx, "Schreibe ein Gedicht in 100 Wörtern über Softwareentwicklung.")
	if err != nil {
		slog.Error("failed to get response", slog.Any("err", err.Error()))
		return
	}
	fmt.Println(resp)
}
