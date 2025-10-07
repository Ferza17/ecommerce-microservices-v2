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
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *productUseCase) CreateProduct(ctx context.Context, requestId string, req *pbProduct.CreateProductRequest) (*empty.Empty, error) {
	var (
		err error
	)

	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "ProductUseCase.CreateProduct")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	now, err := util.GetNowWithTimeZone(pkgContext.CtxValueAsiaJakarta)
	if err != nil {
		u.logger.Error("error getting now with timezone ", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	product := &orm.Product{
		ID:          uuid.NewString(),
		Name:        req.GetName(),
		Price:       req.GetPrice(),
		Stock:       int64(req.Stock),
		Description: req.GetDescription(),
		Image:       req.GetImage(),
		Uom:         req.GetUom(),
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	// SENT TO EVENT STORE
	payload, err := proto.Marshal(product.ToProto())
	if err != nil {
		u.logger.Error("ProductUseCase.CreateProduct", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	if err = u.eventUseCase.AppendEvent(ctx, &pbEvent.Event{
		XId:           primitive.NewObjectID().Hex(),
		AggregateId:   product.ID,
		AggregateType: "products", // TODO: Move To Enum
		EventType:     config.Get().BrokerKafkaTopicProducts.ProductCreated,
		Version:       1,
		Timestamp:     timestamppb.New(now),
		SagaId:        requestId,
		Payload:       payload,
	}); err != nil {
		u.logger.Error("ProductUseCase.CreateProduct", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &empty.Empty{}, nil
}

func (u *productUseCase) ConfirmCreateProduct(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error {
	var (
		err error
	)
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "ProductUseCase.ConfirmCreateProduct")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	savedEvent, err := u.eventMongoDBRepository.FindEventBySagaID(ctx, req.SagaId)
	if err != nil {
		u.logger.Error("ProductUseCase.ConfirmCreateProduct", zap.String("requestId", requestId), zap.Error(err))
		return err
	}

	var product pbProduct.Product
	if err = proto.Unmarshal(savedEvent.Payload, &product); err != nil {
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

	return nil
}

func (u *productUseCase) CompensateCreateProduct(ctx context.Context, requestId string, req *pbEvent.ReserveEvent) error {
	var (
		err error
	)
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "ProductUseCase.CompensateCreateProduct")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	if err = u.eventMongoDBRepository.DeleteEventBySagaId(ctx, req.SagaId); err != nil {
		u.logger.Error("ProductUseCase.CompensateCreateProduct", zap.String("requestId", requestId), zap.Error(err))
		return err
	}

	//TODO: PUBLISH TO COMPENSATE EVENT THAT PUBLISH to TOPIC snapshot-product-product_created

	return nil
}
