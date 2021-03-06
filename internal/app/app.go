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
		cfg      configs.Config
		logger   *logger.Logger
		repo     Repository
		cron     *cron.Cron
		cache    *Cache
		Channels *Channels
	}
	// Cache stores data loaded on app start
	Cache struct {
		scales map[string]*domain.Scale
	}
	// Channels holds app channels for async jobs
	Channels struct {
		UpdateUserCorrelationsChan chan domain.UpdateCorrelationsArgs
	}
	// Repository stores data
	Repository interface {
		// users
		CreateUser(*domain.User) (*domain.User, error)
		GetUserByUsername(string) (*domain.User, error)
		GetUserByID(int) (*domain.User, error)
		GetUserCount() (int, error)

		// indicators
		CreateIndicator(*domain.Indicator) (*domain.Indicator, error)
		GetIndicatorByID(int) (*domain.Indicator, error)
		GetIndicators(domain.GetIndicatorsArgs) ([]*domain.Indicator, error)

		// scales
		GetScales() ([]*domain.Scale, error)

		// datasets
		CreateDataset(*domain.Dataset) (*domain.Dataset, error)
		GetDatasetByID(id int, observationsLimit int, granularity string) (*domain.Dataset, error)
		GetUserIndicatorDataset(*domain.User, *domain.Indicator) (*domain.Dataset, error)
		GetOrCreateUserIndicatorDataset(*domain.User, *domain.Indicator) (*domain.Dataset, error)
		GetUserDatasets(userID int, withShared bool, observationsLimit int, granularity string) ([]*domain.Dataset, error) // to be deprecated
		GetDatasets(domain.GetDatasetsArgs) ([]*domain.Dataset, error)

		// observations
		CreateObservation(*domain.Observation) (*domain.Observation, error)
		CreateOrUpdateObservation(*domain.Observation) (*domain.Observation, error)
		// UpdateObservation

		// correlations
		CreateOrUpdateCorrelation(*domain.Correlation) (*domain.Correlation, error)
		GetUserCorrelations(userID int) ([]*domain.Correlation, error)
		GetCorrelation(domain.GetCorrelationArgs) (*domain.Correlation, error)
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

	channels, err := app.buildChannels()
	if err != nil {
		return nil, err
	}
	app.Channels = channels

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

func (a *App) buildChannels() (*Channels, error) {
	updateCorrelationsChan := make(chan domain.UpdateCorrelationsArgs, a.cfg.App.UpdateCorrelationsBuffer)
	return &Channels{
		UpdateUserCorrelationsChan: updateCorrelationsChan,
	}, nil
}
