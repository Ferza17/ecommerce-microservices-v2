package usecase

import (
	"context"
	"errors"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/token"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (u *authUseCase) AuthUserVerifyAccessControl(ctx context.Context, requestId string, req *pb.AuthUserVerifyAccessControlRequest) (*pb.AuthUserVerifyAccessControlResponse, error) {
	var (
		isValid = false
		tx      = u.postgresSQL.GormDB.Begin()
	)

	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "AuthUseCase.AuthUserVerifyAccessControl")
	defer span.End()

	// Claim Token
	claim, err := token.ValidateJWTToken(req.Token, token.DefaultRefreshTokenConfig())
	if err != nil {
		tx.Rollback()
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	user, err := u.userPostgresqlRepository.FindUserById(ctx, requestId, claim.UserId, tx)
	if err != nil {
		tx.Rollback()
		u.logger.Error("AuthUseCase.AuthUserVerifyAccessControl", zap.String("requestId", requestId), zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.Unauthenticated, "user not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	if req.FullMethodName != nil {
		isValid, err = u.accessControlUseCase.IsHasRPCAccess(ctx, requestId, user.Role.Role, *req.FullMethodName)
		if err != nil {
			tx.Rollback()
			u.logger.Error("AuthUseCase.IsHasRPCAccess", zap.String("requestId", requestId), zap.Error(err))
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
	}

	if req.HttpMethod != nil && req.HttpUrl != nil {
		isValid, err = u.accessControlUseCase.IsHasHTTPAccess(ctx, requestId, user.Role.Role, *req.HttpMethod, *req.HttpUrl)
		if err != nil {
			tx.Rollback()
			u.logger.Error("AuthUseCase.IsHasRPCAccess", zap.String("requestId", requestId), zap.Error(err))
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
	}

	tx.Commit()
	return &pb.AuthUserVerifyAccessControlResponse{
		IsValid: isValid,
		User:    user.ToProto(),
	}, nil

}
