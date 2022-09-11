// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: services/secauthn/secauthn.proto

package secauthn

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SecauthnClient is the client API for Secauthn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecauthnClient interface {
	LogonUser(ctx context.Context, in *LogonUserRequest, opts ...grpc.CallOption) (*LogonUserResponse, error)
}

type secauthnClient struct {
	cc grpc.ClientConnInterface
}

func NewSecauthnClient(cc grpc.ClientConnInterface) SecauthnClient {
	return &secauthnClient{cc}
}

func (c *secauthnClient) LogonUser(ctx context.Context, in *LogonUserRequest, opts ...grpc.CallOption) (*LogonUserResponse, error) {
	out := new(LogonUserResponse)
	err := c.cc.Invoke(ctx, "/services.secauthn.Secauthn/LogonUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecauthnServer is the server API for Secauthn service.
// All implementations must embed UnimplementedSecauthnServer
// for forward compatibility
type SecauthnServer interface {
	LogonUser(context.Context, *LogonUserRequest) (*LogonUserResponse, error)
	mustEmbedUnimplementedSecauthnServer()
}

// UnimplementedSecauthnServer must be embedded to have forward compatible implementations.
type UnimplementedSecauthnServer struct {
}

func (UnimplementedSecauthnServer) LogonUser(context.Context, *LogonUserRequest) (*LogonUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogonUser not implemented")
}
func (UnimplementedSecauthnServer) mustEmbedUnimplementedSecauthnServer() {}

// UnsafeSecauthnServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecauthnServer will
// result in compilation errors.
type UnsafeSecauthnServer interface {
	mustEmbedUnimplementedSecauthnServer()
}

func RegisterSecauthnServer(s grpc.ServiceRegistrar, srv SecauthnServer) {
	s.RegisterService(&Secauthn_ServiceDesc, srv)
}

func _Secauthn_LogonUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogonUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecauthnServer).LogonUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.secauthn.Secauthn/LogonUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecauthnServer).LogonUser(ctx, req.(*LogonUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Secauthn_ServiceDesc is the grpc.ServiceDesc for Secauthn service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Secauthn_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.secauthn.Secauthn",
	HandlerType: (*SecauthnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogonUser",
			Handler:    _Secauthn_LogonUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/secauthn/secauthn.proto",
}
