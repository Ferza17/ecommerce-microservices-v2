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

func (u *paymentUseCase) FindPaymentById(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByIdRequest) (*paymentRpc.Payment, error) {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentUseCase.FindPaymentById")
	defer span.End()

	// Begin transaction
	tx := u.postgres.GormDB.Begin()

	// Call the repository's FindPaymentById method
	payment, err := u.paymentRepository.FindPaymentById(ctx, requestId, request.Id, tx)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error(fmt.Sprintf("Payment not found for  RequestId: %s", requestId))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		u.logger.Error(fmt.Sprintf("error for RequestId: %s", requestId))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if payment == nil {
		tx.Rollback()
		return nil, status.Error(codes.NotFound, "payment Not Found")
	}

	tx.Commit()
	return payment.ToProto(), nil
}
