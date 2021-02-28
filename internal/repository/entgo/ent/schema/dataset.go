package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Dataset holds the schema definition for the Dataset entity.
type Dataset struct {
	ent.Schema
}

// Fields of the Dataset.
func (Dataset) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("shared").Default(false),
		field.String("source").Nillable().Optional(),
	}
}

// Edges of the Dataset.
func (Dataset) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("left", Correlation.Type),
		edge.To("right", Correlation.Type),
		edge.To("observations", Observation.Type),
		// belongs to
		edge.From("indicator", Indicator.Type).Ref("datasets").Required().Unique(),
		edge.From("user", User.Type).Ref("datasets").Unique(),
	}
}

func (Dataset) Indexes() []ent.Index {
	return []ent.Index{
		// one dataset for one indicator for user
		index.Edges("indicator").Edges("user").Unique(),
	}
}

func (Dataset) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
