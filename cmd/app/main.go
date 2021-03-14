package main

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/DanielTitkov/correlateme-server/cmd/app/prepare"
	"github.com/DanielTitkov/correlateme-server/internal/app"
	"github.com/DanielTitkov/correlateme-server/internal/configs"
	"github.com/DanielTitkov/correlateme-server/internal/job"
	"github.com/DanielTitkov/correlateme-server/internal/logger"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/lib/pq"
)

func main() {
	logger := logger.NewLogger()
	defer logger.Sync()
	logger.Info("starting service", "")

	args := os.Args[1:]
	if len(args) < 1 {
		logger.Fatal("failed to load config", errors.New("config path is not provided"))
	}
	configPath := args[0]
	logger.Info("loading config from "+configPath, "")

	cfg, err := configs.ReadConfigs(configPath)
	if err != nil {
		logger.Fatal("failed to load config", err)
	}
	logger.Info("loaded config", "")

	db, err := ent.Open(cfg.DB.Driver, cfg.DB.URI)
	if err != nil {
		logger.Fatal("failed connecting to database", err)
	}
	defer db.Close()
	logger.Info("connected to database", cfg.DB.Driver+", "+cfg.DB.URI)

	err = prepare.Migrate(context.Background(), db) // run db migration
	if err != nil {
		logger.Fatal("failed creating schema resources", err)
	}
	logger.Info("migrations done", "")

	repo := entgo.NewEntgoRepository(db, logger)

	app, err := app.NewApp(cfg, logger, repo)
	if err != nil {
		logger.Fatal("failed creating app", err)
	}
	j := job.New(cfg, logger, app)
	go j.ListenUpdateUserCorrelationsChannel()
	go j.ListenUpdateDatasetAggregationsChannel()

	// prometheus
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	server := prepare.NewServer(cfg, logger, app)
	logger.Fatal("failed to start server", server.Start(cfg.Server.GetAddress()))
}
