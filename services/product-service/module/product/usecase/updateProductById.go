package usecase

import (
	"context"

	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/event"
	pbProduct "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/util"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *productUseCase) UpdateProductById(ctx context.Context, requestId string, req *pbProduct.UpdateProductByIdRequest) (*empty.Empty, error) {
	var (
		err error
	)
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "ProductUseCase.UpdateProductById")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	now, err := util.GetNowWithTimeZone(pkgContext.CtxValueAsiaJakarta)
	if err != nil {
		u.logger.Error("ProductUseCase.UpdateProductById", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	existingProduct, err := u.productPgsqlRepository.FindProductById(ctx, req.Id, nil)
	if err != nil {
		u.logger.Error("ProductUseCase.UpdateProductById", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.NotFound, "Product not found")
	}

	// Partial UPDATE
	if req.Name != nil && existingProduct.Name != *req.Name {
		existingProduct.Name = *req.Name
	}

	if req.Description != nil && existingProduct.Description != *req.Description {
		existingProduct.Description = *req.Description
	}

	if req.Image != nil && existingProduct.Image != *req.Image {
		existingProduct.Image = *req.Image
	}

	if req.Uom != nil && existingProduct.Uom != *req.Uom {
		existingProduct.Uom = *req.Uom
	}

	if req.Stock != nil && existingProduct.Stock != int64(*req.Stock) {
		existingProduct.Stock = int64(*req.Stock)
	}

	if req.Price != nil && existingProduct.Price != *req.Price {
		existingProduct.Price = *req.Price
	}
	existingProduct.UpdatedAt = &now

	// SENT TO EVENT STORE
	payload, err := proto.Marshal(existingProduct.ToProto())
	if err != nil {
		u.logger.Error("ProductUseCase.UpdateProductById", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	if err = u.eventUseCase.AppendEvent(ctx, &pbEvent.Event{
		XId:           primitive.NewObjectID().Hex(),
		AggregateId:   existingProduct.ID,
		AggregateType: "products", // TODO: Move To Enum
		EventType:     config.Get().BrokerKafkaTopicProducts.ProductUpdated,
		Version:       1,
		Timestamp:     timestamppb.New(now),
		SagaId:        requestId,
		Payload:       payload,
	}); err != nil {
		u.logger.Error("ProductUseCase.UpdateProductById", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &empty.Empty{}, nil
}

func (u *productUseCase) ConfirmUpdateProductById(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error {
	var (
		err error
	)
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "ProductUseCase.ConfirmUpdateProductById")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	savedEvents, err := u.eventMongoDBRepository.FindEventsBySagaID(ctx, req.SagaId)
	if err != nil {
		u.logger.Error("ProductUseCase.ConfirmCreateProduct", zap.String("requestId", requestId), zap.Error(err))
		return err
	}

	for _, event := range savedEvents {
		var product pbProduct.Product
		if err = proto.Unmarshal(event.Payload, &product); err != nil {
			u.logger.Error("ProductUseCase.ConfirmCreateProduct", zap.String("requestId", requestId), zap.Error(err))
			return err
		}

		if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkProduct.PgProducts, product.Id, kafka.JSON_SCHEMA, orm.ProductFromProto(&product)); err != nil {
			u.logger.Error("ProductUseCase.ConfirmCreateProduct", zap.String("requestId", requestId), zap.Error(err))
			return err
		}

		if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkProduct.EsProducts, product.Id, kafka.JSON_SCHEMA, orm.ProductFromProto(&product)); err != nil {
			u.logger.Error("ProductUseCase.ConfirmCreateProduct", zap.String("requestId", requestId), zap.Error(err))
			return err
		}
	}

	return nil
}

func (u *productUseCase) CompensateUpdateProductById(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error {
	var (
		err error
	)
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "ProductUseCase.CompensateUpdateProductById")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	if err = u.eventMongoDBRepository.DeleteEventBySagaId(ctx, req.SagaId); err != nil {
		u.logger.Error("ProductUseCase.CompensateUpdateProductById", zap.String("requestId", requestId), zap.Error(err))
		return err
	}

	//TODO: PUBLISH TO COMPENSATE EVENT THAT PUBLISH to TOPIC snapshot-product-product_updated

	return nil
}
