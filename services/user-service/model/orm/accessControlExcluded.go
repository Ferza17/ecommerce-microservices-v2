package orm

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
)

type AccessControlExcluded struct {
	ID             string    `gorm:"primaryKey;type:varchar(255)" json:"id"`
	FullMethodName string    `gorm:"type:varchar(255);not null" json:"full_method_name"`
	HttpUrl        string    `gorm:"type:varchar(255);not null" json:"http_url"`
	HttpMethod     string    `gorm:"type:varchar(255);not null" json:"http_method"`
	EventType      string    `gorm:"type:varchar(255);" json:"event_type"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (a *AccessControlExcluded) TableName() string {
	return "access_control_excluded"
}

// ToProto converts AccessControlExcluded to a protobuf message.
func (a *AccessControlExcluded) ToProto() (*pb.AccessControlExcluded, error) {
	return &pb.AccessControlExcluded{
		Id:             a.ID,
		FullMethodName: a.FullMethodName,
		HttpUrl:        a.HttpUrl,
		HttpMethod:     a.HttpMethod,
		CreatedAt:      timestamppb.New(a.CreatedAt),
		UpdatedAt:      timestamppb.New(a.UpdatedAt),
	}, nil
}

// AccessControlExcludedFromProto converts a protobuf message to AccessControlExcluded.
func AccessControlExcludedFromProto(protoMsg *pb.AccessControlExcluded) *AccessControlExcluded {
	excluded := &AccessControlExcluded{
		ID:             protoMsg.Id,
		HttpUrl:        protoMsg.HttpUrl,
		HttpMethod:     protoMsg.HttpMethod,
		FullMethodName: protoMsg.FullMethodName,
	}

	if protoMsg.CreatedAt != nil {
		excluded.CreatedAt = protoMsg.CreatedAt.AsTime()
	}

	if protoMsg.UpdatedAt != nil {
		excluded.UpdatedAt = protoMsg.UpdatedAt.AsTime()
	}

	return excluded
}
