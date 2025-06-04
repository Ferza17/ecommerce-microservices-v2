package grpc

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
	paymentPresenter "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/presenter"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
		logger     logger.IZapLogger
	}
)

// NewGrpcServer creates and returns a new instance of GrpcServer.
func NewGrpcServer(logger logger.IZapLogger, options ...grpc.ServerOption) IGrpcServer {
	grpcServer := grpc.NewServer(options...) // Initialize gRPC server with any provided options
	return &GrpcServer{
		address:    config.Get().RpcHost,
		port:       config.Get().RpcPort,
		grpcServer: grpcServer,
		logger:     logger,
	}
}

func (s *GrpcServer) Serve() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.address, s.port))
	if err != nil {
		s.logger.Error(fmt.Sprintf("Err Listen : %v", err))
	}

	paymentRpc.RegisterPaymentServiceServer(
		s.grpcServer,
		paymentPresenter.ProvidePaymentPresenter(),
	)

	// Enable Reflection to Evans grpc client
	reflection.Register(s.grpcServer)
	if err = s.grpcServer.Serve(listen); err != nil {
		s.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
	}

	s.logger.Info(fmt.Sprintf("Server listen at port %s", s.port))

	return
}

func (s *GrpcServer) GracefulStop() {
	s.grpcServer.GracefulStop()
}
