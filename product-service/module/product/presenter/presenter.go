package presenter

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
)

type ProductGrpcPresenter struct {
	pb.UnimplementedProductServiceServer

	ProductUseCase productUseCase.IProductUseCase
}

func NewProductGrpcPresenter(productUseCase productUseCase.IProductUseCase) *ProductGrpcPresenter {
	return &ProductGrpcPresenter{
		ProductUseCase: productUseCase,
	}
}
