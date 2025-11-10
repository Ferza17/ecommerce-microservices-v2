package usecase

import (
	"context"

	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/kafka"
	pbProduct "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/util"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkProduct.PgProducts, existingProduct.ID, kafka.JSON_SCHEMA, existingProduct); err != nil {
		u.logger.Error("ProductUseCase.UpdateProductById", zap.String("requestId", requestId), zap.Error(err))
		return nil, err
	}

	if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkProduct.EsProducts, existingProduct.ID, kafka.JSON_SCHEMA, existingProduct); err != nil {
		u.logger.Error("ProductUseCase.UpdateProductById", zap.String("requestId", requestId), zap.Error(err))
		return nil, err
	}

	return &empty.Empty{}, nil
}
