package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// UserSettings holds the schema definition for the UserSettings entity.
type UserSettings struct {
	ent.Schema
}

// Fields of the UserSettings.
func (UserSettings) Fields() []ent.Field {
	return []ent.Field{
		// field.JSON("style", domain.UserSettings{}),
	}
}

// Edges of the UserSettings.
func (UserSettings) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		// belongs to
		edge.From("user", User.Type).Ref("settings").Unique().Required(),
	}
}

func (UserSettings) Indexes() []ent.Index {
	return []ent.Index{}
}

func (UserSettings) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}
