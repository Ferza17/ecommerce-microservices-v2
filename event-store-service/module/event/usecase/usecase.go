package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/pb"
	eventRepository "github.com/ferza17/ecommerce-microservices-v2/event-store-service/module/event/repository/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg"
)

type (
	IEventUseCase interface {
		CreateEventStore(ctx context.Context, requestId string, req *pb.EventStore) (*pb.CreateEventStoreResponse, error)
	}

	eventUseCase struct {
		eventRepository eventRepository.IEventRepository
		logger          pkg.IZapLogger
	}
)

func NewEventStoreUseCase(eventStoreRepository eventRepository.IEventRepository, logger pkg.IZapLogger) IEventUseCase {
	return &eventUseCase{
		eventRepository: eventStoreRepository,
		logger:          logger,
	}
}
