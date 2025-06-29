package context

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const CtxKeyAuthorization = "authorization"

func SetAuthorizationToContext(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, CtxKeyAuthorization, token)
}

func GetAuthorizationFromContext(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(CtxKeyAuthorization).(string)
	return token, ok
}

func SetAuthorizationToMetadata(ctx context.Context, token string) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		authMd := metadata.Pairs(CtxKeyAuthorization, token)
		return metadata.NewIncomingContext(ctx, authMd)
	}

	md.Append(CtxKeyAuthorization, token)
	return metadata.NewIncomingContext(ctx, md)
}

func GetAuthorizationFromMetadata(ctx context.Context) (string, error) {
	md := metadata.ValueFromIncomingContext(ctx, CtxKeyAuthorization)
	if len(md) == 0 {
		return "", status.Error(codes.InvalidArgument, "metadata not found")
	}
	return md[0], nil
}
