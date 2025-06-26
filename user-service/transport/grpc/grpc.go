package grpc

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	authPresenter "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/presenter"
	userPresenter "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/presenter"
	"github.com/google/wire"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type (
	Server struct {
		address                 string
		port                    string
		grpcServer              *grpc.Server
		logger                  logger.IZapLogger
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		authPresenter           *authPresenter.AuthPresenter
		userPresenter           *userPresenter.UserPresenter
	}
)

var Set = wire.NewSet(NewServer)

func NewServer(
	logger logger.IZapLogger,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	authPresenter *authPresenter.AuthPresenter,
	userPresenter *userPresenter.UserPresenter,
) *Server {
	return &Server{
		address:                 config.Get().RpcHost,
		port:                    config.Get().RpcPort,
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
		authPresenter:           authPresenter,
		userPresenter:           userPresenter,
	}
}

func (srv *Server) Serve() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", srv.address, srv.port))
	if err != nil {
		log.Fatalln(err)
	}
	opts := []grpc.ServerOption{
		//grpc.ChainUnaryInterceptor(
		//grpcMiddleware.ChainUnaryServer(
		//	otgrpc.OpenTracingServerInterceptor(srv.tracer),
		//),
		//),
	}
	srv.grpcServer = grpc.NewServer(opts...)
	userRpc.RegisterUserServiceServer(
		srv.grpcServer,
		srv.userPresenter,
	)

	userRpc.RegisterAuthServiceServer(
		srv.grpcServer,
		srv.authPresenter,
	)

	// Mark the service as healthy
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(srv.grpcServer, healthServer)
	healthServer.SetServingStatus(config.Get().ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

	// Enable Reflection to Evans grpc client
	reflection.Register(srv.grpcServer)
	if err = srv.grpcServer.Serve(listen); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
	}
}

func (srv *Server) GracefulStop() {
	srv.grpcServer.GracefulStop()
}
