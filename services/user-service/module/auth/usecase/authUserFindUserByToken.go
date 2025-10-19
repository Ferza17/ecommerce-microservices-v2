package usecase

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/token"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *authUseCase) AuthUserFindUserByToken(ctx context.Context, requestId string, req *pb.AuthUserFindUserByTokenRequest) (*pb.AuthUserFindUserByTokenResponse, error) {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "AuthUseCase.AuthUserFindUserByToken")
	defer span.End()
	tx := u.postgresSQL.GormDB().Begin()
	claimedToken, err := token.ValidateJWTToken(req.Token, token.DefaultRefreshTokenConfig())
	if err != nil {
		tx.Rollback()
		u.logger.Error("AuthUseCase.FindUserByToken", zap.String("requestId", requestId), zap.Error(errors.New("error parsing token")))
		return nil, token.MapErrorToGrpcStatus(err)
	}

	user, err := u.userPostgresqlRepository.FindUserById(ctx, requestId, claimedToken.UserId, tx)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error("AuthUseCase.FindUserByToken", zap.String("requestId", requestId), zap.Error(errors.New("user not found")))
			return nil, status.Error(codes.NotFound, fmt.Sprintf("requestId : %s , user does not exist", requestId))
		}

		u.logger.Error("AuthUseCase.FindUserByToken", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Build Response
	tx.Commit()
	return &pb.AuthUserFindUserByTokenResponse{
		Status:  "success",
		Message: "FindUserByToken",
		Data: &pb.AuthUserFindUserByTokenResponse_AuthUserFindUserByTokenResponseData{
			User: user.ToProto(),
		},
	}, nil
}
