package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	pkgWorker "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/worker"
	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type (
	GrpcServer struct {
		address    string
		port       string
		workerPool *pkgWorker.WorkerPool
		grpcServer *grpc.Server
		logger     logger.IZapLogger
	}
)

var Set = wire.NewSet(NewGrpcServer)

func NewGrpcServer(
	logger logger.IZapLogger) *GrpcServer {
	return &GrpcServer{
		workerPool: pkgWorker.NewWorkerPool(
			fmt.Sprintf("GRPC SERVER ON %s:%s", config.Get().ConfigServiceNotification.RpcHost, config.Get().ConfigServiceNotification.RpcPort),
			1,
		),
		address:    config.Get().ConfigServiceNotification.RpcHost,
		port:       config.Get().ConfigServiceNotification.RpcPort,
		grpcServer: grpc.NewServer(),
		logger:     logger,
	}
}

func (s *GrpcServer) Serve(ctx context.Context) error {
	s.workerPool.Start()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		s.logger.Error(fmt.Sprintf("Err Listen : %v", err))
	}

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s.grpcServer, healthServer)
	healthServer.SetServingStatus(config.Get().ConfigServiceNotification.ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

	log.Printf("Starting gRPC server on %s:%s", s.address, s.port)
	// Enable Reflection to Evans grpc client
	reflection.Register(s.grpcServer)
	if err = s.grpcServer.Serve(listen); err != nil {
		s.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
	}

	<-ctx.Done()
	s.grpcServer.GracefulStop()
	s.workerPool.Stop()
	return nil
}

func (s *GrpcServer) GracefulStop() {
	s.grpcServer.GracefulStop()
}
