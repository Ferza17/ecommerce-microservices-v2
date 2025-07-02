package http

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	authInterceptor "github.com/ferza17/ecommerce-microservices-v2/user-service/interceptor/auth"
	loggerInterceptor "github.com/ferza17/ecommerce-microservices-v2/user-service/interceptor/logger"
	metricInterceptor "github.com/ferza17/ecommerce-microservices-v2/user-service/interceptor/metric"
	requestIdInterceptor "github.com/ferza17/ecommerce-microservices-v2/user-service/interceptor/requestid"
	telemetryInterceptor "github.com/ferza17/ecommerce-microservices-v2/user-service/interceptor/telemetry"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	accessControlUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/usecase"
	authPresenter "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/presenter"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	userPresenter "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/presenter"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/metric"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/response"
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
)

type (
	Server struct {
		address                 string
		port                    string
		server                  *http.Server
		logger                  logger.IZapLogger
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		authPresenter           *authPresenter.AuthPresenter
		userPresenter           *userPresenter.UserPresenter

		// For Middleware
		accessControlUseCase accessControlUseCase.IAccessControlUseCase
		authUseCase          authUseCase.IAuthUseCase
	}
)

var Set = wire.NewSet(NewServer)

func NewServer(
	logger logger.IZapLogger,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	authPresenter *authPresenter.AuthPresenter,
	userPresenter *userPresenter.UserPresenter,
	accessControlUseCase accessControlUseCase.IAccessControlUseCase,
	authUseCase authUseCase.IAuthUseCase,
) *Server {
	return &Server{
		address:                 config.Get().UserServiceHttpHost,
		port:                    config.Get().UserServiceHttpPort,
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
		authPresenter:           authPresenter,
		userPresenter:           userPresenter,
		accessControlUseCase:    accessControlUseCase,
		authUseCase:             authUseCase,
	}
}

func (s *Server) Serve(ctx context.Context) error {
	// Create Gorilla mux router
	router := mux.NewRouter()

	// Create grpc-gateway mux for gRPC-HTTP gateway
	gwMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
		runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			switch key {
			case "Authorization",
				"Content-Type",
				"Accept",
				"traceparent",
				"tracestate",
				"baggage",
				pkgContext.CtxKeyRequestID,
				pkgContext.CtxKeyAuthorization:
				return key, true
			default:
				return "", false
			}
		}),
	)

	// For Matrics
	prometheus.MustRegister(pkgMetric.HttpRequests)
	prometheus.MustRegister(pkgMetric.HttpDuration)

	// Register gRPC-HTTP gateway handlers
	if err := pb.RegisterUserServiceHandlerServer(ctx, gwMux, s.userPresenter); err != nil {
		return fmt.Errorf("failed to register gRPC gateway handlers: %w", err)
	}

	if err := pb.RegisterAuthServiceHandlerServer(ctx, gwMux, s.authPresenter); err != nil {
		return fmt.Errorf("failed to register gRPC gateway handlers: %w", err)
	}

	// Mount the gRPC gateway with JWT middleware wrapping

	// Health check endpoint
	router.HandleFunc("/v1/user/check", func(w http.ResponseWriter, r *http.Request) {
		response.WriteSuccessResponse(w, http.StatusOK, []byte(`{"status": "ok", "service": "product-service"}`))
		return
	}).Methods("GET")

	// Route for api documentation
	swaggerJSONPath := "./docs/v1/user/service.swagger.json"
	router.HandleFunc("/docs/v1/user/service.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, swaggerJSONPath)
	}).Methods("GET")

	router.PathPrefix("/v1/user/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/docs/v1/user/service.swagger.json"), // 👈 URL must match your exposed JSON
	))

	router.Use(requestIdInterceptor.RequestIDHTTPMiddleware())
	router.Use(telemetryInterceptor.TelemetryHTTPMiddleware(s.telemetryInfrastructure))
	router.Use(metricInterceptor.MetricHTTPMiddleware())
	router.Use(loggerInterceptor.LoggerHTTPMiddleware(s.logger))
	router.Use(authInterceptor.AuthHTTPMiddleware(s.logger, s.accessControlUseCase, s.authUseCase))

	router.PathPrefix("/v1/user").Handler(gwMux)

	// Create an HTTP server instance
	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.address, s.port),
		Handler: router,
	}

	log.Printf("Starting HTTP server on %s:%s", s.address, s.port)

	// ListenAndServe returns http.ErrServerClosed when gracefully shutdown
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("HTTP server failed to start: %w", err)
	}

	return nil
}
