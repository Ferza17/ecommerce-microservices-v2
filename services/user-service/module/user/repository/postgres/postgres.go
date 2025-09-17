package postgres

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type (
	IUserPostgresqlRepository interface {
		FindUserById(ctx context.Context, requestId string, id string, tx *gorm.DB) (*orm.User, error)
		FindUserByEmail(ctx context.Context, requestId string, email string, tx *gorm.DB) (*orm.User, error)
	}

	userPostgresqlRepository struct {
		postgresSQLInfrastructure postgres.IPostgresSQL
		telemetryInfrastructure   telemetryInfrastructure.ITelemetryInfrastructure
		logger                    logger.IZapLogger
	}
)

var Set = wire.NewSet(NewUserPostgresqlRepository)

func NewUserPostgresqlRepository(
	postgresSQLInfrastructure postgres.IPostgresSQL,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger) IUserPostgresqlRepository {
	c := &userPostgresqlRepository{
		postgresSQLInfrastructure: postgresSQLInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		logger:                    logger,
	}
	return c
}
