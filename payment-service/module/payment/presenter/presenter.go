package presenter

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
)

type (
	IPaymentPresenter interface {
	}

	paymentPresenter struct {
		paymentUseCase          usecase.IPaymentUseCase
		telemetryInfrastructure telemetry.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

// NewPaymentPresenter creates a new instance of paymentPresenter.
func NewPaymentPresenter(
	paymentUseCase usecase.IPaymentUseCase,
	telemetryInfrastructure telemetry.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IPaymentPresenter {
	return &paymentPresenter{
		paymentUseCase:          paymentUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
