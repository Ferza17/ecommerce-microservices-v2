package grpc

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/user/v1"

	authPresenter "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/presenter"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/presenter"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
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
		logger                  pkg.IZapLogger
		userUseCase             usecase.IUserUseCase
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		authUseCase             authUseCase.IAuthUseCase
	}
)

func NewServer(dependency *bootstrap.Bootstrap) *Server {
	return &Server{
		address:                 config.Get().RpcHost,
		port:                    config.Get().RpcPort,
		logger:                  dependency.Logger,
		userUseCase:             dependency.UserUseCase,
		telemetryInfrastructure: dependency.TelemetryInfrastructure,
		authUseCase:             dependency.AuthUseCase,
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
		presenter.NewUserPresenter(srv.userUseCase, srv.telemetryInfrastructure, srv.logger),
	)

	userRpc.RegisterAuthServiceServer(
		srv.grpcServer,
		authPresenter.NewAuthPresenter(srv.authUseCase, srv.telemetryInfrastructure, srv.logger),
	)

	// Enable Reflection to Evans grpc client
	reflection.Register(srv.grpcServer)
	if err = srv.grpcServer.Serve(listen); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
	}
}

func (srv *Server) GracefulStop() {
	srv.grpcServer.GracefulStop()
}
