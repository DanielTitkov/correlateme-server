package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/correlateme-server/internal/domain"
)

// IndicatorParams holds the schema definition for the IndicatorParams entity.
type IndicatorParams struct {
	ent.Schema
}

// Fields of the IndicatorParams.
func (IndicatorParams) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("value_mapping", map[string]string{}).Optional(),
		field.JSON("value_params", domain.IndicatorValueParams{}),
	}
}

// Edges of the IndicatorParams.
func (IndicatorParams) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		// belongs to
		edge.From("indicator", Indicator.Type).Ref("indicator_params").Unique().Required(),
	}
}

func (IndicatorParams) Indexes() []ent.Index {
	return []ent.Index{}
}

func (IndicatorParams) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}
