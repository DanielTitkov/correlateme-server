package entgo

import (
	"context"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/dataset"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/indicator"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/user"
)

func (r *EntgoRepository) CreateDataset(d *domain.Dataset) (*domain.Dataset, error) {
	dataset, err := r.client.Dataset.
		Create().
		SetUserID(d.User.ID).
		SetIndicatorID(d.Indicator.ID).
		SetShared(d.Shared).
		SetNillableSource(&d.Source).
		Save(context.TODO())
	if err != nil {
		return nil, err
	}

	return entToDomainDataset(dataset), nil
}

func (r *EntgoRepository) GetDatasetByID(id int) (*domain.Dataset, error) {
	ds, err := r.client.Dataset.
		Query().
		WithIndicator(func(q *ent.IndicatorQuery) {
			q.WithAuthor()
			q.WithScale()
		}).
		WithUser().
		Where(dataset.IDEQ(id)).
		Only(context.TODO())
	if err != nil {
		return nil, err
	}

	// safe because dataset must have an indicator edge
	return entToDomainDataset(ds), nil
}

func (r *EntgoRepository) GetUserIndicatorDataset(u *domain.User, ind *domain.Indicator) (*domain.Dataset, error) {
	ds, err := r.client.Dataset.
		Query().
		WithIndicator(func(q *ent.IndicatorQuery) {
			q.WithAuthor()
			q.WithScale()
		}).
		WithUser().
		Where(dataset.And(
			dataset.HasIndicatorWith(indicator.IDEQ(ind.ID)),
			dataset.HasUserWith(user.IDEQ(u.ID)),
		)).
		Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return entToDomainDataset(ds), nil
}

func (r *EntgoRepository) GetOrCreateUserIndicatorDataset(u *domain.User, ind *domain.Indicator) (*domain.Dataset, error) {
	ds, err := r.GetUserIndicatorDataset(u, ind)
	if err == nil {
		return ds, nil
	}

	if !ent.IsNotFound(err) {
		return nil, err
	}

	dataset, err := r.client.Dataset.
		Create().
		SetUserID(u.ID).
		SetIndicatorID(ind.ID).
		Save(context.TODO())
	if err != nil {
		return nil, err
	}

	return entToDomainDataset(dataset), nil
}

func entToDomainDataset(d *ent.Dataset) *domain.Dataset {
	var user *domain.User
	if d.Edges.User != nil {
		user = entToDomainUser(d.Edges.User)
	}

	var indicator *domain.Indicator
	if d.Edges.Indicator != nil {
		indicator = entToDomainIndicator(d.Edges.Indicator)
	}

	return &domain.Dataset{
		ID:           d.ID,
		Shared:       d.Shared,
		User:         user,      // required
		Indicator:    indicator, // required
		Observations: nil,
		CreateTime:   d.CreateTime,
		UpdateTime:   d.UpdateTime,
		// Source: *d.Source, // not used by now
	}
}
