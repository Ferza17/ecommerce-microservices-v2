package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
)

func (u *userUseCase) FindUserById(ctx context.Context, requestId string, req *pb.FindUserByIdRequest) (*pb.User, error) {
	fetchedUser, err := u.userPostgresqlRepository.FindUserById(ctx, requestId, req.Id)
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error finding user by id: %v", requestId, err))
		return nil, err
	}
	return &pb.User{
		Id:    fetchedUser.ID,
		Name:  fetchedUser.Name,
		Email: fetchedUser.Email,
	}, nil
}
