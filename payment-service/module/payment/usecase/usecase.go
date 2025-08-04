package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	productService "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/product"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/shipping"
	userService "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/user"

	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/repository"
	paymentProviderRepository "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/repository"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IPaymentUseCase interface {
		CreatePayment(ctx context.Context, requestId string, request *paymentRpc.CreatePaymentRequest) (*paymentRpc.CreatePaymentResponse, error)
		PaymentOrderDelayedCancelled(ctx context.Context, requestId string, request *paymentRpc.PaymentOrderDelayedCancelledRequest) error

		FindPaymentById(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByIdRequest) (*paymentRpc.FindPaymentByIdResponse, error)
		FindPaymentByUserIdAndStatus(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByUserIdAndStatusRequest) (*paymentRpc.Payment, error)
	}

	paymentUseCase struct {
		paymentRepository         repository.IPaymentRepository
		paymentProviderRepository paymentProviderRepository.IPaymentProviderRepository
		rabbitmqInfrastructure    rabbitmq.IRabbitMQInfrastructure
		telemetryInfrastructure   telemetry.ITelemetryInfrastructure
		logger                    logger.IZapLogger
		postgres                  *postgresql.PostgresSQL
		shippingService           shipping.IShippingService
		userService               userService.IUserService
		productService            productService.IProductService
	}
)

// Set is a Wire provider set for Payment use case dependencies
var Set = wire.NewSet(
	NewPaymentUseCase,
)

func NewPaymentUseCase(
	paymentRepository repository.IPaymentRepository,
	paymentProviderRepository paymentProviderRepository.IPaymentProviderRepository,
	rabbitmqInfrastructure rabbitmq.IRabbitMQInfrastructure,
	telemetryInfrastructure telemetry.ITelemetryInfrastructure,
	logger logger.IZapLogger,
	postgres *postgresql.PostgresSQL,
	shippingService shipping.IShippingService,
	userService userService.IUserService,
	productService productService.IProductService,
) IPaymentUseCase {
	return &paymentUseCase{
		paymentRepository:         paymentRepository,
		paymentProviderRepository: paymentProviderRepository,
		rabbitmqInfrastructure:    rabbitmqInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		logger:                    logger,
		postgres:                  postgres,
		shippingService:           shippingService,
		userService:               userService,
		productService:            productService,
	}
}
