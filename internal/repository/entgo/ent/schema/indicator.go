package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Indicator holds the schema definition for the Indicator entity.
type Indicator struct {
	ent.Schema
}

// Fields of the Indicator.
func (Indicator) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Unique().NotEmpty(),
		field.String("title").NotEmpty(),
		field.String("description").Optional(),
		field.Bool("active").Default(true),
		field.Bool("built_in").Default(false),
		field.Bool("external").Default(false),
	}
}

// Edges of the Indicator.
func (Indicator) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("datasets", Dataset.Type),
		// belongs to
		edge.From("author", User.Type).Ref("indicators").Unique(),
		edge.From("scale", Scale.Type).Ref("indicators").Unique().Required(),
	}
}

func (Indicator) Indexes() []ent.Index {
	// TODO: unique title for user
	return nil
}

func (Indicator) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
