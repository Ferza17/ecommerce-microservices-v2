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
		err       error
		requestId = pkgContext.GetRequestIDFromContext(ctx)
	)
	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "UserConsumer.FindUserByEmail")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	if err = proto.Unmarshal(message.Value, &request); err != nil {
		c.logger.Info(fmt.Sprintf("proto.Unmarshal: %v", err))
		return err
	}

	if _, err = c.paymentUseCase.CreatePayment(ctx, requestId, &request); err != nil {
		c.logger.Error("Payment Order Create Failed", zap.Error(err))
		return err
	}

	return nil
}

func (c *paymentConsumer) CompensateSnapshotPaymentsPaymentOrderCreated(ctx context.Context, message *kafka.Message) error {
	var (
		err error
	)
	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "UserConsumer.FindUserByEmail")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	return nil
}

func (c *paymentConsumer) ConfirmSnapshotPaymentsPaymentOrderCreated(ctx context.Context, message *kafka.Message) error {
	var (
		err error
	)
	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "UserConsumer.FindUserByEmail")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	return nil
}
