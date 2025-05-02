package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
)

func (p *UserPresenter) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	res, err := p.UserUseCase.CreateUser(context.Background(), "abcde", req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
