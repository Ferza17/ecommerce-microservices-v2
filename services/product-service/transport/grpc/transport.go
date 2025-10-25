package grpc

import (
	"context"
	"fmt"

	"github.com/alitto/pond/v2"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	userService "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/service/user"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	loggerInterceptor "github.com/ferza17/ecommerce-microservices-v2/product-service/interceptor/logger"
	requestIdInterceptor "github.com/ferza17/ecommerce-microservices-v2/product-service/interceptor/requestid"
	telemetryInterceptor "github.com/ferza17/ecommerce-microservices-v2/product-service/interceptor/telemetry"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"log"
	"net"

	"github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/presenter"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type (
	Transport struct {
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

var Set = wire.NewSet(NewTransport)

func NewTransport(
	logger logger.IZapLogger,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	productPresenter *presenter.ProductPresenter,
	userService userService.IUserService,
) *Transport {
	return &Transport{
		address:                 config.Get().ConfigServiceProduct.RpcHost,
		port:                    config.Get().ConfigServiceProduct.RpcPort,
		productPresenter:        productPresenter,
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
		userService:             userService,
	}
}

func (srv *Transport) Serve(ctx context.Context) error {
	pool := pond.NewPool(10, pond.WithContext(ctx), pond.WithQueueSize(1000), pond.WithNonBlocking(true))
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", srv.address, srv.port))
	if err != nil {
		log.Fatalln(err)
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			requestIdInterceptor.RequestIDRPCInterceptor(),
			telemetryInterceptor.TelemetryRPCInterceptor(srv.telemetryInfrastructure),
			loggerInterceptor.LoggerRPCInterceptor(srv.logger),
			//authInterceptor.AuthRPCUnaryInterceptor(srv.logger),
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
	healthServer.SetServingStatus(config.Get().ConfigServiceProduct.ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

	// Enable Reflection to Evans grpc client
	reflection.Register(srv.grpcServer)
	task := pool.SubmitErr(func() error {
		if err = srv.grpcServer.Serve(listen); err != nil {
			srv.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
			return err
		}
		return nil
	})
	if err = task.Wait(); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
		return err
	}
	srv.grpcServer.GracefulStop()
	return nil
}

func (srv *Transport) GracefulStop() {
	srv.grpcServer.GracefulStop()
}
