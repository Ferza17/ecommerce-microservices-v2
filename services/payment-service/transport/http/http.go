package http

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	userService "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/user"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	authInterceptor "github.com/ferza17/ecommerce-microservices-v2/payment-service/interceptor/auth"
	loggerInterceptor "github.com/ferza17/ecommerce-microservices-v2/payment-service/interceptor/logger"
	requestIdInterceptor "github.com/ferza17/ecommerce-microservices-v2/payment-service/interceptor/requestid"
	telemetryInterceptor "github.com/ferza17/ecommerce-microservices-v2/payment-service/interceptor/telemetry"
	pb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	paymentPresenter "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/presenter"
	paymentProviderPresenter "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/presenter"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/response"
	pkgWorker "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/worker"
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
)

type HttpServer struct {
	address                  string
	port                     string
	workerPool               *pkgWorker.WorkerPool
	paymentPresenter         paymentPresenter.IPaymentPresenter
	paymentProviderPresenter paymentProviderPresenter.IPaymentProviderPresenter

	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	userService             userService.IUserService

	server *http.Server
	logger logger.IZapLogger
}

var Set = wire.NewSet(NewHttpServer)

// NewHttpServer creates and returns a new instance of HttpServer with all dependencies.
func NewHttpServer(
	logger logger.IZapLogger,
	paymentPresenter paymentPresenter.IPaymentPresenter,
	paymentProviderPresenter paymentProviderPresenter.IPaymentProviderPresenter,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	userService userService.IUserService,
) *HttpServer {
	return &HttpServer{
		address: config.Get().ConfigServicePayment.HttpHost,
		port:    config.Get().ConfigServicePayment.HttpPort,
		workerPool: pkgWorker.NewWorkerPool(
			fmt.Sprintf("HTTP SERVER ON %s:%s", config.Get().ConfigServicePayment.HttpHost, config.Get().ConfigServicePayment.HttpPort),
			2,
		),
		paymentPresenter:         paymentPresenter,
		paymentProviderPresenter: paymentProviderPresenter,
		logger:                   logger,
		telemetryInfrastructure:  telemetryInfrastructure,
		userService:              userService,
	}
}

func (s *HttpServer) Serve(ctx context.Context) error {
	s.workerPool.Start()

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

	// Register Payment Provider
	if err := pb.RegisterPaymentProviderServiceHandlerServer(ctx, gwMux, s.paymentProviderPresenter); err != nil {
		s.logger.Error(fmt.Sprintf("failed to register payment provider service handler server : %s", zap.Error(err).String))
		return err
	}

	// Register Payment
	if err := pb.RegisterPaymentServiceHandlerServer(ctx, gwMux, s.paymentPresenter); err != nil {
		s.logger.Error(fmt.Sprintf("failed to register payment service handler server : %s", zap.Error(err).String))
		return err
	}

	// Health check endpoint
	router.HandleFunc("/v1/payment/check", func(w http.ResponseWriter, r *http.Request) {
		response.WriteSuccessResponse(w, http.StatusOK, []byte(`{"status": "ok", "service": "product-service"}`))
		return
	}).Methods("GET")

	swaggerJSONPath := "./docs/v1/payment/service.swagger.json"

	router.HandleFunc("/docs/v1/payment/service.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, swaggerJSONPath)
	}).Methods("GET")

	router.PathPrefix("/v1/payment/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/docs/v1/payment/service.swagger.json"), // 👈 URL must match your exposed JSON
	))

	router.Use(requestIdInterceptor.RequestIDHTTPMiddleware())
	router.Use(telemetryInterceptor.TelemetryHTTPMiddleware(s.telemetryInfrastructure))
	router.Use(loggerInterceptor.LoggerHTTPMiddleware(s.logger))
	router.Use(authInterceptor.AuthHTTPMiddleware(s.logger, s.userService))

	router.PathPrefix("/v1/payment").Handler(gwMux)

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
