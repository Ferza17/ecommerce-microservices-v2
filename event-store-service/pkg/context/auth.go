package context

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const CtxKeyAuthorization = "authorization"

func SetTokenAuthorizationToContext(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, CtxKeyAuthorization, token)
}

func GetTokenAuthorizationFromContext(ctx context.Context) string {
	token, _ := ctx.Value(CtxKeyAuthorization).(string)
	return token
}

func SetTokenAuthorizationToMetadata(ctx context.Context, token string) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		authMd := metadata.Pairs(CtxKeyAuthorization, token)
		return metadata.NewIncomingContext(ctx, authMd)
	}

	md.Append(CtxKeyAuthorization, token)
	return metadata.NewIncomingContext(ctx, md)
}

func GetTokenAuthorizationFromMetadata(ctx context.Context) (string, error) {
	md := metadata.ValueFromIncomingContext(ctx, CtxKeyAuthorization)
	if len(md) == 0 {
		return "", status.Error(codes.InvalidArgument, "metadata not found")
	}
	return md[0], nil
}
