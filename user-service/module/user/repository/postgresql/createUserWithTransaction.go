package postgresql

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"gorm.io/gorm"
)

func (r *userPostgresqlRepository) CreateUserWithTransaction(ctx context.Context, requestId string, req *orm.User, tx *gorm.DB) (string, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.CreateUserWithTransaction")
	defer span.End()
	if err := tx.WithContext(ctx).
		Table(userTable).
		Create(req).
		Error; err != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s , error creating user: %v", requestId, err))
		return "", err
	}

	return req.ID, nil
}
