package auth

import (
	"errors"
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
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			requestId := pkgContext.GetRequestIDFromContext(ctx)

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

			// Add token to request context
			ctx = pkgContext.SetTokenAuthorizationToContext(ctx, tokenHeader)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
