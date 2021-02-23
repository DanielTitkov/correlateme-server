package main

import (
	"context"
	"errors"
	"os"

	"github.com/DanielTitkov/correlateme-server/cmd/app/prepare"
	"github.com/DanielTitkov/correlateme-server/internal/app"
	"github.com/DanielTitkov/correlateme-server/internal/configs"
	"github.com/DanielTitkov/correlateme-server/internal/logger"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"

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

	app := app.NewApp(cfg, logger, repo)

	server := prepare.NewServer(cfg, logger, app)
	logger.Fatal("failed to start server", server.Start(cfg.Server.GetAddress()))
}
