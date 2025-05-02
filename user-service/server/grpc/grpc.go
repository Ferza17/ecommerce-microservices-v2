package grpc

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/connector"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type (
	Server struct {
		address             string
		port                string
		listener            *net.Listener
		grpcServer          *grpc.Server
		logger              pkg.IZapLogger
		mongoDBConnector    *connector.MongodbConnector
		postgresqlConnector *connector.PostgresqlConnector
	}

	Option func(server *Server)
)

func NewServer(address, port string, option ...Option) *Server {
	s := &Server{
		address: address,
		port:    port,
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
		srv.logger.Error(fmt.Sprintf("failed to serve", zap.Error(err)))
	}
}

func (srv *Server) GracefulStop() {
	srv.grpcServer.GracefulStop()
}

func (srv *Server) setup() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", srv.address, srv.port))
	if err != nil {
		srv.logger.Error(fmt.Sprintf("failed to serve", zap.Error(err)))
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
