package grpc

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/product/v1"

	"github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/presenter"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
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
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	}
	srv.grpcServer = grpc.NewServer(opts...)
	productRpc.RegisterProductServiceServer(
		srv.grpcServer,
		presenter.NewProductGrpcPresenter(srv.dependency.ProductUseCase, srv.dependency.TelemetryInfrastructure, srv.dependency.Logger),
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
