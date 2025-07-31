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
		if _, err := u.paymentRepository.CreatePaymentItem(ctx, orm.PaymentItemFromProto(item), tx); err != nil {
			tx.Rollback() // Roll back the transaction on error
			u.logger.Error(fmt.Sprintf("Failed to create payment item, requestId: %s, error: %v", requestId, err))
			return nil, fmt.Errorf("failed to create payment item: %w", err)
		}
	}

	// Create Shipping with RPC
	// TODO: Change to RabbitMQ
	if _, err = u.shippingService.CreateShipping(ctx, requestId, &shippingRpc.CreateShippingRequest{
		UserId:             user.Data.User.Id,
		PaymentId:          paymentID,
		ShippingProviderId: request.ShippingProviderId,
	}); err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("Failed to create shipping, requestId: %s, error: %v", requestId, err))
		return nil, fmt.Errorf("failed to create shipping: %w", err)
	}

	// Publish to payment.order.direct.created
	reqNotificationOrderCreated := &notificationRpc.SendEmailPaymentOrderCreateRequest{
		Email:            user.Data.User.Email,
		Payment:          nil, //TODO: Fill this argument
		NotificationType: notificationRpc.NotificationTypeEnum_NOTIFICATION_EMAIL_PAYMENT_ORDER_CREATED,
	}
	message, err := proto.Marshal(reqNotificationOrderCreated)
	if err != nil {
		u.logger.Error("AuthUseCase.SentOTP", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, config.Get().ExchangePaymentDirect, config.Get().QueuePaymentOrderCreated, message); err != nil {
		u.logger.Error("Failed to publish notification", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Publish to Payment.Order.Delayed.Cancelled
	delayedMessage, err := proto.Marshal(&paymentRpc.PaymentOrderDelayedCancelledRequest{
		Id: paymentID,
	})

	if err = u.rabbitmqInfrastructure.PublishDelayedMessage(ctx, requestId, config.Get().ExchangePaymentDelayed, config.Get().QueuePaymentOrderDelayedCancelled, delayedMessage, config.Get().PaymentOrderCancelledInMs); err != nil {
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
