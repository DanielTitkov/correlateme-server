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
		cache    *Cache
		Cron     *cron.Cron
		Channels *Channels
	}
	// Cache stores data loaded on app start
	Cache struct {
		scales map[string]*domain.Scale
	}
	// Channels holds app channels for async jobs
	Channels struct {
		UpdateUserCorrelationsChan    chan domain.UpdateCorrelationsArgs
		UpdateDatasetAggregationsChan chan domain.UpdateAggregationsArgs
	}
	// Repository stores data
	Repository interface {
		// users
		UserCount() (int, error)
		CreateUser(*domain.User) (*domain.User, error)
		GetUserByUsername(string) (*domain.User, error)
		GetUserByID(int) (*domain.User, error)
		GetUserCount() (int, error)

		// indicators
		IndicatorCount() (int, error)
		CreateIndicator(*domain.Indicator) (*domain.Indicator, error)
		UpdateIndicator(*domain.Indicator) (*domain.Indicator, error)
		GetIndicatorByID(int) (*domain.Indicator, error)
		GetIndicatorByCode(string) (*domain.Indicator, error)
		GetIndicators(domain.GetIndicatorsArgs) ([]*domain.Indicator, error)

		// scales
		ScaleCount() (int, error)
		GetScales() ([]*domain.Scale, error)
		GetScaleByType(string) (*domain.Scale, error)
		CreateScale(domain.Scale) (*domain.Scale, error)

		// datasets
		DatasetCount() (int, error)
		CreateDataset(*domain.Dataset) (*domain.Dataset, error)
		GetDatasetByID(id int, observationsLimit int, granularity string) (*domain.Dataset, error)
		GetUserIndicatorDataset(*domain.User, *domain.Indicator) (*domain.Dataset, error)
		GetOrCreateUserIndicatorDataset(*domain.User, *domain.Indicator) (*domain.Dataset, error)
		GetUserDatasets(userID int, withShared bool, observationsLimit int, granularity string) ([]*domain.Dataset, error) // to be deprecated
		GetDatasets(domain.GetDatasetsArgs) ([]*domain.Dataset, error)

		// observations
		ObservationCount() (int, error)
		CreateObservation(*domain.Observation) (*domain.Observation, error)
		CreateOrUpdateObservation(*domain.Observation) (*domain.Observation, error)
		// UpdateObservation

		// correlations
		CorrelationCount() (int, error)
		CreateOrUpdateCorrelation(*domain.Correlation) (*domain.Correlation, error)
		GetUserCorrelations(userID int, granularity string) ([]*domain.Correlation, error)
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
		Cron:   c,
	}

	// TODO: maybe move it out from the constructor
	err := app.initScales()
	if err != nil {
		return nil, err
	}
	cache, err := app.buildCache()
	if err != nil {
		return nil, err
	}
	app.cache = cache

	// indicators can be set up only after scales are loaded and cached
	err = app.initBuiltinIndicators()
	if err != nil {
		return nil, err
	}

	channels, err := app.makeChannels()
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

func (a *App) makeChannels() (*Channels, error) {
	updateCorrelationsChan := make(chan domain.UpdateCorrelationsArgs, a.cfg.App.UpdateCorrelationsBuffer)
	updateAggregationsChan := make(chan domain.UpdateAggregationsArgs, a.cfg.App.UpdateAggregationsBuffer)
	return &Channels{
		UpdateUserCorrelationsChan:    updateCorrelationsChan,
		UpdateDatasetAggregationsChan: updateAggregationsChan,
	}, nil
}

func (a *App) loadPresets() error {
	err := a.initScales()
	if err != nil {
		return err
	}

	err = a.initBuiltinIndicators()
	if err != nil {
		return err
	}

	return nil
}
