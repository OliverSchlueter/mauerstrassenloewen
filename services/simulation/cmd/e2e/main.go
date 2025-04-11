package main

import (
	"context"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/middleware"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"github.com/OliverSchlueter/mauerstrassenloewen/simulation/internal/backend"
	"github.com/OliverSchlueter/mauerstrassenloewen/simulation/internal/fflags"
	"github.com/go-pg/pg/v10"
	"github.com/justinas/alice"
	"github.com/nats-io/nats.go"
	"github.com/questdb/go-questdb-client/v3"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
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
		Service:      "simulation",
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

	// Setup MongoDB
	mc, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		slog.Error("Could not connect to MongoDB", slog.Any("err", err.Error()))
		os.Exit(1)
	}
	err = mc.Ping(context.Background(), readpref.Primary())
	if err != nil {
		slog.Error("Could not ping MongoDB", slog.Any("err", err.Error()))
		os.Exit(1)
	}
	mdb := mc.Database("msl_simulation")

	// Setup QuestDB
	qdbQ := pg.Connect(&pg.Options{
		Addr:     "localhost:8812",
		User:     "admin",
		Password: "quest",
	})
	err = qdbQ.Ping(context.Background())
	if err != nil {
		slog.Error("Could not ping QuestDB Query client", slog.Any("err", err.Error()))
		os.Exit(1)
	}

	qdbI, err := questdb.NewLineSender(
		context.Background(),
		questdb.WithHttp(),
		questdb.WithAddress("localhost:9005"),
		questdb.WithBasicAuth("admin", "quest"),
	)
	if err != nil {
		slog.Error("Could not connect to QuestDB ingestion client", slog.Any("err", err.Error()))
		os.Exit(1)
	}

	mux := http.NewServeMux()
	port := "8080"

	backend.Start(backend.Configuration{
		Mux:              mux,
		Nats:             nc,
		Mongo:            mdb,
		QuestDBIngestion: &qdbI,
		QuestDBQuery:     qdbQ,
	})

	go func() {
		chain := alice.New(
			middleware.Logging,
			middleware.RecoveryMiddleware,
		).Then(mux)

		err := http.ListenAndServe(":"+port, chain)
		if err != nil {
			slog.Error("Could not start server on port "+port, slog.Any("err", err.Error()))
			os.Exit(1)
		}
	}()

	slog.Info(fmt.Sprintf("Simulation service started on http://localhost:%s", port))

	// Wait for a signal to exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	switch <-sig {
	case os.Interrupt:
		slog.Info("Received interrupt signal, shutting down...")

		if err := qdbI.Flush(context.Background()); err != nil {
			slog.Error("Could not flush QuestDB ingestion client", slog.Any("err", err.Error()))
		}
		if err := qdbI.Close(context.Background()); err != nil {
			slog.Error("Could not close QuestDB ingestion client", slog.Any("err", err.Error()))
		}
		if err := qdbQ.Close(); err != nil {
			slog.Error("Could not close QuestDB query client", slog.Any("err", err.Error()))
		}

		if err := mc.Disconnect(context.Background()); err != nil {
			slog.Error("Could not disconnect from MongoDB", slog.Any("err", err.Error()))
		}

		nc.Close()

		slog.Info("Shutdown complete")
	}
}
