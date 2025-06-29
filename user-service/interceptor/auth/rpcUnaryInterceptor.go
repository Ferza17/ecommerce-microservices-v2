package auth

import (
	"context"
	accessControlUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/usecase"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"go.uber.org/zap"
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

		// Validate is excluded method
		isExcluded, _ := accessControlUseCase.IsExcludedRPC(ctx, pkgContext.GetRequestIDFromContext(ctx), info.FullMethod)
		if err != nil {
			logger.Error("Interceptor.AuthRPCUnaryInterceptor", zap.Error(status.Error(codes.PermissionDenied, err.Error())))
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}

		// Bypass if excluded methods
		if isExcluded {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			logger.Error("Interceptor.AuthRPCUnaryInterceptor", zap.Error(status.Error(codes.Unauthenticated, "missing metadata")))
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		// Extract authorization header
		authHeaders := md.Get(pkgContext.CtxKeyAuthorization)
		if len(authHeaders) == 0 {
			logger.Error("Interceptor.AuthRPCUnaryInterceptor", zap.Error(status.Error(codes.Unauthenticated, "no authorization header")))
			return nil, status.Error(codes.Unauthenticated, "missing authorization header")
		}

		authHeader := authHeaders[0]
		if !strings.HasPrefix(authHeader, "Bearer ") {
			logger.Error("Interceptor.AuthRPCUnaryInterceptor", zap.Error(status.Error(codes.Unauthenticated, "invalid authorization header")))
			return nil, status.Error(codes.Unauthenticated, "invalid authorization header format")
		}

		tokenHeader := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenHeader == "" {
			logger.Error("Interceptor.AuthRPCUnaryInterceptor", zap.Error(status.Error(codes.Unauthenticated, "invalid authorization header")))
			return nil, status.Error(codes.Unauthenticated, "missing token")
		}

		ctx = pkgContext.SetAuthorizationToContext(ctx, tokenHeader)
		return handler(ctx, req)
	}
}
