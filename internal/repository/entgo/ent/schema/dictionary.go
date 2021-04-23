package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Dictionary holds the schema definition for the Dictionary entity.
type Dictionary struct {
	ent.Schema
}

// Fields of the Dictionary.
func (Dictionary) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Immutable().Unique().Positive(),
		field.String("code").NotEmpty().Unique(),
		field.String("description").Optional(),
	}
}

// Edges of the Dictionary.
func (Dictionary) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("entries", DictionaryEntry.Type),
		// belongs to
	}
}
