package grpc

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	userService "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/service/user"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	authInterceptor "github.com/ferza17/ecommerce-microservices-v2/product-service/interceptor/auth"
	loggerInterceptor "github.com/ferza17/ecommerce-microservices-v2/product-service/interceptor/logger"
	requestIdInterceptor "github.com/ferza17/ecommerce-microservices-v2/product-service/interceptor/requestid"
	telemetryInterceptor "github.com/ferza17/ecommerce-microservices-v2/product-service/interceptor/telemetry"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/presenter"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type (
	GrpcTransport struct {
		address    string
		port       string
		grpcServer *grpc.Server

		logger                  logger.IZapLogger
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		productPresenter        *presenter.ProductPresenter

		// For Middleware
		userService userService.IUserService
	}
)

var Set = wire.NewSet(NewServer)

func NewServer(
	logger logger.IZapLogger,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	productPresenter *presenter.ProductPresenter,
	userService userService.IUserService,
) *GrpcTransport {
	return &GrpcTransport{
		address:                 config.Get().ProductServiceRpcHost,
		port:                    config.Get().ProductServiceRpcPort,
		productPresenter:        productPresenter,
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
		userService:             userService,
	}
}

func (srv *GrpcTransport) Serve() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", srv.address, srv.port))
	if err != nil {
		log.Fatalln(err)
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			requestIdInterceptor.RequestIDRPCInterceptor(),
			telemetryInterceptor.TelemetryRPCInterceptor(srv.telemetryInfrastructure),
			loggerInterceptor.LoggerRPCInterceptor(srv.logger),
			authInterceptor.AuthRPCUnaryInterceptor(srv.logger, srv.userService),
		),
	}
	srv.grpcServer = grpc.NewServer(opts...)

	productRpc.RegisterProductServiceServer(
		srv.grpcServer,
		srv.productPresenter,
	)

	// Mark the service as healthy
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(srv.grpcServer, healthServer)
	healthServer.SetServingStatus(config.Get().UserServiceServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

	log.Printf("Starting gRPC server on %s:%s", srv.address, srv.port)

	// Enable Reflection to Evans grpc client
	reflection.Register(srv.grpcServer)
	if err = srv.grpcServer.Serve(listen); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
	}
}

func (srv *GrpcTransport) GracefulStop() {
	srv.grpcServer.GracefulStop()
}
