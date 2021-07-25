package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Pet holds the schema definition for the User entity.
type Pet struct {
	ent.Schema
}

// Pet of the User.
func (Pet) Fields() []ent.Field {
	return nil
}

// Pet of the User.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
	}
}
