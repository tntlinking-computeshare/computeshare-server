// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: api/compute/v1/storage_s3.proto

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
	StorageS3_CreateS3Key_FullMethodName           = "/api.server.compute.v1.StorageS3/CreateS3Key"
	StorageS3_GetUserS3User_FullMethodName         = "/api.server.compute.v1.StorageS3/GetUserS3User"
	StorageS3_CreateBucket_FullMethodName          = "/api.server.compute.v1.StorageS3/CreateBucket"
	StorageS3_DeleteBucket_FullMethodName          = "/api.server.compute.v1.StorageS3/DeleteBucket"
	StorageS3_EmptyBucket_FullMethodName           = "/api.server.compute.v1.StorageS3/EmptyBucket"
	StorageS3_ListBucket_FullMethodName            = "/api.server.compute.v1.StorageS3/ListBucket"
	StorageS3_S3StorageInBucketList_FullMethodName = "/api.server.compute.v1.StorageS3/S3StorageInBucketList"
	StorageS3_S3StorageUploadFile_FullMethodName   = "/api.server.compute.v1.StorageS3/S3StorageUploadFile"
	StorageS3_S3StorageMkdir_FullMethodName        = "/api.server.compute.v1.StorageS3/S3StorageMkdir"
	StorageS3_S3StorageDownload_FullMethodName     = "/api.server.compute.v1.StorageS3/S3StorageDownload"
	StorageS3_S3StorageDelete_FullMethodName       = "/api.server.compute.v1.StorageS3/S3StorageDelete"
)

// StorageS3Client is the client API for StorageS3 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageS3Client interface {
	CreateS3Key(ctx context.Context, in *CreateS3KeyRequest, opts ...grpc.CallOption) (*CreateS3KeyReply, error)
	GetUserS3User(ctx context.Context, in *GetS3UserRequest, opts ...grpc.CallOption) (*GetS3UserReply, error)
	CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*CreateBucketReply, error)
	DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*DeleteBucketReply, error)
	EmptyBucket(ctx context.Context, in *EmptyBucketRequest, opts ...grpc.CallOption) (*EmptyBucketReply, error)
	ListBucket(ctx context.Context, in *ListBucketRequest, opts ...grpc.CallOption) (*ListBucketReply, error)
	S3StorageInBucketList(ctx context.Context, in *S3StorageInBucketListRequest, opts ...grpc.CallOption) (*S3StorageInBucketListReply, error)
	S3StorageUploadFile(ctx context.Context, in *S3StorageUploadFileRequest, opts ...grpc.CallOption) (*S3StorageUploadFileReply, error)
	S3StorageMkdir(ctx context.Context, in *S3StorageMkdirRequest, opts ...grpc.CallOption) (*S3StorageMkdirReply, error)
	S3StorageDownload(ctx context.Context, in *S3StorageDownloadRequest, opts ...grpc.CallOption) (*S3StorageDownloadReply, error)
	S3StorageDelete(ctx context.Context, in *S3StorageDeleteRequest, opts ...grpc.CallOption) (*S3StorageDeleteReply, error)
}

type storageS3Client struct {
	cc grpc.ClientConnInterface
}

func NewStorageS3Client(cc grpc.ClientConnInterface) StorageS3Client {
	return &storageS3Client{cc}
}

