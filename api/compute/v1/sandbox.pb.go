// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: api/compute/v1/sandbox.proto

package v1

import (
	v1 "github.com/mohaijiang/computeshare-server/api/network_mapping/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateSandboxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instance       *CreateInstanceRequest            `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	NetworkMapping []*v1.CreateNetworkMappingRequest `protobuf:"bytes,2,rep,name=networkMapping,proto3" json:"networkMapping,omitempty"`
}

func (x *CreateSandboxRequest) Reset() {
	*x = CreateSandboxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_compute_v1_sandbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSandboxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSandboxRequest) ProtoMessage() {}

func (x *CreateSandboxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_compute_v1_sandbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSandboxRequest.ProtoReflect.Descriptor instead.
func (*CreateSandboxRequest) Descriptor() ([]byte, []int) {
	return file_api_compute_v1_sandbox_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSandboxRequest) GetInstance() *CreateInstanceRequest {
	if x != nil {
		return x.Instance
	}
	return nil
}

func (x *CreateSandboxRequest) GetNetworkMapping() []*v1.CreateNetworkMappingRequest {
	if x != nil {
		return x.NetworkMapping
	}
	return nil
}

type CreateSandboxReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                                       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string                                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *CreateSandboxReply_CreateSandboxReply_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateSandboxReply) Reset() {
	*x = CreateSandboxReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_compute_v1_sandbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSandboxReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSandboxReply) ProtoMessage() {}

