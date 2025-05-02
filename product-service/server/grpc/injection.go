package grpc

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	productPresenter "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/presenter"
	productpgRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/postgresql"
	productEventStoreRepository "github.com/ferza17/ecommerce-microservices-v2/product-service/module/productEventStore/repository/mongodb"
	productEventStoreUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/productEventStore/usecase"

	productUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
)

func (srv *Server) RegisterService() {

	// Register Repository & UseCase
	newProductEventStoreRepository := productEventStoreRepository.NewProductEventStoreRepository(srv.mongoDBConnector, srv.logger)
	newProductEventStoreUseCase := productEventStoreUseCase.NewProductEventStoreUseCase(newProductEventStoreRepository, srv.logger)

	newProductPgRepo := productpgRepo.NewProductPostgresqlRepository(srv.postgresqlConnector, srv.logger)
	newProductUseCase := productUseCase.NewProductUseCase(newProductPgRepo, newProductEventStoreUseCase, srv.logger)

	// CreateUser Service, Service can be multiple
	pb.RegisterProductServiceServer(srv.grpcServer, productPresenter.NewProductGrpcPresenter(newProductUseCase))
}
