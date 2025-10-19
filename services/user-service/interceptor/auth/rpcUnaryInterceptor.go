package auth

import (
	"context"
	"strings"

	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/token"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthRPCUnaryInterceptor(
	logger logger.IZapLogger,
) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		// EXCLUDED Full Method
		if info.FullMethod == pb.AuthService_AuthUserRegister_FullMethodName ||
			info.FullMethod == pb.AuthService_AuthUserLoginByEmailAndPassword_FullMethodName ||
			info.FullMethod == pb.AuthService_AuthUserVerifyOtp_FullMethodName ||
			info.FullMethod == grpc_health_v1.Health_Check_FullMethodName {
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

		if _, err = token.ValidateJWTToken(tokenHeader, token.DefaultRefreshTokenConfig()); err != nil {
			logger.Error("Interceptor.AuthRPCUnaryInterceptor", zap.Error(status.Error(codes.Unauthenticated, "invalid authorization header")))
			return nil, token.MapErrorToGrpcStatus(err)
		}

		ctx = pkgContext.SetTokenAuthorizationToContext(ctx, tokenHeader)
		return handler(ctx, req)
	}
}