func (x *CreateSandboxReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_compute_v1_sandbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSandboxReply.ProtoReflect.Descriptor instead.
func (*CreateSandboxReply) Descriptor() ([]byte, []int) {
	return file_api_compute_v1_sandbox_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSandboxReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateSandboxReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateSandboxReply) GetData() *CreateSandboxReply_CreateSandboxReply_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateSandboxReply_CreateSandboxReply_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId      string                                             `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	NetworkMappings []*CreateSandboxReply_CreateSandbox_NetworkMapping `protobuf:"bytes,2,rep,name=network_mappings,json=networkMappings,proto3" json:"network_mappings,omitempty"`
}

func (x *CreateSandboxReply_CreateSandboxReply_Data) Reset() {
	*x = CreateSandboxReply_CreateSandboxReply_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_compute_v1_sandbox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSandboxReply_CreateSandboxReply_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSandboxReply_CreateSandboxReply_Data) ProtoMessage() {}

func (x *CreateSandboxReply_CreateSandboxReply_Data) ProtoReflect() protoreflect.Message {
	mi := &file_api_compute_v1_sandbox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSandboxReply_CreateSandboxReply_Data.ProtoReflect.Descriptor instead.
func (*CreateSandboxReply_CreateSandboxReply_Data) Descriptor() ([]byte, []int) {
	return file_api_compute_v1_sandbox_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CreateSandboxReply_CreateSandboxReply_Data) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *CreateSandboxReply_CreateSandboxReply_Data) GetNetworkMappings() []*CreateSandboxReply_CreateSandbox_NetworkMapping {
	if x != nil {
		return x.NetworkMappings
	}
	return nil
}

type CreateSandboxReply_CreateSandbox_NetworkMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ComputerPort int32  `protobuf:"varint,3,opt,name=computer_port,json=computerPort,proto3" json:"computer_port,omitempty"`
	ServerIp     string `protobuf:"bytes,4,opt,name=server_ip,json=serverIp,proto3" json:"server_ip,omitempty"`
	ServerPort   int32  `protobuf:"varint,5,opt,name=server_port,json=serverPort,proto3" json:"server_port,omitempty"`
}

func (x *CreateSandboxReply_CreateSandbox_NetworkMapping) Reset() {
	*x = CreateSandboxReply_CreateSandbox_NetworkMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_compute_v1_sandbox_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSandboxReply_CreateSandbox_NetworkMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSandboxReply_CreateSandbox_NetworkMapping) ProtoMessage() {}

func (x *CreateSandboxReply_CreateSandbox_NetworkMapping) ProtoReflect() protoreflect.Message {
	mi := &file_api_compute_v1_sandbox_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSandboxReply_CreateSandbox_NetworkMapping.ProtoReflect.Descriptor instead.
func (*CreateSandboxReply_CreateSandbox_NetworkMapping) Descriptor() ([]byte, []int) {
	return file_api_compute_v1_sandbox_proto_rawDescGZIP(), []int{1, 1}
}

func (x *CreateSandboxReply_CreateSandbox_NetworkMapping) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateSandboxReply_CreateSandbox_NetworkMapping) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSandboxReply_CreateSandbox_NetworkMapping) GetComputerPort() int32 {
	if x != nil {
		return x.ComputerPort
	}
	return 0
}

func (x *CreateSandboxReply_CreateSandbox_NetworkMapping) GetServerIp() string {
	if x != nil {
		return x.ServerIp
	}
	return ""
}

func (x *CreateSandboxReply_CreateSandbox_NetworkMapping) GetServerPort() int32 {
	if x != nil {
		return x.ServerPort
	}
	return 0
}

var File_api_compute_v1_sandbox_proto protoreflect.FileDescriptor

var file_api_compute_v1_sandbox_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x48, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x0e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x22,
	0xf0, 0x03, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x55, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x41, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x5f, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0xac, 0x01, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x5f, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x71, 0x0a, 0x10, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x46, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x5f, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0xa5, 0x01, 0x0a, 0x1c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x5f, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f,
	0x72, 0x74, 0x32, 0x8a, 0x01, 0x0a, 0x07, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x12, 0x7f,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x12,
	0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61,
	0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6e, 0x64, 0x62,
	0x6f, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a,
	0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x42,
	0x4f, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x6f, 0x68, 0x61, 0x69, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_compute_v1_sandbox_proto_rawDescOnce sync.Once
	file_api_compute_v1_sandbox_proto_rawDescData = file_api_compute_v1_sandbox_proto_rawDesc
)

func file_api_compute_v1_sandbox_proto_rawDescGZIP() []byte {
	file_api_compute_v1_sandbox_proto_rawDescOnce.Do(func() {
		file_api_compute_v1_sandbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_compute_v1_sandbox_proto_rawDescData)
	})
	return file_api_compute_v1_sandbox_proto_rawDescData
}

var file_api_compute_v1_sandbox_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_compute_v1_sandbox_proto_goTypes = []interface{}{
	(*CreateSandboxRequest)(nil),                            // 0: api.server.compute.v1.CreateSandboxRequest
	(*CreateSandboxReply)(nil),                              // 1: api.server.compute.v1.CreateSandboxReply
	(*CreateSandboxReply_CreateSandboxReply_Data)(nil),      // 2: api.server.compute.v1.CreateSandboxReply.CreateSandboxReply_Data
	(*CreateSandboxReply_CreateSandbox_NetworkMapping)(nil), // 3: api.server.compute.v1.CreateSandboxReply.CreateSandbox_NetworkMapping
	(*CreateInstanceRequest)(nil),                           // 4: api.server.compute.v1.CreateInstanceRequest
	(*v1.CreateNetworkMappingRequest)(nil),                  // 5: api.server.network_mapping.v1.CreateNetworkMappingRequest
}
var file_api_compute_v1_sandbox_proto_depIdxs = []int32{
	4, // 0: api.server.compute.v1.CreateSandboxRequest.instance:type_name -> api.server.compute.v1.CreateInstanceRequest
	5, // 1: api.server.compute.v1.CreateSandboxRequest.networkMapping:type_name -> api.server.network_mapping.v1.CreateNetworkMappingRequest
	2, // 2: api.server.compute.v1.CreateSandboxReply.data:type_name -> api.server.compute.v1.CreateSandboxReply.CreateSandboxReply_Data
	3, // 3: api.server.compute.v1.CreateSandboxReply.CreateSandboxReply_Data.network_mappings:type_name -> api.server.compute.v1.CreateSandboxReply.CreateSandbox_NetworkMapping
	0, // 4: api.server.compute.v1.Sandbox.CreateSandbox:input_type -> api.server.compute.v1.CreateSandboxRequest
	1, // 5: api.server.compute.v1.Sandbox.CreateSandbox:output_type -> api.server.compute.v1.CreateSandboxReply
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_compute_v1_sandbox_proto_init() }
func file_api_compute_v1_sandbox_proto_init() {
	if File_api_compute_v1_sandbox_proto != nil {
		return
	}
	file_api_compute_v1_compute_instance_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_compute_v1_sandbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSandboxRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_compute_v1_sandbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSandboxReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_compute_v1_sandbox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSandboxReply_CreateSandboxReply_Data); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_compute_v1_sandbox_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSandboxReply_CreateSandbox_NetworkMapping); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_compute_v1_sandbox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_compute_v1_sandbox_proto_goTypes,
		DependencyIndexes: file_api_compute_v1_sandbox_proto_depIdxs,
		MessageInfos:      file_api_compute_v1_sandbox_proto_msgTypes,
	}.Build()
	File_api_compute_v1_sandbox_proto = out.File
	file_api_compute_v1_sandbox_proto_rawDesc = nil
	file_api_compute_v1_sandbox_proto_goTypes = nil
	file_api_compute_v1_sandbox_proto_depIdxs = nil
}
