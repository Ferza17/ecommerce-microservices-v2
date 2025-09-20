package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	notificationPb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/notification"
	paymentPb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	productPb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/product"
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

func (u *paymentUseCase) CreatePayment(ctx context.Context, requestId string, request *paymentPb.CreatePaymentRequest) (*paymentPb.CreatePaymentResponse, error) {
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

	// VALIDATE PRODUCT AND SET AMOUNT
	var (
		productIds    []string
		mapProductQty = map[string]int32{}              // product_id as a key
		mapProduct    = map[string]*productPb.Product{} // product_id as a key
		totalAmount   = float64(0)
	)

	for _, item := range request.Items {
		mapProductQty[item.ProductId] = item.Qty
		productIds = append(productIds, item.ProductId)
	}

	fetchProducts, err := u.productService.FindProductsWithPagination(ctx, requestId, &productPb.FindProductsWithPaginationRequest{
		Ids:   productIds,
		Names: nil,
		Page:  1,
		Limit: int32(len(productIds)),
	})
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("failed to fetch products: %v", err))
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	if fetchProducts.Data == nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("failed to fetch products: %v", err))
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	for _, datum := range fetchProducts.Data.Data {
		productQty, ok := mapProductQty[datum.Id]
		if !ok {
			tx.Rollback()
			u.logger.Error(fmt.Sprintf("failed to fetch product qty with id : %s : %v", datum.Id, err))
			return nil, fmt.Errorf("failed to fetch product qty with id %s : %w", datum.Id, err)
		}
		mapProduct[datum.Id] = datum
		totalAmount += float64(productQty) * datum.Price
	}

	now := time.Now()
	payment := &orm.Payment{
		ID:         uuid.NewString(),
		Code:       util.GenerateInvoiceCode(),
		TotalPrice: totalAmount,
		Status:     paymentPb.PaymentStatus_PENDING.String(),
		ProviderID: request.ProviderId,
		UserID:     request.UserId,
	}
	if err = u.kafkaInfrastructure.PublishWithJsonSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkPgPayment.Payments, payment.ID, payment); err != nil {
		u.logger.Error(fmt.Sprintf("Error publishing event to kafka for payment creation: %s", err.Error()))
		return nil, status.Errorf(codes.Internal, "Error publishing event to kafka for payment creation: %s", err.Error())
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
		productQty, ok := mapProductQty[item.ProductId]
		if !ok {
			tx.Rollback()
			u.logger.Error(fmt.Sprintf("failed to fetch product qty with id : %s : %v", item.ProductId, err))
			return nil, fmt.Errorf("failed to fetch product qty with id %s : %w", item.ProductId, err)
		}

		product, ok := mapProduct[item.ProductId]
		if !ok {
			tx.Rollback()
			u.logger.Error(fmt.Sprintf("failed to fetch product with id : %s : %v", item.ProductId, err))
			return nil, fmt.Errorf("failed to fetch product with id %s : %w", item.ProductId, err)
		}

		// Call a repository to create the PaymentItem
		paymentItem := &orm.PaymentItem{
			ID:          uuid.NewString(),
			ProductID:   item.ProductId,
			Amount:      float64(productQty) * product.Price,
			Qty:         item.Qty,
			PaymentID:   payment.ID,
			CreatedAt:   &now,
			UpdatedAt:   &now,
			DiscardedAt: nil,
		}

		if err = u.kafkaInfrastructure.PublishWithJsonSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkPgPayment.PaymentItems, paymentItem.ID, paymentItem); err != nil {
			u.logger.Error(fmt.Sprintf("Error publishing event to kafka for payment creation: %s", err.Error()))
			return nil, status.Errorf(codes.Internal, "Error publishing event to kafka for payment creation: %s", err.Error())
		}
	}

	// Publish to shipping.created
	messages, err := proto.Marshal(&shippingRpc.CreateShippingRequest{
		UserId:             user.Data.User.Id,
		PaymentId:          payment.ID,
		ShippingProviderId: request.ShippingProviderId,
	})
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("Failed to marshal CreateShipping request, requestId: %s, error: %v", requestId, err))
		return nil, fmt.Errorf("failed to marshal CreateShipping request: %w", err)
	}
	if err = u.kafkaInfrastructure.Publish(ctx, config.Get().BrokerKafkaTopicShippings.ShippingCreated, uuid.NewString(), messages); err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("Failed to publish CreateShipping request, requestId: %s, error: %v", requestId, err))
	}

	// Publish to Notification Payment Order Created
	if messages, err = proto.Marshal(&notificationPb.SendEmailPaymentOrderCreateRequest{
		Email:            user.Data.User.Email,
		PaymentId:        payment.ID,
		NotificationType: notificationPb.NotificationTypeEnum_NOTIFICATION_EMAIL_PAYMENT_ORDER_CREATED,
	}); err != nil {
		u.logger.Error("NotificationTypeEnum_NOTIFICATION_EMAIL_PAYMENT_ORDER_CREATED", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err = u.kafkaInfrastructure.Publish(ctx, config.Get().BrokerKafkaTopicNotifications.PaymentOrderCreated, fmt.Sprintf("%s:%s", user.Data.User.Email, payment.ID), messages); err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("Failed to publish SendEmailPaymentOrderCreateRequest request, requestId: %s, error: %v", requestId, err))
	}

	// TODO: Publish to Payment.Order.Delayed.Cancelled
	//if messages, err = proto.Marshal(&paymentPb.PaymentOrderDelayedCancelledRequest{
	//	Id: paymentID,
	//}); err != nil {
	//	tx.Rollback()
	//	u.logger.Error(fmt.Sprintf("Failed to marshal PaymentOrderDelayedCancelledRequest request, requestId: %s, error: %v", requestId, err))
	//	return nil, fmt.Errorf("failed to marshal PaymentOrderDelayedCancelledRequest request: %w", err)
	//}
	//
	//if err = u.rabbitmqInfrastructure.PublishDelayedMessage(ctx, requestId, config.Get().ExchangePaymentDelayed, config.Get().QueuePaymentOrderDelayedCancelled, messages, config.Get().PaymentOrderCancelledInMs); err != nil {
	//	tx.Rollback()
	//	u.logger.Error(fmt.Sprintf("Failed to publish event payment.delayed.cancelled, requestId: %s, error: %v", requestId, err))
	//	return nil, fmt.Errorf("failed to commit transaction: %w", err)
	//}

	// Commit the transaction
	if err = tx.Commit().Error; err != nil {
		u.logger.Error(fmt.Sprintf("Failed to commit transaction, requestId: %s, error: %v", requestId, err))
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &paymentPb.CreatePaymentResponse{
		Message: "CreatePayment",
		Status:  "success",
		Data: &paymentPb.CreatePaymentResponse_CreatePaymentResponseData{
			Id: payment.ID,
		},
	}, nil
}
