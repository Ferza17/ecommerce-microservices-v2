package orm

import (
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"time"
)

// AccessControl represents the access_control table
type AccessControl struct {
	ID             string    `gorm:"primaryKey;type:varchar(255)" json:"id"`
	ServiceName    string    `gorm:"type:varchar(255);not null" json:"service_name"`
	FullMethodName string    `gorm:"type:varchar(255);not null" json:"full_method_name"`
	RoleID         string    `gorm:"type:varchar(255);index;not null" json:"role_id"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (AccessControl) TableName() string {
	return "access_controls"
}

func (ac *AccessControl) ToProto() *pb.AccessControl {
	return &pb.AccessControl{
		Id:             ac.ID,
		ServiceName:    ac.ServiceName,
		FullMethodName: ac.FullMethodName,
		RoleId:         ac.RoleID,
	}
}

func AccessControlFromProto(proto *pb.AccessControl) *AccessControl {
	return &AccessControl{
		ID:             proto.Id,
		ServiceName:    proto.ServiceName,
		FullMethodName: proto.FullMethodName,
		RoleID:         proto.RoleId,
	}
}
