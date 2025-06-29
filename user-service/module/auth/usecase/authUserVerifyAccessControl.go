package usecase

import (
	"context"
	"errors"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (u *authUseCase) AuthUserVerifyAccessControl(ctx context.Context, requestId string, req *pb.AuthUserVerifyAccessControlRequest) (*pb.AuthUserVerifyAccessControlResponse, error) {
	tx := u.postgresSQL.GormDB.Begin()

	// Claim Token
	isValid := false
	authClaimedToken, err := u.AuthUserFindUserByToken(ctx, requestId, &pb.AuthUserFindUserByTokenRequest{
		Token: req.GetToken(),
	})
	if err != nil {
		tx.Rollback()
		u.logger.Error("AuthUseCase.AuthUserVerifyAccessControl", zap.String("requestId", requestId), zap.Error(err))
		return nil, err
	}

	if req.FullMethodName != nil {
		// Search cache If Exists
		isValid, err = u.accessControlRedisRepository.GetAccessControlRPC(ctx, requestId, authClaimedToken.Role.String(), req.GetFullMethodName())
		if err != nil && !errors.Is(err, redis.Nil) {
			tx.Rollback()
			u.logger.Error("AuthUseCase.AuthUserVerifyAccessControl", zap.String("requestId", requestId), zap.Error(err))
			return nil, status.Error(codes.Internal, "Internal Error")
		}

		if !isValid {
			// Search on Repository
			acl, err := u.accessControlPostgresqlRepository.FindAccessControlByRoleIdAndFullMethodName(ctx, requestId, authClaimedToken.Role.Id, req.GetFullMethodName(), tx)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				tx.Rollback()
				u.logger.Error("AuthUseCase.AuthUserVerifyAccessControl", zap.String("requestId", requestId), zap.Error(err))
				return nil, status.Error(codes.Internal, "Internal Error")
			}

			if acl != nil {
				isValid = true
			}
		}

	}

	if req.HttpMethod != nil && req.HttpUrl != nil {
		// Search cache If Exists
		isValid, err = u.accessControlRedisRepository.GetAccessControlHTTP(ctx, requestId, authClaimedToken.Role.String(), req.GetHttpMethod(), req.GetHttpUrl())
		if err != nil && !errors.Is(err, redis.Nil) {
			tx.Rollback()
			u.logger.Error("AuthUseCase.AuthUserVerifyAccessControl", zap.String("requestId", requestId), zap.Error(err))
			return nil, status.Error(codes.Internal, "Internal Error")
		}

		if !isValid {
			// Search on Repository
			acl, err := u.accessControlPostgresqlRepository.FindAccessControlByRoleIdAndHttpMethodAndHttpUrl(ctx, requestId, authClaimedToken.Role.Id, req.GetHttpMethod(), req.GetHttpUrl(), tx)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				tx.Rollback()
				u.logger.Error("AuthUseCase.AuthUserVerifyAccessControl", zap.String("requestId", requestId), zap.Error(err))
				return nil, status.Error(codes.Internal, "Internal Error")
			}

			if acl != nil {
				isValid = true
			}
		}
	}

	tx.Commit()
	return &pb.AuthUserVerifyAccessControlResponse{
		IsValid:        isValid,
		User:           nil,
		Role:           nil,
		AccessControls: nil,
	}, nil

}
