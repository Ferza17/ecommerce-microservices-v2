package usecase

import (
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	rolePostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/repository/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IRoleUseCase interface{}

	roleUseCase struct {
		rolePostgresqlRepository  rolePostgresqlRepository.IRolePostgresqlRepository
		kafkaInfrastructure       kafkaInfrastructure.IKafkaInfrastructure
		postgresSQLInfrastructure postgres.IPostgresSQL
		telemetryInfrastructure   telemetryInfrastructure.ITelemetryInfrastructure
		logger                    logger.IZapLogger
	}
)

var Set = wire.NewSet(NewRoleUseCase)

func NewRoleUseCase(
	rolePostgresqlRepository rolePostgresqlRepository.IRolePostgresqlRepository,
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	postgresSQLInfrastructure postgres.IPostgresSQL,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IRoleUseCase {
	return &roleUseCase{
		rolePostgresqlRepository:  rolePostgresqlRepository,
		kafkaInfrastructure:       kafkaInfrastructure,
		postgresSQLInfrastructure: postgresSQLInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		logger:                    logger,
	}
}
