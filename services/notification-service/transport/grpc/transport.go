package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/alitto/pond/v2"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type (
	Transport struct {
		address    string
		port       string
		grpcServer *grpc.Server
		logger     logger.IZapLogger
	}
)

var Set = wire.NewSet(NewTransport)

func NewTransport(
	logger logger.IZapLogger) *Transport {
	return &Transport{
		address:    config.Get().ConfigServiceNotification.RpcHost,
		port:       config.Get().ConfigServiceNotification.RpcPort,
		grpcServer: grpc.NewServer(),
		logger:     logger,
	}
}

func (s *Transport) Serve(ctx context.Context) error {
	pool := pond.NewPool(10, pond.WithContext(ctx), pond.WithQueueSize(1000), pond.WithNonBlocking(true))
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
	task := pool.SubmitErr(func() error {
		if err = s.grpcServer.Serve(listen); err != nil {
			s.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
			return err
		}
		return nil
	})
	if err = task.Wait(); err != nil {
		s.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
		return err
	}
	return nil
}

func (s *Transport) GracefulStop() {
	s.grpcServer.GracefulStop()
}
