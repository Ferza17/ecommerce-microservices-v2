package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/connector/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"gorm.io/gorm"
)

type (
	IProductPostgresqlRepository interface {
		FindProductById(ctx context.Context, id string, tx *gorm.DB) (*pb.ProductORM, error)
		CreateProduct(ctx context.Context, product *pb.ProductORM, tx *gorm.DB) (string, error)
		// Transaction
		OpenTransactionWithContext(ctx context.Context) *gorm.DB
	}

	ProductPostgresqlRepository struct {
		connector *postgresql.PostgresqlConnector
	}
)

func NewProductPostgresqlRepository(connector *postgresql.PostgresqlConnector) IProductPostgresqlRepository {
	return &ProductPostgresqlRepository{
		connector: connector,
	}
}

func (r *ProductPostgresqlRepository) OpenTransactionWithContext(ctx context.Context) *gorm.DB {
	return r.connector.GormDB.WithContext(ctx).Begin()
}
