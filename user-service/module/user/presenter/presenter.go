package presenter

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
)

type UserPresenter struct {
	pb.UnimplementedUserServiceServer

	UserUseCase usecase.IUserUseCase
}

func NewUserPresenter(userUseCase usecase.IUserUseCase) *UserPresenter {
	return &UserPresenter{
		UserUseCase: userUseCase,
	}
}
