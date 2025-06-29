package requestid

import (
	"context"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func RequestIDRPCInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		reqID := ""
		md := metadata.ValueFromIncomingContext(ctx, pkgContext.CtxKeyRequestID)
		if len(md) == 0 {
			reqID = uuid.NewString()
		}

		ctx = pkgContext.SetRequestIDToContext(ctx, reqID)
		ctx = pkgContext.SetRequestIDToMetadata(ctx, reqID)
		return handler(ctx, req)
	}
}
