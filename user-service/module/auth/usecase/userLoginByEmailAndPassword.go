package usecase

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (u *authUseCase) UserLoginByEmailAndPassword(ctx context.Context, requestId string, req *pb.UserLoginByEmailAndPasswordRequest) (*pb.UserLoginByEmailAndPasswordResponse, error) {
	var (
		err        error
		tx         = u.userPostgresqlRepository.OpenTransactionWithContext(ctx)
		now        = time.Now().UTC()
		eventStore = &pb.EventStore{
			RequestId:     requestId,
			Service:       enum.UserService.String(),
			EventType:     enum.USER_LOGIN.String(),
			Status:        enum.SUCCESS.String(),
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
	)

	defer func(err error, eventStore *pb.EventStore) {
		payload, err := util.ConvertStructToProtoStruct(req)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error converting struct to proto struct: %s", err.Error()))
		}
		eventStore.Payload = payload

		eventStoreMessage, err := proto.Marshal(eventStore)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		}

		if err != nil {
			tx.Rollback()
			eventStore.Status = enum.FAILED.String()
		}

		if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, enum.EventExchange, enum.EVENT_CREATED, eventStoreMessage); err != nil {
			u.logger.Error(fmt.Sprintf("error creating product event store: %s", err.Error()))
			return
		}
		tx.Commit()
	}(err, eventStore)

	user, err := u.userPostgresqlRepository.FindUserByEmailWithTransaction(ctx, requestId, req.Email, tx)
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error finding user by email and password: %v", requestId, err))
		return nil, err
	}

	reqHashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error hashing password: %v", requestId, err))
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword(reqHashedPassword, []byte(req.Password)); err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error comparing password: %v", requestId, err))
		return nil, err
	}

	accessToken, err := pkg.GenerateToken(pkg.Claim{
		UserID:    user.ID,
		CreatedAt: &now,
		StandardClaims: jwt.StandardClaims{
			Audience:  enum.UserService.String(),
			ExpiresAt: now.Add(config.Get().JwtAccessTokenExpirationTime).Unix(),
		},
	}, config.Get().JwtAccessTokenSecret)
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error generating token: %v", requestId, err))
		return nil, err
	}

	refreshToken, err := pkg.GenerateToken(pkg.Claim{
		UserID:    user.ID,
		CreatedAt: &now,
		StandardClaims: jwt.StandardClaims{
			Audience:  enum.UserService.String(),
			ExpiresAt: now.Add(config.Get().JwtRefreshTokenExpirationTime).Unix(),
		},
	}, config.Get().JwtRefreshTokenSecret)
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error generating refresh token: %v", requestId, err))
	}

	// TODO: Implement User
	//1. send Token (Send To Notification Service)
	//2. send Refresh Token (Send To Notification Service)
	//3. Send Email to Notification Service

	tx.Commit()
	return &pb.UserLoginByEmailAndPasswordResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}
