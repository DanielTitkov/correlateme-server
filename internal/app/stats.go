package app

import "github.com/DanielTitkov/correlateme-server/internal/domain"

func (a *App) ServiceStats() (*domain.ServiceStats, error) {
	users, err := a.repo.UserCount()
	if err != nil {
		return nil, err
	}

	indicators, err := a.repo.IndicatorCount()
	if err != nil {
		return nil, err
	}

	datasets, err := a.repo.DatasetCount()
	if err != nil {
		return nil, err
	}

	scales, err := a.repo.ScaleCount()
	if err != nil {
		return nil, err
	}

	observations, err := a.repo.ObservationCount()
	if err != nil {
		return nil, err
	}

	correlations, err := a.repo.CorrelationCount()
	if err != nil {
		return nil, err
	}

	return &domain.ServiceStats{
		Users:        users,
		Scales:       scales,
		Indicators:   indicators,
		Datasets:     datasets,
		Observations: observations,
		Correlations: correlations,
	}, nil
}
