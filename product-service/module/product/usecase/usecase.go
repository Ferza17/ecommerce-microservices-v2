package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	productRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/postgresql"
)

type (
	IProductUseCase interface {
		FindProductById(ctx context.Context, req *pb.FindProductByIdRequest) (res *pb.Product, err error)
	}

	ProductUseCase struct {
		ProductPgsqlRepository productRepo.IProductPostgresqlRepository
	}
)

func NewProductUseCase(productPgsqlRepo *productRepo.ProductPostgresqlRepository) *ProductUseCase {
	return &ProductUseCase{
		ProductPgsqlRepository: productPgsqlRepo,
	}
}
