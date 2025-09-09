package postgres

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgres"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type (
	IProductPostgresqlRepository interface {
		// FindAllProductForElasticIndex
		// This Function is only use for init index product's data. do not use on Real Use Case Scenario
		FindAllProductForElasticIndex(ctx context.Context, tx *gorm.DB) ([]*orm.Product, error)

		FindProductById(ctx context.Context, id string, tx *gorm.DB) (*orm.Product, error)
		CreateProduct(ctx context.Context, product *orm.Product, tx *gorm.DB) (string, error)
		DeleteProductById(ctx context.Context, id string, tx *gorm.DB) error
		UpdateProductById(ctx context.Context, product *orm.Product, tx *gorm.DB) (*orm.Product, error)
	}

	ProductPostgresqlRepository struct {
		postgresSQLInfrastructure *postgres.PostgresSQL
		telemetryInfrastructure   telemetryInfrastructure.ITelemetryInfrastructure
		logger                    logger.IZapLogger
	}
)

var Set = wire.NewSet(NewProductPostgresqlRepository)

func NewProductPostgresqlRepository(
	postgresSQLInfrastructure *postgres.PostgresSQL,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IProductPostgresqlRepository {
	return &ProductPostgresqlRepository{
		postgresSQLInfrastructure: postgresSQLInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		logger:                    logger,
	}
}
