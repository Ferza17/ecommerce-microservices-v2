package orm

type (
	Product struct {
		Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
		Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
		Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
		Uom         string  `protobuf:"bytes,4,opt,name=uom,proto3" json:"uom,omitempty"`
		Image       string  `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
		Price       float64 `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
		Stock       int64   `protobuf:"varint,7,opt,name=stock,proto3" json:"stock,omitempty"`
		CreatedAt   int64   `protobuf:"varint,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
		UpdatedAt   int64   `protobuf:"varint,9,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
		DiscardedAt int64   `protobuf:"varint,10,opt,name=discardedAt,proto3" json:"discardedAt,omitempty"`
	}
)
