package usecase

import (
	"context"

	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/kafka"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/event"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *eventUseCase) AppendEventEnvelope(ctx context.Context, request *pbEvent.EventEnvelope) error {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "EventUseCase.AppendEvent")
	defer span.End()

	metadata := map[string]string{}
	if pkgContext.GetRequestIDFromContext(ctx) != "" {
		metadata[pkgContext.CtxKeyRequestID] = pkgContext.GetRequestIDFromContext(ctx)
	}
	if pkgContext.GetTokenAuthorizationFromContext(ctx) != "" {
		metadata[pkgContext.CtxKeyAuthorization] = pkgContext.GetTokenAuthorizationFromContext(ctx)
	}
	carrier := u.telemetryInfrastructure.InjectSpanToTextMapPropagator(ctx)
	for k, v := range carrier {
		metadata[k] = v
	}
	request.Metadata = metadata

	// Send to an Event Store
	if err := u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkMongoEvent.EventEnvelopes, request.XId, kafka.PROTOBUF_SCHEMA, request); err != nil {
		u.logger.Error("eventUseCase.AppendEvent", zap.String("requestId", request.CorrelationId), zap.Error(err))
		return status.Error(codes.Internal, "internal server error")
	}
	return nil
}
