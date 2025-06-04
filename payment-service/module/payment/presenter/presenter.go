package presenter

import (
	"context"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"

	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
)

type (
	IPaymentPresenter interface {
		FindPaymentById(ctx context.Context, request *paymentRpc.FindPaymentByIdRequest) (*paymentRpc.Payment, error)
		FindPaymentByUserIdAndStatus(ctx context.Context, request *paymentRpc.FindPaymentByUserIdAndStatusRequest) (*paymentRpc.FindPaymentByUserIdAndStatusRequest, error)
	}

	paymentPresenter struct {
		paymentRpc.UnimplementedPaymentServiceServer
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
