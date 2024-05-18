package pkg

import "context"

func GetContextValue(ctx context.Context, key string) interface{} {
	return ctx.Value(key)
}
