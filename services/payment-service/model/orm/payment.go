package orm

import (
	"time"

	pb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Payment GORM model
type Payment struct {
	ID          string     `gorm:"primaryKey;type:varchar(255)" json:"id"`
	Code        string     `gorm:"type:varchar(100);uniqueIndex;not null" json:"code"`
	TotalPrice  float64    `gorm:"not null" json:"total_price"`
	Status      string     `gorm:"not null" json:"status"`
	ProviderID  string     `gorm:"type:varchar(255);not null" json:"provider_id"`
	UserID      string     `gorm:"type:varchar(255);not null" json:"user_id"`
	CreatedAt   *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DiscardedAt *time.Time `gorm:"index" json:"-"`

	PaymentItems    []*PaymentItem `gorm:"foreignKey:PaymentID;references:ID" json:"-"`
	PaymentProvider *Provider      `gorm:"foreignKey:ProviderID;references:ID" json:"-"`
}

func (Payment) TableName() string {
	return "payments"
}

func (p *Payment) ToProto() *pb.Payment {
	if p == nil {
		return nil
	}
	return &pb.Payment{
		Id:         p.ID,
		Code:       p.Code,
		TotalPrice: p.TotalPrice,
		Status: func() pb.PaymentStatus {
			if p.Status != "" {
				if s, ok := pb.PaymentStatus_value[p.Status]; ok {
					return pb.PaymentStatus(s)
				}
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
		Items: func() []*pb.PaymentItem {
			if p.PaymentItems != nil {
				return PaymentItemsToProto(p.PaymentItems)
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
		PaymentItems: func() []*PaymentItem {
			var items []*PaymentItem
			if proto.Items != nil {
				items = PaymentItemsFromProto(proto.Items)
			}
			return items
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
