package entgo

import (
	"context"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
)

func (r *EntgoRepository) CreateDetectionInstanceJob(i *domain.DetectionJobInstance) (*domain.DetectionJobInstance, error) {
	ins, err := r.client.DetectionJobInstance.
		Create().
		SetDetectionJobID(i.DetectionJobID).
		SetStartedAt(i.StartedAt).
		SetFinishedAt(i.FinishedAt).
		Save(context.TODO())

	if err != nil {
		return nil, err
	}

	i.ID = ins.ID
	return i, nil
}
