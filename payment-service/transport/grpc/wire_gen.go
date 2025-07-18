// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package grpc

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/service/user"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/presenter"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/repository"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/module/payment/usecase"
	presenter2 "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/presenter"
	repository2 "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/repository"
	usecase2 "github.com/ferza17/ecommerce-microservices-v2/payment-service/module/provider/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
)

// Injectors from wire.go:

// ProvideGrpcServer wires all dependencies for IGrpcServer
func ProvideGrpcServer() IGrpcServer {
	iZapLogger := logger.NewZapLogger()
	postgresSQL := postgresql.NewPostgresqlInfrastructure(iZapLogger)
	iTelemetryInfrastructure := telemetry.NewTelemetry(iZapLogger)
	iPaymentRepository := repository.NewPaymentRepository(postgresSQL, iTelemetryInfrastructure, iZapLogger)
	iRabbitMQInfrastructure := rabbitmq.NewRabbitMQInfrastructure(iTelemetryInfrastructure, iZapLogger)
	iPaymentUseCase := usecase.NewPaymentUseCase(iPaymentRepository, iRabbitMQInfrastructure, iTelemetryInfrastructure, iZapLogger, postgresSQL)
	iUserService := user.NewUserService(iZapLogger)
	iPaymentPresenter := presenter.NewPaymentPresenter(iPaymentUseCase, iTelemetryInfrastructure, iUserService, iZapLogger)
	iPaymentProviderRepository := repository2.NewPaymentProviderRepository(postgresSQL, iTelemetryInfrastructure, iZapLogger)
	iPaymentProviderUseCase := usecase2.NewPaymentProviderUseCase(iPaymentProviderRepository, iRabbitMQInfrastructure, iTelemetryInfrastructure, postgresSQL, iZapLogger)
	iPaymentProviderPresenter := presenter2.NewPaymentProviderPresenter(iPaymentProviderUseCase, iTelemetryInfrastructure, iUserService, iZapLogger)
	iGrpcServer := NewGrpcServer(iZapLogger, iPaymentPresenter, iPaymentProviderPresenter, iTelemetryInfrastructure, iUserService)
	return iGrpcServer
}
