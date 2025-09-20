package orm

import (
	pb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type PaymentItem struct {
	ID          string     `gorm:"primaryKey;type:varchar(255)" json:"id"`
	ProductID   string     `gorm:"type:varchar(255);not null" json:"product_id"`
	Amount      float64    `gorm:"not null" json:"amount"`
	Qty         int32      `gorm:"not null" json:"qty"`
	PaymentID   string     `gorm:"type:varchar(255);not null" json:"payment_id"`
	CreatedAt   *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DiscardedAt *time.Time `gorm:"index" json:"-"`
}

func (PaymentItem) TableName() string {
	return "payment_items"
}

func (p *PaymentItem) ToProto() *pb.PaymentItem {
	return &pb.PaymentItem{
		Id:        p.ID,
		ProductId: p.ProductID,
		Amount:    p.Amount,
		Qty:       p.Qty,
		PaymentId: p.PaymentID,
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

func PaymentItemFromProto(proto *pb.PaymentItem) *PaymentItem {
	return &PaymentItem{
		ID:        proto.Id,
		ProductID: proto.ProductId,
		Amount:    proto.Amount,
		Qty:       proto.Qty,
		PaymentID: proto.PaymentId,
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

func PaymentItemsFromProto(protos []*pb.PaymentItem) []*PaymentItem {
	var items []*PaymentItem
	for _, proto := range protos {
		items = append(items, PaymentItemFromProto(proto))
	}
	return items
}

func PaymentItemsToProto(items []*PaymentItem) []*pb.PaymentItem {
	var protos []*pb.PaymentItem
	for _, item := range items {
		protos = append(protos, item.ToProto())
	}
	return protos
}
