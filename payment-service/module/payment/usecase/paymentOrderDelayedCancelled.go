package usecase

import (
	"context"
	"fmt"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (u *paymentUseCase) PaymentOrderDelayedCancelled(ctx context.Context, requestId string, request *paymentRpc.PaymentOrderDelayedCancelledRequest) error {
	var (
		err error
		tx  = u.postgres.GormDB.Begin()
	)

	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentUseCase.CreatePayment")
	defer span.End()

	payment, err := u.paymentRepository.LockPaymentByIdWithTransaction(ctx, requestId, request.Id, tx)
	if err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			u.logger.Error(fmt.Sprintf("Payment not found for  RequestId: %s", requestId))
			return status.Error(codes.NotFound, err.Error())
		}

		u.logger.Error(fmt.Sprintf("Failed to lock payment, requestId: %s, error: %v", requestId, err))
		return status.Error(codes.Internal, err.Error())
	}

	if payment.Status == paymentRpc.PaymentStatus_SUCCESS.String() {
		tx.Rollback()
		u.logger.Info(fmt.Sprintf("payment already success, requestId: %s, error: %v", requestId, err))
		return nil
	}

	if err = u.paymentRepository.UpdatePaymentStatusByIdWithTransaction(ctx, requestId, payment.ID, paymentRpc.PaymentStatus_FAILED.String(), tx); err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("Failed to update payment status, requestId: %s, error: %v", requestId, err))
		return status.Error(codes.Internal, err.Error())
	}

	tx.Commit()
	return nil
}
