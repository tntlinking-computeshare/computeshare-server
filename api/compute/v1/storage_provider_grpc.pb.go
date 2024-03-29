// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: api/compute/v1/storage_provider.proto

package v1

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

const (
	StorageProvider_CreateStorageProvider_FullMethodName = "/api.server.compute.v1.StorageProvider/CreateStorageProvider"
	StorageProvider_DeleteStorageProvider_FullMethodName = "/api.server.compute.v1.StorageProvider/DeleteStorageProvider"
	StorageProvider_GetStorageProvider_FullMethodName    = "/api.server.compute.v1.StorageProvider/GetStorageProvider"
	StorageProvider_ListStorageProvider_FullMethodName   = "/api.server.compute.v1.StorageProvider/ListStorageProvider"
)

// StorageProviderClient is the client API for StorageProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageProviderClient interface {
	CreateStorageProvider(ctx context.Context, in *CreateStorageProviderRequest, opts ...grpc.CallOption) (*CreateStorageProviderReply, error)
	DeleteStorageProvider(ctx context.Context, in *DeleteStorageProviderRequest, opts ...grpc.CallOption) (*DeleteStorageProviderReply, error)
	GetStorageProvider(ctx context.Context, in *GetStorageProviderRequest, opts ...grpc.CallOption) (*GetStorageProviderReply, error)
	ListStorageProvider(ctx context.Context, in *ListStorageProviderRequest, opts ...grpc.CallOption) (*ListStorageProviderReply, error)
}

type storageProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageProviderClient(cc grpc.ClientConnInterface) StorageProviderClient {
	return &storageProviderClient{cc}
}

func (c *storageProviderClient) CreateStorageProvider(ctx context.Context, in *CreateStorageProviderRequest, opts ...grpc.CallOption) (*CreateStorageProviderReply, error) {
	out := new(CreateStorageProviderReply)
	err := c.cc.Invoke(ctx, StorageProvider_CreateStorageProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageProviderClient) DeleteStorageProvider(ctx context.Context, in *DeleteStorageProviderRequest, opts ...grpc.CallOption) (*DeleteStorageProviderReply, error) {
	out := new(DeleteStorageProviderReply)
	err := c.cc.Invoke(ctx, StorageProvider_DeleteStorageProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageProviderClient) GetStorageProvider(ctx context.Context, in *GetStorageProviderRequest, opts ...grpc.CallOption) (*GetStorageProviderReply, error) {
	out := new(GetStorageProviderReply)
	err := c.cc.Invoke(ctx, StorageProvider_GetStorageProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageProviderClient) ListStorageProvider(ctx context.Context, in *ListStorageProviderRequest, opts ...grpc.CallOption) (*ListStorageProviderReply, error) {
	out := new(ListStorageProviderReply)
	err := c.cc.Invoke(ctx, StorageProvider_ListStorageProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageProviderServer is the server API for StorageProvider service.
// All implementations must embed UnimplementedStorageProviderServer
// for forward compatibility
type StorageProviderServer interface {
	CreateStorageProvider(context.Context, *CreateStorageProviderRequest) (*CreateStorageProviderReply, error)
	DeleteStorageProvider(context.Context, *DeleteStorageProviderRequest) (*DeleteStorageProviderReply, error)
	GetStorageProvider(context.Context, *GetStorageProviderRequest) (*GetStorageProviderReply, error)
	ListStorageProvider(context.Context, *ListStorageProviderRequest) (*ListStorageProviderReply, error)
	mustEmbedUnimplementedStorageProviderServer()
}

// UnimplementedStorageProviderServer must be embedded to have forward compatible implementations.
type UnimplementedStorageProviderServer struct {
}

func (UnimplementedStorageProviderServer) CreateStorageProvider(context.Context, *CreateStorageProviderRequest) (*CreateStorageProviderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStorageProvider not implemented")
}
func (UnimplementedStorageProviderServer) DeleteStorageProvider(context.Context, *DeleteStorageProviderRequest) (*DeleteStorageProviderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStorageProvider not implemented")
}
func (UnimplementedStorageProviderServer) GetStorageProvider(context.Context, *GetStorageProviderRequest) (*GetStorageProviderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorageProvider not implemented")
}
func (UnimplementedStorageProviderServer) ListStorageProvider(context.Context, *ListStorageProviderRequest) (*ListStorageProviderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStorageProvider not implemented")
}
func (UnimplementedStorageProviderServer) mustEmbedUnimplementedStorageProviderServer() {}

// UnsafeStorageProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageProviderServer will
// result in compilation errors.
type UnsafeStorageProviderServer interface {
	mustEmbedUnimplementedStorageProviderServer()
}

func RegisterStorageProviderServer(s grpc.ServiceRegistrar, srv StorageProviderServer) {
	s.RegisterService(&StorageProvider_ServiceDesc, srv)
}

func _StorageProvider_CreateStorageProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStorageProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageProviderServer).CreateStorageProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageProvider_CreateStorageProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageProviderServer).CreateStorageProvider(ctx, req.(*CreateStorageProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageProvider_DeleteStorageProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStorageProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageProviderServer).DeleteStorageProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageProvider_DeleteStorageProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageProviderServer).DeleteStorageProvider(ctx, req.(*DeleteStorageProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageProvider_GetStorageProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStorageProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageProviderServer).GetStorageProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageProvider_GetStorageProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageProviderServer).GetStorageProvider(ctx, req.(*GetStorageProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageProvider_ListStorageProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStorageProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageProviderServer).ListStorageProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageProvider_ListStorageProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageProviderServer).ListStorageProvider(ctx, req.(*ListStorageProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageProvider_ServiceDesc is the grpc.ServiceDesc for StorageProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.server.compute.v1.StorageProvider",
	HandlerType: (*StorageProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStorageProvider",
			Handler:    _StorageProvider_CreateStorageProvider_Handler,
		},
		{
			MethodName: "DeleteStorageProvider",
			Handler:    _StorageProvider_DeleteStorageProvider_Handler,
		},
		{
			MethodName: "GetStorageProvider",
			Handler:    _StorageProvider_GetStorageProvider_Handler,
		},
		{
			MethodName: "ListStorageProvider",
			Handler:    _StorageProvider_ListStorageProvider_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/compute/v1/storage_provider.proto",
}
