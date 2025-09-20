package consumer

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	pb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (c *paymentConsumer) SnapshotPaymentsPaymentOrderCreated(ctx context.Context, message *kafka.Message) error {
	var (
		request   pb.CreatePaymentRequest
		requestId = pkgContext.GetRequestIDFromContext(ctx)
	)

	if err := proto.Unmarshal(message.Value, &request); err != nil {
		c.logger.Info(fmt.Sprintf("proto.Unmarshal: %v", err))
		return err
	}

	if _, err := c.paymentUseCase.CreatePayment(ctx, requestId, &request); err != nil {
		c.logger.Error("Payment Order Create Failed", zap.Error(err))
		return err
	}

	return nil
}
