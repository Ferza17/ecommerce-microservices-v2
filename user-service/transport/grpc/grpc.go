package grpc

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	authInterceptor "github.com/ferza17/ecommerce-microservices-v2/user-service/interceptor/auth"
	loggerInterceptor "github.com/ferza17/ecommerce-microservices-v2/user-service/interceptor/logger"
	metricInterceptor "github.com/ferza17/ecommerce-microservices-v2/user-service/interceptor/metric"
	requestIdInterceptor "github.com/ferza17/ecommerce-microservices-v2/user-service/interceptor/requestid"
	telemetryInterceptor "github.com/ferza17/ecommerce-microservices-v2/user-service/interceptor/telemetry"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	accessControlUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/accessControl/usecase"
	authPresenter "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/presenter"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	userPresenter "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/presenter"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	pkgWorker "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/worker"
	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
)

type (
	Server struct {
		address                 string
		port                    string
		workerPool              *pkgWorker.WorkerPool
		grpcServer              *grpc.Server
		logger                  logger.IZapLogger
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		authPresenter           *authPresenter.AuthPresenter
		userPresenter           *userPresenter.UserPresenter

		// For Middleware
		accessControlUseCase accessControlUseCase.IAccessControlUseCase
		authUseCase          authUseCase.IAuthUseCase
	}
)

var Set = wire.NewSet(NewServer)

func NewServer(
	logger logger.IZapLogger,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	authPresenter *authPresenter.AuthPresenter,
	userPresenter *userPresenter.UserPresenter,
	accessControlUseCase accessControlUseCase.IAccessControlUseCase,
	authUseCase authUseCase.IAuthUseCase,
) *Server {
	//TODO: Add workers from consul config
	return &Server{
		address: config.Get().UserServiceRpcHost,
		port:    config.Get().UserServiceRpcPort,
		workerPool: pkgWorker.NewWorkerPool(
			fmt.Sprintf("GRPC SERVER ON %s:%s", config.Get().UserServiceRpcHost, config.Get().UserServiceRpcPort),
			2,
		),
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
		authPresenter:           authPresenter,
		userPresenter:           userPresenter,
		accessControlUseCase:    accessControlUseCase,
		authUseCase:             authUseCase,
	}
}

func (srv *Server) Serve(ctx context.Context) error {
	srv.workerPool.Start()

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", srv.address, srv.port))
	if err != nil {
		srv.logger.Error("failed to listen", zap.Error(err))
		return err
	}
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			requestIdInterceptor.RequestIDRPCInterceptor(),
			metricInterceptor.MetricUnaryInterceptor(),
			telemetryInterceptor.TelemetryRPCInterceptor(srv.telemetryInfrastructure),
			loggerInterceptor.LoggerRPCInterceptor(srv.logger),
			authInterceptor.AuthRPCUnaryInterceptor(srv.logger, srv.accessControlUseCase, srv.authUseCase),
		),
	}

	srv.grpcServer = grpc.NewServer(opts...)

	// Register all services first
	userRpc.RegisterUserServiceServer(
		srv.grpcServer,
		srv.userPresenter,
	)

	userRpc.RegisterAuthServiceServer(
		srv.grpcServer,
		srv.authPresenter,
	)

	// Register health service
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(srv.grpcServer, healthServer)
	healthServer.SetServingStatus(config.Get().UserServiceServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

	// IMPORTANT: Register reflection AFTER all services are registered
	if config.Get().Env != enum.CONFIG_ENV_PROD {
		reflection.Register(srv.grpcServer)
	}

	if err = srv.grpcServer.Serve(listen); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
		return err
	}

	<-ctx.Done()
	srv.grpcServer.GracefulStop()
	srv.workerPool.Stop()
	return nil
}
