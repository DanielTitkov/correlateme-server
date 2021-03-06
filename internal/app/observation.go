package app

import (
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

	_, err = a.repo.CreateOrUpdateObservation(&domain.Observation{
		Value:   args.Value,
		Dataset: dataset,
		Date:    args.Date,
	})
	if err != nil {
		return err
	}

	go func() {
		// TODO: add timeout
		metrics.UnprocessedUpdateCorrelationsRequests.Add(1)
		a.Channels.UpdateUserCorrelationsChan <- domain.UpdateCorrelationsArgs{
			UserID:     args.UserID,
			WithShared: true,
			Method:     "auto",
		}
	}()

	return nil
}

func (a *App) UpdateAggregations(args domain.UpdateAggregationsArgs) error {
	// get dataset with observations
	dataset, err := a.repo.GetDatasetByID(args.DatasetID, a.cfg.App.MaxMonthAggregationObservations, "day")
	if err != nil {
		return err
	}

	// calculate aggregated values
	valueMap := make(map[string][]float64)
	for _, obs := range dataset.Observations {
		y, w := obs.Date.ISOWeek()
		week := helper.PairOfIDsToString(y, w)
		valueMap[week] = append(valueMap[week], obs.Value)
	}

	var aggregatedObs []domain.Observation
	for week, values := range valueMap {
		var sum float64 = 0
		for _, value := range values {
			sum += value
		}
		mean := sum / float64(len(values))
		y, w, err := helper.StringToPairOfIDs(week)
		if err != nil {
			return err // TODO: maybe save error, not exit right now
		}

		weekStart := helper.WeekStart(y, w)
		aggregatedObs = append(aggregatedObs, domain.Observation{
			Value:       mean,
			Dataset:     dataset,
			Date:        &weekStart,
			Granularity: domain.GranularityWeek,
		})
	}

	// save observations
	for _, obs := range aggregatedObs { // TODO: change to bulk? but ent doens't support that as yet
		_, err := a.repo.CreateOrUpdateObservation(&obs)
		if err != nil {
			return err // TODO: maybe save error, not exit right now
		}
	}

	// TODO: add month

	return nil
}
