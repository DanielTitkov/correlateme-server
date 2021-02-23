package entgo

import (
	"context"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/anomaly"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/detectionjob"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/detectionjobinstance"
)

func (r *EntgoRepository) CreateAnomaly(a *domain.Anomaly) (*domain.Anomaly, error) {
	anom, err := r.client.Anomaly.
		Create().
		SetDetectionJobInstanceID(a.DetectionJobInstanceID).
		SetProcessed(a.Processed).
		SetType(a.Type).
		SetValue(a.Value).
		SetPeriodStart(a.PeriodStart).
		SetPeriodEnd(a.PeriodEnd).
		Save(context.TODO())

	if err != nil {
		return nil, err
	}

	a.ID = anom.ID
	return a, nil
}

func (r *EntgoRepository) FilterAnomalies(args *domain.FilterAnomaliesArgs) ([]*domain.Anomaly, error) {
	query := r.client.Anomaly.Query().WithDetectionJobInstance(
		func(q *ent.DetectionJobInstanceQuery) {
			q.WithDetectionJob()
		},
	)

	if args.JobID != 0 {
		query = query.Where(
			anomaly.HasDetectionJobInstanceWith(
				detectionjobinstance.HasDetectionJobWith(
					detectionjob.IDEQ(args.JobID),
				),
			),
		)
	}

	if args.Processed != nil {
		query = query.Where(anomaly.ProcessedEQ(*args.Processed))
	}

	anoms, err := query.All(context.TODO())
	if err != nil {
		return []*domain.Anomaly{}, err
	}

	var res []*domain.Anomaly
	for _, anom := range anoms {
		res = append(res, &domain.Anomaly{
			ID:                     anom.ID,
			DetectionJobInstanceID: anom.Edges.DetectionJobInstance.ID,
			DetectionJobID:         anom.Edges.DetectionJobInstance.Edges.DetectionJob.ID,
			Type:                   anom.Type,
			Value:                  anom.Value,
			Processed:              anom.Processed,
			PeriodStart:            anom.PeriodStart,
			PeriodEnd:              anom.PeriodEnd,
		})
	}

	return res, nil
}

func (r *EntgoRepository) SetAnomalyStatus(anomalyID int, processed bool) error {
	_, err := r.client.Anomaly.
		UpdateOneID(anomalyID).
		SetProcessed(processed).
		Save(context.TODO())

	return err
}
