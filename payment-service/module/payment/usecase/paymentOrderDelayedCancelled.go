package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (u *paymentUseCase) PaymentOrderDelayedCancelled(ctx context.Context, requestId string, request *paymentRpc.PaymentOrderDelayedCancelledRequest) error {
	var (
		err error
		tx  *gorm.DB
	)

	// Start tracing the use case
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindPaymentProviders")

	defer func(err error, tx *gorm.DB) {
		defer span.End()
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}(err, tx)

	payment, err := u.paymentRepository.LockPaymentByIdWithTransaction(ctx, requestId, request.Id, tx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.logger.Error(fmt.Sprintf("Payment not found for  RequestId: %s", requestId))
			return status.Error(codes.NotFound, err.Error())
		}

		u.logger.Error(fmt.Sprintf("Failed to lock payment, requestId: %s, error: %v", requestId, err))
		return status.Error(codes.Internal, err.Error())
	}

	if payment.Status == enum.PaymentStatusSuccess {
		u.logger.Info(fmt.Sprintf("payment already success, requestId: %s, error: %v", requestId, err))
		return nil
	}

	if err = u.paymentRepository.UpdatePaymentStatusByIdWithTransaction(ctx, requestId, payment.ID, enum.PaymentStatusFailed, tx); err != nil {
		u.logger.Error(fmt.Sprintf("Failed to update payment status, requestId: %s, error: %v", requestId, err))
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