func (c *storageS3Client) CreateS3Key(ctx context.Context, in *CreateS3KeyRequest, opts ...grpc.CallOption) (*CreateS3KeyReply, error) {
	out := new(CreateS3KeyReply)
	err := c.cc.Invoke(ctx, StorageS3_CreateS3Key_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageS3Client) GetUserS3User(ctx context.Context, in *GetS3UserRequest, opts ...grpc.CallOption) (*GetS3UserReply, error) {
	out := new(GetS3UserReply)
	err := c.cc.Invoke(ctx, StorageS3_GetUserS3User_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageS3Client) CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*CreateBucketReply, error) {
	out := new(CreateBucketReply)
	err := c.cc.Invoke(ctx, StorageS3_CreateBucket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageS3Client) DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*DeleteBucketReply, error) {
	out := new(DeleteBucketReply)
	err := c.cc.Invoke(ctx, StorageS3_DeleteBucket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageS3Client) EmptyBucket(ctx context.Context, in *EmptyBucketRequest, opts ...grpc.CallOption) (*EmptyBucketReply, error) {
	out := new(EmptyBucketReply)
	err := c.cc.Invoke(ctx, StorageS3_EmptyBucket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageS3Client) ListBucket(ctx context.Context, in *ListBucketRequest, opts ...grpc.CallOption) (*ListBucketReply, error) {
	out := new(ListBucketReply)
	err := c.cc.Invoke(ctx, StorageS3_ListBucket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageS3Client) S3StorageInBucketList(ctx context.Context, in *S3StorageInBucketListRequest, opts ...grpc.CallOption) (*S3StorageInBucketListReply, error) {
	out := new(S3StorageInBucketListReply)
	err := c.cc.Invoke(ctx, StorageS3_S3StorageInBucketList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageS3Client) S3StorageUploadFile(ctx context.Context, in *S3StorageUploadFileRequest, opts ...grpc.CallOption) (*S3StorageUploadFileReply, error) {
	out := new(S3StorageUploadFileReply)
	err := c.cc.Invoke(ctx, StorageS3_S3StorageUploadFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageS3Client) S3StorageMkdir(ctx context.Context, in *S3StorageMkdirRequest, opts ...grpc.CallOption) (*S3StorageMkdirReply, error) {
	out := new(S3StorageMkdirReply)
	err := c.cc.Invoke(ctx, StorageS3_S3StorageMkdir_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageS3Client) S3StorageDownload(ctx context.Context, in *S3StorageDownloadRequest, opts ...grpc.CallOption) (*S3StorageDownloadReply, error) {
	out := new(S3StorageDownloadReply)
	err := c.cc.Invoke(ctx, StorageS3_S3StorageDownload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageS3Client) S3StorageDelete(ctx context.Context, in *S3StorageDeleteRequest, opts ...grpc.CallOption) (*S3StorageDeleteReply, error) {
	out := new(S3StorageDeleteReply)
	err := c.cc.Invoke(ctx, StorageS3_S3StorageDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageS3Server is the server API for StorageS3 service.
// All implementations must embed UnimplementedStorageS3Server
// for forward compatibility
type StorageS3Server interface {
	CreateS3Key(context.Context, *CreateS3KeyRequest) (*CreateS3KeyReply, error)
	GetUserS3User(context.Context, *GetS3UserRequest) (*GetS3UserReply, error)
	CreateBucket(context.Context, *CreateBucketRequest) (*CreateBucketReply, error)
	DeleteBucket(context.Context, *DeleteBucketRequest) (*DeleteBucketReply, error)
	EmptyBucket(context.Context, *EmptyBucketRequest) (*EmptyBucketReply, error)
	ListBucket(context.Context, *ListBucketRequest) (*ListBucketReply, error)
	S3StorageInBucketList(context.Context, *S3StorageInBucketListRequest) (*S3StorageInBucketListReply, error)
	S3StorageUploadFile(context.Context, *S3StorageUploadFileRequest) (*S3StorageUploadFileReply, error)
	S3StorageMkdir(context.Context, *S3StorageMkdirRequest) (*S3StorageMkdirReply, error)
	S3StorageDownload(context.Context, *S3StorageDownloadRequest) (*S3StorageDownloadReply, error)
	S3StorageDelete(context.Context, *S3StorageDeleteRequest) (*S3StorageDeleteReply, error)
	mustEmbedUnimplementedStorageS3Server()
}

// UnimplementedStorageS3Server must be embedded to have forward compatible implementations.
type UnimplementedStorageS3Server struct {
}

func (UnimplementedStorageS3Server) CreateS3Key(context.Context, *CreateS3KeyRequest) (*CreateS3KeyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateS3Key not implemented")
}
func (UnimplementedStorageS3Server) GetUserS3User(context.Context, *GetS3UserRequest) (*GetS3UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserS3User not implemented")
}
func (UnimplementedStorageS3Server) CreateBucket(context.Context, *CreateBucketRequest) (*CreateBucketReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBucket not implemented")
}
func (UnimplementedStorageS3Server) DeleteBucket(context.Context, *DeleteBucketRequest) (*DeleteBucketReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBucket not implemented")
}
func (UnimplementedStorageS3Server) EmptyBucket(context.Context, *EmptyBucketRequest) (*EmptyBucketReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmptyBucket not implemented")
}
func (UnimplementedStorageS3Server) ListBucket(context.Context, *ListBucketRequest) (*ListBucketReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBucket not implemented")
}
func (UnimplementedStorageS3Server) S3StorageInBucketList(context.Context, *S3StorageInBucketListRequest) (*S3StorageInBucketListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method S3StorageInBucketList not implemented")
}
func (UnimplementedStorageS3Server) S3StorageUploadFile(context.Context, *S3StorageUploadFileRequest) (*S3StorageUploadFileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method S3StorageUploadFile not implemented")
}
func (UnimplementedStorageS3Server) S3StorageMkdir(context.Context, *S3StorageMkdirRequest) (*S3StorageMkdirReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method S3StorageMkdir not implemented")
}
func (UnimplementedStorageS3Server) S3StorageDownload(context.Context, *S3StorageDownloadRequest) (*S3StorageDownloadReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method S3StorageDownload not implemented")
}
func (UnimplementedStorageS3Server) S3StorageDelete(context.Context, *S3StorageDeleteRequest) (*S3StorageDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method S3StorageDelete not implemented")
}
func (UnimplementedStorageS3Server) mustEmbedUnimplementedStorageS3Server() {}

// UnsafeStorageS3Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageS3Server will
// result in compilation errors.
type UnsafeStorageS3Server interface {
	mustEmbedUnimplementedStorageS3Server()
}

func RegisterStorageS3Server(s grpc.ServiceRegistrar, srv StorageS3Server) {
	s.RegisterService(&StorageS3_ServiceDesc, srv)
}

func _StorageS3_CreateS3Key_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateS3KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageS3Server).CreateS3Key(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageS3_CreateS3Key_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageS3Server).CreateS3Key(ctx, req.(*CreateS3KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageS3_GetUserS3User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetS3UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageS3Server).GetUserS3User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageS3_GetUserS3User_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageS3Server).GetUserS3User(ctx, req.(*GetS3UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageS3_CreateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageS3Server).CreateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageS3_CreateBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageS3Server).CreateBucket(ctx, req.(*CreateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageS3_DeleteBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageS3Server).DeleteBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageS3_DeleteBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageS3Server).DeleteBucket(ctx, req.(*DeleteBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageS3_EmptyBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageS3Server).EmptyBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageS3_EmptyBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageS3Server).EmptyBucket(ctx, req.(*EmptyBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageS3_ListBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageS3Server).ListBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageS3_ListBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageS3Server).ListBucket(ctx, req.(*ListBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageS3_S3StorageInBucketList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S3StorageInBucketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageS3Server).S3StorageInBucketList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageS3_S3StorageInBucketList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageS3Server).S3StorageInBucketList(ctx, req.(*S3StorageInBucketListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageS3_S3StorageUploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S3StorageUploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageS3Server).S3StorageUploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageS3_S3StorageUploadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageS3Server).S3StorageUploadFile(ctx, req.(*S3StorageUploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageS3_S3StorageMkdir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S3StorageMkdirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageS3Server).S3StorageMkdir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageS3_S3StorageMkdir_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageS3Server).S3StorageMkdir(ctx, req.(*S3StorageMkdirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageS3_S3StorageDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S3StorageDownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageS3Server).S3StorageDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageS3_S3StorageDownload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageS3Server).S3StorageDownload(ctx, req.(*S3StorageDownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageS3_S3StorageDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S3StorageDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageS3Server).S3StorageDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageS3_S3StorageDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageS3Server).S3StorageDelete(ctx, req.(*S3StorageDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageS3_ServiceDesc is the grpc.ServiceDesc for StorageS3 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageS3_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.server.compute.v1.StorageS3",
	HandlerType: (*StorageS3Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateS3Key",
			Handler:    _StorageS3_CreateS3Key_Handler,
		},
		{
			MethodName: "GetUserS3User",
			Handler:    _StorageS3_GetUserS3User_Handler,
		},
		{
			MethodName: "CreateBucket",
			Handler:    _StorageS3_CreateBucket_Handler,
		},
		{
			MethodName: "DeleteBucket",
			Handler:    _StorageS3_DeleteBucket_Handler,
		},
		{
			MethodName: "EmptyBucket",
			Handler:    _StorageS3_EmptyBucket_Handler,
		},
		{
			MethodName: "ListBucket",
			Handler:    _StorageS3_ListBucket_Handler,
		},
		{
			MethodName: "S3StorageInBucketList",
			Handler:    _StorageS3_S3StorageInBucketList_Handler,
		},
		{
			MethodName: "S3StorageUploadFile",
			Handler:    _StorageS3_S3StorageUploadFile_Handler,
		},
		{
			MethodName: "S3StorageMkdir",
			Handler:    _StorageS3_S3StorageMkdir_Handler,
		},
		{
			MethodName: "S3StorageDownload",
			Handler:    _StorageS3_S3StorageDownload_Handler,
		},
		{
			MethodName: "S3StorageDelete",
			Handler:    _StorageS3_S3StorageDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/compute/v1/storage_s3.proto",
}
