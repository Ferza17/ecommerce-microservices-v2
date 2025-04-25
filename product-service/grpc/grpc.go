package grpc

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

type (
	Server struct {
		address    string
		port       string
		listener   *net.Listener
		grpcServer *grpc.Server
		logger     *zap.Logger
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

func (srv *Server) setup() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", srv.address, srv.port))
	if err != nil {
		log.Fatalln(err)
	}
	opts := []grpc.ServerOption{
		//grpc.ChainUnaryInterceptor(
		//grpcMiddleware.ChainUnaryServer(
		//	otgrpc.OpenTracingServerInterceptor(srv.tracer),
		//	middleware.UnaryRegisterTracerContext(srv.tracer),
		//	middleware.UnaryRegisterRedisContext(srv.redisClient),
		//	middleware.UnaryRegisterPostgresSQLContext(srv.postgresClient),
		//	middleware.UnaryRegisterCassandraDBContext(srv.cassandraSession),
		//	middleware.UnaryRegisterRabbitMQAmqpContext(srv.rabbitMQConnection),
		//	middleware.UnaryRegisterElasticsearchContext(srv.elasticsearchClient),
		//	product.UnaryRegisterProductUseCaseContext(),
		//),
		//),
	}
	srv.grpcServer = grpc.NewServer(opts...)
	srv.listener = &listen
	//srv.RegisterService()
}
