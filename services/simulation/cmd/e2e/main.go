package main

import (
	"common/middleware"
	"context"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/simulation/internal/backend"
	"github.com/OliverSchlueter/mauerstrassenloewen/simulation/internal/fflags"
	"github.com/OliverSchlueter/sloki/sloki"
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
	lokiLevel := slog.LevelInfo
	if !fflags.SendLogsToLoki.IsEnabled() {
		lokiLevel = 100_0000
	}
	lokiService := sloki.NewService(sloki.Configuration{
		URL:          "http://localhost:3100/loki/api/v1/push",
		Service:      "simulation",
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
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
