package presenter

import (
	"context"
	userService "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/user"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IPaymentProviderPresenter interface {
		FindPaymentProviderById(ctx context.Context, request *paymentRpc.FindPaymentProviderByIdRequest) (*paymentRpc.Provider, error)
		FindPaymentProviders(ctx context.Context, request *paymentRpc.FindPaymentProvidersRequest) (*paymentRpc.FindPaymentProvidersResponse, error)
	}

	paymentProviderPresenter struct {
		paymentRpc.UnimplementedPaymentProviderServiceServer
		paymentProviderUseCase  usecase.IPaymentProviderUseCase
		telemetryInfrastructure telemetry.ITelemetryInfrastructure
		userService             userService.IUserService
		logger                  logger.IZapLogger
	}
)

// Set is a Wire provider set for PaymentProvider presenter dependencies
var Set = wire.NewSet(
	NewPaymentProviderPresenter,
)

// NewPaymentProviderPresenter creates and returns a new instance of IPaymentProviderPresenter.
func NewPaymentProviderPresenter(
	paymentProviderUseCase usecase.IPaymentProviderUseCase,
	telemetryInfrastructure telemetry.ITelemetryInfrastructure,
	userService userService.IUserService,
	logger logger.IZapLogger,
) IPaymentProviderPresenter {
	return &paymentProviderPresenter{
		paymentProviderUseCase:  paymentProviderUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		userService:             userService,
		logger:                  logger,
	}
}
