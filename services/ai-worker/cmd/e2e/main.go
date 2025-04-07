package main

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/chatbot"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/fflags"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/ollama"
	"github.com/OliverSchlueter/sloki/sloki"
	"github.com/nats-io/nats.go"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fflags.EndToEndEnvironment.Enable()
	//fflags.SendLogsToLoki.Enable()

	// Setup logging
	lokiLevel := slog.LevelInfo
	if !fflags.SendLogsToLoki.IsEnabled() {
		lokiLevel = 100_0000
	}
	lokiService := sloki.NewService(sloki.Configuration{
		URL:          "http://localhost:3100/loki/api/v1/push",
		Service:      "backend",
		ConsoleLevel: slog.LevelDebug,
		LokiLevel:    lokiLevel,
	})
	slog.SetDefault(slog.New(lokiService))

	// Setup NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		slog.Error("Could not connect to NATS", slog.Any("err", err.Error()))
		os.Exit(1)
	}

	// Setup ollama client
	oc, err := ollama.NewClient(ollama.Configuration{
		BaseURL: "http://localhost:11434",
		Model:   "deepseek-r1:14b",
	})
	if err != nil {
		slog.Error("failed to create ollama client", slog.Any("err", err.Error()))
		return
	}

	chatbotService := chatbot.NewService(chatbot.Configuration{
		Nats:   nc,
		Ollama: oc,
	})
	if err := chatbotService.Register(); err != nil {
		slog.Error("failed to register chatbot service", slog.Any("err", err.Error()))
		return
	}

	slog.Info("AI worker started")

	// Wait for a signal to exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
