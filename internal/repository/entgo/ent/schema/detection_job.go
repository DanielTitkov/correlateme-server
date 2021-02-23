package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// DetectionJob holds the schema definition for the DetectionJob entity.
type DetectionJob struct {
	ent.Schema
}

// Fields of the DetectionJob.
func (DetectionJob) Fields() []ent.Field {
	return []ent.Field{
		field.String("schedule").Nillable().Optional(), // TODO: add proper validation
		field.String("method").NotEmpty(),              // TODO: add enum
		field.String("site_id").NotEmpty(),
		field.String("metric").NotEmpty(),
		field.String("attribute").NotEmpty(),
		field.String("time_ago").NotEmpty(),
		field.String("time_step").NotEmpty(),
		field.String("description").Optional(),
	}
}

// Edges of the DetectionJob.
func (DetectionJob) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("instance", DetectionJobInstance.Type),
	}
}

// Indexes of the DetectionJob.
func (DetectionJob) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of the DetectionJob.
func (DetectionJob) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
