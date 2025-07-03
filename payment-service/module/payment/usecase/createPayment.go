package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/util"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func (u *paymentUseCase) CreatePayment(ctx context.Context, requestId string, request *paymentRpc.CreatePaymentRequest) error {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.CreatePayment")
	defer span.End()

	// Begin transaction
	tx := u.paymentRepository.OpenTransactionWithContext(ctx)

	// Map the request into the ORM Payment model
	payment := &orm.Payment{
		ID:         uuid.NewString(),
		Code:       util.GenerateInvoiceCode(),
		TotalPrice: request.Amount,
		Status:     paymentRpc.PaymentStatus_PENDING.String(),
		ProviderID: request.ProviderId,
		UserID:     request.UserId,
	}

	// Call repository to create the Payment
	paymentID, err := u.paymentRepository.CreatePayment(ctx, requestId, payment, tx)
	if err != nil {
		tx.Rollback() // Roll back the transaction on error
		u.logger.Error(fmt.Sprintf("Failed to create payment, requestId: %s, error: %v", requestId, err))
		return fmt.Errorf("failed to create payment: %w", err)
	}

	// Process PaymentItems
	for _, item := range request.Items {
		// Call a repository to create the PaymentItem
		if _, err := u.paymentRepository.CreatePaymentItem(ctx, orm.PaymentItemFromProto(item), tx); err != nil {
			tx.Rollback() // Roll back the transaction on error
			u.logger.Error(fmt.Sprintf("Failed to create payment item, requestId: %s, error: %v", requestId, err))
			return fmt.Errorf("failed to create payment item: %w", err)
		}
	}

	// Publish to Payment.Order.Delayed.Cancelled
	delayedMessage, err := proto.Marshal(&paymentRpc.PaymentOrderDelayedCancelledRequest{
		Id: paymentID,
	})

	if err = u.rabbitmqInfrastructure.PublishDelayedMessage(ctx, requestId, config.Get().ExchangePaymentDelayed, config.Get().QueuePaymentOrderDelayedCancelled, delayedMessage, config.Get().PaymentOrderCancelledInMs); err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("Failed to publish event payment.delayed.cancelled, requestId: %s, error: %v", requestId, err))
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Commit the transaction
	if err = tx.Commit().Error; err != nil {
		u.logger.Error(fmt.Sprintf("Failed to commit transaction, requestId: %s, error: %v", requestId, err))
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Log success
	u.logger.Info(fmt.Sprintf("Successfully created payment and payment items, requestId: %s, paymentId: %s", requestId, paymentID))
	return nil
}
