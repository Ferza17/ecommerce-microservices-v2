package auth

import (
	"errors"
	"fmt"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	accessControlUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/usecase"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/response"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strings"
)

// AuthHTTPMiddleware  returns HTTP middleware for JWT authentication
func AuthHTTPMiddleware(
	logger logger.IZapLogger,
	accessControlUseCase accessControlUseCase.IAccessControlUseCase,
	authUseCase authUseCase.IAuthUseCase,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			requestId := pkgContext.GetRequestIDFromContext(ctx)
			method := strings.ToLower(r.Method)
			url := strings.ToLower(r.URL.Path)

			route, _ := mux.CurrentRoute(r).GetPathTemplate()

			log.Printf("Request Path: %s, Matched Path Template: %s", r.URL.Path, route)

			// Validate is excluded method
			isExcluded, _ := accessControlUseCase.IsExcludedHTTP(
				ctx,
				pkgContext.GetRequestIDFromContext(ctx),
				method,
				url,
			)
			// Bypass if excluded methods
			if isExcluded {
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
			acl, err := authUseCase.AuthUserVerifyAccessControl(
				ctx,
				pkgContext.GetRequestIDFromContext(ctx),
				&pb.AuthUserVerifyAccessControlRequest{
					Token:      tokenHeader,
					HttpUrl:    &url,
					HttpMethod: &method,
				},
			)
			if err != nil {
				logger.Error(fmt.Sprintf("invalid authorization access control"))
				response.WriteErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", errors.New("invalid authorization access control"))
				return
			}

			if !acl.IsValid {
				logger.Error(fmt.Sprintf("Permission Denied"))
				response.WriteErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", errors.New("permission denied"))
				return
			}

			// Add token to request context
			ctx = pkgContext.SetTokenAuthorizationToContext(ctx, tokenHeader)
			ctx = pkgContext.SetTokenAuthorizationToMetadata(ctx, tokenHeader)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
