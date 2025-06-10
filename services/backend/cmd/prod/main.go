package main

import (
	"context"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/backend"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/fflags"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/middleware"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/middleware/metricshandler"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"github.com/justinas/alice"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	lokiUrlEnv       = "LOKI_URL"
	natsUrlEnv       = "NATS_URL"
	natsAuthTokenEnv = "NATS_AUTH_TOKEN"
	mongoUrlEnv      = "MONGODB_URL"
)

func main() {
	ctx := context.Background()

	// Setup feature flags
	fflags.EndToEndEnvironment.Disable()

	lokiService := sloki.NewService(sloki.Configuration{
		URL:          mustGetLokiURL(),
		Service:      "backend",
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

	// Setup MongoDB
	mc, err := mongo.Connect(options.Client().ApplyURI(mustGetEnvStr(mongoUrlEnv)))
	if err != nil {
		slog.Error("Could not connect to MongoDB", sloki.WrapError(err))
		os.Exit(1)
	}
	err = mc.Ping(ctx, readpref.Primary())
	if err != nil {
		slog.Error("Could not ping MongoDB", sloki.WrapError(err))
		os.Exit(1)
	}
	mdb := mc.Database("msl_backend")

	// Start the web server
	mux := &http.ServeMux{}
	port := "8082"

	appCfg := backend.Configuration{
		Mux:     mux,
		Nats:    nc,
		MongoDB: mdb,
	}
	authMiddleware := backend.Start(appCfg)

	metricshandler.Register(mux, "")
	middleware.RegisterPrometheusHttpLogging()

	go func() {
		chain := alice.New(
			middleware.CORS,
			middleware.Logging,
			authMiddleware,
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
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	switch <-sig {
	case os.Interrupt:
		slog.Info("Received interrupt signal, shutting down...")

		if err := mc.Disconnect(context.Background()); err != nil {
			slog.Error("Could not disconnect from MongoDB", sloki.WrapError(err))
		}

		nc.Close()

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
