package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// DictionaryEntry holds the schema definition for the DictionaryEntry entity.
type DictionaryEntry struct {
	ent.Schema
}

// Fields of the DictionaryEntry.
func (DictionaryEntry) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Immutable().Unique().Positive(),
		field.String("code").NotEmpty(),
		field.String("description").Optional(),
	}
}

// Edges of the DictionaryEntry.
func (DictionaryEntry) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		// edge.To("left", Correlation.Type),
		// edge.To("right", Correlation.Type),
		// edge.To("observations", Observation.Type),
		// edge.To("dataset_params", DatasetParams.Type).Unique(),
		// belongs to
		edge.From("dictionary", Dictionary.Type).Ref("entries").Unique().Required(),
	}
}

func (DictionaryEntry) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").Edges("dictionary").Unique(),
	}
}
