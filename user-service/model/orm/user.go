package orm

import (
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
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
	CreatedAt   *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DiscardedAt *time.Time `gorm:"index" json:"discarded_at,omitempty"`

	Role *Role `gorm:"foreignKey:RoleID" json:"role"`
}

// TableName methods to specify custom table names if needed
func (User) TableName() string {
	return "users"
}

func (u *User) ToProto() *pb.User {
	proto := &pb.User{
		Id:         u.ID,
		Name:       u.Name,
		Email:      u.Email,
		Password:   u.Password,
		IsVerified: u.IsVerified,
	}

	if u.CreatedAt != nil {
		proto.CreatedAt = timestamppb.New(*u.CreatedAt)
	}

	if u.UpdatedAt != nil {
		proto.UpdatedAt = timestamppb.New(*u.UpdatedAt)
	}

	if u.Role.ID != "" {
		proto.Role = u.Role.ToProto()
	}

	if u.DiscardedAt != nil {
		proto.DiscardedAt = timestamppb.New(*u.DiscardedAt)
	}

	return proto
}

func UserFromProto(proto *pb.User) *User {
	user := &User{
		ID:         proto.Id,
		Name:       proto.Name,
		Email:      proto.Email,
		Password:   proto.Password,
		IsVerified: proto.IsVerified,
	}

	if proto.CreatedAt != nil {
		t := proto.CreatedAt.AsTime()
		user.CreatedAt = &t
	}

	if proto.UpdatedAt != nil {
		t := proto.UpdatedAt.AsTime()
		user.UpdatedAt = &t
	}

	if proto.DiscardedAt != nil {
		discardedTime := proto.DiscardedAt.AsTime()
		user.DiscardedAt = &discardedTime
	}

	if proto.Role != nil {
		user.Role = RoleFromProto(proto.Role)
		user.RoleID = proto.Role.Id
	}

	return user
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
