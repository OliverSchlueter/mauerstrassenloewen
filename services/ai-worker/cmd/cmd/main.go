package main

import (
	"context"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/ollama"
	"log/slog"
	"os"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	ctx := context.Background()

	client, err := ollama.NewClient(ollama.Configuration{
		BaseURL: "http://localhost:11434",
		Model:   "deepseek-r1:14b",
	})
	if err != nil {
		slog.Error("failed to create ollama client", "error", err)
		os.Exit(1)
	}

	resp, err := client.Chat(ctx, "Schreibe ein Gedicht in 100 Wörtern über Softwareentwicklung.")
	if err != nil {
		slog.Error("failed to get response", "error", err)
		os.Exit(1)
	}
	fmt.Println(resp)

}
