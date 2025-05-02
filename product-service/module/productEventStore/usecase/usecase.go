package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/module/productEventStore/repository/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
)

type (
	IProductEventStoreUseCase interface {
		CreateProductEventStore(ctx context.Context, requestId string, req *pb.CreateProductEventStoreRequest) (*pb.CreateProductEventStoreResponse, error)
		CreateProductEventStoreWithTransaction(ctx context.Context, requestId string, req *pb.CreateProductEventStoreRequest) (*pb.CreateProductEventStoreResponse, error)
	}

	productEventStoreUseCase struct {
		productEventStoreRepository mongodb.IProductEventStoreRepository
		logger                      pkg.IZapLogger
	}
)

func NewProductEventStoreUseCase(productEventStoreRepository mongodb.IProductEventStoreRepository, logger pkg.IZapLogger) IProductEventStoreUseCase {
	return &productEventStoreUseCase{
		productEventStoreRepository: productEventStoreRepository,
		logger:                      logger,
	}
}
