package entext

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/jackc/pgtype"
)

// AuditMixin struct
type AuditMixin struct {
	AuditEntity interface{}
}

// Fields Default audit fields
func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Int("created_by_id").StorageKey("created_by"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int("updated_by_id").Optional().StorageKey("updated_by"),
	}
}

func (m AuditMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("created_by", m.AuditEntity).Required().Unique().Field("created_by_id"),
		edge.To("updated_by", m.AuditEntity).Unique().Field("updated_by_id"),
	}
}

func (AuditMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		AuditHook,
	}
}

func (AuditMixin) Indexes() []ent.Index {
	return nil
}

func (AuditMixin) Policy() ent.Policy {
	return nil
}

func (AuditMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{
			Audited: true,
		},
	}
}

// AuditMixinSelf struct is used for the user table to reference itself
type AuditMixinSelf struct {
	AuditMixin
}

// Fields Default audit fields
func (AuditMixinSelf) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Int("created_by_id").Optional().StorageKey("created_by"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int("updated_by_id").Optional().StorageKey("updated_by"),
	}
}

func (m AuditMixinSelf) Edges() []ent.Edge {
	return []ent.Edge{
		// Removed until https://github.com/ent/ent/issues/1768 is resolved
		// edge.To("created_by", m.AuditEntity).
		// 	Unique().
		// 	Field("created_by_id"),
		// edge.To("updated_by", m.AuditEntity).
		// 	Unique().
		// 	Field("created_by_id"),
	}
}

// PointInTimeMixin struct
type PointInTimeMixin struct{}

// Fields Default audit fields
func (PointInTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("active").
			Default(true),
		field.Other("duration", &pgtype.Tstzrange{}).
			Annotations(&entsql.Annotation{
				Default: "tstzrange(now(), 'infinity', '[)')",
			}).
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "tstzrange",
			}),
	}
}

func (PointInTimeMixin) Edges() []ent.Edge {
	return nil
}

func (PointInTimeMixin) Hooks() []ent.Hook {
	return nil
}

func (PointInTimeMixin) Indexes() []ent.Index {
	return nil
}

func (PointInTimeMixin) Policy() ent.Policy {
	return nil
}

func (PointInTimeMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{
			PointInTime: true,
		},
	}
}
