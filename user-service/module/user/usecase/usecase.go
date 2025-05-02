package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgresql"
	userEventStoreUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/userEventStore/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
)

type (
	IUserUseCase interface {
		CreateUser(ctx context.Context, requestId string, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
		FindUserById(ctx context.Context, requestId string, req *pb.FindUserByIdRequest) (*pb.User, error)
		FindUserByEmailAndPassword(ctx context.Context, requestId string, req *pb.FindUserByEmailAndPasswordRequest) (*pb.User, error)
	}

	userUseCase struct {
		userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository
		userEventStoreUseCase    userEventStoreUseCase.IUserEventStoreUseCase
		logger                   pkg.IZapLogger
	}
)

func NewUserUseCase(userPostgresqlRepository userPostgresqlRepository.IUserPostgresqlRepository, userEventStoreUseCase userEventStoreUseCase.IUserEventStoreUseCase, logger pkg.IZapLogger) IUserUseCase {
	return &userUseCase{
		userPostgresqlRepository: userPostgresqlRepository,
		userEventStoreUseCase:    userEventStoreUseCase,
		logger:                   logger,
	}
}
