package app

import (
	"github.com/DanielTitkov/correlateme-server/internal/configs"
	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/logger"
	"github.com/robfig/cron"
)

type (
	// App combines services and holds business logic
	App struct {
		cfg    configs.Config
		logger *logger.Logger
		repo   Repository
		cron   *cron.Cron
	}
	// Repository stores data
	Repository interface {
		// users
		CreateUser(*domain.User) (*domain.User, error)
		GetUserByUsername(username string) (*domain.User, error)
		GetUserCount() (int, error)
	}
)

func NewApp(
	cfg configs.Config,
	logger *logger.Logger,
	repo Repository,
) *App {
	c := cron.New()
	c.Start()

	return &App{
		cfg:    cfg,
		logger: logger,
		repo:   repo,
		cron:   c,
	}
}
