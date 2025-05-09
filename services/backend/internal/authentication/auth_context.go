package authentication

import (
	"context"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/usermanagement"
)

type userContextKey struct{}

func writeUser(ctx context.Context, u *usermanagement.User) context.Context {
	return context.WithValue(ctx, userContextKey{}, u)
}

func UserFromCtx(ctx context.Context) *usermanagement.User {
	v := ctx.Value(userContextKey{})
	if v == nil {
		return nil
	}

	return v.(*usermanagement.User)
}
