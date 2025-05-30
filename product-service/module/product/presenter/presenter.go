package presenter

import (
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/product/v1"

	productUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
)

type ProductGrpcPresenter struct {
	productRpc.UnimplementedProductServiceServer

	productUseCase          productUseCase.IProductUseCase
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	logger                  pkg.IZapLogger
}

func NewProductGrpcPresenter(
	productUseCase productUseCase.IProductUseCase,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger) *ProductGrpcPresenter {
	return &ProductGrpcPresenter{
		productUseCase:          productUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
