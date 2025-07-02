package grpc

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type (
	GrpcServer struct {
		address string
		port    string

		grpcServer *grpc.Server
		logger     logger.IZapLogger
	}
)

var Set = wire.NewSet(NewGrpcServer)

func NewGrpcServer(logger logger.IZapLogger) *GrpcServer {
	return &GrpcServer{
		address:    config.Get().NotificationServiceRpcHost,
		port:       config.Get().NotificationServiceRpcPort,
		grpcServer: grpc.NewServer(),
		logger:     logger,
	}
}

func (s *GrpcServer) Serve() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		s.logger.Error(fmt.Sprintf("Err Listen : %v", err))
	}

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s.grpcServer, healthServer)
	healthServer.SetServingStatus(config.Get().NotificationServiceServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

	log.Printf("Starting gRPC server on %s:%s", s.address, s.port)
	// Enable Reflection to Evans grpc client
	reflection.Register(s.grpcServer)
	if err = s.grpcServer.Serve(listen); err != nil {
		s.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
	}
	return
}

func (s *GrpcServer) GracefulStop() {
	s.grpcServer.GracefulStop()
}
