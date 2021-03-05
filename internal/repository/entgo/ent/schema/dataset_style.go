package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/correlateme-server/internal/domain"
)

// DatasetStyle holds the schema definition for the DatasetStyle entity.
type DatasetStyle struct {
	ent.Schema
}

// Fields of the DatasetStyle.
func (DatasetStyle) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("style", domain.DatasetStyle{}),
	}
}

// Edges of the DatasetStyle.
func (DatasetStyle) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		// belongs to
		edge.From("dataset", Dataset.Type).Ref("style").Unique().Required(),
	}
}

func (DatasetStyle) Indexes() []ent.Index {
	return []ent.Index{}
}

func (DatasetStyle) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}
