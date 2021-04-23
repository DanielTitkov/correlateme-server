package main

import (
	"context"
	"errors"
	"log"
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
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("failed to load config", errors.New("config path is not provided"))
	}
	configPath := args[0]
	log.Println("loading config from "+configPath, "")

	cfg, err := configs.ReadConfigs(configPath)
	if err != nil {
		log.Fatal("failed to load config", err)
	}
	log.Println("loaded config")

	logger := logger.NewLogger(cfg.Env)
	defer logger.Sync()
	logger.Info("starting service", "")

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
	err = app.Cron.AddFunc(cfg.Job.GatherAndSendServiceStatsSchedule, j.GatherAndSendServiceStats)
	if err != nil {
		logger.Fatal("failed to set up cron job", err)
	}

	// prometheus
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":2112", nil)
		if err != nil {
			logger.Error("failed to start metrics handler", err)
		}
	}()

	server := prepare.NewServer(cfg, logger, app)
	logger.Fatal("failed to start server", server.Start(cfg.Server.GetAddress()))
}
