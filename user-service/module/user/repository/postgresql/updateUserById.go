package postgresql

import (
	"context"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/user/v1"

	"gorm.io/gorm"
	"time"
)

func (r *userPostgresqlRepository) UpdateUserByIdWithTransaction(ctx context.Context, requestId string, req *userRpc.UpdateUserByIdRequest, tx *gorm.DB) (string, error) {
	var (
		now      = time.Now().UTC()
		buildReq = map[string]any{}
	)
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.UpdateUserByIdWithTransaction")
	defer span.End()
	if req.Password != nil {
		buildReq["password"] = req.GetPassword()
	}
	if req.Name != nil {
		buildReq["name"] = req.GetName()
	}
	if req.Email != nil {
		buildReq["email"] = req.GetEmail()
	}

	buildReq["updated_at"] = now

	if err := tx.WithContext(ctx).
		Table(userTable).
		Where("id = ?", req.Id).Updates(buildReq).Error; err != nil {
	}

	return req.Id, nil
}
