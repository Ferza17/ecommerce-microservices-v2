package presenter

import (
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
)

type AuthPresenter struct {
	pb.UnimplementedAuthServiceServer

	authUseCase             usecase.IAuthUseCase
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	logger                  pkg.IZapLogger
}

func NewAuthPresenter(
	authUseCase usecase.IAuthUseCase,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger) *AuthPresenter {
	return &AuthPresenter{
		authUseCase:             authUseCase,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
