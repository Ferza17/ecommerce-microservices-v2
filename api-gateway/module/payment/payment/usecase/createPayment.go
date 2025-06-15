package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/config"
	eventRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/event/v1"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/payment/v1"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/product/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *paymentUseCase) CretePayment(ctx context.Context, requestId string, request *paymentRpc.CreatePaymentRequest) error {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "PaymentUseCase.CretePayment")
	defer span.End()

	var (
		err        error = nil
		eventStore       = &eventRpc.EventStore{
			RequestId:     requestId,
			Service:       config.Get().ProductServiceName,
			EventType:     config.Get().QueueProductCreated,
			Status:        config.Get().CommonSagaStatusPending,
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
		totalPrice   = 0.0
		mapProductID = make(map[string]bool)
		productIds   = make([]string, 0)
	)

	// Sent Event after finish executing function
	defer func(err error, eventStore *eventRpc.EventStore) {
		if err != nil {
			eventStore.Status = config.Get().CommonSagaStatusFailed
		}

		eventStoreMessage, err := proto.Marshal(eventStore)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
			return
		}

		if err = u.rabbitMQ.Publish(ctx, requestId, config.Get().ExchangeEvent, config.Get().QueueEventCreated, eventStoreMessage); err != nil {
			u.logger.Error(fmt.Sprintf("error creating product event store: %s", err.Error()))
			return
		}
	}(err, eventStore)

	for _, item := range request.Items {
		mapProductID[item.ProductId] = true
		productIds = append(productIds, item.ProductId)
	}

	productsResp, err := u.productSvc.FindProductsWithPagination(ctx, requestId, &productRpc.FindProductsWithPaginationRequest{
		Ids:   productIds,
		Page:  1,
		Limit: int32(len(request.Items)),
	})
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding products : %v", err))
		return err
	}

	for _, product := range productsResp.Data {
		if _, ok := mapProductID[product.Id]; !ok {
			u.logger.Error(fmt.Sprintf("product not found : %s", product.Id))
			return fmt.Errorf("product not found : %s", product.Id)
		}
		totalPrice += product.Price
	}

	// Validate Provider Id
	provider, err := u.providerSvc.FindPaymentProviderById(ctx, requestId, &paymentRpc.FindPaymentProviderByIdRequest{
		Id: request.ProviderId,
	})
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding payment provider by id : %v", err))
		return err
	}

	if provider == nil {
		u.logger.Error(fmt.Sprintf("payment provider not found : %s", request.ProviderId))
		return fmt.Errorf("payment provider not found : %s", request.ProviderId)
	}

	// Publish to Queue.Payment.Order.Created
	request.Amount = totalPrice
	message, err := proto.Marshal(request)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		return err
	}

	if err = u.rabbitMQ.Publish(ctx, requestId, config.Get().ExchangePayment, config.Get().QueuePaymentOrderCreated, message); err != nil {
		u.logger.Error(fmt.Sprintf("error publishing message to rabbitmq: %s", err.Error()))
		return err
	}

	return nil
}
