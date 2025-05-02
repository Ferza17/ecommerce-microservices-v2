package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
)

func (u *userUseCase) FindUserById(ctx context.Context, requestId string, req *pb.FindUserByIdRequest) (*pb.User, error) {
	//TODO implement me
	panic("implement me")
}
