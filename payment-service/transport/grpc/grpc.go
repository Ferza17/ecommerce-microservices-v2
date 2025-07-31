package grpc

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/shipping"
	userService "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/user"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	authInterceptor "github.com/ferza17/ecommerce-microservices-v2/payment-service/interceptor/auth"
	loggerInterceptor "github.com/ferza17/ecommerce-microservices-v2/payment-service/interceptor/logger"
	requestIdInterceptor "github.com/ferza17/ecommerce-microservices-v2/payment-service/interceptor/requestid"
	telemetryInterceptor "github.com/ferza17/ecommerce-microservices-v2/payment-service/interceptor/telemetry"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	paymentPresenter "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/presenter"
	paymentProviderPresenter "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/presenter"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	pkgWorker "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/worker"
	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
)

type (
	IGrpcServer interface {
		Serve(ctx context.Context) error
		GracefulStop()
	}

	GrpcServer struct {
		address                  string
		port                     string
		workerPool               *pkgWorker.WorkerPool
		paymentPresenter         paymentPresenter.IPaymentPresenter
		paymentProviderPresenter paymentProviderPresenter.IPaymentProviderPresenter

		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		userService             userService.IUserService
		shippingService         shipping.IShippingService

		grpcServer *grpc.Server
		logger     logger.IZapLogger
	}
)

// NewGrpcServer creates and returns a new instance of GrpcServer with all dependencies.
func NewGrpcServer(
	logger logger.IZapLogger,
	paymentPresenter paymentPresenter.IPaymentPresenter,
	paymentProviderPresenter paymentProviderPresenter.IPaymentProviderPresenter,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	userService userService.IUserService,
) IGrpcServer {
	return &GrpcServer{
		address: config.Get().PaymentServiceRpcHost,
		port:    config.Get().PaymentServiceRpcPort,
		workerPool: pkgWorker.NewWorkerPool(
			fmt.Sprintf("GRPC SERVER ON %s:%s", config.Get().PaymentServiceRpcHost, config.Get().PaymentServiceRpcPort),
			2,
		),
		paymentPresenter:         paymentPresenter,
		paymentProviderPresenter: paymentProviderPresenter,
		telemetryInfrastructure:  telemetryInfrastructure,
		userService:              userService,
		logger:                   logger,
	}
}

// Set is a Wire provider set for GrpcServer dependencies.
var Set = wire.NewSet(
	NewGrpcServer,
)

func (s *GrpcServer) Serve(ctx context.Context) error {
	s.workerPool.Start()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		s.logger.Error(fmt.Sprintf("Err Listen : %v", err))
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			requestIdInterceptor.RequestIDRPCInterceptor(),
			telemetryInterceptor.TelemetryRPCInterceptor(s.telemetryInfrastructure),
			loggerInterceptor.LoggerRPCInterceptor(s.logger),
			authInterceptor.AuthRPCUnaryInterceptor(s.logger, s.userService),
		),
	}
	s.grpcServer = grpc.NewServer(opts...)

	paymentRpc.RegisterPaymentServiceServer(s.grpcServer, s.paymentPresenter)
	paymentRpc.RegisterPaymentProviderServiceServer(s.grpcServer, s.paymentProviderPresenter)

	// Mark the service as healthy
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s.grpcServer, healthServer)
	healthServer.SetServingStatus(config.Get().PaymentServiceServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

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
