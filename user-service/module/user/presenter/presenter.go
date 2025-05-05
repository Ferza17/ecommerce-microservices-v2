package presenter

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
)

type UserPresenter struct {
	pb.UnimplementedUserServiceServer

	userUseCase usecase.IUserUseCase
	logger      pkg.IZapLogger
}

func NewUserPresenter(userUseCase usecase.IUserUseCase, logger pkg.IZapLogger) *UserPresenter {
	return &UserPresenter{
		userUseCase: userUseCase,
		logger:      logger,
	}
}
