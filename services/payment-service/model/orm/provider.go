package orm

import (
	pb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// Provider GORM model
type Provider struct {
	ID          string     `gorm:"primaryKey;type:varchar(255)" json:"id"`
	Name        string     `gorm:"type:varchar(255);not null" json:"name"`
	Method      string     `gorm:"not null" json:"method"`
	CreatedAt   *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DiscardedAt *time.Time `gorm:"index" json:"-"`
}

func (Provider) TableName() string {
	return "payment_providers"
}

func (p *Provider) ToProto() *pb.Provider {
	return &pb.Provider{
		Id:   p.ID,
		Name: p.Name,
		Method: func() pb.ProviderMethod {
			if p.Method != "" {
				return pb.ProviderMethod(pb.ProviderMethod_value[p.Method])
			}
			return pb.ProviderMethod_BANK
		}(),
		CreatedAt: func() *timestamppb.Timestamp {
			if p.CreatedAt != nil {
				return timestamppb.New(*p.CreatedAt)
			}
			return nil
		}(),
		UpdatedAt: func() *timestamppb.Timestamp {
			if p.UpdatedAt != nil {
				return timestamppb.New(*p.UpdatedAt)
			}
			return nil
		}(),
		DiscardedAt: func() *timestamppb.Timestamp {
			if p.DiscardedAt != nil {
				return timestamppb.New(*p.DiscardedAt)
			}
			return nil
		}(),
	}
}

func ProviderFromProto(proto *pb.Provider) *Provider {
	return &Provider{
		ID:     proto.Id,
		Name:   proto.Name,
		Method: proto.Method.String(),
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
	}
}

func ProvidersToProto(providers []*Provider) []*pb.Provider {
	var protos []*pb.Provider
	for _, provider := range providers {
		protos = append(protos, provider.ToProto())
	}
	return protos
}

func ProvidersFromProto(protos []*pb.Provider) []*Provider {
	var providers []*Provider
	for _, proto := range protos {
		providers = append(providers, ProviderFromProto(proto))
	}
	return providers
}
