package requestid

import (
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/google/uuid"
	"net/http"
)

func RequestIDHTTPMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := r.Header.Get(pkgContext.CtxKeyRequestID)
			if reqID == "" {
				reqID = uuid.NewString()
			}

			ctx := r.Context()
			ctx = pkgContext.SetRequestIDToContext(ctx, reqID)
			ctx = pkgContext.SetRequestIDToMetadata(ctx, reqID)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
