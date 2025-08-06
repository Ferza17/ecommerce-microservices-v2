package http

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/temporal"
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
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/response"
	pkgWorker "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/worker"
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"net/http"
)

type (
	Server struct {
		address                 string
		port                    string
		workerPool              *pkgWorker.WorkerPool
		server                  *http.Server
		logger                  logger.IZapLogger
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		temporal                temporal.ITemporalInfrastructure
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
	temporal temporal.ITemporalInfrastructure,
	authPresenter *authPresenter.AuthPresenter,
	userPresenter *userPresenter.UserPresenter,
	accessControlUseCase accessControlUseCase.IAccessControlUseCase,
	authUseCase authUseCase.IAuthUseCase,
) *Server {
	//TODO: Add workers from consul config
	return &Server{
		address: config.Get().UserServiceHttpHost,
		port:    config.Get().UserServiceHttpPort,
		workerPool: pkgWorker.NewWorkerPool(
			fmt.Sprintf("HTTP SERVER ON %s:%s", config.Get().UserServiceHttpHost, config.Get().UserServiceHttpPort),
			2,
		),
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
		temporal:                temporal,
		authPresenter:           authPresenter,
		userPresenter:           userPresenter,
		accessControlUseCase:    accessControlUseCase,
		authUseCase:             authUseCase,
	}
}

func (s *Server) Serve(ctx context.Context) error {
	s.workerPool.Start()
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
				pkgContext.ContextKeyTracerparent,
				"tracestate",
				"baggage",
				pkgContext.CtxKeyRequestID,
				pkgContext.CtxKeyAuthorization:
				return key, true
			default:
				return "", false
			}
		}),
		runtime.WithErrorHandler(response.CustomErrorHandler),
		// Other useful options
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{}),
	)

	// Register gRPC-HTTP gateway handlers
	if err := pb.RegisterUserServiceHandlerServer(ctx, gwMux, s.userPresenter); err != nil {
		return fmt.Errorf("failed to register gRPC gateway handlers: %w", err)
	}

	if err := pb.RegisterAuthServiceHandlerServer(ctx, gwMux, s.authPresenter); err != nil {
		return fmt.Errorf("failed to register gRPC gateway handlers: %w", err)
	}

	// Health check endpoint
	router.HandleFunc("/v1/user/check", func(w http.ResponseWriter, r *http.Request) {
		response.WriteSuccessResponse(w, http.StatusOK, &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"status": &structpb.Value{
					Kind: &structpb.Value_StringValue{StringValue: "OK"},
				},
				"service": &structpb.Value{
					Kind: &structpb.Value_StringValue{StringValue: config.Get().UserServiceServiceName},
				},
			},
		})
		return
	}).Methods("GET")

	// Route for api documentation
	swaggerJSONPath := "./docs/v1/user/service.swagger.json"
	router.HandleFunc("/docs/v1/user/service.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, swaggerJSONPath)
	}).Methods("GET")

	router.PathPrefix("/v1/user/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/docs/v1/user/service.swagger.json"),
	))

	router.Use(requestIdInterceptor.RequestIDHTTPMiddleware())
	router.Use(telemetryInterceptor.TelemetryHTTPMiddleware(s.telemetryInfrastructure))
	router.Use(metricInterceptor.MetricHTTPMiddleware(s.telemetryInfrastructure))
	router.Use(loggerInterceptor.LoggerHTTPMiddleware(s.logger))
	router.Use(authInterceptor.AuthHTTPMiddleware(s.logger, s.accessControlUseCase, s.authUseCase))

	router.PathPrefix("/v1/user").Handler(gwMux)
	router.PathPrefix("/v1/user").
		Subrouter()

	// Setup Temporal, FAIL FORGET !
	go func() {
		if err := s.temporal.Start(); err != nil {
			s.logger.Error("failed to start temporal server", zap.Error(err))
			return
		}
	}()

	// Create an HTTP server instance
	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.address, s.port),
		Handler: router,
	}

	// ListenAndServe returns http.ErrServerClosed when gracefully shutdown
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("HTTP server failed to start: %w", err)
	}

	<-ctx.Done()
	s.workerPool.Stop()
	return nil
}
