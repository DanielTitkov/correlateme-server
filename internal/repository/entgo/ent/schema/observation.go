package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/DanielTitkov/correlateme-server/internal/domain"
)

// Observation holds the schema definition for the Observation entity.
type Observation struct {
	ent.Schema
}

// Fields of the Observation.
func (Observation) Fields() []ent.Field {
	return []ent.Field{
		field.Float("value"),
		field.Time("date").SchemaType(map[string]string{
			dialect.Postgres: "date",
			dialect.MySQL:    "DATE",
		}),
		field.Enum("granularity").Values(
			domain.GranularityDay,
			domain.GranularityWeek,
			domain.GranularityMonth,
		).Default(domain.GranularityDay).Immutable(),
	}
}

// Edges of the Observation.
func (Observation) Edges() []ent.Edge {
	return []ent.Edge{
		// has

		// belongs to
		edge.From("dataset", Dataset.Type).Unique().Required().Ref("observations"),
	}
}

func (Observation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("date", "granularity").Edges("dataset").Unique(),
	}
}

func (Observation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
