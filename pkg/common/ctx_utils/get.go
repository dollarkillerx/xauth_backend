package ctx_utils

import (
	"context"
)

func Get(ctx context.Context, key string) string {
	role, _ := ctx.Value(key).(string)
	return role
}
