package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/connector"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
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
		connector *connector.PostgresqlConnector
	}
)

func NewProductPostgresqlRepository(connector *connector.PostgresqlConnector) IProductPostgresqlRepository {
	return &ProductPostgresqlRepository{
		connector: connector,
	}
}

func (r *ProductPostgresqlRepository) OpenTransactionWithContext(ctx context.Context) *gorm.DB {
	return r.connector.GormDB.WithContext(ctx).Begin()
}
