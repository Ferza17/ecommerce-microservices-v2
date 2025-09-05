package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	loggerInterceptor "github.com/ferza17/ecommerce-microservices-v2/event-store-service/interceptor/logger"
	requestIdInterceptor "github.com/ferza17/ecommerce-microservices-v2/event-store-service/interceptor/requestid"
	telemetryInterceptor "github.com/ferza17/ecommerce-microservices-v2/event-store-service/interceptor/telemetry"
	pb "github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/gen/v1/event"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/presenter"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/logger"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/worker"
	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type (
	Transport struct {
		address                 string
		port                    string
		workerPool              *worker.WorkerPool
		server                  *grpc.Server
		logger                  logger.IZapLogger
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
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
		address: config.Get().EventStoreServiceRpcHost,
		port:    config.Get().EventStoreServiceRpcPort,
		workerPool: worker.NewWorkerPool(
			fmt.Sprintf("GRPC SERVER ON %s:%s", config.Get().EventStoreServiceRpcHost, config.Get().EventStoreServiceRpcPort),
			2),
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
		eventPresenter:          eventPresenter,
	}
}

func (s *Transport) Serve(ctx context.Context) error {
	s.workerPool.Start()
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.address, s.port))
	if err != nil {
		s.logger.Error("failed to listen", zap.Error(err))
		return err
	}
	s.server = grpc.NewServer([]grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			requestIdInterceptor.RequestIDRPCInterceptor(),
			telemetryInterceptor.TelemetryRPCInterceptor(s.telemetryInfrastructure),
			loggerInterceptor.LoggerRPCInterceptor(s.logger),
			//authInterceptor.AuthRPCUnaryInterceptor(s.logger),
		),
	}...)

	// TODO: Implement Server
	pb.RegisterEventStoreServer(s.server, s.eventPresenter)

	// Mark the service as healthy
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s.server, healthServer)
	healthServer.SetServingStatus(config.Get().EventStoreServiceServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

	reflection.Register(s.server)
	if err = s.server.Serve(listen); err != nil {
		s.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
		return err
	}

	<-ctx.Done()
	s.server.GracefulStop()
	s.workerPool.Stop()
	return nil
}
