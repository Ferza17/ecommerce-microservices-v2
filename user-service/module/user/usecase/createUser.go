package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/utils"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (u *userUseCase) CreateUser(ctx context.Context, requestId string, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	tx := u.userPostgresqlRepository.OpenTransactionWithContext(ctx)
	now := time.Now().UTC()
	createUserEventStoreReq := &pb.CreateUserEventStoreRequest{
		SagaId:    requestId,
		Entity:    "user",
		EventType: enum.USER_CREATED.String(),
		Status:    enum.SUCCESS.String(),
		Payload: &pb.UserState{
			Name:      &req.Name,
			Email:     &req.Email,
			CreatedAt: timestamppb.New(now),
			UpdatedAt: timestamppb.New(now),
		},
	}

	hashedPassword, err := utils.Hashed(req.Password)
	if err != nil {
		defer tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error hashing password: %v", requestId, err))
		createUserEventStoreReq.Status = enum.FAILED.String()
		if _, err = u.userEventStoreUseCase.CreateUserEventStore(ctx, requestId, createUserEventStoreReq); err != nil {
			u.logger.Error(fmt.Sprintf("requestId : %s , error creating userEventStore event: %v", requestId, err))
		}
		return nil, err
	}
	createUserEventStoreReq.Payload.Password = &hashedPassword

	result, err := u.userPostgresqlRepository.CreateUserWithTransaction(ctx, requestId, &orm.User{
		ID:          uuid.NewString(),
		Name:        req.Name,
		Email:       req.Email,
		Password:    hashedPassword,
		CreatedAt:   &now,
		UpdatedAt:   &now,
		DiscardedAt: nil,
	}, tx)

	if err != nil {
		defer tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error creating user: %v", requestId, err))
		createUserEventStoreReq.Status = enum.FAILED.String()
		if _, err = u.userEventStoreUseCase.CreateUserEventStore(ctx, requestId, createUserEventStoreReq); err != nil {
			u.logger.Error(fmt.Sprintf("requestId : %s , error creating userEventStore event: %v", requestId, err))
		}
		return nil, err
	}

	createUserEventStoreReq.EntityId = result
	if _, err = u.userEventStoreUseCase.CreateUserEventStore(ctx, requestId, createUserEventStoreReq); err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error creating userEventStore event: %v", requestId, err))
		return nil, err
	}

	tx.Commit()
	return &pb.CreateUserResponse{
		Id: result,
	}, nil
}
