package orm

import (
	pb "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/product"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Product struct {
	ID          string     `gorm:"column:id;primaryKey" json:"id"`
	Name        string     `gorm:"column:name" json:"name"`
	Price       float64    `gorm:"column:price" json:"price"`
	Stock       int64      `gorm:"column:stock" json:"stock"`
	Description string     `gorm:"column:description" json:"description"`
	Image       string     `gorm:"column:image" json:"image"`
	Uom         string     `gorm:"column:uom" json:"uom"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	DiscardedAt *time.Time `gorm:"column:discarded_at" json:"discarded_at,omitempty"`
}

func (Product) TableName() string {
	return "products"
}

// ToProto converts Go struct to protobuf
func (p *Product) ToProto() *pb.Product {
	product := &pb.Product{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Uom:         p.Uom,
		Image:       p.Image,
		Price:       p.Price,
		Stock:       p.Stock,
	}

	if p.CreatedAt == nil {
		product.CreatedAt = timestamppb.New(*p.CreatedAt)
	}

	if p.DiscardedAt == nil {
		product.DiscardedAt = timestamppb.New(*p.DiscardedAt)
	}

	if p.UpdatedAt == nil {
		product.UpdatedAt = timestamppb.New(*p.UpdatedAt)
	}

	return product
}

// FromProto converts protobuf to Go struct
func ProductFromProto(proto *pb.Product) *Product {
	p := &Product{}
	p.ID = proto.Id
	p.Name = proto.Name
	p.Description = proto.Description
	p.Uom = proto.Uom
	p.Image = proto.Image
	p.Price = proto.Price
	p.Stock = proto.Stock

	if proto.CreatedAt != nil {
		t := proto.CreatedAt.AsTime()
		p.CreatedAt = &t
	}

	if proto.UpdatedAt != nil {
		t := proto.UpdatedAt.AsTime()
		p.UpdatedAt = &t
	}

	if proto.DiscardedAt != nil {
		t := proto.DiscardedAt.AsTime()
		p.DiscardedAt = &t
	}

	return p
}

func ProductsToProto(products []*Product) []*pb.Product {
	var protos []*pb.Product
	for _, product := range products {
		protos = append(protos, product.ToProto())
	}
	return protos
}

func RolesFromProto(protos []*pb.Product) []*Product {
	var products []*Product
	for _, proto := range protos {
		products = append(products, ProductFromProto(proto))
	}
	return products
}
