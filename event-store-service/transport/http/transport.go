package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	authInterceptor "github.com/ferza17/ecommerce-microservices-v2/event-store-service/interceptor/auth"
	loggerInterceptor "github.com/ferza17/ecommerce-microservices-v2/event-store-service/interceptor/logger"
	requestIdInterceptor "github.com/ferza17/ecommerce-microservices-v2/event-store-service/interceptor/requestid"
	telemetryInterceptor "github.com/ferza17/ecommerce-microservices-v2/event-store-service/interceptor/telemetry"
	pb "github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/gen/v1/event"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/presenter"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/logger"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/response"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/worker"
	"github.com/google/wire"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

type (
	Transport struct {
		address                 string
		port                    string
		workerPool              *worker.WorkerPool
		server                  *http.Server
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
		// Presenter
		eventPresenter presenter.IEventPresenter
	}
)

var Set = wire.NewSet(NewTransport)

func NewTransport(
	logger logger.IZapLogger,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	eventPresenter presenter.IEventPresenter,
) *Transport {
	return &Transport{
		address: config.Get().EventStoreServiceHttpHost,
		port:    config.Get().EventStoreServiceHttpPort,
		workerPool: worker.NewWorkerPool(
			fmt.Sprintf("HTTP SERVER ON %s:%s", config.Get().EventStoreServiceHttpHost, config.Get().EventStoreServiceHttpPort),
			2),
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
		eventPresenter:          eventPresenter,
	}
}

func (s *Transport) Serve(ctx context.Context) error {
	s.workerPool.Start()
	router := mux.NewRouter()
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

	if err := pb.RegisterEventStoreHandlerServer(ctx, gwMux, s.eventPresenter); err != nil {
		return fmt.Errorf("failed to register gRPC gateway handlers: %w", err)
	}

	// Health check endpoint
	router.HandleFunc("/v1/event-store/check", func(w http.ResponseWriter, r *http.Request) {
		response.WriteSuccessResponse(w, http.StatusOK, &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"status": &structpb.Value{
					Kind: &structpb.Value_StringValue{StringValue: "OK"},
				},
				"service": &structpb.Value{
					Kind: &structpb.Value_StringValue{StringValue: config.Get().EventStoreServiceServiceName},
				},
			},
		})
		return
	}).Methods("GET")

	router.Use(requestIdInterceptor.RequestIDHTTPMiddleware())
	router.Use(telemetryInterceptor.TelemetryHTTPMiddleware(s.telemetryInfrastructure))
	router.Use(loggerInterceptor.LoggerHTTPMiddleware(s.logger))
	router.Use(authInterceptor.AuthHTTPMiddleware(s.logger))
	router.PathPrefix("/v1/event-store").Handler(gwMux)

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
