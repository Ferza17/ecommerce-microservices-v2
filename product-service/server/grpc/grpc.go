package grpc

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type (
	Server struct {
		address                  string
		port                     string
		listener                 *net.Listener
		grpcServer               *grpc.Server
		logger                   pkg.IZapLogger
		postgresqlInfrastructure postgresql.IPostgreSQLInfrastructure
		rabbitmqInfrastructure   rabbitmq.IRabbitMQInfrastructure
	}

	Option func(server *Server)
)

func NewServer(option ...Option) *Server {
	s := &Server{
		address: config.Get().RpcHost,
		port:    config.Get().RpcPort,
	}
	for _, o := range option {
		o(s)
	}
	s.setup()
	return s
}

func (srv *Server) Serve() {
	// Enable Reflection to Evans grpc client
	reflection.Register(srv.grpcServer)
	if err := srv.grpcServer.Serve(*srv.listener); err != nil {
		srv.logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
	}
}

func (srv *Server) GracefulStop() {
	srv.grpcServer.GracefulStop()
}

func (srv *Server) setup() {
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
	srv.listener = &listen
	srv.RegisterService()
}
