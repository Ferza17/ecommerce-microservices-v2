package postgres

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	"gorm.io/gorm"
)

// FindAllProductForElasticIndex
// This Function is only use for init index product's data. do not use on Real Use Case Scenario
func (r *ProductPostgresqlRepository) FindAllProductForElasticIndex(ctx context.Context, tx *gorm.DB) ([]*orm.Product, error) {
	var (
		products []*orm.Product
	)

	if err := tx.WithContext(ctx).
		Table("products").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
