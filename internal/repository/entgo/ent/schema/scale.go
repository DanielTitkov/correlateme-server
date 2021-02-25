package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Scale holds the schema definition for the Scale entity.
type Scale struct {
	ent.Schema
}

// Fields of the Scale.
func (Scale) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Unique().NotEmpty(),
		field.String("title").Unique().NotEmpty(),
		field.String("description").Unique().NotEmpty(),
	}
}

// Edges of the Scale.
func (Scale) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("indicators", Indicator.Type),
	}
}

func (Scale) Indexes() []ent.Index {
	return nil
}

func (Scale) Mixin() []ent.Mixin {
	return nil
}
