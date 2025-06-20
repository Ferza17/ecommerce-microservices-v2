// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: user/v1/authServices.proto

package gen

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
	AuthService_UserLogoutByToken_FullMethodName = "/user_v1.AuthService/UserLogoutByToken"
	AuthService_UserVerifyOtp_FullMethodName     = "/user_v1.AuthService/UserVerifyOtp"
	AuthService_FindUserByToken_FullMethodName   = "/user_v1.AuthService/FindUserByToken"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	UserLogoutByToken(ctx context.Context, in *UserLogoutByTokenRequest, opts ...grpc.CallOption) (*UserLogoutByTokenResponse, error)
	UserVerifyOtp(ctx context.Context, in *UserVerifyOtpRequest, opts ...grpc.CallOption) (*UserVerifyOtpResponse, error)
	FindUserByToken(ctx context.Context, in *FindUserByTokenRequest, opts ...grpc.CallOption) (*User, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) UserLogoutByToken(ctx context.Context, in *UserLogoutByTokenRequest, opts ...grpc.CallOption) (*UserLogoutByTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLogoutByTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_UserLogoutByToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserVerifyOtp(ctx context.Context, in *UserVerifyOtpRequest, opts ...grpc.CallOption) (*UserVerifyOtpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserVerifyOtpResponse)
	err := c.cc.Invoke(ctx, AuthService_UserVerifyOtp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) FindUserByToken(ctx context.Context, in *FindUserByTokenRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, AuthService_FindUserByToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations should embed UnimplementedAuthServiceServer
// for forward compatibility.
type AuthServiceServer interface {
	UserLogoutByToken(context.Context, *UserLogoutByTokenRequest) (*UserLogoutByTokenResponse, error)
	UserVerifyOtp(context.Context, *UserVerifyOtpRequest) (*UserVerifyOtpResponse, error)
	FindUserByToken(context.Context, *FindUserByTokenRequest) (*User, error)
}

// UnimplementedAuthServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) UserLogoutByToken(context.Context, *UserLogoutByTokenRequest) (*UserLogoutByTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogoutByToken not implemented")
}
func (UnimplementedAuthServiceServer) UserVerifyOtp(context.Context, *UserVerifyOtpRequest) (*UserVerifyOtpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserVerifyOtp not implemented")
}
func (UnimplementedAuthServiceServer) FindUserByToken(context.Context, *FindUserByTokenRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserByToken not implemented")
}
func (UnimplementedAuthServiceServer) testEmbeddedByValue() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_UserLogoutByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogoutByTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserLogoutByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UserLogoutByToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserLogoutByToken(ctx, req.(*UserLogoutByTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserVerifyOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserVerifyOtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserVerifyOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UserVerifyOtp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserVerifyOtp(ctx, req.(*UserVerifyOtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_FindUserByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserByTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).FindUserByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_FindUserByToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).FindUserByToken(ctx, req.(*FindUserByTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogoutByToken",
			Handler:    _AuthService_UserLogoutByToken_Handler,
		},
		{
			MethodName: "UserVerifyOtp",
			Handler:    _AuthService_UserVerifyOtp_Handler,
		},
		{
			MethodName: "FindUserByToken",
			Handler:    _AuthService_FindUserByToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/authServices.proto",
}
