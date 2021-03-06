package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/DanielTitkov/correlateme-server/internal/domain"
)

// Correlation holds the schema definition for the Correlation entity.
type Correlation struct {
	ent.Schema
}

// Fields of the Correlation.
func (Correlation) Fields() []ent.Field {
	return []ent.Field{
		field.Float("coef"),
		field.Float("p"),
		field.Float("r2"),
		field.String("type").NotEmpty().Immutable(),
		field.Enum("granularity").Values(
			domain.GranularityDay,
			domain.GranularityWeek,
			domain.GranularityMonth,
		).Default(domain.GranularityDay),
	}
}

// Edges of the Correlation.
func (Correlation) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		// belongs to
		edge.From("left", Dataset.Type).Ref("left").Required().Unique(),
		edge.From("right", Dataset.Type).Ref("right").Required().Unique(),
	}
}

func (Correlation) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("left", "right").Unique(),
		index.Edges("right", "left").Unique(),
	}
}

func (Correlation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
