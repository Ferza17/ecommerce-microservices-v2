package usecase

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (u *userUseCase) FindUserByEmail(ctx context.Context, requestId string, request *pb.FindUserByEmailRequest) (*pb.FindUserByEmailResponse, error) {
	var (
		tx = u.postgresSQLInfrastructure.GormDB().Begin()
	)
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "UserUseCase.FindUserByEmail")
	defer span.End()

	user, err := u.userPostgresqlRepository.FindUserByEmail(ctx, requestId, request.Email, tx)
	if err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			u.logger.Error("AuthUseCase.FindUserByEmail", zap.String("requestId", requestId), zap.Error(errors.New("user not found")))
			return nil, status.Error(codes.NotFound, err.Error())
		}

		u.logger.Error(fmt.Sprintf("requestId : %s , error finding user by email : %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	tx.Commit()
	return &pb.FindUserByEmailResponse{
		Status:  "success",
		Message: "FindUserByEmail",
		Data: &pb.FindUserByEmailResponse_FindUserByEmailResponseData{
			User: user.ToProto(),
		},
	}, nil
}
