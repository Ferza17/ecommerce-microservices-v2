package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/alitto/pond/v2"
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
	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type (
	Transport struct {
		address                  string
		port                     string
		paymentPresenter         paymentPresenter.IPaymentPresenter
		paymentProviderPresenter paymentProviderPresenter.IPaymentProviderPresenter

		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		userService             userService.IUserService
		shippingService         shipping.IShippingService

		grpcServer *grpc.Server
		logger     logger.IZapLogger
	}
)

func NewTransport(
	logger logger.IZapLogger,
	paymentPresenter paymentPresenter.IPaymentPresenter,
	paymentProviderPresenter paymentProviderPresenter.IPaymentProviderPresenter,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	userService userService.IUserService,
) *Transport {
	return &Transport{
		address:                  config.Get().ConfigServicePayment.RpcHost,
		port:                     config.Get().ConfigServicePayment.RpcPort,
		paymentPresenter:         paymentPresenter,
		paymentProviderPresenter: paymentProviderPresenter,
		telemetryInfrastructure:  telemetryInfrastructure,
		userService:              userService,
		logger:                   logger,
	}
}

// Set is a Wire provider set for GrpcServer dependencies.
var Set = wire.NewSet(
	NewTransport,
)

func (s *Transport) Serve(ctx context.Context) error {
	pool := pond.NewPool(10, pond.WithContext(ctx), pond.WithQueueSize(1000), pond.WithNonBlocking(true))
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
	healthServer.SetServingStatus(config.Get().ConfigServicePayment.ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

	// IMPORTANT: Register reflection AFTER all services are registered
	//if config.Get().Env != enum.CONFIG_ENV_PROD {
	reflection.Register(s.grpcServer)
	//}

	if err = s.grpcServer.Serve(listen); err != nil {
		s.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
	}

	task := pool.SubmitErr(func() error {
		if err = s.grpcServer.Serve(listen); err != nil {
			s.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
		}
		return nil
	})
	if err = task.Wait(); err != nil {
		s.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
		return err
	}
	s.grpcServer.GracefulStop()
	return nil
}

func (s *Transport) GracefulStop() {
	s.grpcServer.GracefulStop()
}
