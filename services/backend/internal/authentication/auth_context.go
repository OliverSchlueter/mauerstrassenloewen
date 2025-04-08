package authentication

import "context"

type authenticatedContextKey struct{}

func writeIsAuthenticated(ctx context.Context) context.Context {
	return context.WithValue(ctx, authenticatedContextKey{}, true)
}

func IsAuthenticated(ctx context.Context) bool {
	v := ctx.Value(authenticatedContextKey{})
	if v == nil {
		return false
	}

	return v.(bool)
}
