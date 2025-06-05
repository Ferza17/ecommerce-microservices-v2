package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
)

type (
	IPaymentProviderPresenter interface {
		FindPaymentProviders(ctx context.Context, request *paymentRpc.FindPaymentProvidersRequest) (*paymentRpc.FindPaymentProvidersResponse, error)
	}

	paymentProviderPresenter struct {
		paymentRpc.UnimplementedPaymentProviderServiceServer
		paymentProviderUseCase  usecase.IPaymentProviderUseCase
		telemetryInfrastructure telemetry.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

// NewPaymentProviderPresenter creates and returns a new instance of IPaymentProviderPresenter.
func NewPaymentProviderPresenter(
	paymentProviderUseCase usecase.IPaymentProviderUseCase,
	telemetryInfrastructure telemetry.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IPaymentProviderPresenter {
	return &paymentProviderPresenter{
		paymentProviderUseCase:  paymentProviderUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
