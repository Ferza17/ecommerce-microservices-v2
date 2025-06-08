package usecase

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/repository"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IPaymentProviderUseCase interface {
	}

	paymentProviderUseCase struct {
		paymentProviderRepository repository.IPaymentProviderRepository
		rabbitmqInfrastructure    rabbitmq.IRabbitMQInfrastructure
		telemetryInfrastructure   telemetry.ITelemetryInfrastructure
		logger                    logger.IZapLogger
	}
)

// NewPaymentProviderUseCase creates and returns a new instance of IPaymentProvider.
func NewPaymentProviderUseCase(
	paymentProviderRepository repository.IPaymentProviderRepository,
	rabbitmqInfrastructure rabbitmq.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetry.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IPaymentProviderUseCase {
	return &paymentProviderUseCase{
		paymentProviderRepository: paymentProviderRepository,
		rabbitmqInfrastructure:    rabbitmqInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		logger:                    logger,
	}
}

// Set is a Wire provider set for PaymentProvider use case dependencies
var Set = wire.NewSet(
	NewPaymentProviderUseCase,
)
