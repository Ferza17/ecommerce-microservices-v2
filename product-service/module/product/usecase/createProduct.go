package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	eventRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/event/v1"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/product/v1"

	"github.com/ferza17/ecommerce-microservices-v2/product-service/util"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (u *productUseCase) CreateProduct(ctx context.Context, requestId string, req *productRpc.CreateProductRequest) (*productRpc.CreateProductResponse, error) {
	var (
		ctxTimeout, cancel = context.WithTimeout(ctx, 5*time.Second)
		err                error
		tx                 = u.productPgsqlRepository.OpenTransactionWithContext(ctxTimeout)
		now                = time.Now().UTC()
		eventStore         = &eventRpc.EventStore{
			RequestId:     requestId,
			Service:       config.Get().ServiceName,
			EventType:     config.Get().QueueProductCreated,
			Status:        config.Get().CommonSagaStatusSuccess,
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
	)
	ctxTimeout, span := u.telemetryInfrastructure.Tracer(ctxTimeout, "UseCase.CreateProduct")

	defer func(err error, eventStore *eventRpc.EventStore) {
		defer cancel()
		defer span.End()
		payload, err := util.ConvertStructToProtoStruct(req)
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

		if err = u.rabbitmqInfrastructure.Publish(ctxTimeout, requestId, config.Get().ExchangeEvent, config.Get().QueueEventCreated, eventStoreMessage); err != nil {
			u.logger.Error(fmt.Sprintf("error creating product event store: %s", err.Error()))
			return
		}
	}(err, eventStore)

	result, err := u.productPgsqlRepository.CreateProduct(ctxTimeout, &orm.Product{
		ID:          uuid.NewString(),
		Name:        req.GetName(),
		Price:       req.GetPrice(),
		Stock:       req.GetStock(),
		Description: req.GetDescription(),
		Image:       req.GetImage(),
		Uom:         req.GetUom(),
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}, tx)
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error creating product: %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	tx.Commit()
	return &productRpc.CreateProductResponse{
		Id: result,
	}, nil
}
