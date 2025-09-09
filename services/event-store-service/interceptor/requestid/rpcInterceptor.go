package requestid

import (
	"context"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/context"
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
		} else {
			reqID = md[0]
		}

		ctx = pkgContext.SetRequestIDToContext(ctx, reqID)
		ctx = pkgContext.SetRequestIDToMetadata(ctx, reqID)

		defer func() {
			// Define response header metadata
			header := metadata.Pairs(
				pkgContext.CtxKeyRequestID, reqID,
			)

			if err == nil {
				// Send metadata as response headers
				if err = grpc.SetHeader(ctx, header); err != nil {
					return
				}
				return
			}

		}()

		return handler(ctx, req)
	}
}
