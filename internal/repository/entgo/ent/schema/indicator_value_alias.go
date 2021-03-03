package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// IndicatorValueAlias holds the schema definition for the IndicatorValueAlias entity.
type IndicatorValueAlias struct {
	ent.Schema
}

// Fields of the IndicatorValueAlias.
func (IndicatorValueAlias) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("value_mapping", map[float64]string{}).Optional(),
	}
}

// Edges of the IndicatorValueAlias.
func (IndicatorValueAlias) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		// belongs to
		edge.From("indicator", Indicator.Type).Ref("indicator_value_alias").Unique().Required(),
	}
}

func (IndicatorValueAlias) Indexes() []ent.Index {
	return []ent.Index{}
}

func (IndicatorValueAlias) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}
