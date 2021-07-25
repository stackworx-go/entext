package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	// stackworxcontrib "gitlab.com/stackworx.io/go/entcontrib.git/internal/example/ent"
)

type UserStatus struct {
	ent.Schema
}

func (UserStatus) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
	}
}

func (UserStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Required().Unique().Ref("statuses").Field("user_id"),
	}
}

func (UserStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin,
		PointInTimeMixin,
	}
}
