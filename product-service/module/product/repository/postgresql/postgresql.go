package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgresql"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
	"gorm.io/gorm"
)

type (
	IProductPostgresqlRepository interface {
		FindProductById(ctx context.Context, id string, tx *gorm.DB) (*orm.Product, error)
		CreateProduct(ctx context.Context, product *orm.Product, tx *gorm.DB) (string, error)
		DeleteProductById(ctx context.Context, id string, tx *gorm.DB) error
		UpdateProductById(ctx context.Context, product *orm.Product, tx *gorm.DB) (*orm.Product, error)

		// Transaction
		OpenTransactionWithContext(ctx context.Context) *gorm.DB
	}

	ProductPostgresqlRepository struct {
		postgreSQLInfrastructure postgresql.IPostgreSQLInfrastructure
		telemetryInfrastructure  telemetryInfrastructure.ITelemetryInfrastructure
		logger                   pkg.IZapLogger
	}
)

func NewProductPostgresqlRepository(
	infrastructure postgresql.IPostgreSQLInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger) IProductPostgresqlRepository {
	return &ProductPostgresqlRepository{
		postgreSQLInfrastructure: infrastructure,
		telemetryInfrastructure:  telemetryInfrastructure,
		logger:                   logger,
	}
}

func (r *ProductPostgresqlRepository) OpenTransactionWithContext(ctx context.Context) *gorm.DB {
	return r.postgreSQLInfrastructure.GormDB().WithContext(ctx).Begin()
}
