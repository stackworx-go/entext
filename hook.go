package entext

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
)

// AuditHook uses the context methods to exract the current user id and correct set it for create and update operations
// The created at and updated at values are also popuated if they are not already set
func AuditHook(next ent.Mutator) ent.Mutator {
	type AuditLogger interface {
		SetCreatedAt(time.Time)
		CreatedAt() (value time.Time, exists bool)
		SetCreatedByID(int)
		CreatedByID() (id int, exists bool)
		SetUpdatedAt(time.Time)
		UpdatedAt() (value time.Time, exists bool)
		SetUpdatedByID(int)
		UpdatedByID() (id int, exists bool)
	}
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(AuditLogger)
		if !ok {
			return nil, fmt.Errorf("unexpected audit-log call from mutation type %T", m)
		}
		userId := UserIDFromContext(ctx)
		if userId != nil {
			switch op := m.Op(); {
			case op.Is(ent.OpCreate):
				ml.SetCreatedAt(time.Now())
				if _, exists := ml.CreatedByID(); !exists {
					ml.SetCreatedByID(*userId)
				}
			case op.Is(ent.OpUpdateOne | ent.OpUpdate):
				ml.SetUpdatedAt(time.Now())
				if _, exists := ml.UpdatedByID(); !exists {
					ml.SetUpdatedByID(*userId)
				}
			}
		}
		return next.Mutate(ctx, m)
	})
}
