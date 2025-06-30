package helper

import "context"

type contextKey string

const userIdKey contextKey = "userId"

func ContextWithUserId(ctx context.Context, userId int) context.Context {
	return context.WithValue(ctx, userIdKey, userId)
}

func GetUserIdFromContext(ctx context.Context) int {
	userId, _ := ctx.Value(userIdKey).(int)
	return userId
}
