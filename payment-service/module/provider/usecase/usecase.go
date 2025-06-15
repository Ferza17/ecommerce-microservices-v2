package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/repository"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IPaymentProviderUseCase interface {
		FindPaymentProviderById(ctx context.Context, requestId string, request *paymentRpc.FindPaymentProviderByIdRequest) (*paymentRpc.Provider, error)
		FindPaymentProviders(ctx context.Context, requestId string, request *paymentRpc.FindPaymentProvidersRequest) (*paymentRpc.FindPaymentProvidersResponse, error)
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
