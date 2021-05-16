package app

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/helper"
	"github.com/DanielTitkov/correlateme-server/internal/service/metrics"
)

// CreateOrUpdateObservation checks if there is a dataset for user/indicator.
// If none, creates one and adds observation to it, otherwise adds to existing dataset.
func (a *App) CreateOrUpdateObservation(args domain.CreateOrUpdateObservationArgs) error {
	user, err := a.repo.GetUserByID(args.UserID)
	if err != nil {
		return err
	}

	indicator, err := a.repo.GetIndicatorByID(args.IndicatorID)
	if err != nil {
		return err
	}

	dataset, err := a.repo.GetOrCreateUserIndicatorDataset(user, indicator)
	if err != nil {
		return err
	}

	obs := &domain.Observation{
		Value:       args.Value,
		Dataset:     dataset,
		Date:        args.Date,
		Granularity: domain.GranularityDay, // only daily observations can be created by user
	}

	if err := a.validateObservation(indicator, obs); err != nil {
		return err
	}

	_, err = a.repo.CreateOrUpdateObservation(obs)
	if err != nil {
		return err
	}

	go func() {
		// TODO: add timeout
		// update correlations for all user datasets
		metrics.UnprocessedUpdateCorrelationsRequests.Add(1)
		a.Channels.UpdateUserCorrelationsChan <- domain.UpdateCorrelationsArgs{
			UserID:      args.UserID,
			WithShared:  true,
			Method:      "auto",
			Granularity: domain.GranularityDay,
		}
	}()

	go func() {
		// TODO: add timeout
		// update aggregations for current dataset
		metrics.UnprocessedUpdateAggregationsRequests.Add(1)
		a.Channels.UpdateDatasetAggregationsChan <- domain.UpdateAggregationsArgs{
			UserID:    args.UserID,
			DatasetID: dataset.ID,
		}
	}()

	return nil
}

func (a *App) UpdateAggregations(args domain.UpdateAggregationsArgs) error {
	// get dataset with observations
	dataset, err := a.repo.GetDatasetByID(args.DatasetID, a.cfg.App.MaxMonthAggregationObservations, domain.GranularityDay)
	if err != nil {
		return err
	}

	var includeZeroValues bool
	if dataset.Params != nil {
		includeZeroValues = dataset.Params.Aggregation.IncludeZeroValues
	}

	// calculate aggregated values
	// by week
	weekValueMap := mapObservationsValues(dataset.Observations, domain.GranularityWeek, includeZeroValues)
	weekAggregatedObs, err := aggregateObservations(weekValueMap, dataset, domain.GranularityWeek)
	if err != nil {
		return err
	}

	// by month
	monthValueMap := mapObservationsValues(dataset.Observations, domain.GranularityMonth, includeZeroValues)
	monthAggregatedObs, err := aggregateObservations(monthValueMap, dataset, domain.GranularityMonth)
	if err != nil {
		return err
	}

	// save observations
	aggregatedObs := append(weekAggregatedObs, monthAggregatedObs...)
	for _, obs := range aggregatedObs { // TODO: change to bulk? but ent doens't support that as yet
		_, err := a.repo.CreateOrUpdateObservation(&obs)
		if err != nil {
			return err // TODO: maybe save error, not exit right now
		}
	}

	return nil
}

func (a *App) validateObservation(ind *domain.Indicator, obs *domain.Observation) error {
	if ind.ValueParams == nil {
		return nil
	}

	if obs.Value > ind.ValueParams.Max {
		return fmt.Errorf(
			"value must be not greater than indicator maximum: got %f while max = %f",
			obs.Value, ind.ValueParams.Max,
		)
	}

	if obs.Value < ind.ValueParams.Min {
		return fmt.Errorf(
			"value must be greater than indicator minimum: got %f while min = %f",
			obs.Value, ind.ValueParams.Min,
		)
	}

	if ind.ValueParams.Step != 0 && math.Mod(obs.Value, ind.ValueParams.Step) != 0 {
		return fmt.Errorf(
			"value must be multiple of indicator step: step = %f",
			ind.ValueParams.Step,
		)
	}

	return nil
}

func mapObservationsValues(observations []*domain.Observation, gran string, includeZeroValues bool) map[string][]float64 {
	valueMap := make(map[string][]float64)

	var getIdx func(time.Time) (int, int)
	switch gran {
	case domain.GranularityMonth:
		getIdx = helper.MonthIdx
	case domain.GranularityWeek:
		getIdx = helper.WeekIdx
	default:
		return nil
	}

	for _, obs := range observations {
		if !includeZeroValues && obs.Value == 0 {
			continue
		}
		idx := helper.PairOfIDsToString(getIdx(*obs.Date))
		valueMap[idx] = append(valueMap[idx], obs.Value)
	}

	return valueMap
}

func aggregateObservations(valueMap map[string][]float64, dataset *domain.Dataset, gran string) ([]domain.Observation, error) {
	var aggregatedObs []domain.Observation

	var periodStart func(int, int) time.Time
	switch gran {
	case domain.GranularityMonth:
		periodStart = helper.MonthStart
	case domain.GranularityWeek:
		periodStart = helper.WeekStart
	default:
		return nil, errors.New("got unknown granularity level: " + gran)
	}

	for idx, values := range valueMap {
		var sum float64 = 0
		for _, value := range values {
			sum += value
		}
		mean := sum / float64(len(values))
		y, d, err := helper.StringToPairOfIDs(idx)
		if err != nil {
			return nil, err // TODO: maybe save error, not exit right away
		}

		startDate := periodStart(y, d)
		aggregatedObs = append(aggregatedObs, domain.Observation{
			Value:       mean,
			Dataset:     dataset,
			Date:        &startDate,
			Granularity: gran,
		})
	}

	return aggregatedObs, nil
}

// this is needed to get correct (asc) observations order with one request with limit
func orderObservationsAsc(observations []*domain.Observation) []*domain.Observation {
	sort.Slice(observations, func(i, j int) bool {
		return observations[i].Date.Before(*observations[j].Date)
	})
	return observations
}
