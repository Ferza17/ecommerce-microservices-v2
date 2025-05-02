package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	productRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/postgresql"
	productEventStoreUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/productEventStore/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
)

type (
	IProductUseCase interface {
		FindProductById(ctx context.Context, requestId string, req *pb.FindProductByIdRequest) (*pb.Product, error)
		FindProductsWithPagination(ctx context.Context, requestId string, req *pb.FindProductsWithPaginationRequest) (*pb.FindProductsWithPaginationResponse, error)
		CreateProduct(ctx context.Context, requestId string, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error)
		UpdateProductById(ctx context.Context, requestId string, req *pb.UpdateProductByIdRequest) (*pb.Product, error)
		DeleteProductById(ctx context.Context, requestId string, req *pb.DeleteProductByIdRequest) (*pb.DeleteProductByIdResponse, error)
	}

	productUseCase struct {
		productPgsqlRepository   productRepo.IProductPostgresqlRepository
		productEventStoreUseCase productEventStoreUseCase.IProductEventStoreUseCase
		logger                   pkg.IZapLogger
	}
)

func NewProductUseCase(productPgsqlRepository productRepo.IProductPostgresqlRepository, productEventStoreUseCase productEventStoreUseCase.IProductEventStoreUseCase, logger pkg.IZapLogger) IProductUseCase {
	return &productUseCase{
		productPgsqlRepository:   productPgsqlRepository,
		productEventStoreUseCase: productEventStoreUseCase,
		logger:                   logger,
	}

}
