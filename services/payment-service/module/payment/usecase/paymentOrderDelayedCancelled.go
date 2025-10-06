package usecase

import (
	"context"
	"fmt"

	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/kafka"
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

	if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkPgPayment.Payments, payment.ID, kafka.JSON_SCHEMA, payment); err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("Error publishing event to kafka for payment creation: %s", err.Error()))
		return status.Errorf(codes.Internal, "Error publishing event to kafka for payment creation: %s", err.Error())
	}

	tx.Commit()
	return nil
}
