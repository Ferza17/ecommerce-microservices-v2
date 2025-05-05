package presenter

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
)

type ProductGrpcPresenter struct {
	pb.UnimplementedProductServiceServer

	productUseCase productUseCase.IProductUseCase
	logger         pkg.IZapLogger
}

func NewProductGrpcPresenter(productUseCase productUseCase.IProductUseCase, logger pkg.IZapLogger) *ProductGrpcPresenter {
	return &ProductGrpcPresenter{
		productUseCase: productUseCase,
		logger:         logger,
	}
}
