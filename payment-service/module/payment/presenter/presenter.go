package presenter

import (
	"context"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"github.com/google/wire"

	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
)

type (
	IPaymentPresenter interface {
		FindPaymentById(ctx context.Context, request *paymentRpc.FindPaymentByIdRequest) (*paymentRpc.Payment, error)
		FindPaymentByUserIdAndStatus(ctx context.Context, request *paymentRpc.FindPaymentByUserIdAndStatusRequest) (*paymentRpc.Payment, error)
	}

	paymentPresenter struct {
		paymentRpc.UnimplementedPaymentServiceServer
		paymentUseCase          usecase.IPaymentUseCase
		telemetryInfrastructure telemetry.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

// Set is a Wire provider set for Payment presenter dependencies
var Set = wire.NewSet(
	NewPaymentPresenter,
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
