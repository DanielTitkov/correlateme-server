package entgo

import (
	"context"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/dataset"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/indicator"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/observation"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/scale"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/user"
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

	return entToDomainIndicator(ind), nil
}

func (r *EntgoRepository) GetIndicatorByID(id int) (*domain.Indicator, error) {
	ind, err := r.client.Indicator.
		Query().
		WithAuthor().
		WithScale().
		Where(indicator.IDEQ(id)).
		Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return entToDomainIndicator(ind), nil
}

func (r *EntgoRepository) GetIndicators(args domain.GetIndicatorsArgs) ([]*domain.Indicator, error) {
	query := r.client.Indicator.Query().WithAuthor().WithScale()

	filter := args.Filter
	if filter.ID != nil {
		query = query.Where(indicator.IDIn(filter.ID...))
	}

	if filter.Code != nil {
		query = query.Where(indicator.CodeIn(filter.Code...))
	}

	if filter.Title != nil {
		query = query.Where(indicator.TitleIn(filter.Title...))
	}

	if filter.Active != nil {
		query = query.Where(indicator.ActiveEQ(*filter.Active))
	}

	if filter.BuiltIn != nil {
		query = query.Where(indicator.BuiltInEQ(*filter.BuiltIn))
	}

	if filter.External != nil {
		query = query.Where(indicator.ExternalEQ(*filter.External))
	}

	if filter.AuthorID != nil {
		query = query.Where(indicator.HasAuthorWith(user.IDEQ(*filter.AuthorID)))
	}

	if filter.ScaleType != nil {
		query = query.Where(indicator.HasScaleWith(scale.TypeEQ(*filter.ScaleType)))
	}

	if args.WithDataset {
		query = query.WithDatasets(func(q *ent.DatasetQuery) {
			q.Where(dataset.HasUserWith(user.IDEQ(args.UserID)))
			if args.ObservationLimit > 0 {
				q.WithObservations(func(q *ent.ObservationQuery) {
					if args.Granularity != "" {
						q.Where(observation.GranularityEQ(observation.Granularity(args.Granularity)))
					} else {
						q.Where(observation.GranularityEQ(observation.Granularity(domain.GranularityDay)))
					}
					q.Order(ent.Desc(observation.FieldDate))
					q.Limit(int(args.ObservationLimit))
				})
			}
		})
	}

	inds, err := query.All(context.TODO())
	if err != nil {
		return nil, err
	}

	var res []*domain.Indicator
	for _, ind := range inds {
		res = append(res, entToDomainIndicator(ind))
	}

	return res, nil
}

func entToDomainIndicator(ind *ent.Indicator) *domain.Indicator {
	var scale *domain.Scale
	if ind.Edges.Scale != nil {
		scale = entToDomainScale(ind.Edges.Scale)
	}

	var author *domain.User
	if ind.Edges.Author != nil {
		author = entToDomainUser(ind.Edges.Author)
	}

	var dataset *domain.Dataset
	if ind.Edges.Datasets != nil && len(ind.Edges.Datasets) == 1 {
		dataset = entToDomainDataset(ind.Edges.Datasets[0])
	}

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
		UserDataset: dataset,
		CreateTime:  ind.CreateTime,
		UpdateTime:  ind.UpdateTime,
	}
}
