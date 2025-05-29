package presenter

import (
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/telemetry"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/auth/usecase"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
	"net/http"
)

type (
	IAuthPresenter interface {
		UserLoginByEmailAndPassword(w http.ResponseWriter, r *http.Request)
		UserLogoutByToken(w http.ResponseWriter, r *http.Request)
		CreateUser(w http.ResponseWriter, r *http.Request)
		UserVerifyOtp(w http.ResponseWriter, r *http.Request)
	}
	authPresenter struct {
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		authUseCase             authUseCase.IAuthUseCase
		userUseCase             userUseCase.IUserUseCase
		logger                  pkg.IZapLogger
	}
)

func NewAuthPresenter(
	authUseCase authUseCase.IAuthUseCase,
	userUseCase userUseCase.IUserUseCase,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger,
) IAuthPresenter {
	return &authPresenter{
		telemetryInfrastructure: telemetryInfrastructure,
		authUseCase:             authUseCase,
		userUseCase:             userUseCase,
		logger:                  logger,
	}
}
