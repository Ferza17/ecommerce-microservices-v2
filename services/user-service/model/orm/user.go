package orm

import (
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	schema "github.com/ferza17/ecommerce-microservices-v2/user-service/model/schema/kafka-schema"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type User struct {
	ID          string     `gorm:"primaryKey;type:varchar(255)" json:"id"`
	Name        string     `gorm:"type:varchar(255);not null" json:"name"`
	Email       string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password    string     `gorm:"type:varchar(255);not null" json:"password"`
	IsVerified  bool       `gorm:"default:false" json:"is_verified"`
	RoleID      string     `gorm:"type:varchar(255);index" json:"role_id"`
	CreatedAt   *time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime" json:"-"`
	DiscardedAt *time.Time `gorm:"index" json:"-"`

	Role *Role `gorm:"foreignKey:RoleID" json:"-"`
}

// TableName methods to specify custom table names if needed
func (User) TableName() string {
	return "users"
}

func (u *User) ToProto() *pb.User {
	return &pb.User{
		Id:         u.ID,
		Name:       u.Name,
		Email:      u.Email,
		Password:   u.Password,
		IsVerified: u.IsVerified,
		CreatedAt: func() *timestamppb.Timestamp {
			if u.CreatedAt != nil {
				return timestamppb.New(*u.CreatedAt)
			}
			return nil
		}(),
		UpdatedAt: func() *timestamppb.Timestamp {
			if u.UpdatedAt != nil {
				return timestamppb.New(*u.UpdatedAt)
			}
			return nil
		}(),
		Role: func() *pb.Role {
			if u.Role != nil && u.Role.ID != "" {
				return u.Role.ToProto()
			}
			return nil
		}(),
		DiscardedAt: func() *timestamppb.Timestamp {
			if u.DiscardedAt != nil {
				return timestamppb.New(*u.DiscardedAt)
			}
			return nil
		}(),
	}
}

func UserFromProto(proto *pb.User) *User {
	return &User{
		ID:         proto.Id,
		Name:       proto.Name,
		Email:      proto.Email,
		Password:   proto.Password,
		IsVerified: proto.IsVerified,
		CreatedAt: func() *time.Time {
			if proto.CreatedAt != nil {
				t := proto.CreatedAt.AsTime()
				return &t
			}
			return nil
		}(),
		UpdatedAt: func() *time.Time {
			if proto.UpdatedAt != nil {
				t := proto.UpdatedAt.AsTime()
				return &t
			}
			return nil
		}(),
		DiscardedAt: func() *time.Time {
			if proto.DiscardedAt != nil {
				t := proto.DiscardedAt.AsTime()
				return &t
			}
			return nil
		}(),
		RoleID: func() string {
			if proto.Role != nil {
				return proto.Role.Id
			}
			return ""
		}(),
		Role: func() *Role {
			if proto.Role != nil {
				return RoleFromProto(proto.Role)
			}
			return nil
		}(),
	}
}

// UsersToProto
func UsersToProto(users []*User) []*pb.User {
	var protos []*pb.User
	for _, user := range users {
		protos = append(protos, user.ToProto())
	}
	return protos
}

// UsersFromProto Converts a slice of pb.User protos to a slice of User structs.
func UsersFromProto(protos []*pb.User) []*User {
	var users []*User
	for _, proto := range protos {
		users = append(users, UserFromProto(proto))
	}
	return users
}

// TO Kafka Sink Schema
func (u *User) ToKafkaSinkSchema() schema.KafkaSinkSchema {
	discardedAt := ""
	if u.DiscardedAt != nil {
		discardedAt = u.DiscardedAt.Format(time.RFC3339)
	}
	return schema.KafkaSinkSchema{
		Schema: schema.Schema{
			Type: "json",
			Fields: []schema.Fields{
				{
					Name:     "id",
					Type:     "string",
					Optional: false,
				},
				{
					Name:     "name",
					Type:     "string",
					Optional: false,
				},
				{
					Name:     "email",
					Type:     "string",
					Optional: false,
				},
				{
					Name:     "password",
					Type:     "string",
					Optional: false,
				},
				{
					Name:     "is_verified",
					Type:     "boolean",
					Optional: false,
				},
				{
					Name:     "role_id",
					Type:     "string",
					Optional: false,
				},
				{
					Name:     "created_at",
					Type:     "string",
					Optional: false,
				},
				{
					Name:     "updated_at",
					Type:     "string",
					Optional: false,
				},
				{
					Name:     "discarded_at",
					Type:     "string",
					Optional: true,
				},
			},
		},
		Payload: map[string]interface{}{
			"id":           u.ID,
			"name":         u.Name,
			"email":        u.Email,
			"password":     u.Password,
			"is_verified":  u.IsVerified,
			"role_id":      u.RoleID,
			"created_at":   u.CreatedAt.Format(time.RFC3339),
			"updated_at":   u.UpdatedAt.Format(time.RFC3339),
			"discarded_at": discardedAt,
		},
	}
}
