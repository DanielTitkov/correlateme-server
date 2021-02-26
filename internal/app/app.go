package app

import (
	"fmt"

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
		cache  *Cache
	}
	// Cache stores data loaded on app start
	Cache struct {
		scales map[string]*domain.Scale
	}
	// Repository stores data
	Repository interface {
		// users
		CreateUser(*domain.User) (*domain.User, error)
		GetUserByUsername(username string) (*domain.User, error)
		GetUserCount() (int, error)

		// indicators
		CreateIndicator(*domain.Indicator) (*domain.Indicator, error)
		FilterIndicators(domain.FilterIndicatorsArgs) ([]*domain.Indicator, error)

		// scales
		GetScales() ([]*domain.Scale, error)
	}
)

func NewApp(
	cfg configs.Config,
	logger *logger.Logger,
	repo Repository,
) (*App, error) {
	c := cron.New()
	c.Start()

	app := App{
		cfg:    cfg,
		logger: logger,
		repo:   repo,
		cron:   c,
	}

	cache, err := app.buildCache()
	if err != nil {
		return nil, err
	}
	app.cache = cache

	return &app, nil
}

func (a *App) buildCache() (*Cache, error) {
	a.logger.Info("loading scales", "")
	scales, err := a.repo.GetScales()
	if err != nil {
		return nil, err
	}

	scaleMap := make(map[string]*domain.Scale)
	for _, scale := range scales {
		scaleMap[scale.Type] = scale
	}
	a.logger.Info("loaded scales", "")

	cache := Cache{
		scales: scaleMap,
	}

	a.logger.Info("created cache", fmt.Sprintf("%+v", cache))
	return &cache, nil
}
