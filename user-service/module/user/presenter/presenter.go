package presenter

import (
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	"github.com/google/wire"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
)

type UserPresenter struct {
	pb.UnimplementedUserServiceServer

	userUseCase             usecase.IUserUseCase
	authUseCase             authUseCase.IAuthUseCase
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	logger                  logger.IZapLogger
}

var Set = wire.NewSet(NewUserPresenter)

func NewUserPresenter(
	userUseCase usecase.IUserUseCase,
	authUseCase authUseCase.IAuthUseCase,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger) *UserPresenter {
	return &UserPresenter{
		userUseCase:             userUseCase,
		authUseCase:             authUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
