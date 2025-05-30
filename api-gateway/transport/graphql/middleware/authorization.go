package middleware

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/telemetry"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/user/v1"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/auth/service"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strings"
)

func Authorization(svc service.IAuthService, tele telemetry.ITelemetryInfrastructure) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
			ctx, span := tele.Tracer(ctx, "Middleware.Authentication")
			parts := strings.Split(r.Header.Get("Authorization"), "Bearer ")
			token := ""
			if len(parts) == 2 && strings.ToLower(parts[1]) != "" {
				token = parts[1]
			} else {
				span.RecordError(fmt.Errorf("invalid authorization header"))
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
				enum.XRequestIDHeader.String(): ctx.Value(enum.XRequestIDHeader.String()).(string),
			}))
			// access user service find user by token
			user, err := svc.FindUserByToken(ctx, ctx.Value(enum.XRequestIDHeader.String()).(string), &userRpc.FindUserByTokenRequest{
				Token: token,
			})
			if user.Id == "" {
				span.RecordError(err)
				http.Error(w, "invalid access token", http.StatusUnauthorized)
				return
			}
			ctx = context.WithValue(ctx, enum.ContextKeyUserID, user.Id)
			span.End()
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
