package usecase

import (
	"context"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *eventUseCase) AppendEvent(ctx context.Context, request *pbEvent.Event) error {
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

	existingEvent, err := u.eventMongoDBRepository.FindEventByAggregateIDAndAggregateType(ctx, request.AggregateId, request.AggregateType)
	if err != nil && err != mongo.ErrNoDocuments {
		u.logger.Error("eventUseCase.AppendEvent", zap.String("requestId", request.SagaId), zap.Error(err))
		return status.Error(codes.Internal, "internal server error")
	}

	// DEFINE VERSION
	if existingEvent != nil {
		request.Version = existingEvent.Version + 1
	} else {
		request.Version = 1
	}

	// Send to an Event Store
	if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkMongoEvent.User, request.XId, kafka.PROTOBUF_SCHEMA, request); err != nil {
		u.logger.Error("eventUseCase.AppendEvent", zap.String("requestId", request.SagaId), zap.Error(err))
		return status.Error(codes.Internal, "internal server error")
	}
	return nil
}
