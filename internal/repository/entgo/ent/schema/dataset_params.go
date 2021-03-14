package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/correlateme-server/internal/domain"
)

// DatasetParams holds the schema definition for the DatasetParams entity.
type DatasetParams struct {
	ent.Schema
}

// Fields of the DatasetParams.
func (DatasetParams) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("style", domain.DatasetStyle{}),
		field.JSON("aggregation", domain.DatasetAggregation{}),
	}
}

// Edges of the DatasetParams.
func (DatasetParams) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		// belongs to
		edge.From("dataset", Dataset.Type).Ref("dataset_params").Unique().Required(),
	}
}

func (DatasetParams) Indexes() []ent.Index {
	return []ent.Index{}
}

func (DatasetParams) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}
