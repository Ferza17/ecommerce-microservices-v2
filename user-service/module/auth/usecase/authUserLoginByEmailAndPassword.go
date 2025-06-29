package usecase

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (u *authUseCase) AuthUserLoginByEmailAndPassword(ctx context.Context, requestId string, req *pb.AuthUserLoginByEmailAndPasswordRequest) (*pb.AuthUserLoginByEmailAndPasswordResponse, error) {
	tx := u.postgresSQL.GormDB.Begin()
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.AuthUserLoginByEmailAndPassword")
	defer span.End()

	user, err := u.userPostgresqlRepository.FindUserByEmail(ctx, requestId, req.Email, tx)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error("AuthUseCase.AuthUserLoginByEmailAndPassword", zap.String("requestId", requestId), zap.Error(errors.New("user not found")))
			return nil, status.Error(codes.NotFound, err.Error())
		}

		u.logger.Error(fmt.Sprintf("requestId : %s , error finding user by email : %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	reqHashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		u.logger.Error("AuthUseCase.AuthUserLoginByEmailAndPassword", zap.String("requestId", requestId), zap.Error(errors.New("error hashing password")))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = bcrypt.CompareHashAndPassword(reqHashedPassword, []byte(req.Password)); err != nil {
		tx.Rollback()
		u.logger.Error("AuthUseCase.AuthUserLoginByEmailAndPassword", zap.String("requestId", requestId), zap.Error(errors.New("error comparing password")))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Generate Access & Refresh Token
	accessToken, refreshToken, err := u.GenerateAccessToken(ctx, requestId, user)
	if err != nil {
		tx.Rollback()
		u.logger.Error("AuthUseCase.AuthUserLoginByEmailAndPassword", zap.String("requestId", requestId), zap.Error(errors.New("error generating token")))
		return nil, err
	}

	// Build Response
	resp := &pb.AuthUserLoginByEmailAndPasswordResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	tx.Commit()
	return resp, nil
}
