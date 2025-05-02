package postgresql

import (
	"context"
	"gorm.io/gorm"
)

func (r *ProductPostgresqlRepository) DeleteProductById(ctx context.Context, id string, tx *gorm.DB) error {
	//TODO implement me
	panic("implement me")
}
