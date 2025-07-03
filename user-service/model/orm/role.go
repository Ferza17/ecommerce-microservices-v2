package orm

import (
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"time"
)

// Role represents the roles table
type Role struct {
	ID        string    `gorm:"primaryKey;type:varchar(255)" json:"id"`
	Role      string    `gorm:"type:varchar(50);not null" json:"role"` // EnumRole as int32
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	AccessControls []*AccessControl `gorm:"foreignKey:RoleID" json:"access_controls"`
}

func (Role) TableName() string {
	return "roles"
}

func RoleFromProto(proto *pb.Role) *Role {
	return &Role{
		ID:   proto.Id,
		Role: proto.Role.String(),
		AccessControls: func() []*AccessControl {
			if proto.AccessControls == nil {
				return AccessControlsFromProto(proto.AccessControls)
			}
			return nil
		}(),
	}
}

func (r *Role) ToProto() *pb.Role {
	proto := &pb.Role{
		Id: r.ID,
	}

	if r.Role != "" {
		role, _ := pb.EnumRole_value[r.Role]
		proto.Role = pb.EnumRole(role)
	}

	for _, ac := range r.AccessControls {
		proto.AccessControls = append(proto.AccessControls, ac.ToProto())
	}

	return proto
}

func RolesToProto(roles []*Role) []*pb.Role {
	var protos []*pb.Role
	for _, role := range roles {
		protos = append(protos, role.ToProto())
	}
	return protos
}

func RolesFromProto(protos []*pb.Role) []*Role {
	var roles []*Role
	for _, proto := range protos {
		roles = append(roles, RoleFromProto(proto))
	}
	return roles
}
