package usecase

import (
	"context"
	"errors"
	"fmt"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (u *paymentUseCase) FindPaymentByUserIdAndStatus(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByUserIdAndStatusRequest) (*paymentRpc.Payment, error) {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentUseCase.CreatePayment")
	defer span.End()

	// Begin transaction
	tx := u.postgres.GormDB.Begin()

	// Call the repository method
	payment, err := u.paymentRepository.FindPaymentByUserIdAndStatus(ctx, requestId, request.UserId, request.Status.String(), tx)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error(fmt.Sprintf("Payment not found for UserID: %s, Status: %s, RequestId: %s",
				request.GetUserId(), request.Status, requestId))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		u.logger.Error(fmt.Sprintf("Failed to fetch payment for UserID: %s, Status: %s, RequestId: %s, Error: %v",
			request.GetUserId(), request.Status, requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	tx.Commit()
	return payment.ToProto(), nil

}
