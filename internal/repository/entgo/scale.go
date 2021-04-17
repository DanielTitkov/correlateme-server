package entgo

import (
	"context"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/scale"
)

func (r *EntgoRepository) ScaleCount() (int, error) {
	return r.client.Scale.Query().Count(context.TODO())
}

func (r *EntgoRepository) GetScales() ([]*domain.Scale, error) {
	scales, err := r.client.Scale.Query().All(context.TODO())
	if err != nil {
		return nil, err
	}

	res := []*domain.Scale{}
	for _, s := range scales {
		res = append(res, entToDomainScale(s))
	}

	return res, nil
}

func (r *EntgoRepository) GetScaleByType(scaleType string) (*domain.Scale, error) {
	scale, err := r.client.Scale.Query().
		Where(scale.TypeEQ(scaleType)).
		Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return entToDomainScale(scale), nil
}

func (r *EntgoRepository) CreateScale(s domain.Scale) (*domain.Scale, error) {
	scale, err := r.client.Scale.
		Create().
		SetType(s.Type).
		SetTitle(s.Title).
		SetDescription(s.Description).
		Save(context.TODO())

	if err != nil {
		return nil, err
	}
	return entToDomainScale(scale), nil
}

func entToDomainScale(s *ent.Scale) *domain.Scale {
	return &domain.Scale{
		ID:          s.ID,
		Type:        s.Type,
		Title:       s.Title,
		Description: s.Description,
	}
}
