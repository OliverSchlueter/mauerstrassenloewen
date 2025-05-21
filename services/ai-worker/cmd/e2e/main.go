package main

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/chatbot"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/chatbot/store"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/fflags"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/ollama"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/tools"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"github.com/nats-io/nats.go"
	"github.com/qdrant/go-client/qdrant"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fflags.EndToEndEnvironment.Enable()
	//fflags.SendLogsToLoki.Enable()

	// Setup logging
	lokiService := sloki.NewService(sloki.Configuration{
		URL:          "http://localhost:3100/loki/api/v1/push",
		Service:      "ai-worker",
		ConsoleLevel: slog.LevelDebug,
		LokiLevel:    slog.LevelInfo,
		EnableLoki:   fflags.SendLogsToLoki.IsEnabled(),
	})
	slog.SetDefault(slog.New(lokiService))

	// Setup NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		slog.Error("Could not connect to NATS", slog.Any("err", err.Error()))
		os.Exit(1)
	}

	// Setup tool service
	tc := tools.NewService()

	// Setup ollama client
	oc, err := ollama.NewClient(ollama.Configuration{
		BaseURL:        "http://localhost:11434",
		Model:          "deepseek-r1:14b",
		EmbeddingModel: "TOOD",
		Tools:          *tc,
	})
	if err != nil {
		slog.Error("failed to create ollama client", slog.Any("err", err.Error()))
		return
	}

	// Setup qdrant client
	qc, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6333,
	})
	if err != nil {
		slog.Error("failed to create qdrant client", slog.Any("err", err.Error()))
		return
	}

	chatbotStore := store.NewStore(store.Configuration{})
	chatbotService := chatbot.NewService(chatbot.Configuration{
		Store:  chatbotStore,
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

	slog.Info("AI worker shutting down")

	nc.Close()

	if err := qc.Close(); err != nil {
		slog.Error("failed to close qdrant client", slog.Any("err", err.Error()))
	}
}
