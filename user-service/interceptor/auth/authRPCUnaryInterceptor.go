package auth

import (
	"context"
	"fmt"
	accessControlUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

func AuthRPCUnaryInterceptor(
	logger logger.IZapLogger,
	accessControlUseCase accessControlUseCase.IAccessControlUseCase,
) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		//TODO:
		// 1. Add Interceptor RequestID
		// 2. Get RequestID From context

		// Validate is excluded method
		isExcluded, err := accessControlUseCase.IsExcludedUrl(ctx, "", info.FullMethod)
		if err != nil {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}

		if isExcluded {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			logger.Error(fmt.Sprintf("missing metadata"))
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		// Extract authorization header
		authHeaders := md.Get("authorization")
		if len(authHeaders) == 0 {
			logger.Error(fmt.Sprintf("missing authorization header"))
			return nil, status.Error(codes.Unauthenticated, "missing authorization header")
		}

		authHeader := authHeaders[0]
		if !strings.HasPrefix(authHeader, "Bearer ") {
			logger.Error("Invalid authorization header format")
			return nil, status.Error(codes.Unauthenticated, "invalid authorization header format")
		}

		tokenHeader := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenHeader == "" {
			logger.Error("Missing token")
			return nil, status.Error(codes.Unauthenticated, "missing token")
		}

		ctx = token.SetTokenToContext(ctx, tokenHeader)
		return handler(ctx, req)
	}
}
