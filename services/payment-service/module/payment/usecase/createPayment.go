package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/event"
	notificationPb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/notification"
	paymentPb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	productPb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/product"
	shippingRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/shipping"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"

	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/util"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func (u *paymentUseCase) CreatePayment(ctx context.Context, requestId string, request *paymentPb.CreatePaymentRequest) (*paymentPb.CreatePaymentResponse, error) {
	var (
		err error
	)
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentUseCase.CreatePayment")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	// VALIDATE PAYMENT PROVIDER
	if _, err = u.paymentProviderRepository.FindPaymentProviderById(ctx, requestId, request.ProviderId, nil); err != nil {
		u.logger.Error(fmt.Sprintf("payment provider id not found, provider_id: %s ,requestId: %s, error: %v", request.ProviderId, requestId, err))
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	// VALIDATE SHIPPING PROVIDER
	if _, err = u.shippingService.GetShippingProviderById(ctx, requestId, &shippingRpc.GetShippingProviderByIdRequest{
		Id: request.ShippingProviderId,
	}); err != nil {
		u.logger.Error(fmt.Sprintf("shipping provider id not found, shipping_provider_id: %s ,requestId: %s, error: %v", request.ShippingProviderId, requestId, err))
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

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

	user, err := u.userService.AuthUserFindUserByToken(ctx, requestId, &userRpc.AuthUserFindUserByTokenRequest{
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
		UserID:       request.UserId,
		PaymentItems: []*orm.PaymentItem{},
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

		// Send to topic product update for updating product stock
		if err = u.kafkaInfrastructure.Publish(ctx, config.Get().BrokerKafkaTopicProducts.ProductUpdated, product.Id, kafka.PROTOBUF_SCHEMA, product); err != nil {
			u.logger.Error(fmt.Sprintf("failed to publish product updated event: %v", err))
			return nil, fmt.Errorf("failed to publish product updated event: %w", err)
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
	// Publish to Shipping Created
	if err = u.kafkaInfrastructure.Publish(ctx, config.Get().BrokerKafkaTopicShippings.ShippingCreated, uuid.NewString(), kafka.PROTOBUF_SCHEMA, &shippingRpc.CreateShippingRequest{
		UserId:             user.Data.User.Id,
		PaymentId:          payment.ID,
		ShippingProviderId: request.ShippingProviderId,
	}); err != nil {
		u.logger.Error(fmt.Sprintf("Failed to publish CreateShipping request, requestId: %s, error: %v", requestId, err))
	}

	// Publish to Notification Payment Order Created
	if err = u.kafkaInfrastructure.Publish(ctx, config.Get().BrokerKafkaTopicNotifications.PaymentOrderCreated, payment.ID, kafka.PROTOBUF_SCHEMA, &notificationPb.SendEmailPaymentOrderCreateRequest{
		Email:            user.Data.User.Email,
		PaymentId:        payment.ID,
		NotificationType: notificationPb.NotificationTypeEnum_NOTIFICATION_EMAIL_PAYMENT_ORDER_CREATED,
	}); err != nil {
		u.logger.Error(fmt.Sprintf("Failed to publish SendEmailPaymentOrderCreateRequest request, requestId: %s, error: %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Append Event
	payload, err := proto.Marshal(payment.ToProto())
	if err != nil {
		u.logger.Error("PaymentUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	// SENT TO EVENT STORE
	if err = u.eventUseCase.AppendEvent(ctx, &pbEvent.Event{
		XId:           primitive.NewObjectID().Hex(),
		AggregateId:   payment.ID,
		AggregateType: "payments", // TODO: Move To Enum
		EventType:     config.Get().BrokerKafkaTopicPayments.PaymentOrderCreated,
		Version:       1,
		Timestamp:     timestamppb.New(now),
		SagaId:        requestId,
		Payload:       payload,
	}); err != nil {
		u.logger.Error("PaymentUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
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

func (u *paymentUseCase) ConfirmCreatePayment(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error {
	var (
		err error
	)
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentUseCase.ConfirmCreatePayment")
	defer func() {
		if err != nil {
			// TODO: Publish to compensate Topics
			span.RecordError(err)
		}
		span.End()
	}()

	savedEvent, err := u.eventMongoDBRepository.FindEventBySagaIDAndAggregateType(ctx, req.SagaId, "payments")
	if err != nil {
		u.logger.Error("PaymentUseCase.ConfirmCreatePayment", zap.String("requestId", requestId), zap.Error(err))
		return err
	}

	var payment paymentPb.Payment
	if err = proto.Unmarshal(savedEvent.Payload, &payment); err != nil {
		u.logger.Error(fmt.Sprintf("Failed to unmarshal event: %s", err.Error()))
		return err
	}
	if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkPgPayment.Payments, payment.Id, kafka.JSON_SCHEMA, orm.PaymentFromProto(&payment)); err != nil {
		u.logger.Error(fmt.Sprintf("Error publishing event to kafka for payment creation: %s", err.Error()))
		return err
	}

	var (
		payload []byte
	)
	for _, item := range payment.Items {
		payload, err = json.Marshal(item)
		if err != nil {
			u.logger.Error(fmt.Sprintf("Failed to marshal event: %s", err.Error()))
			return err
		}
		if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkPgPayment.PaymentItems, item.Id, kafka.JSON_SCHEMA, payload); err != nil {
			u.logger.Error(fmt.Sprintf("Error publishing event to kafka for payment item creation: %s", err.Error()))
			return err
		}
	}

	// 1. Publish to topic product updated confirm
	if err = u.kafkaInfrastructure.Publish(ctx, config.Get().BrokerKafkaTopicProducts.ConfirmProductUpdated, req.SagaId, kafka.PROTOBUF_SCHEMA, &pbEvent.ReserveEvent{
		SagaId:        req.SagaId,
		AggregateType: "products",
	}); err != nil {
		u.logger.Error(fmt.Sprintf("Error publishing event to kafka for reserve event: %s", err.Error()))
		return err
	}

	// 2. Publish to topic shipping created confirm
	if err = u.kafkaInfrastructure.Publish(ctx, config.Get().BrokerKafkaTopicShippings.ConfirmShippingCreated, req.SagaId, kafka.PROTOBUF_SCHEMA, &pbEvent.ReserveEvent{
		SagaId:        req.SagaId,
		AggregateType: "shippings",
	}); err != nil {
		u.logger.Error(fmt.Sprintf("Error publishing event to kafka for reserve event: %s", err.Error()))
		return err
	}

	return nil
}

func (u *paymentUseCase) CompensateCreatePayment(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error {
	var (
		err error
	)
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentUseCase.CompensateCreatePayment")
	defer func() {
		if err != nil {
			// TODO: Publish to compensate Topics
			span.RecordError(err)
		}
		span.End()
	}()

	if err = u.eventMongoDBRepository.DeleteEventBySagaId(ctx, req.SagaId); err != nil && err != mongo.ErrNoDocuments {
		u.logger.Error("PaymentUseCase.CompensateCreatePayment", zap.Error(err))
		return err
	}

	// 1. Publish to topic product updated confirm
	if err = u.kafkaInfrastructure.Publish(ctx, config.Get().BrokerKafkaTopicProducts.CompensateProductUpdated, req.SagaId, kafka.PROTOBUF_SCHEMA, &pbEvent.ReserveEvent{
		SagaId:        req.SagaId,
		AggregateType: "products",
	}); err != nil {
		u.logger.Error(fmt.Sprintf("Error publishing event to kafka for reserve event: %s", err.Error()))
		return err
	}

	// 2. Publish to topic shipping created confirm
	if err = u.kafkaInfrastructure.Publish(ctx, config.Get().BrokerKafkaTopicShippings.CompensateShippingCreated, req.SagaId, kafka.PROTOBUF_SCHEMA, &pbEvent.ReserveEvent{
		SagaId:        req.SagaId,
		AggregateType: "shippings",
	}); err != nil {
		u.logger.Error(fmt.Sprintf("Error publishing event to kafka for reserve event: %s", err.Error()))
		return err
	}

	return nil
}
