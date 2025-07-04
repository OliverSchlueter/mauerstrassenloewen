package main

import (
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/backend"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/fflags"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/ollama"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/tools"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/middleware"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"github.com/justinas/alice"
	"github.com/nats-io/nats.go"
	"github.com/qdrant/go-client/qdrant"
	"log/slog"
	"net/http"
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
		slog.Error("Could not connect to NATS", sloki.WrapError(err))
		os.Exit(1)
	}

	// Setup tool service
	ts := tools.NewService()

	// Setup qdrant client
	qc, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6333,
	})
	if err != nil {
		slog.Error("failed to create qdrant client", sloki.WrapError(err))
		return
	}

	// Setup ollama client
	oc, err := ollama.NewClient(ollama.Configuration{
		BaseURL:        "http://localhost:11434",
		Model:          "deepseek-r1:14b",
		EmbeddingModel: "TOOD",
		Tools:          ts,
		QC:             qc,
	})
	if err != nil {
		sloki.WrapError(err)
		slog.Error("failed to create ollama client", sloki.WrapError(err))
		return
	}

	mux := http.NewServeMux()
	port := "8085"

	backend.Start(backend.Configuration{
		Mux:          mux,
		Nats:         nc,
		OllamaClient: oc,
	})

	go func() {
		chain := alice.New(
			middleware.CORS,
			middleware.Logging,
			middleware.RecoveryMiddleware,
		).Then(mux)

		err := http.ListenAndServe(":"+port, chain)
		if err != nil {
			slog.Error("Could not start server on port "+port, sloki.WrapError(err))
			os.Exit(1)
		}
	}()

	slog.Info(fmt.Sprintf("Started server on http://localhost:%s\n", port))

	// Wait for a signal to exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	switch <-sig {
	case os.Interrupt:
		slog.Info("Received interrupt signal, shutting down...")

		nc.Close()

		if err := qc.Close(); err != nil {
			slog.Error("failed to close qdrant client", sloki.WrapError(err))
		}

		slog.Info("Shutdown complete")
	}
}
