package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// DetectionJobInstance holds the schema definition for the DetectionJobInstance entity.
type DetectionJobInstance struct {
	ent.Schema
}

// Fields of the DetectionJobInstance.
func (DetectionJobInstance) Fields() []ent.Field {
	return []ent.Field{
		field.Time("started_at").Nillable().Optional(),
		field.Time("finished_at").Nillable().Optional(),
	}
}

// Edges of the DetectionJobInstance.
func (DetectionJobInstance) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("anomalies", Anomaly.Type),
		// belongs to
		edge.From("detection_job", DetectionJob.Type).Ref("instance").Unique().Required(),
	}
}

// Mixin of the DetectionJobInstance.
func (DetectionJobInstance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
