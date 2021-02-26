package entgo

import (
	"context"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
)

func (r *EntgoRepository) GetScales() ([]*domain.Scale, error) {
	res := []*domain.Scale{}

	scales, err := r.client.Scale.Query().All(context.TODO())
	if err != nil {
		return res, err
	}

	for _, s := range scales {
		res = append(res, entToDomainScale(s))
	}

	return res, nil
}

func entToDomainScale(s *ent.Scale) *domain.Scale {
	return &domain.Scale{
		ID:          s.ID,
		Type:        s.Type,
		Title:       s.Title,
		Description: s.Description,
	}
}
