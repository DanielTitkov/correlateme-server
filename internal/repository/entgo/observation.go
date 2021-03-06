package entgo

import (
	"context"
	"errors"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/dataset"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/observation"
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

func (r *EntgoRepository) CreateOrUpdateObservation(o *domain.Observation) (*domain.Observation, error) {
	if o.Date == nil {
		return nil, errors.New("observation date is required")
	}

	obs, err := r.client.Observation.
		Query().
		Where(observation.And(
			observation.HasDatasetWith(dataset.IDEQ(o.Dataset.ID)),
			observation.DateEQ(*o.Date),
		)).
		Only(context.TODO())
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		// create observation
		obs, err = r.client.Observation.
			Create().
			SetValue(o.Value).
			SetDatasetID(o.Dataset.ID).
			SetDate(*o.Date).
			Save(context.TODO())
		if err != nil {
			return nil, err
		}
		return entToDomainObservation(obs), nil
	}

	// update observation
	obs, err = obs.Update().
		SetValue(o.Value).
		Save(context.TODO())
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
		ID:          obs.ID,
		Value:       obs.Value,
		Dataset:     dataset,
		Date:        &obs.Date,
		Granularity: obs.Granularity.String(),
		CreateTime:  obs.CreateTime,
		UpdateTime:  obs.UpdateTime,
	}
}
