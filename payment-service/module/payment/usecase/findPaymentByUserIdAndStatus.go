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
	// Start tracing the use case
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindPaymentProviders")
	defer span.End()

	// Call the repository method
	payment, err := u.paymentRepository.FindPaymentByUserIdAndStatus(ctx, requestId, request.Id, request.Status.String())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error(fmt.Sprintf("Payment not found for UserID: %s, Status: %s, RequestId: %s",
				request.Id, request.Status, requestId))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		u.logger.Error(fmt.Sprintf("Failed to fetch payment for UserID: %s, Status: %s, RequestId: %s, Error: %v",
			request.Id, request.Status, requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Return the response
	return payment.ToProto(), nil

}
