package grpc

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
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
		dependency *bootstrap.Bootstrap
	}

	Option func(server *GrpcTransport)
)

func NewServer(dependency *bootstrap.Bootstrap) *GrpcTransport {
	return &GrpcTransport{
		address:    config.Get().RpcHost,
		port:       config.Get().RpcPort,
		dependency: dependency,
	}
}

func (srv *GrpcTransport) Serve() {
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
	pb.RegisterProductServiceServer(
		srv.grpcServer,
		presenter.NewProductGrpcPresenter(srv.dependency.ProductUseCase, srv.dependency.Logger),
	)

	// Enable Reflection to Evans grpc client
	reflection.Register(srv.grpcServer)
	if err = srv.grpcServer.Serve(listen); err != nil {
		srv.dependency.Logger.Error(fmt.Sprintf("failed to serve : %s", zap.Error(err).String))
	}
}

func (srv *GrpcTransport) GracefulStop() {
	srv.grpcServer.GracefulStop()
}
