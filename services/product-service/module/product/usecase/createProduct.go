package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/util"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *productUseCase) CreateProduct(ctx context.Context, requestId string, req *productRpc.CreateProductRequest) (*empty.Empty, error) {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "ProductUseCase.CreateProduct")
	defer span.End()

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

	if err = u.kafkaInfrastructure.PublishWithJsonSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkProduct.PgProducts, product.ID, product); err != nil {
		u.logger.Error(fmt.Sprintf("Error publishing event to kafka for product creation: %s", err.Error()))
		return nil, status.Errorf(codes.Internal, "Error publishing event to kafka for product creation: %s", err.Error())
	}

	if err = u.kafkaInfrastructure.PublishWithJsonSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkProduct.EsProducts, product.ID, product); err != nil {
		u.logger.Error(fmt.Sprintf("Error publishing event to kafka for product creation: %s", err.Error()))
		return nil, status.Errorf(codes.Internal, "Error publishing event to kafka for product creation: %s", err.Error())
	}

	return &empty.Empty{}, nil
}
