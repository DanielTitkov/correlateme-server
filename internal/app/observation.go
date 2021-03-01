package app

import (
	"github.com/DanielTitkov/correlateme-server/internal/domain"
)

// CreateObservation checks if there is a dataset for user/indicator.
// If none, creates one and adds observation to it, otherwise adds to existing dataset.
func (a *App) CreateObservation(args domain.CreateObservationArgs) error {
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

	_, err = a.repo.CreateObservation(&domain.Observation{
		Value:   args.Value,
		Dataset: dataset,
		Date:    args.Date,
	})
	if err != nil {
		return err
	}

	return nil
}
