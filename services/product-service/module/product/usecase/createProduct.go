package usecase

import (
	"context"

	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	pbProduct "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/util"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkProduct.PgProducts, product.ID, kafka.JSON_SCHEMA, product); err != nil {
		u.logger.Error("ProductUseCase.CreateProduct", zap.String("requestId", requestId), zap.Error(err))
		return nil, err
	}

	if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkProduct.EsProducts, product.ID, kafka.JSON_SCHEMA, product); err != nil {
		u.logger.Error("ProductUseCase.CreateProduct", zap.String("requestId", requestId), zap.Error(err))
		return nil, err
	}

	return &empty.Empty{}, nil
}
