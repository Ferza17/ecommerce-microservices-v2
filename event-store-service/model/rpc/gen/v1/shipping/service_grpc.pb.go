// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: v1/shipping/service.proto

package shipping

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ShippingProviderService_CreateShippingProvider_FullMethodName  = "/shipping.ShippingProviderService/CreateShippingProvider"
	ShippingProviderService_GetShippingProviderById_FullMethodName = "/shipping.ShippingProviderService/GetShippingProviderById"
	ShippingProviderService_UpdateShippingProvider_FullMethodName  = "/shipping.ShippingProviderService/UpdateShippingProvider"
	ShippingProviderService_DeleteShippingProvider_FullMethodName  = "/shipping.ShippingProviderService/DeleteShippingProvider"
	ShippingProviderService_ListShippingProviders_FullMethodName   = "/shipping.ShippingProviderService/ListShippingProviders"
)

// ShippingProviderServiceClient is the client API for ShippingProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShippingProviderServiceClient interface {
	CreateShippingProvider(ctx context.Context, in *CreateShippingProviderRequest, opts ...grpc.CallOption) (*CreateShippingProviderResponse, error)
	GetShippingProviderById(ctx context.Context, in *GetShippingProviderByIdRequest, opts ...grpc.CallOption) (*GetShippingProviderByIdResponse, error)
	UpdateShippingProvider(ctx context.Context, in *UpdateShippingProviderRequest, opts ...grpc.CallOption) (*UpdateShippingProviderResponse, error)
	DeleteShippingProvider(ctx context.Context, in *DeleteShippingProviderRequest, opts ...grpc.CallOption) (*DeleteShippingProviderResponse, error)
	ListShippingProviders(ctx context.Context, in *ListShippingProvidersRequest, opts ...grpc.CallOption) (*ListShippingProvidersResponse, error)
}

type shippingProviderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingProviderServiceClient(cc grpc.ClientConnInterface) ShippingProviderServiceClient {
	return &shippingProviderServiceClient{cc}
}

