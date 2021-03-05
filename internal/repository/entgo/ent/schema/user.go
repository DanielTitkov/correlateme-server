package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("password_hash"),
		field.Bool("service").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("indicators", Indicator.Type),
		edge.To("datasets", Dataset.Type),
		edge.To("settings", UserSettings.Type),
	}
}

func (User) Indexes() []ent.Index {
	return nil
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
