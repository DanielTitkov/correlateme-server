package entgo

import (
	"context"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/correlation"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/dataset"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/observation"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/user"
)

func (r *EntgoRepository) CorrelationCount() (int, error) {
	return r.client.Correlation.Query().Count(context.TODO())
}

func (r *EntgoRepository) CreateOrUpdateCorrelation(c *domain.Correlation) (*domain.Correlation, error) {
	corr, err := r.client.Correlation.
		Query().
		Where(
			// ent doesn't allow to create index with custom expressions
			// so this logic is needed to prevent duplicates
			correlation.Or(
				correlation.And(
					correlation.HasLeftWith(dataset.IDEQ(c.Left.ID)),
					correlation.HasRightWith(dataset.IDEQ(c.Right.ID)),
					correlation.GranularityEQ(correlation.Granularity(c.Granularity)),
				),
				correlation.And(
					correlation.HasLeftWith(dataset.IDEQ(c.Right.ID)),
					correlation.HasRightWith(dataset.IDEQ(c.Left.ID)),
					correlation.GranularityEQ(correlation.Granularity(c.Granularity)),
				),
			)).
		Only(context.TODO())
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		// create correlation
		corr, err = r.client.Correlation.
			Create().
			SetLeftID(c.Left.ID).
			SetRightID(c.Right.ID).
			SetCoef(c.Coef).
			SetType(c.Type).
			SetP(c.P).
			SetR2(c.R2).
			SetGranularity(correlation.Granularity(c.Granularity)).
			Save(context.TODO())
		if err != nil {
			return nil, err
		}
		return entToDomainCorrelation(corr), nil
	}

	// update correlation
	corr, err = corr.Update().
		SetCoef(c.Coef).
		SetP(c.P).
		SetR2(c.R2).
		Save(context.TODO())
	if err != nil {
		return nil, err
	}

	return entToDomainCorrelation(corr), nil
}

func (r *EntgoRepository) GetUserCorrelations(userID int, granularity string) ([]*domain.Correlation, error) {
	corrs, err := r.client.Correlation.
		Query().
		Where(correlation.Or( // TODO: maybe And?
			correlation.HasLeftWith(dataset.HasUserWith(user.IDEQ(userID))),
			correlation.HasRightWith(dataset.HasUserWith(user.IDEQ(userID))),
		)).
		Where(correlation.GranularityEQ(correlation.Granularity(granularity))).
		WithLeft(func(q *ent.DatasetQuery) {
			q.WithIndicator()
		}).
		WithRight(func(q *ent.DatasetQuery) {
			q.WithIndicator()
		}).
		All(context.TODO())
	if err != nil {
		return nil, err
	}

	var res []*domain.Correlation
	for _, corr := range corrs {
		res = append(res, entToDomainCorrelation(corr))
	}

	return res, nil
}

func (r *EntgoRepository) GetCorrelation(args domain.GetCorrelationArgs) (*domain.Correlation, error) {
	query := r.client.Correlation.
		Query().
		Where(correlation.And(
			correlation.IDEQ(args.ID),
			correlation.HasLeftWith(dataset.HasUserWith(user.IDEQ(args.UserID))),
			correlation.HasRightWith(dataset.HasUserWith(user.IDEQ(args.UserID))),
		))

	// need to know correlation granularity.
	// TODO: this is ugly, maybe there is a better way?
	corr, err := query.Only(context.TODO())
	if err != nil {
		return nil, err
	}

	if args.WithDatasets {
		var dsQuery func(q *ent.DatasetQuery)
		if args.WithObservations {
			dsQuery = func(q *ent.DatasetQuery) {
				q.WithObservations(func(q *ent.ObservationQuery) {
					q.Limit(args.ObservationLimit)
					q.Where(observation.GranularityEQ(observation.Granularity(corr.Granularity)))
					q.Order(ent.Desc(observation.FieldDate))
				})
			}
		}
		query.WithLeft(dsQuery)
		query.WithRight(dsQuery)
	}

	corr, err = query.Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return entToDomainCorrelation(corr), nil
}

func entToDomainCorrelation(corr *ent.Correlation) *domain.Correlation {
	var left, right *domain.Dataset
	if corr.Edges.Left != nil {
		left = entToDomainDataset(corr.Edges.Left)
	}

	if corr.Edges.Right != nil {
		right = entToDomainDataset(corr.Edges.Right)
	}

	return &domain.Correlation{
		ID:          corr.ID,
		Left:        left,
		Right:       right,
		Coef:        corr.Coef,
		P:           corr.P,
		R2:          corr.R2,
		Type:        corr.Type,
		Granularity: corr.Granularity.String(),
		UpdateTime:  corr.UpdateTime,
	}
}
