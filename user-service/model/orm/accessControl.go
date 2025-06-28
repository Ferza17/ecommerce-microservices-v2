package orm

import (
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// AccessControl represents the access_control table
type AccessControl struct {
	ID             string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	FullMethodName string `gorm:"type:varchar(255);not null" json:"full_method_name"`
	HttpUrl        string `gorm:"type:varchar(255);not null" json:"http_url"`
	HttpMethod     string `gorm:"type:varchar(255);not null" json:"http_method"`
	EventType      string `gorm:"type:varchar(255);" json:"event_type"`

	RoleID    string    `gorm:"type:varchar(255);index;not null" json:"role_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (AccessControl) TableName() string {
	return "access_controls"
}

func (ac *AccessControl) ToProto() *pb.AccessControl {
	return &pb.AccessControl{
		Id:             ac.ID,
		FullMethodName: ac.FullMethodName,
		EventType:      ac.EventType,
		HttpUrl:        ac.HttpUrl,
		HttpMethod:     ac.HttpMethod,
		RoleId:         ac.RoleID,
		CreatedAt:      timestamppb.New(ac.CreatedAt),
		UpdatedAt:      timestamppb.New(ac.UpdatedAt),
	}
}

func AccessControlFromProto(proto *pb.AccessControl) *AccessControl {
	return &AccessControl{
		ID:             proto.Id,
		FullMethodName: proto.FullMethodName,
		EventType:      proto.EventType,
		HttpUrl:        proto.HttpUrl,
		RoleID:         proto.RoleId,
	}
}
