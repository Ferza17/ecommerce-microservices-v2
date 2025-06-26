package token

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const ctxKeyToken = "CONTEXT_KEY_TOKEN"

func SetTokenToContext(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, ctxKeyToken, token)
}

func GetTokenFromContext(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(ctxKeyToken).(string)
	return token, ok
}

func SetTokenToMetadata(ctx context.Context, token string) context.Context {
	return metadata.NewIncomingContext(ctx, metadata.New(map[string]string{ctxKeyToken: token}))
}

func GetTokenFromMetadata(ctx context.Context) (string, error) {
	md := metadata.ValueFromIncomingContext(ctx, ctxKeyToken)
	if len(md) == 0 {
		return "", status.Error(codes.InvalidArgument, "metadata not found")
	}
	return md[0], nil
}
