package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/userEventStore/repository/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
)

type (
	IUserEventStoreUseCase interface {
		CreateUserEventStore(ctx context.Context, requestId string, req *pb.CreateUserEventStoreRequest) (*pb.CreateUserEventStoreResponse, error)
		CreateUserEventStoreWithTransaction(ctx context.Context, requestId string, req *pb.CreateUserEventStoreRequest) (*pb.CreateUserEventStoreResponse, error)
	}

	userEventStoreUseCase struct {
		userEventStoreRepository mongodb.IUserEventStoreRepository
		logger                   pkg.IZapLogger
	}
)

func NewUserEventStoreUseCase(userEventStoreRepository mongodb.IUserEventStoreRepository, logger pkg.IZapLogger) IUserEventStoreUseCase {
	return &userEventStoreUseCase{
		userEventStoreRepository: userEventStoreRepository,
		logger:                   logger,
	}
}
