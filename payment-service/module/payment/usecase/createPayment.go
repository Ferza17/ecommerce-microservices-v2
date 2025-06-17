package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	eventRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/event/v1"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/util"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *paymentUseCase) CreatePayment(ctx context.Context, requestId string, request *paymentRpc.CreatePaymentRequest) error {
	var (
		err        error
		eventStore = &eventRpc.EventStore{
			RequestId:     requestId,
			Service:       config.Get().ServiceName,
			EventType:     config.Get().QueueEventCreated,
			Status:        config.Get().CommonSagaStatusSuccess,
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
	)

	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.CreatePayment")
	defer span.End()

	// Begin transaction
	tx := u.paymentRepository.OpenTransactionWithContext(ctx)

	defer func(err error, eventStore *eventRpc.EventStore) {
		defer span.End()
		payload, err := util.ConvertStructToProtoStruct(request)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error converting struct to proto struct: %s", err.Error()))
		}
		eventStore.Payload = payload

		eventStoreMessage, err := proto.Marshal(eventStore)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		}

		if err != nil {
			eventStore.Status = config.Get().CommonSagaStatusFailed
		}

		if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, config.Get().ExchangeEvent, config.Get().QueueEventCreated, eventStoreMessage); err != nil {
			u.logger.Error(fmt.Sprintf("error creating product event store: %s", err.Error()))
			return
		}
	}(err, eventStore)

	// Map the request into the ORM Payment model
	payment := &orm.Payment{
		ID:         uuid.NewString(),
		Code:       util.GenerateInvoiceCode(),
		TotalPrice: request.Amount,
		Status:     enum.PaymentStatusPending,
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
		paymentItem := &orm.PaymentItem{
			ID:        uuid.NewString(),
			PaymentID: paymentID, // Associate with the Payment
			ProductID: item.ProductId,
			Amount:    item.Amount,
			Qty:       item.Qty,
		}

		// Call a repository to create the PaymentItem
		if _, err := u.paymentRepository.CreatePaymentItem(ctx, paymentItem, tx); err != nil {
			tx.Rollback() // Roll back the transaction on error
			u.logger.Error(fmt.Sprintf("Failed to create payment item, requestId: %s, paymentItemId: %s, error: %v", requestId, paymentItem.ID, err))
			return fmt.Errorf("failed to create payment item: %w", err)
		}
	}

	// Publish to Payment.Order.Delayed.Cancelled
	delayedMessage, err := proto.Marshal(&paymentRpc.PaymentOrderDelayedCancelledRequest{
		Id: paymentID,
	})

	if err = u.rabbitmqInfrastructure.PublishDelayedMessage(ctx, requestId, config.Get().ExchangePayment, config.Get().QueuePaymentOrderDelayedCancelled, delayedMessage, config.Get().PaymentOrderCancelledInMs); err != nil {
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
