package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Anomaly holds the schema definition for the Anomaly entity.
type Anomaly struct {
	ent.Schema
}

// Fields of the Anomaly.
func (Anomaly) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").NotEmpty(),
		field.Float("value"),
		field.Bool("processed").Default(false),
		field.Time("period_start"),
		field.Time("period_end"),
	}
}

// Edges of the Anomaly.
func (Anomaly) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("detection_job_instance", DetectionJobInstance.Type).Ref("anomalies").Unique().Required(),
	}
}

// Mixin of the Anomaly.
func (Anomaly) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
