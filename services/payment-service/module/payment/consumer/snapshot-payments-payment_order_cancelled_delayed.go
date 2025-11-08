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

func (c *paymentConsumer) SnapshotPaymentsPaymentOrderCancelledDelayed(ctx context.Context, message *pbEvent.EventEnvelope) error {
	var (
		request   pb.PaymentOrderDelayedCancelledRequest
		requestId = pkgContext.GetRequestIDFromContext(ctx)
	)

	if err := util.Base64URLToProtobuf(message.Payload, &request); err != nil {
		c.logger.Info(fmt.Sprintf("util.Base64URLToProtobuf: %v", err))
		return err
	}

	if err := c.paymentUseCase.PaymentOrderDelayedCancelled(ctx, requestId, &request); err != nil {
		c.logger.Error("Payment Order Delayed Cancel Failed", zap.Error(err))
		return err
	}

	return nil
}
