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

func (u *paymentUseCase) FindPaymentByUserIdAndStatus(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByUserIdAndStatusRequest) (*paymentRpc.Payment, error) {
	// Start tracing the use case
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindPaymentProviders")
	defer span.End()

	// Parse the payment status from request
	paymentStatus, err := enum.ProtoToPaymentStatus(request.Status)
	if err != nil {
		u.logger.Error(fmt.Sprintf("Invalid payment status: %s, RequestId: %s, Error: %v", request.Status, requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Call the repository method
	payment, err := u.paymentRepository.FindPaymentByUserIdAndStatus(ctx, requestId, request.UserId, paymentStatus)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error(fmt.Sprintf("Payment not found for UserID: %s, Status: %s, RequestId: %s",
				request.UserId, request.Status, requestId))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		u.logger.Error(fmt.Sprintf("Failed to fetch payment for UserID: %s, Status: %s, RequestId: %s, Error: %v",
			request.UserId, request.Status, requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Map the ORM payment to the RPC response
	response := &paymentRpc.Payment{
		Id:         payment.ID,
		UserId:     payment.UserID,
		TotalPrice: payment.TotalPrice,
		Status:     request.Status,
		Provider: &paymentRpc.Provider{
			Id:   payment.Provider.ID,
			Name: payment.Provider.Name,
		},
		CreatedAt: timestamppb.New(payment.CreatedAt),
		UpdatedAt: timestamppb.New(payment.UpdatedAt),
	}

	// Return the response
	return response, nil

}
