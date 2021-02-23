package app

import "github.com/DanielTitkov/correlateme-server/internal/domain"

func (a *App) CreateAnomaly(anomaly *domain.Anomaly) error {
	_, err := a.repo.CreateAnomaly(anomaly)
	return err
}

func (a *App) ListAnomalies(args *domain.FilterAnomaliesArgs) ([]*domain.Anomaly, error) {
	return a.repo.FilterAnomalies(args)
}

func (a *App) SetAnomalyStatus(anomalyID int, processed bool) error {
	return a.repo.SetAnomalyStatus(anomalyID, processed)
}
