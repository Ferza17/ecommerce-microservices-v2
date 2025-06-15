package grpc

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
	paymentPresenter "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/presenter"
	paymentProviderPresenter "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/presenter"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
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

		paymentPresenter         paymentPresenter.IPaymentPresenter
		paymentProviderPresenter paymentProviderPresenter.IPaymentProviderPresenter

		grpcServer *grpc.Server
		logger     logger.IZapLogger
	}
)

// NewGrpcServer creates and returns a new instance of GrpcServer with all dependencies.
func NewGrpcServer(
	logger logger.IZapLogger,
	paymentPresenter paymentPresenter.IPaymentPresenter,
	paymentProviderPresenter paymentProviderPresenter.IPaymentProviderPresenter,
	options ...grpc.ServerOption,
) IGrpcServer {
	grpcServer := grpc.NewServer(options...)

	host := config.Get().RpcHost
	_ = host
	return &GrpcServer{
		address:                  config.Get().RpcHost,
		port:                     config.Get().RpcPort,
		paymentPresenter:         paymentPresenter,
		paymentProviderPresenter: paymentProviderPresenter,
		grpcServer:               grpcServer,
		logger:                   logger,
	}
}

// Set is a Wire provider set for GrpcServer dependencies.
var Set = wire.NewSet(
	NewGrpcServer,
	ProvideGrpcServerOptions,
)

func ProvideGrpcServerOptions() []grpc.ServerOption {
	// Add any default server options you need here
	return []grpc.ServerOption{
		// Example options:
		// grpc.MaxRecvMsgSize(1024 * 1024 * 4), // 4MB
		// grpc.MaxSendMsgSize(1024 * 1024 * 4), // 4MB
	}
}

func (s *GrpcServer) Serve() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.address, s.port))
	if err != nil {
		s.logger.Error(fmt.Sprintf("Err Listen : %v", err))
	}

	paymentRpc.RegisterPaymentServiceServer(s.grpcServer, s.paymentPresenter)
	paymentRpc.RegisterPaymentProviderServiceServer(s.grpcServer, s.paymentProviderPresenter)

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
