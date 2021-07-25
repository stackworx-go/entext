package schema

import (
	"entgo.io/ent/schema/mixin"
	"github.com/stackworx-go/entext"
)

type EmptyMixin struct {
	mixin.Schema
}

var (
	AuditMixin = entext.AuditMixin{
		AuditEntity: User.Type,
	}
	AuditMixinSelf = entext.AuditMixinSelf{
		AuditMixin: AuditMixin,
	}
	PointInTimeMixin = entext.PointInTimeMixin{}
)
