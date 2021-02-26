package entgo

import (
	"context"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
)

func (r *EntgoRepository) CreateIndicator(i *domain.Indicator) (*domain.Indicator, error) {
	ind, err := r.client.Indicator.
		Create().
		SetCode(i.Code).
		SetAuthorID(i.Author.ID).
		SetScaleID(i.Scale.ID).
		SetActive(i.Active).
		SetBuiltIn(i.BuiltIn).
		SetExternal(i.External).
		SetDescription(i.Description).
		SetTitle(i.Title).
		Save(context.TODO())
	if err != nil {
		return nil, err
	}

	return entToDomainIndicator(ind, i.Scale, i.Author), nil
}

func (r *EntgoRepository) FilterIndicators(filter domain.FilterIndicatorsArgs) ([]*domain.Indicator, error) {
	return []*domain.Indicator{}, nil
}

func entToDomainIndicator(ind *ent.Indicator, scale *domain.Scale, author *domain.User) *domain.Indicator {
	return &domain.Indicator{
		ID:          ind.ID,
		Code:        ind.Code,
		Author:      author,
		Scale:       scale,
		Title:       ind.Title,
		Description: ind.Description,
		Active:      ind.Active,
		BuiltIn:     ind.BuiltIn,
		External:    ind.External,
		CreateTime:  ind.CreateTime,
		UpdateTime:  ind.UpdateTime,
	}
}
