package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/notification"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	shippingRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/shipping"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/user"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/util"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func (u *paymentUseCase) CreatePayment(ctx context.Context, requestId string, request *paymentRpc.CreatePaymentRequest) (*paymentRpc.CreatePaymentResponse, error) {
	tx := u.postgres.GormDB.Begin()
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentUseCase.CreatePayment")
	defer span.End()

	// VALIDATE PAYMENT PROVIDER
	if _, err := u.paymentProviderRepository.FindPaymentProviderById(ctx, requestId, request.ProviderId, tx); err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("payment provider id not found, provider_id: %s ,requestId: %s, error: %v", request.ProviderId, requestId, err))
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	// VALIDATE SHIPPING PROVIDER
	if _, err := u.shippingService.GetShippingProviderById(ctx, requestId, &shippingRpc.GetShippingProviderByIdRequest{
		Id: request.ShippingProviderId,
	}); err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("shipping provider id not found, shipping_provider_id: %s ,requestId: %s, error: %v", request.ShippingProviderId, requestId, err))
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	now := time.Now()
	paymentID, err := u.paymentRepository.CreatePayment(ctx, requestId, &orm.Payment{
		ID:         uuid.NewString(),
		Code:       util.GenerateInvoiceCode(),
		TotalPrice: request.Amount,
		Status:     paymentRpc.PaymentStatus_PENDING.String(),
		ProviderID: request.ProviderId,
		UserID:     request.UserId,
	}, tx)
	if err != nil {
		tx.Rollback() // Roll back the transaction on error
		u.logger.Error(fmt.Sprintf("Failed to create payment, requestId: %s, error: %v", requestId, err))
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	user, err := u.userService.AuthUserFindUserByToken(ctx, requestId, &userRpc.AuthUserFindUserByTokenRequest{
		Token: pkgContext.GetTokenAuthorizationFromContext(ctx),
	})
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("Failed to find user by provided token, requestId: %s, error: %v", requestId, err))
		return nil, fmt.Errorf("failed to find user by provided token: %w", err)
	}

	// Process PaymentItems
	for _, item := range request.Items {
		// Call a repository to create the PaymentItem
		if _, err = u.paymentRepository.CreatePaymentItem(ctx, &orm.PaymentItem{
			ID:          uuid.NewString(),
			ProductID:   item.ProductId,
			Amount:      0,
			Qty:         item.Qty,
			PaymentID:   paymentID,
			CreatedAt:   &now,
			UpdatedAt:   &now,
			DiscardedAt: nil,
		}, tx); err != nil {
			tx.Rollback() // Roll back the transaction on error
			u.logger.Error(fmt.Sprintf("Failed to create payment item, requestId: %s, error: %v", requestId, err))
			return nil, fmt.Errorf("failed to create payment item: %w", err)
		}
	}

	// Publish to shipping.created
	messages, err := proto.Marshal(&shippingRpc.CreateShippingRequest{
		UserId:             user.Data.User.Id,
		PaymentId:          paymentID,
		ShippingProviderId: request.ShippingProviderId,
	})
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("Failed to marshal CreateShipping request, requestId: %s, error: %v", requestId, err))
		return nil, fmt.Errorf("failed to marshal CreateShipping request: %w", err)
	}
	if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, config.Get().ExchangeShipping, config.Get().QueueShippingCreated, messages); err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("Failed to publish CreateShipping request, requestId: %s, error: %v", requestId, err))
		return nil, fmt.Errorf("failed to publish CreateShipping request: %w", err)
	}

	// Publish to Notification Payment Order Created
	if messages, err = proto.Marshal(&notificationRpc.SendEmailPaymentOrderCreateRequest{
		Email:            user.Data.User.Email,
		PaymentId:        paymentID,
		NotificationType: notificationRpc.NotificationTypeEnum_NOTIFICATION_EMAIL_PAYMENT_ORDER_CREATED,
	}); err != nil {
		u.logger.Error("NotificationTypeEnum_NOTIFICATION_EMAIL_PAYMENT_ORDER_CREATED", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, config.Get().ExchangeNotification, config.Get().QueueNotificationEmailPaymentOrderCreated, messages); err != nil {
		u.logger.Error("Failed to publish notification", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Publish to Payment.Order.Delayed.Cancelled
	if messages, err = proto.Marshal(&paymentRpc.PaymentOrderDelayedCancelledRequest{
		Id: paymentID,
	}); err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("Failed to marshal PaymentOrderDelayedCancelledRequest request, requestId: %s, error: %v", requestId, err))
		return nil, fmt.Errorf("failed to marshal PaymentOrderDelayedCancelledRequest request: %w", err)
	}

	if err = u.rabbitmqInfrastructure.PublishDelayedMessage(ctx, requestId, config.Get().ExchangePaymentDelayed, config.Get().QueuePaymentOrderDelayedCancelled, messages, config.Get().PaymentOrderCancelledInMs); err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("Failed to publish event payment.delayed.cancelled, requestId: %s, error: %v", requestId, err))
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Commit the transaction
	if err = tx.Commit().Error; err != nil {
		u.logger.Error(fmt.Sprintf("Failed to commit transaction, requestId: %s, error: %v", requestId, err))
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	tx.Commit()
	return &paymentRpc.CreatePaymentResponse{
		Message: "CreatePayment",
		Status:  "success",
		Data: &paymentRpc.CreatePaymentResponse_CreatePaymentResponseData{
			Id: paymentID,
		},
	}, nil
}
