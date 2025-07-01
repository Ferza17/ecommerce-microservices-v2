package presenter

import (
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	"github.com/google/wire"

	productUseCase "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
)

type ProductPresenter struct {
	productRpc.UnimplementedProductServiceServer

	productUseCase          productUseCase.IProductUseCase
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	logger                  logger.IZapLogger
}

var Set = wire.NewSet(NewProductPresenter)

func NewProductPresenter(
	productUseCase productUseCase.IProductUseCase,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger) *ProductPresenter {
	return &ProductPresenter{
		productUseCase:          productUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
