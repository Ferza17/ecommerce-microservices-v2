package auth

import (
	"errors"
	userService "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/service/user"
	pb "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/response"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// AuthHTTPMiddleware  returns HTTP middleware for JWT authentication
func AuthHTTPMiddleware(
	logger logger.IZapLogger,
	userService userService.IUserService,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			requestId := pkgContext.GetRequestIDFromContext(ctx)
			method := strings.ToLower(r.Method)
			url := strings.ToLower(r.URL.Path)

			// Validate is an excluded method
			authExcluded, err := userService.AuthServiceVerifyIsExcluded(ctx, pkgContext.GetRequestIDFromContext(ctx), &pb.AuthServiceVerifyIsExcludedRequest{
				HttpUrl:    &url,
				HttpMethod: &method,
			})
			if err != nil {
				logger.Error("Interceptor.AuthHTTPMiddleware", zap.String("requestId", requestId), zap.Error(errors.New("no url found on access control list")))
				response.WriteErrorResponse(w, http.StatusUnauthorized, "UNAUTHENTICATED", errors.New("no url found on access control list"))
				return
			}

			// Bypass if excluded methods
			if authExcluded.IsExcluded {
				next.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				logger.Error("Interceptor.AuthHTTPMiddleware", zap.String("requestId", requestId), zap.Error(errors.New("no authorization header")))
				response.WriteErrorResponse(w, http.StatusUnauthorized, "UNAUTHENTICATED", errors.New("no authorization header"))
				return
			}

			if !strings.HasPrefix(authHeader, "Bearer ") {
				logger.Error("Interceptor.AuthHTTPMiddleware", zap.String("requestId", requestId), zap.Error(errors.New("invalid authorization header format")))
				response.WriteErrorResponse(w, http.StatusUnauthorized, "UNAUTHENTICATED", errors.New("invalid authorization header format"))
				return
			}

			tokenHeader := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenHeader == "" {
				logger.Error("Interceptor.AuthHTTPMiddleware", zap.String("requestId", requestId), zap.Error(errors.New("invalid authorization header format")))
				response.WriteErrorResponse(w, http.StatusUnauthorized, "UNAUTHENTICATED", errors.New("invalid authorization header format"))
				return
			}

			// Access Control Authorization
			acl, err := userService.AuthUserVerifyAccessControl(
				ctx,
				requestId,
				&pb.AuthUserVerifyAccessControlRequest{
					Token:      tokenHeader,
					HttpUrl:    &url,
					HttpMethod: &method,
				},
			)
			if err != nil {
				logger.Error("Interceptor.AuthHTTPMiddleware", zap.String("requestId", requestId), zap.Error(errors.New("invalid authorization access control")))
				response.WriteErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", errors.New("invalid authorization access control"))
				return
			}

			if !acl.IsValid {
				logger.Error("Interceptor.AuthHTTPMiddleware", zap.String("requestId", requestId), zap.Error(errors.New("permission denied")))
				response.WriteErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", errors.New("permission denied"))
				return
			}

			// Add token to request context
			ctx = pkgContext.SetTokenAuthorizationToContext(ctx, tokenHeader)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
