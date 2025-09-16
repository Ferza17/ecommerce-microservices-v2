package context

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const CtxKeyRequestID = "X-Request-Id"

func SetRequestIDToContext(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, CtxKeyRequestID, requestId)
}

func GetRequestIDFromContext(ctx context.Context) string {
	reqId, ok := ctx.Value(CtxKeyRequestID).(string)
	if !ok {
		return ""
	}
	return reqId
}

func SetRequestIDToMetadata(ctx context.Context, requestId string) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		authMd := metadata.Pairs(CtxKeyRequestID, requestId)
		return metadata.NewIncomingContext(ctx, authMd)
	}

	md.Append(CtxKeyRequestID, requestId)
	return metadata.NewIncomingContext(ctx, md)
}

func GetTokenFromMetadata(ctx context.Context) (string, error) {
	md := metadata.ValueFromIncomingContext(ctx, CtxKeyRequestID)
	if len(md) == 0 {
		return "", status.Error(codes.InvalidArgument, "metadata not found")
	}
	return md[0], nil
}
