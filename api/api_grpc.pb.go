// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// SyncTreeClient is the client API for SyncTree service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncTreeClient interface {
	UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*Response, error)
	UserSend(ctx context.Context, in *UserSendRequest, opts ...grpc.CallOption) (*Response, error)
}

type syncTreeClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncTreeClient(cc grpc.ClientConnInterface) SyncTreeClient {
	return &syncTreeClient{cc}
}

func (c *syncTreeClient) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserSend(ctx context.Context, in *UserSendRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncTreeServer is the server API for SyncTree service.
// All implementations must embed UnimplementedSyncTreeServer
// for forward compatibility
type SyncTreeServer interface {
	UserCreate(context.Context, *UserCreateRequest) (*Response, error)
	UserSend(context.Context, *UserSendRequest) (*Response, error)
	mustEmbedUnimplementedSyncTreeServer()
}

// UnimplementedSyncTreeServer must be embedded to have forward compatible implementations.
type UnimplementedSyncTreeServer struct {
}

func (UnimplementedSyncTreeServer) UserCreate(context.Context, *UserCreateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreate not implemented")
}
func (UnimplementedSyncTreeServer) UserSend(context.Context, *UserSendRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSend not implemented")
}
func (UnimplementedSyncTreeServer) mustEmbedUnimplementedSyncTreeServer() {}

// UnsafeSyncTreeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyncTreeServer will
// result in compilation errors.
type UnsafeSyncTreeServer interface {
	mustEmbedUnimplementedSyncTreeServer()
}

func RegisterSyncTreeServer(s grpc.ServiceRegistrar, srv SyncTreeServer) {
	s.RegisterService(&SyncTree_ServiceDesc, srv)
}

func _SyncTree_UserCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserCreate(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserSend(ctx, req.(*UserSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SyncTree_ServiceDesc is the grpc.ServiceDesc for SyncTree service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncTree_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.SyncTree",
	HandlerType: (*SyncTreeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserCreate",
			Handler:    _SyncTree_UserCreate_Handler,
		},
		{
			MethodName: "UserSend",
			Handler:    _SyncTree_UserSend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
