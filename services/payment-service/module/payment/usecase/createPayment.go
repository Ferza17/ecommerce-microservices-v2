package usecase

import (
	"context"
	"fmt"

	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	eventPb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/event"
	notificationPb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/notification"
	paymentPb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	productPb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/product"
	shippingPb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/shipping"
	userPb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/util"
	"github.com/google/uuid"
)

func (u *paymentUseCase) CreatePayment(ctx context.Context, requestId string, request *paymentPb.CreatePaymentRequest) (*paymentPb.CreatePaymentResponse, error) {
	var (
		causationId = pkgContext.GetCausationIdFromContext(ctx)
		err         error
	)
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentUseCase.CreatePayment")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	// VALIDATE PAYMENT PROVIDER
	paymentProvider, err := u.paymentProviderRepository.FindPaymentProviderById(ctx, requestId, request.ProviderId, nil)
	if err != nil {
		u.logger.Error(fmt.Sprintf("payment provider id not found, provider_id: %s ,requestId: %s, error: %v", request.ProviderId, requestId, err))
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "provider id not found")
		}
		return nil, status.Errorf(codes.Internal, "invalid payment provider")
	}

	// VALIDATE SHIPPING PROVIDER
	//if _, err = u.shippingService.GetShippingProviderById(ctx, requestId, &shippingPb.GetShippingProviderByIdRequest{
	//	Id: request.ShippingProviderId,
	//}); err != nil {
	//	u.logger.Error(fmt.Sprintf("shipping provider id not found, shipping_provider_id: %s ,requestId: %s, error: %v", request.ShippingProviderId, requestId, err))
	//	return nil, err
	//}

	// VALIDATE PRODUCT AND SET AMOUNT
	var (
		productIds []string
		mapProduct = map[string]*productPb.Product{} // product_id as a key
		now        = time.Now()
	)

	for _, item := range request.Items {
		productIds = append(productIds, item.ProductId)
	}

	fetchProducts, err := u.productService.FindProductsWithPagination(ctx, requestId, &productPb.FindProductsWithPaginationRequest{
		Ids:   productIds,
		Page:  1,
		Limit: int32(len(productIds)),
	})
	if err != nil {
		u.logger.Error(fmt.Sprintf("failed to fetch products: %v", err))
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	if fetchProducts.Data == nil {
		u.logger.Error(fmt.Sprintf("failed to fetch products: %v", err))
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	for _, datum := range fetchProducts.Data.Data {
		mapProduct[datum.Id] = datum
	}

	user, err := u.userService.AuthUserFindUserByToken(ctx, requestId, &userPb.AuthUserFindUserByTokenRequest{
		Token: pkgContext.GetTokenAuthorizationFromContext(ctx),
	})
	if err != nil {
		u.logger.Error(fmt.Sprintf("Failed to find user by provided token, requestId: %s, error: %v", requestId, err))
		return nil, fmt.Errorf("failed to find user by provided token: %w", err)
	}

	payment := &orm.Payment{
		ID:           uuid.NewString(),
		Code:         util.GenerateInvoiceCode(),
		TotalPrice:   0,
		Status:       paymentPb.PaymentStatus_PENDING.String(),
		ProviderID:   request.ProviderId,
		UserID:       user.Data.User.Id,
		PaymentItems: []*orm.PaymentItem{},
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}
	// Process PaymentItems
	for _, item := range request.Items {
		product, ok := mapProduct[item.ProductId]
		if !ok {
			u.logger.Error(fmt.Sprintf("failed to fetch product with id : %s : %v", item.ProductId, err))
			return nil, fmt.Errorf("failed to fetch product with id %s : %w", item.ProductId, err)
		}

		stock := product.Stock - int64(item.Qty)
		if stock < 0 {
			u.logger.Error(fmt.Sprintf("invalid product qty stock with id : %s : %v", item.ProductId, err))
			return nil, fmt.Errorf("invalid product qty stock with id : %s : %w", item.ProductId, err)
		}
		product.Stock = stock

		// Send to topic product update for updating product stock, move with outbox
		if err = u.eventUseCase.AppendEventEnvelope(ctx, &eventPb.EventEnvelope{
			XId:           primitive.NewObjectID().Hex(),
			EventType:     config.Get().BrokerKafkaTopicProducts.ProductUpdated,
			AggregateType: eventPb.AggregateType_PRODUCT,
			AggregateId:   product.Id,
			Version:       0,
			OccurredAt:    timestamppb.New(now),
			CorrelationId: requestId,
			CausationId:   &causationId,
		}, product); err != nil {
			u.logger.Error(fmt.Sprintf("failed to append event envelope: %v", err))
			return nil, fmt.Errorf("failed to append event envelope: %w", err)
		}

		amount := float64(item.Qty) * product.Price
		payment.TotalPrice += amount

		payment.PaymentItems = append(payment.PaymentItems, &orm.PaymentItem{
			ID:          uuid.NewString(),
			ProductID:   item.ProductId,
			Amount:      amount,
			Qty:         item.Qty,
			PaymentID:   payment.ID,
			CreatedAt:   &now,
			UpdatedAt:   &now,
			DiscardedAt: nil,
		})

	}

	// Publish to Shipping Created, with outbox
	if err = u.eventUseCase.AppendEventEnvelope(ctx, &eventPb.EventEnvelope{
		XId:           primitive.NewObjectID().Hex(),
		EventType:     config.Get().BrokerKafkaTopicShippings.ShippingCreated,
		AggregateType: eventPb.AggregateType_SHIPPING,
		AggregateId:   uuid.NewString(),
		Version:       0,
		OccurredAt:    timestamppb.New(now),
		CorrelationId: requestId,
		CausationId:   &causationId,
	}, &shippingPb.CreateShippingRequest{
		UserId:             user.Data.User.Id,
		PaymentId:          payment.ID,
		ShippingProviderId: request.ShippingProviderId,
	}); err != nil {
		u.logger.Error(fmt.Sprintf("failed to append event envelope: %v", err))
		return nil, fmt.Errorf("failed to append event envelope: %w", err)
	}

	// Publish to Notification Payment Order Created, with outbox
	if err = u.eventUseCase.AppendEventEnvelope(ctx, &eventPb.EventEnvelope{
		XId:           primitive.NewObjectID().Hex(),
		EventType:     config.Get().BrokerKafkaTopicNotifications.EmailPaymentOrderCreated,
		AggregateType: eventPb.AggregateType_NOTIFICATION,
		AggregateId:   uuid.NewString(),
		Version:       0,
		OccurredAt:    timestamppb.New(now),
		CorrelationId: requestId,
		CausationId:   &causationId,
	}, &notificationPb.SendEmailPaymentOrderCreateRequest{
		Email:            user.Data.User.Email,
		Payment:          payment.ToProto(),
		PaymentProvider:  paymentProvider.ToProto(),
		NotificationType: notificationPb.NotificationTypeEnum_NOTIFICATION_EMAIL_PAYMENT_ORDER_CREATED,
	}); err != nil {
		u.logger.Error(fmt.Sprintf("failed to append event envelope: %v", err))
		return nil, fmt.Errorf("failed to append event envelope: %w", err)
	}

	if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkPgPayment.Payments, payment.ID, kafka.JSON_SCHEMA, payment); err != nil {
		u.logger.Error(fmt.Sprintf("Error publishing event to kafka for payment creation: %s", err.Error()))
		return nil, err
	}

	for _, item := range payment.PaymentItems {
		if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkPgPayment.PaymentItems, item.ID, kafka.JSON_SCHEMA, item); err != nil {
			u.logger.Error(fmt.Sprintf("Error publishing event to kafka for payment item creation: %s", err.Error()))
			return nil, err
		}
	}

	// TODO: Publish to Payment.Order.Delayed.Cancelled
	//if messages, err = proto.Marshal(&paymentPb.PaymentOrderDelayedCancelledRequest{
	//	Id: paymentID,
	//}); err != nil {
	//
	//	u.logger.Error(fmt.Sprintf("Failed to marshal PaymentOrderDelayedCancelledRequest request, requestId: %s, error: %v", requestId, err))
	//	return nil, fmt.Errorf("failed to marshal PaymentOrderDelayedCancelledRequest request: %w", err)
	//}
	//
	//if err = u.rabbitmqInfrastructure.PublishDelayedMessage(ctx, requestId, config.Get().ExchangePaymentDelayed, config.Get().QueuePaymentOrderDelayedCancelled, messages, config.Get().PaymentOrderCancelledInMs); err != nil {
	//
	//	u.logger.Error(fmt.Sprintf("Failed to publish event payment.delayed.cancelled, requestId: %s, error: %v", requestId, err))
	//	return nil, fmt.Errorf("failed to commit transaction: %w", err)
	//}

	return &paymentPb.CreatePaymentResponse{
		Message: "CreatePayment",
		Status:  "success",
		Data: &paymentPb.CreatePaymentResponse_CreatePaymentResponseData{
			Id: payment.ID,
		},
	}, nil
}
