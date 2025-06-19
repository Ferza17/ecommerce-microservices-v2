package grpc

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type (
	IGrpcServer interface {
		Serve()
		GracefulStop()
	}

	GrpcServer struct {
		address string
		port    string

		grpcServer *grpc.Server
		logger     pkg.IZapLogger
	}
)

func NewGrpcServer(logger pkg.IZapLogger) IGrpcServer {
	return &GrpcServer{
		address:    config.Get().RpcHost,
		port:       config.Get().RpcPort,
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
	// Mark the service as healthy
	healthServer.SetServingStatus(config.Get().ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

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
