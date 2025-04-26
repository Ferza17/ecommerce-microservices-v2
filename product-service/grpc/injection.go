package grpc

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	productPresenter "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/presenter"
	productpgRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/postgresql"

	productUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
)

func (srv *Server) RegisterService() {

	// Register Repository & UseCase
	newProductPgRepo := productpgRepo.NewProductPostgresqlRepository(srv.postgresqlConnector)
	newProductUseCase := productUseCase.NewProductUseCase(newProductPgRepo)

	// CreateUser Service, Service can be multiple
	pb.RegisterProductServiceServer(srv.grpcServer, productPresenter.NewProductGrpcPresenter(newProductUseCase))
}
