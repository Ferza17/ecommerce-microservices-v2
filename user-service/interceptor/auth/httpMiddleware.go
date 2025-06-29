package auth

import (
	"errors"
	"fmt"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/response"
	"net/http"
	"strings"
)

// AuthHTTPMiddleware  returns HTTP middleware for JWT authentication
func AuthHTTPMiddleware(
	logger logger.IZapLogger,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				logger.Error(fmt.Sprintf("missing header"))
				response.WriteErrorResponse(w, http.StatusUnauthorized, "UNAUTHENTICATED", errors.New("UNAUTHENTICATED"))
				return
			}

			if !strings.HasPrefix(authHeader, "Bearer ") {
				logger.Error(fmt.Sprintf("invalid authorization header format"))
				response.WriteErrorResponse(w, http.StatusUnauthorized, "UNAUTHENTICATED", errors.New("invalid authorization header format"))
				return
			}

			tokenHeader := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenHeader == "" {
				logger.Error(fmt.Sprintf("invalid authorization header format"))
				response.WriteErrorResponse(w, http.StatusUnauthorized, "UNAUTHENTICATED", errors.New("invalid authorization header format"))
				return
			}

			// Add token to request context
			ctx := pkgContext.SetAuthorizationToContext(r.Context(), tokenHeader)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
