package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
	}

}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("statuses", UserStatus.Type),
		edge.To("pets", Pet.Type),
		edge.To("groups", Group.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixinSelf,
	}
}
