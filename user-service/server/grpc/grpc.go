package grpc

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
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
		address     string
		port        string
		grpcServer  *grpc.Server
		logger      pkg.IZapLogger
		userUseCase usecase.IUserUseCase
	}
)

func NewServer(dependency *bootstrap.Bootstrap) *Server {
	return &Server{
		address:     config.Get().RpcHost,
		port:        config.Get().RpcPort,
		logger:      dependency.Logger,
		userUseCase: dependency.UserUseCase,
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
	pb.RegisterUserServiceServer(
		srv.grpcServer,
		presenter.NewUserPresenter(srv.userUseCase, srv.logger),
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
