package grpc

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/connector/postgresql"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type (
	Server struct {
		address             string
		port                string
		listener            *net.Listener
		grpcServer          *grpc.Server
		logger              *zap.Logger
		postgresqlConnector *postgresql.PostgresqlConnector
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
		log.Fatalln(err)
	}
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
