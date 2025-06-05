package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func (u *paymentUseCase) FindPaymentById(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByIdRequest) (*paymentRpc.Payment, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindPaymentById")
	defer span.End()

	// Call the repository's FindPaymentById method
	payment, err := u.paymentRepository.FindPaymentById(ctx, requestId, request.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error(fmt.Sprintf("Payment not found for  RequestId: %s", requestId))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		u.logger.Error(fmt.Sprintf("error for RequestId: %s", requestId))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if payment == nil {
		return nil, status.Error(codes.NotFound, "payment Not Found")
	}

	paymentStatus, err := enum.PaymentStatusToProto(payment.Status)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error parse status : %v", err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	// Convert the orm.Payment model to a paymentRpc.Payment response
	response := &paymentRpc.Payment{
		Id:         payment.ID,
		Code:       payment.Code,
		TotalPrice: payment.TotalPrice,
		Status:     paymentStatus,
		Provider: &paymentRpc.Provider{
			Id:   payment.Provider.ID,
			Name: payment.Provider.Name,
		},
		CreatedAt: timestamppb.New(payment.CreatedAt),
		UpdatedAt: timestamppb.New(payment.UpdatedAt),
	}

	return response, nil
}
