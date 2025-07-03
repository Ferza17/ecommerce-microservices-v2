package orm

import (
	pb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// Payment GORM model
type Payment struct {
	ID          string     `gorm:"primaryKey;type:varchar(255)"`
	Code        string     `gorm:"type:varchar(100);uniqueIndex;not null"`
	TotalPrice  float64    `gorm:"not null"`
	Status      string     `gorm:"not null"`
	ProviderID  string     `gorm:"type:varchar(255);not null"`
	UserID      string     `gorm:"type:varchar(255);not null"`
	CreatedAt   *time.Time `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime"`
	DiscardedAt *time.Time `gorm:"index"`
}

func (Payment) TableName() string {
	return "payments"
}

func (p *Payment) ToProto() *pb.Payment {
	return &pb.Payment{
		Id:         p.ID,
		Code:       p.Code,
		TotalPrice: p.TotalPrice,
		Status: func() pb.PaymentStatus {
			if p.Status != "" {
				return pb.PaymentStatus(pb.PaymentStatus_value[p.Status])
			}
			return pb.PaymentStatus_FAILED
		}(),
		ProviderId: p.ProviderID,
		UserId:     p.UserID,
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

func PaymentFromProto(proto *pb.Payment) *Payment {
	return &Payment{
		ID:         proto.Id,
		Code:       proto.Code,
		TotalPrice: proto.TotalPrice,
		Status:     proto.GetStatus().String(),
		ProviderID: proto.ProviderId,
		UserID:     proto.UserId,
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

func PaymentsToProto(payments []*Payment) []*pb.Payment {
	var protos []*pb.Payment
	for _, payment := range payments {
		protos = append(protos, payment.ToProto())
	}
	return protos
}

func PaymentsFromProto(protos []*pb.Payment) []*Payment {
	var payments []*Payment
	for _, proto := range protos {
		payments = append(payments, PaymentFromProto(proto))
	}
	return payments
}
