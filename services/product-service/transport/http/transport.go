package http

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	userService "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/service/user"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	authInterceptor "github.com/ferza17/ecommerce-microservices-v2/product-service/interceptor/auth"
	loggerInterceptor "github.com/ferza17/ecommerce-microservices-v2/product-service/interceptor/logger"
	requestIdInterceptor "github.com/ferza17/ecommerce-microservices-v2/product-service/interceptor/requestid"
	telemetryInterceptor "github.com/ferza17/ecommerce-microservices-v2/product-service/interceptor/telemetry"
	pb "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/presenter"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/response"
	pkgWorker "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/worker"
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
)

type (
	Transport struct {
		address    string
		port       string
		server     *http.Server
		workerPool *pkgWorker.WorkerPool

		logger                  logger.IZapLogger
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		productPresenter        *presenter.ProductPresenter

		// For Middleware
		userService userService.IUserService
	}
)

var Set = wire.NewSet(NewServer)

func NewServer(
	logger logger.IZapLogger,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	productPresenter *presenter.ProductPresenter,
	userService userService.IUserService,
) *Transport {
	return &Transport{
		workerPool: pkgWorker.NewWorkerPool(
			fmt.Sprintf("HTTP SERVER ON %s:%s", config.Get().ConfigServiceProduct.HttpHost, config.Get().ConfigServiceProduct.HttpPort),
			1,
		),
		address:                 config.Get().ConfigServiceProduct.HttpHost,
		port:                    config.Get().ConfigServiceProduct.HttpPort,
		productPresenter:        productPresenter,
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
		userService:             userService,
	}
}

func (s *Transport) Serve(ctx context.Context) error {
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
	if err := pb.RegisterProductServiceHandlerServer(ctx, gwMux, s.productPresenter); err != nil {
		return fmt.Errorf("failed to register gRPC gateway handlers: %w", err)
	}

	// Health check endpoint
	router.HandleFunc("/v1/product/check", func(w http.ResponseWriter, r *http.Request) {
		response.WriteSuccessResponse(w, http.StatusOK, []byte(`{"status": "ok", "service": "product-service"}`))
		return
	}).Methods("GET")

	swaggerJSONPath := "./docs/v1/product/service.swagger.json"

	router.HandleFunc("/docs/v1/product/service.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, swaggerJSONPath)
	}).Methods("GET")

	router.PathPrefix("/v1/product/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/docs/v1/product/service.swagger.json"), // 👈 URL must match your exposed JSON
	))

	router.Use(requestIdInterceptor.RequestIDHTTPMiddleware())
	router.Use(telemetryInterceptor.TelemetryHTTPMiddleware(s.telemetryInfrastructure))
	router.Use(loggerInterceptor.LoggerHTTPMiddleware(s.logger))
	router.Use(authInterceptor.AuthHTTPMiddleware(s.logger))

	router.PathPrefix("/v1/product").Handler(gwMux)

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

	<-ctx.Done()
	s.workerPool.Stop()
	return nil
}

func (s *Transport) ServeHttpPrometheusMetricCollector() error {
	handler := http.NewServeMux()
	handler.Handle("/v1/product/metrics", promhttp.Handler())
	log.Printf("Starting HTTP Metric Server on %s:%s", config.Get().ConfigServiceProduct.HttpHost, config.Get().ConfigServiceProduct.HttpPort)
	if err := http.ListenAndServe(
		fmt.Sprintf("%s:%s", config.Get().ConfigServiceProduct.HttpHost, config.Get().ConfigServiceProduct.MetricHttpPort),
		handler,
	); err != nil {
		return err
	}
	return nil
}
