package entgo

import (
	"context"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/correlation"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/dataset"
)

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
				),
				correlation.And(
					correlation.HasLeftWith(dataset.IDEQ(c.Right.ID)),
					correlation.HasRightWith(dataset.IDEQ(c.Left.ID)),
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
			Save(context.TODO())
		if err != nil {
			return nil, err
		}
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

func entToDomainCorrelation(corr *ent.Correlation) *domain.Correlation {
	var left, right *domain.Dataset
	if corr.Edges.Left != nil {
		left = entToDomainDataset(corr.Edges.Left)
	}

	if corr.Edges.Right != nil {
		right = entToDomainDataset(corr.Edges.Right)
	}

	return &domain.Correlation{
		ID:    corr.ID,
		Left:  left,
		Right: right,
		Coef:  corr.Coef,
		P:     corr.P,
		R2:    corr.R2,
		Type:  corr.Type,
	}
}
