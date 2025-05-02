package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
)

func (u *userUseCase) FindUserByEmailAndPassword(ctx context.Context, requestId string, req *pb.FindUserByEmailAndPasswordRequest) (*pb.User, error) {
	//TODO implement me
	panic("implement me")
}
