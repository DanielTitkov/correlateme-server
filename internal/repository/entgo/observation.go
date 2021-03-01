package entgo

import (
	"context"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
)

func (r *EntgoRepository) CreateObservation(o *domain.Observation) (*domain.Observation, error) {
	create := r.client.Observation.
		Create().
		SetValue(o.Value).
		SetDatasetID(o.Dataset.ID)

	// because ent can't validate "empty" go date
	if o.Date != nil {
		create = create.SetDate(*o.Date)
	}

	obs, err := create.Save(context.TODO())
	if err != nil {
		return nil, err
	}

	return entToDomainObservation(obs), nil
}

func entToDomainObservation(obs *ent.Observation) *domain.Observation {
	var dataset *domain.Dataset
	if obs.Edges.Dataset != nil {
		dataset = entToDomainDataset(obs.Edges.Dataset)
	}

	return &domain.Observation{
		ID:         obs.ID,
		Value:      obs.Value,
		Dataset:    dataset,
		Date:       &obs.Date,
		CreateTime: obs.CreateTime,
		UpdateTime: obs.UpdateTime,
	}
}
