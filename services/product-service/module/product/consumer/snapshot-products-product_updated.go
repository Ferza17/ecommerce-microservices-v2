package consumer

import (
	"context"
	"fmt"

	pbEvent "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/event"
	pb "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/util"
	"go.uber.org/zap"
)

func (c *productConsumer) SnapshotProductsProductUpdated(ctx context.Context, message *pbEvent.EventEnvelope) error {
	var (
		request   pb.UpdateProductByIdRequest
		requestId = pkgContext.GetRequestIDFromContext(ctx)
	)

	if err := util.JSONToProto(message.Payload, &request); err != nil {
		c.logger.Info(fmt.Sprintf("util.JSONToProto: %v", err))
		return err
	}

	if _, err := c.productUseCase.UpdateProductById(ctx, requestId, &request); err != nil {
		c.logger.Error("Update Product Failed", zap.Error(err))
		return err
	}

	return nil
}
