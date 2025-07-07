package usecase

import (
	"context"
	"fmt"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"google.golang.org/grpc/codes"

	"golang.org/x/crypto/bcrypt"
)

func (u *userUseCase) UpdateUserById(ctx context.Context, requestId string, req *userRpc.UpdateUserByIdRequest) (*userRpc.UpdateUserByIdResponse, error) {
	var (
		err error
		tx  = u.postgresSQLInfrastructure.GormDB().Begin()
	)
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "UserUseCase.UpdateUserById")
	defer span.End()

	user, err := u.userPostgresqlRepository.FindUserById(ctx, requestId, req.Id, tx)
	if err != nil {
		return nil, err
	}

	// Partial Update
	if req.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			tx.Rollback()
			u.logger.Error(fmt.Sprintf("requestId : %s , error hashing password: %v", requestId, err))
			return nil, err
		}
		newPassword := string(hashedPassword)
		req.Password = &newPassword
	}

	if req.Email != nil && *req.Email != user.Email {
		user.Email = *req.Email
	}

	if req.Name != nil && *req.Name != user.Name {
		user.Name = *req.Name
	}

	if req.IsVerified != nil && *req.IsVerified != user.IsVerified {
		user.IsVerified = *req.IsVerified
	}

	result, err := u.userPostgresqlRepository.UpdateUserById(ctx, requestId, user, tx)
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error updating user: %v", requestId, err))
		return nil, err
	}

	tx.Commit()
	return &userRpc.UpdateUserByIdResponse{
		Error:   "",
		Message: codes.OK.String(),
		Code:    uint32(codes.OK),
		Data: &userRpc.UpdateUserByIdResponse_UpdateUserByIdResponseData{
			Id: result.ID,
		},
	}, nil
}
