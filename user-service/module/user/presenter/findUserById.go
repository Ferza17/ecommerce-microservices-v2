package presenter

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
)

func (p *UserPresenter) FindUserById(context.Context, *pb.FindUserByIdRequest) (*pb.User, error) {
	return nil, nil
}
