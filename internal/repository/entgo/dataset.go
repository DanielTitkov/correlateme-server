package entgo

import (
	"context"
	"errors"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/dataset"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/indicator"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/observation"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/user"
)

func (r *EntgoRepository) DatasetCount() (int, error) {
	return r.client.Dataset.Query().Count(context.TODO())
}

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

func (r *EntgoRepository) GetDatasetByID(id int, observationsLimit int, granularity string) (*domain.Dataset, error) {
	query := r.client.Dataset.
		Query().
		WithDatasetParams().
		WithIndicator(func(q *ent.IndicatorQuery) {
			q.WithAuthor()
			q.WithScale()
		}).
		WithUser().
		Where(dataset.IDEQ(id))

	if observationsLimit > 0 {
		query.WithObservations(func(q *ent.ObservationQuery) {
			q.Limit(observationsLimit)
			if granularity != "" {
				q.Where(observation.GranularityEQ(observation.Granularity(granularity)))
			} else {
				q.Where(observation.GranularityEQ(observation.Granularity(domain.GranularityDay)))
			}
			q.Order(ent.Desc(observation.FieldDate))
		})
	}

	ds, err := query.Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return entToDomainDataset(ds), nil
}

// GetUserDatasets fetches all datasets created or avaliable to user,
// including shared datasets.
func (r *EntgoRepository) GetUserDatasets(userID int, withShared bool, observationsLimit int, granularity string) ([]*domain.Dataset, error) {
	query := r.client.Dataset.Query().WithDatasetParams()

	if observationsLimit > 0 {
		query.WithObservations(func(q *ent.ObservationQuery) {
			q.Limit(observationsLimit)
			if granularity != "" {
				q.Where(observation.GranularityEQ(observation.Granularity(granularity)))
			} else {
				q.Where(observation.GranularityEQ(observation.Granularity(domain.GranularityDay)))
			}
			q.Order(ent.Desc(observation.FieldDate))
		})
	}

	if withShared {
		query.Where(dataset.Or(
			dataset.HasUserWith(user.IDEQ(userID)),
			dataset.SharedEQ(true),
		))
	} else {
		query.Where(dataset.HasUserWith(user.IDEQ(userID)))
	}

	dss, err := query.All(context.TODO())
	if err != nil {
		return nil, err
	}

	var datasets []*domain.Dataset
	for _, ds := range dss {
		datasets = append(datasets, entToDomainDataset(ds))
	}

	return datasets, nil
}

func (r *EntgoRepository) GetDatasets(args domain.GetDatasetsArgs) ([]*domain.Dataset, error) {
	query := r.client.Dataset.Query().WithDatasetParams()

	// edges
	if args.ObservationLimit > 0 {
		// TODO: refactor // DRY
		query.WithObservations(func(q *ent.ObservationQuery) {
			q.Limit(args.ObservationLimit)
			if args.Granularity != "" {
				q.Where(observation.GranularityEQ(observation.Granularity(args.Granularity)))
			} else {
				q.Where(observation.GranularityEQ(observation.Granularity(domain.GranularityDay)))
			}
			q.Order(ent.Desc(observation.FieldDate))
		})
	}

	if args.WithIndicator {
		query.WithIndicator()
	}

	if args.WithUser {
		query.WithUser()
	}

	// ownership
	// TODO: indicator?

	// filter
	if args.Filter.WithShared {
		query.Where(dataset.Or(
			dataset.HasUserWith(user.IDEQ(args.UserID)),
			dataset.SharedEQ(true),
		))
	} else {
		query.Where(dataset.HasUserWith(user.IDEQ(args.UserID)))
	}

	dss, err := query.All(context.TODO())
	if err != nil {
		return nil, err
	}

	var datasets []*domain.Dataset
	for _, ds := range dss {
		datasets = append(datasets, entToDomainDataset(ds))
	}

	return datasets, nil
}

func (r *EntgoRepository) GetUserIndicatorDataset(u *domain.User, ind *domain.Indicator) (*domain.Dataset, error) {
	ds, err := r.client.Dataset.
		Query().
		WithIndicator(func(q *ent.IndicatorQuery) {
			q.WithAuthor()
			q.WithScale()
		}).
		WithUser().
		WithDatasetParams().
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

	// restrict creating obasevations for other users' indicators
	if ind.Author != nil && ind.Author.ID != u.ID {
		return nil, errors.New("forbidden: indicator doesn't belong to user and is not built-in")
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

	var params *domain.DatasetParams
	if d.Edges.DatasetParams != nil {
		params.Style = d.Edges.DatasetParams.Style
		params.Aggregation = d.Edges.DatasetParams.Aggregation
	}

	var observations []*domain.Observation
	if d.Edges.Observations != nil {
		for _, obs := range d.Edges.Observations {
			observations = append(observations, entToDomainObservation(obs))
		}
	}

	return &domain.Dataset{
		ID:           d.ID,
		Shared:       d.Shared,
		User:         user,      // required
		Indicator:    indicator, // required
		Observations: observations,
		Params:       params,
		CreateTime:   d.CreateTime,
		UpdateTime:   d.UpdateTime,
		// Source: *d.Source, // not used by now
	}
}
