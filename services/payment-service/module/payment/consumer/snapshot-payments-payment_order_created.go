package consumer

import (
	"context"
	"fmt"

	pbEvent "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/event"
	pb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/util"
	"go.uber.org/zap"
)

func (c *paymentConsumer) SnapshotPaymentsPaymentOrderCreated(ctx context.Context, message *pbEvent.EventEnvelope) error {
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

	if err = util.Base64URLToProtobuf(message.Payload, &request); err != nil {
		c.logger.Info(fmt.Sprintf("util.Base64URLToProtobuf: %v", err))
		return err
	}

	if _, err = c.paymentUseCase.CreatePayment(ctx, requestId, &request); err != nil {
		c.logger.Error("Payment Order Create Failed", zap.Error(err))
		return err
	}

	return nil
}