func (c *shippingProviderServiceClient) CreateShippingProvider(ctx context.Context, in *CreateShippingProviderRequest, opts ...grpc.CallOption) (*CreateShippingProviderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateShippingProviderResponse)
	err := c.cc.Invoke(ctx, ShippingProviderService_CreateShippingProvider_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingProviderServiceClient) GetShippingProviderById(ctx context.Context, in *GetShippingProviderByIdRequest, opts ...grpc.CallOption) (*GetShippingProviderByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetShippingProviderByIdResponse)
	err := c.cc.Invoke(ctx, ShippingProviderService_GetShippingProviderById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingProviderServiceClient) UpdateShippingProvider(ctx context.Context, in *UpdateShippingProviderRequest, opts ...grpc.CallOption) (*UpdateShippingProviderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateShippingProviderResponse)
	err := c.cc.Invoke(ctx, ShippingProviderService_UpdateShippingProvider_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingProviderServiceClient) DeleteShippingProvider(ctx context.Context, in *DeleteShippingProviderRequest, opts ...grpc.CallOption) (*DeleteShippingProviderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteShippingProviderResponse)
	err := c.cc.Invoke(ctx, ShippingProviderService_DeleteShippingProvider_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingProviderServiceClient) ListShippingProviders(ctx context.Context, in *ListShippingProvidersRequest, opts ...grpc.CallOption) (*ListShippingProvidersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListShippingProvidersResponse)
	err := c.cc.Invoke(ctx, ShippingProviderService_ListShippingProviders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingProviderServiceServer is the server API for ShippingProviderService service.
// All implementations should embed UnimplementedShippingProviderServiceServer
// for forward compatibility.
type ShippingProviderServiceServer interface {
	CreateShippingProvider(context.Context, *CreateShippingProviderRequest) (*CreateShippingProviderResponse, error)
	GetShippingProviderById(context.Context, *GetShippingProviderByIdRequest) (*GetShippingProviderByIdResponse, error)
	UpdateShippingProvider(context.Context, *UpdateShippingProviderRequest) (*UpdateShippingProviderResponse, error)
	DeleteShippingProvider(context.Context, *DeleteShippingProviderRequest) (*DeleteShippingProviderResponse, error)
	ListShippingProviders(context.Context, *ListShippingProvidersRequest) (*ListShippingProvidersResponse, error)
}

// UnimplementedShippingProviderServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShippingProviderServiceServer struct{}

func (UnimplementedShippingProviderServiceServer) CreateShippingProvider(context.Context, *CreateShippingProviderRequest) (*CreateShippingProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShippingProvider not implemented")
}
func (UnimplementedShippingProviderServiceServer) GetShippingProviderById(context.Context, *GetShippingProviderByIdRequest) (*GetShippingProviderByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShippingProviderById not implemented")
}
func (UnimplementedShippingProviderServiceServer) UpdateShippingProvider(context.Context, *UpdateShippingProviderRequest) (*UpdateShippingProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShippingProvider not implemented")
}
func (UnimplementedShippingProviderServiceServer) DeleteShippingProvider(context.Context, *DeleteShippingProviderRequest) (*DeleteShippingProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShippingProvider not implemented")
}
func (UnimplementedShippingProviderServiceServer) ListShippingProviders(context.Context, *ListShippingProvidersRequest) (*ListShippingProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShippingProviders not implemented")
}
func (UnimplementedShippingProviderServiceServer) testEmbeddedByValue() {}

// UnsafeShippingProviderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShippingProviderServiceServer will
// result in compilation errors.
type UnsafeShippingProviderServiceServer interface {
	mustEmbedUnimplementedShippingProviderServiceServer()
}

func RegisterShippingProviderServiceServer(s grpc.ServiceRegistrar, srv ShippingProviderServiceServer) {
	// If the following call pancis, it indicates UnimplementedShippingProviderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ShippingProviderService_ServiceDesc, srv)
}

func _ShippingProviderService_CreateShippingProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShippingProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingProviderServiceServer).CreateShippingProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingProviderService_CreateShippingProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingProviderServiceServer).CreateShippingProvider(ctx, req.(*CreateShippingProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingProviderService_GetShippingProviderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShippingProviderByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingProviderServiceServer).GetShippingProviderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingProviderService_GetShippingProviderById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingProviderServiceServer).GetShippingProviderById(ctx, req.(*GetShippingProviderByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingProviderService_UpdateShippingProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShippingProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingProviderServiceServer).UpdateShippingProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingProviderService_UpdateShippingProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingProviderServiceServer).UpdateShippingProvider(ctx, req.(*UpdateShippingProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingProviderService_DeleteShippingProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShippingProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingProviderServiceServer).DeleteShippingProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingProviderService_DeleteShippingProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingProviderServiceServer).DeleteShippingProvider(ctx, req.(*DeleteShippingProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingProviderService_ListShippingProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListShippingProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingProviderServiceServer).ListShippingProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingProviderService_ListShippingProviders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingProviderServiceServer).ListShippingProviders(ctx, req.(*ListShippingProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShippingProviderService_ServiceDesc is the grpc.ServiceDesc for ShippingProviderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShippingProviderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shipping.ShippingProviderService",
	HandlerType: (*ShippingProviderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShippingProvider",
			Handler:    _ShippingProviderService_CreateShippingProvider_Handler,
		},
		{
			MethodName: "GetShippingProviderById",
			Handler:    _ShippingProviderService_GetShippingProviderById_Handler,
		},
		{
			MethodName: "UpdateShippingProvider",
			Handler:    _ShippingProviderService_UpdateShippingProvider_Handler,
		},
		{
			MethodName: "DeleteShippingProvider",
			Handler:    _ShippingProviderService_DeleteShippingProvider_Handler,
		},
		{
			MethodName: "ListShippingProviders",
			Handler:    _ShippingProviderService_ListShippingProviders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/shipping/service.proto",
}
