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
	"strconv"
	"syscall"
)

const (
	lokiUrlEnv              = "LOKI_URL"
	natsUrlEnv              = "NATS_URL"
	natsAuthTokenEnv        = "NATS_AUTH_TOKEN"
	ollamaUrlEnv            = "OLLAMA_URL"
	ollamaModelEnv          = "OLLAMA_MODEL"
	ollamaEmbeddingModelEnv = "OLLAMA_EMBEDDING_MODEL"
	qdrantHostEnv           = "QDRANT_HOST"
	qdrantPortEnv           = "QDRANT_PORT"
	qdrantAPIKeyEnv         = "QDRANT_API_KEY"
)

func main() {
	fflags.EndToEndEnvironment.Disable()

	// Setup logging
	lokiService := sloki.NewService(sloki.Configuration{
		URL:          mustGetLokiURL(),
		Service:      "ai-worker",
		ConsoleLevel: slog.LevelDebug,
		LokiLevel:    slog.LevelInfo,
		EnableLoki:   fflags.SendLogsToLoki.IsEnabled(),
	})
	slog.SetDefault(slog.New(lokiService))

	// Setup NATS
	nc, err := nats.Connect(mustGetEnvStr(natsUrlEnv), nats.Token(mustGetEnvStr(natsAuthTokenEnv)))
	if err != nil {
		slog.Error("Could not connect to NATS", sloki.WrapError(err))
		os.Exit(1)
	}

	// Setup tool service
	tc := tools.NewService()

	// Setup qdrant client
	qc, err := qdrant.NewClient(&qdrant.Config{
		Host:   mustGetEnvStr(qdrantHostEnv),
		Port:   mustGetEnvInt(qdrantPortEnv),
		APIKey: mustGetEnvStr(qdrantAPIKeyEnv),
	})
	if err != nil {
		slog.Error("failed to create qdrant client", sloki.WrapError(err))
		return
	}

	// Setup ollama client
	oc, err := ollama.NewClient(ollama.Configuration{
		BaseURL:        mustGetEnvStr(ollamaUrlEnv),
		Model:          mustGetEnvStr(ollamaModelEnv),
		EmbeddingModel: mustGetEnvStr(ollamaEmbeddingModelEnv),
		Tools:          tc,
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

func mustGetLokiURL() string {
	lokiURL := os.Getenv(lokiUrlEnv)
	if lokiURL == "" && fflags.SendLogsToLoki.IsEnabled() {
		panic(lokiUrlEnv + " environment variable not set")
	}
	return lokiURL
}

func mustGetEnvStr(env string) string {
	val := os.Getenv(env)
	if val == "" {
		panic(env + " environment variable not set")
	}
	return val
}

func mustGetEnvInt(env string) int {
	val := mustGetEnvStr(env)
	intVal, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Sprintf("Could not convert %s to int: %v", env, err))
	}
	return intVal
}
