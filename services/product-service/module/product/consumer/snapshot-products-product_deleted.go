package consumer

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	pb "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (c *productConsumer) SnapshotProductsProductDeleted(ctx context.Context, message *kafka.Message) error {
	var (
		request   pb.DeleteProductByIdRequest
		requestId = pkgContext.GetRequestIDFromContext(ctx)
	)

	if err := proto.Unmarshal(message.Value, &request); err != nil {
		c.logger.Info(fmt.Sprintf("proto.Unmarshal: %v", err))
		return err
	}

	if _, err := c.productUseCase.DeleteProductById(ctx, requestId, &request); err != nil {
		c.logger.Error("Delete Product Failed", zap.Error(err))
		return err
	}

	return nil
}
