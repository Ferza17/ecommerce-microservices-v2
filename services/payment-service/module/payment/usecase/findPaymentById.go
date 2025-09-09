package usecase

import (
	"context"
	"fmt"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (u *paymentUseCase) FindPaymentById(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByIdRequest) (*paymentRpc.FindPaymentByIdResponse, error) {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentUseCase.FindPaymentById")
	defer span.End()

	// Begin transaction
	tx := u.postgres.GormDB.Begin()

	// Call the repository's FindPaymentById method
	payment, err := u.paymentRepository.FindPaymentById(ctx, requestId, request.Id, tx)
	if err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			u.logger.Error(fmt.Sprintf("Payment not found for  RequestId: %s", requestId))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		u.logger.Error(fmt.Sprintf("error for RequestId: %s", requestId))
		return nil, status.Error(codes.Internal, err.Error())
	}

	tx.Commit()
	return &paymentRpc.FindPaymentByIdResponse{
		Message: "FindPaymentById",
		Status:  "success",
		Data: &paymentRpc.FindPaymentByIdResponse_FindPaymentByIdResponseData{
			Payment:  payment.ToProto(),
			Provider: payment.PaymentProvider.ToProto(),
			PaymentItems: func() []*paymentRpc.PaymentItem {
				paymentItems := make([]*paymentRpc.PaymentItem, len(payment.PaymentItems))
				for i, paymentItem := range payment.PaymentItems {
					paymentItems[i] = paymentItem.ToProto()
				}
				return paymentItems
			}(),
		},
	}, nil
}
