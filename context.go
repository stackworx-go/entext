package entext

import "context"

// WithUserID puts the request ID into the current context.
func WithUserID(ctx context.Context, userId int) context.Context {
	return context.WithValue(ctx, contextUserIDKey, &userId)
}

// UserIDFromContext returns the request ID from the context.
func UserIDFromContext(ctx context.Context) *int {
	v := ctx.Value(contextUserIDKey)
	if v == nil {
		return nil
	}
	return v.(*int)
}

type contextUserIDType struct{}

var contextUserIDKey = &contextUserIDType{}
