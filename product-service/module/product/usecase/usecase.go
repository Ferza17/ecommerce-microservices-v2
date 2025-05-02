package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	productRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/postgresql"
)

type (
	IProductUseCase interface {
		FindProductById(ctx context.Context, req *pb.FindProductByIdRequest) (*pb.Product, error)
		FindProductsWithPagination(ctx context.Context, req *pb.FindProductsWithPaginationRequest) (*pb.FindProductsWithPaginationResponse, error)
		CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error)
		UpdateProductById(ctx context.Context, req *pb.UpdateProductByIdRequest) (*pb.Product, error)
		DeleteProductById(ctx context.Context, req *pb.DeleteProductByIdRequest) (*pb.DeleteProductByIdResponse, error)
	}

	ProductUseCase struct {
		ProductPgsqlRepository productRepo.IProductPostgresqlRepository
	}
)

func NewProductUseCase(productPgsqlRepo productRepo.IProductPostgresqlRepository) IProductUseCase {
	return &ProductUseCase{
		ProductPgsqlRepository: productPgsqlRepo,
	}
}
