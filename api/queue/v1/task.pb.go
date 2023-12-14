// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: api/queue/v1/task.proto

package v1

import (
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

type TaskCmd int32

const (
	TaskCmd_VM_CREATE          TaskCmd = 0  // 创建虚拟机
	TaskCmd_VM_DELETE          TaskCmd = 1  // 删除虚拟机
	TaskCmd_VM_START           TaskCmd = 2  // 启动虚拟机
	TaskCmd_VM_SHUTDOWN        TaskCmd = 3  //关闭虚拟机
	TaskCmd_VM_RESTART         TaskCmd = 4  //关闭虚拟机
	TaskCmd_VM_VNC_CONNECT     TaskCmd = 5  // vnc 连接
	TaskCmd_NAT_PROXY_CREATE   TaskCmd = 6  // nat 代理创建
	TaskCmd_NAT_PROXY_DELETE   TaskCmd = 7  // nat 代理删除
	TaskCmd_NAT_VISITOR_CREATE TaskCmd = 8  // nat 访问创建
	TaskCmd_NAT_VISITOR_DELETE TaskCmd = 9  // nat 访问删除
	TaskCmd_STORAGE_CREATE     TaskCmd = 10 // 存储节点创建（虚拟机)
	TaskCmd_STORAGE_DELETE     TaskCmd = 11 // 存储节点删除（虚拟机）
	TaskCmd_PING               TaskCmd = 99 // ping
)

// Enum value maps for TaskCmd.
var (
	TaskCmd_name = map[int32]string{
		0:  "VM_CREATE",
		1:  "VM_DELETE",
		2:  "VM_START",
		3:  "VM_SHUTDOWN",
		4:  "VM_RESTART",
		5:  "VM_VNC_CONNECT",
		6:  "NAT_PROXY_CREATE",
		7:  "NAT_PROXY_DELETE",
		8:  "NAT_VISITOR_CREATE",
		9:  "NAT_VISITOR_DELETE",
		10: "STORAGE_CREATE",
		11: "STORAGE_DELETE",
		99: "PING",
	}
	TaskCmd_value = map[string]int32{
		"VM_CREATE":          0,
		"VM_DELETE":          1,
		"VM_START":           2,
		"VM_SHUTDOWN":        3,
		"VM_RESTART":         4,
		"VM_VNC_CONNECT":     5,
		"NAT_PROXY_CREATE":   6,
		"NAT_PROXY_DELETE":   7,
		"NAT_VISITOR_CREATE": 8,
		"NAT_VISITOR_DELETE": 9,
		"STORAGE_CREATE":     10,
		"STORAGE_DELETE":     11,
		"PING":               99,
	}
)

func (x TaskCmd) Enum() *TaskCmd {
	p := new(TaskCmd)
	*p = x
	return p
}

func (x TaskCmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskCmd) Descriptor() protoreflect.EnumDescriptor {
	return file_api_queue_v1_task_proto_enumTypes[0].Descriptor()
}

func (TaskCmd) Type() protoreflect.EnumType {
	return &file_api_queue_v1_task_proto_enumTypes[0]
}

func (x TaskCmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskCmd.Descriptor instead.
func (TaskCmd) EnumDescriptor() ([]byte, []int) {
	return file_api_queue_v1_task_proto_rawDescGZIP(), []int{0}
}

type TaskStatus int32

const (
	TaskStatus_CREATED   TaskStatus = 0 //创建
	TaskStatus_EXECUTING TaskStatus = 1 //执行中
	TaskStatus_EXECUTED  TaskStatus = 2 // 执行成功
	TaskStatus_FAILED    TaskStatus = 3 // 执行失败
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "CREATED",
		1: "EXECUTING",
		2: "EXECUTED",
		3: "FAILED",
	}
	TaskStatus_value = map[string]int32{
		"CREATED":   0,
		"EXECUTING": 1,
		"EXECUTED":  2,
		"FAILED":    3,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_queue_v1_task_proto_enumTypes[1].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_api_queue_v1_task_proto_enumTypes[1]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_queue_v1_task_proto_rawDescGZIP(), []int{1}
}

type QueueTaskGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *QueueTaskGetRequest) Reset() {
	*x = QueueTaskGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_queue_v1_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueTaskGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueTaskGetRequest) ProtoMessage() {}

func (x *QueueTaskGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_queue_v1_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueTaskGetRequest.ProtoReflect.Descriptor instead.
func (*QueueTaskGetRequest) Descriptor() ([]byte, []int) {
	return file_api_queue_v1_task_proto_rawDescGZIP(), []int{0}
}

func (x *QueueTaskGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type QueueTaskGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *QueueTaskVo `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *QueueTaskGetResponse) Reset() {
	*x = QueueTaskGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_queue_v1_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueTaskGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueTaskGetResponse) ProtoMessage() {}

func (x *QueueTaskGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_queue_v1_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueTaskGetResponse.ProtoReflect.Descriptor instead.
func (*QueueTaskGetResponse) Descriptor() ([]byte, []int) {
	return file_api_queue_v1_task_proto_rawDescGZIP(), []int{1}
}

func (x *QueueTaskGetResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *QueueTaskGetResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueueTaskGetResponse) GetData() *QueueTaskVo {
	if x != nil {
		return x.Data
	}
	return nil
}

type QueueTaskVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                                             // 任务id
	AgentId string     `protobuf:"bytes,2,opt,name=agentId,proto3" json:"agentId,omitempty"`                                                   // 客户端节点id
	Cmd     TaskCmd    `protobuf:"varint,3,opt,name=cmd,proto3,enum=github.com.mohaijiang.api.queue.v1.TaskCmd" json:"cmd,omitempty"`          // 执行任务枚举
	Params  string     `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`                                                     // 执行参数，nat 网络类型对应 NatProxyCreateVO, 虚拟机类型对应 ComputeInstanceTaskParamVO
	Status  TaskStatus `protobuf:"varint,5,opt,name=status,proto3,enum=github.com.mohaijiang.api.queue.v1.TaskStatus" json:"status,omitempty"` // 任务状态
}

func (x *QueueTaskVo) Reset() {
	*x = QueueTaskVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_queue_v1_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueTaskVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueTaskVo) ProtoMessage() {}

func (x *QueueTaskVo) ProtoReflect() protoreflect.Message {
	mi := &file_api_queue_v1_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueTaskVo.ProtoReflect.Descriptor instead.
func (*QueueTaskVo) Descriptor() ([]byte, []int) {
	return file_api_queue_v1_task_proto_rawDescGZIP(), []int{2}
}

func (x *QueueTaskVo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *QueueTaskVo) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *QueueTaskVo) GetCmd() TaskCmd {
	if x != nil {
		return x.Cmd
	}
	return TaskCmd_VM_CREATE
}

func (x *QueueTaskVo) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (x *QueueTaskVo) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_CREATED
}

type QueueTaskUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AgentId string     `protobuf:"bytes,2,opt,name=agentId,proto3" json:"agentId,omitempty"`
	Status  TaskStatus `protobuf:"varint,3,opt,name=status,proto3,enum=github.com.mohaijiang.api.queue.v1.TaskStatus" json:"status,omitempty"`
}

func (x *QueueTaskUpdateRequest) Reset() {
	*x = QueueTaskUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_queue_v1_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueTaskUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueTaskUpdateRequest) ProtoMessage() {}

func (x *QueueTaskUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_queue_v1_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueTaskUpdateRequest.ProtoReflect.Descriptor instead.
func (*QueueTaskUpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_queue_v1_task_proto_rawDescGZIP(), []int{3}
}

func (x *QueueTaskUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *QueueTaskUpdateRequest) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *QueueTaskUpdateRequest) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_CREATED
}

type QueueTaskUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *QueueTaskUpdateResponse) Reset() {
	*x = QueueTaskUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_queue_v1_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueTaskUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueTaskUpdateResponse) ProtoMessage() {}

func (x *QueueTaskUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_queue_v1_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueTaskUpdateResponse.ProtoReflect.Descriptor instead.
func (*QueueTaskUpdateResponse) Descriptor() ([]byte, []int) {
	return file_api_queue_v1_task_proto_rawDescGZIP(), []int{4}
}

func (x *QueueTaskUpdateResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *QueueTaskUpdateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type NatNetworkMappingTaskParamVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	InstanceId   string `protobuf:"bytes,3,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	InstancePort int32  `protobuf:"varint,4,opt,name=instancePort,proto3" json:"instancePort,omitempty"`
	RemotePort   int32  `protobuf:"varint,5,opt,name=remotePort,proto3" json:"remotePort,omitempty"`
	GatewayId    string `protobuf:"bytes,6,opt,name=gatewayId,proto3" json:"gatewayId,omitempty"`
	GatewayIp    string `protobuf:"bytes,7,opt,name=gatewayIp,proto3" json:"gatewayIp,omitempty"`
	GatewayPort  int32  `protobuf:"varint,8,opt,name=gatewayPort,proto3" json:"gatewayPort,omitempty"`
}

func (x *NatNetworkMappingTaskParamVO) Reset() {
	*x = NatNetworkMappingTaskParamVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_queue_v1_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NatNetworkMappingTaskParamVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NatNetworkMappingTaskParamVO) ProtoMessage() {}

func (x *NatNetworkMappingTaskParamVO) ProtoReflect() protoreflect.Message {
	mi := &file_api_queue_v1_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NatNetworkMappingTaskParamVO.ProtoReflect.Descriptor instead.
func (*NatNetworkMappingTaskParamVO) Descriptor() ([]byte, []int) {
	return file_api_queue_v1_task_proto_rawDescGZIP(), []int{5}
}

func (x *NatNetworkMappingTaskParamVO) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NatNetworkMappingTaskParamVO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NatNetworkMappingTaskParamVO) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *NatNetworkMappingTaskParamVO) GetInstancePort() int32 {
	if x != nil {
		return x.InstancePort
	}
	return 0
}

func (x *NatNetworkMappingTaskParamVO) GetRemotePort() int32 {
	if x != nil {
		return x.RemotePort
	}
	return 0
}

func (x *NatNetworkMappingTaskParamVO) GetGatewayId() string {
	if x != nil {
		return x.GatewayId
	}
	return ""
}

func (x *NatNetworkMappingTaskParamVO) GetGatewayIp() string {
	if x != nil {
		return x.GatewayIp
	}
	return ""
}

func (x *NatNetworkMappingTaskParamVO) GetGatewayPort() int32 {
	if x != nil {
		return x.GatewayPort
	}
	return 0
}

type ComputeInstanceTaskParamVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cpu        int64  `protobuf:"varint,3,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory     int64  `protobuf:"varint,4,opt,name=memory,proto3" json:"memory,omitempty"`
	Image      string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	PublicKey  string `protobuf:"bytes,6,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	Password   string `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	InstanceId string `protobuf:"bytes,8,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
}

func (x *ComputeInstanceTaskParamVO) Reset() {
	*x = ComputeInstanceTaskParamVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_queue_v1_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeInstanceTaskParamVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeInstanceTaskParamVO) ProtoMessage() {}

func (x *ComputeInstanceTaskParamVO) ProtoReflect() protoreflect.Message {
	mi := &file_api_queue_v1_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeInstanceTaskParamVO.ProtoReflect.Descriptor instead.
func (*ComputeInstanceTaskParamVO) Descriptor() ([]byte, []int) {
	return file_api_queue_v1_task_proto_rawDescGZIP(), []int{6}
}

func (x *ComputeInstanceTaskParamVO) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ComputeInstanceTaskParamVO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ComputeInstanceTaskParamVO) GetCpu() int64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *ComputeInstanceTaskParamVO) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *ComputeInstanceTaskParamVO) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ComputeInstanceTaskParamVO) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *ComputeInstanceTaskParamVO) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ComputeInstanceTaskParamVO) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

type StorageSetupTaskParamVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MasterServer string `protobuf:"bytes,2,opt,name=masterServer,proto3" json:"masterServer,omitempty"`
	PublicIp     string `protobuf:"bytes,3,opt,name=publicIp,proto3" json:"publicIp,omitempty"`
	PublicPort   int32  `protobuf:"varint,4,opt,name=publicPort,proto3" json:"publicPort,omitempty"`
	GrpcPort     int32  `protobuf:"varint,5,opt,name=grpcPort,proto3" json:"grpcPort,omitempty"`
}

func (x *StorageSetupTaskParamVO) Reset() {
	*x = StorageSetupTaskParamVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_queue_v1_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageSetupTaskParamVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageSetupTaskParamVO) ProtoMessage() {}

func (x *StorageSetupTaskParamVO) ProtoReflect() protoreflect.Message {
	mi := &file_api_queue_v1_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageSetupTaskParamVO.ProtoReflect.Descriptor instead.
func (*StorageSetupTaskParamVO) Descriptor() ([]byte, []int) {
	return file_api_queue_v1_task_proto_rawDescGZIP(), []int{7}
}

func (x *StorageSetupTaskParamVO) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StorageSetupTaskParamVO) GetMasterServer() string {
	if x != nil {
		return x.MasterServer
	}
	return ""
}

func (x *StorageSetupTaskParamVO) GetPublicIp() string {
	if x != nil {
		return x.PublicIp
	}
	return ""
}

func (x *StorageSetupTaskParamVO) GetPublicPort() int32 {
	if x != nil {
		return x.PublicPort
	}
	return 0
}

func (x *StorageSetupTaskParamVO) GetGrpcPort() int32 {
	if x != nil {
		return x.GrpcPort
	}
	return 0
}

var File_api_queue_v1_task_proto protoreflect.FileDescriptor

var file_api_queue_v1_task_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x68, 0x61, 0x69, 0x6a, 0x69, 0x61, 0x6e, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x13, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x75, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x68, 0x61, 0x69, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd6,
	0x01, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x68, 0x61, 0x69, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43,
	0x6d, 0x64, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x46, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2e, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x68,
	0x61, 0x69, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x68, 0x61, 0x69, 0x6a,
	0x69, 0x61, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x47, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x75, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x84, 0x02,
	0x0a, 0x1c, 0x4e, 0x61, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x56, 0x4f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x6f,
	0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x50, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49,
	0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x49, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50, 0x6f, 0x72,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x50, 0x6f, 0x72, 0x74, 0x22, 0xda, 0x01, 0x0a, 0x1a, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x56, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x22, 0xa5, 0x01, 0x0a, 0x17, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74,
	0x75, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x56, 0x4f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x67, 0x72, 0x70, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x67, 0x72, 0x70, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x2a, 0xf8, 0x01, 0x0a, 0x07, 0x54, 0x61,
	0x73, 0x6b, 0x43, 0x6d, 0x64, 0x12, 0x0d, 0x0a, 0x09, 0x56, 0x4d, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x56, 0x4d, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54,
	0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10,
	0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x4d, 0x5f, 0x53, 0x48, 0x55, 0x54, 0x44, 0x4f, 0x57, 0x4e,
	0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x4d, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x56, 0x4d, 0x5f, 0x56, 0x4e, 0x43, 0x5f, 0x43, 0x4f, 0x4e,
	0x4e, 0x45, 0x43, 0x54, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x41, 0x54, 0x5f, 0x50, 0x52,
	0x4f, 0x58, 0x59, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10,
	0x4e, 0x41, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x10, 0x07, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x41, 0x54, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x54, 0x4f,
	0x52, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x41,
	0x54, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x54, 0x4f, 0x52, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x10, 0x09, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47,
	0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x0b, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49,
	0x4e, 0x47, 0x10, 0x63, 0x2a, 0x42, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x32, 0xcf, 0x02, 0x0a, 0x09, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x99, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x37, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x68, 0x61, 0x69, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x38, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f,
	0x68, 0x61, 0x69, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x12, 0xa5, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x3a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x68, 0x61, 0x69, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x6f, 0x68, 0x61, 0x69, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x1a, 0x0e, 0x2f, 0x76, 0x31, 0x2f,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x42, 0x4d, 0x0a, 0x0e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x68, 0x61, 0x69,
	0x6a, 0x69, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_queue_v1_task_proto_rawDescOnce sync.Once
	file_api_queue_v1_task_proto_rawDescData = file_api_queue_v1_task_proto_rawDesc
)

func file_api_queue_v1_task_proto_rawDescGZIP() []byte {
	file_api_queue_v1_task_proto_rawDescOnce.Do(func() {
		file_api_queue_v1_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_queue_v1_task_proto_rawDescData)
	})
	return file_api_queue_v1_task_proto_rawDescData
}

var file_api_queue_v1_task_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_queue_v1_task_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_queue_v1_task_proto_goTypes = []interface{}{
	(TaskCmd)(0),                         // 0: github.com.mohaijiang.api.queue.v1.TaskCmd
	(TaskStatus)(0),                      // 1: github.com.mohaijiang.api.queue.v1.TaskStatus
	(*QueueTaskGetRequest)(nil),          // 2: github.com.mohaijiang.api.queue.v1.QueueTaskGetRequest
	(*QueueTaskGetResponse)(nil),         // 3: github.com.mohaijiang.api.queue.v1.QueueTaskGetResponse
	(*QueueTaskVo)(nil),                  // 4: github.com.mohaijiang.api.queue.v1.QueueTaskVo
	(*QueueTaskUpdateRequest)(nil),       // 5: github.com.mohaijiang.api.queue.v1.QueueTaskUpdateRequest
	(*QueueTaskUpdateResponse)(nil),      // 6: github.com.mohaijiang.api.queue.v1.QueueTaskUpdateResponse
	(*NatNetworkMappingTaskParamVO)(nil), // 7: github.com.mohaijiang.api.queue.v1.NatNetworkMappingTaskParamVO
	(*ComputeInstanceTaskParamVO)(nil),   // 8: github.com.mohaijiang.api.queue.v1.ComputeInstanceTaskParamVO
	(*StorageSetupTaskParamVO)(nil),      // 9: github.com.mohaijiang.api.queue.v1.StorageSetupTaskParamVO
}
var file_api_queue_v1_task_proto_depIdxs = []int32{
	4, // 0: github.com.mohaijiang.api.queue.v1.QueueTaskGetResponse.data:type_name -> github.com.mohaijiang.api.queue.v1.QueueTaskVo
	0, // 1: github.com.mohaijiang.api.queue.v1.QueueTaskVo.cmd:type_name -> github.com.mohaijiang.api.queue.v1.TaskCmd
	1, // 2: github.com.mohaijiang.api.queue.v1.QueueTaskVo.status:type_name -> github.com.mohaijiang.api.queue.v1.TaskStatus
	1, // 3: github.com.mohaijiang.api.queue.v1.QueueTaskUpdateRequest.status:type_name -> github.com.mohaijiang.api.queue.v1.TaskStatus
	2, // 4: github.com.mohaijiang.api.queue.v1.QueueTask.GetAgentTask:input_type -> github.com.mohaijiang.api.queue.v1.QueueTaskGetRequest
	5, // 5: github.com.mohaijiang.api.queue.v1.QueueTask.UpdateAgentTask:input_type -> github.com.mohaijiang.api.queue.v1.QueueTaskUpdateRequest
	3, // 6: github.com.mohaijiang.api.queue.v1.QueueTask.GetAgentTask:output_type -> github.com.mohaijiang.api.queue.v1.QueueTaskGetResponse
	6, // 7: github.com.mohaijiang.api.queue.v1.QueueTask.UpdateAgentTask:output_type -> github.com.mohaijiang.api.queue.v1.QueueTaskUpdateResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_queue_v1_task_proto_init() }
func file_api_queue_v1_task_proto_init() {
	if File_api_queue_v1_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_queue_v1_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueTaskGetRequest); i {
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
		file_api_queue_v1_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueTaskGetResponse); i {
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
		file_api_queue_v1_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueTaskVo); i {
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
		file_api_queue_v1_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueTaskUpdateRequest); i {
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
		file_api_queue_v1_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueTaskUpdateResponse); i {
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
		file_api_queue_v1_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NatNetworkMappingTaskParamVO); i {
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
		file_api_queue_v1_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeInstanceTaskParamVO); i {
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
		file_api_queue_v1_task_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageSetupTaskParamVO); i {
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
			RawDescriptor: file_api_queue_v1_task_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_queue_v1_task_proto_goTypes,
		DependencyIndexes: file_api_queue_v1_task_proto_depIdxs,
		EnumInfos:         file_api_queue_v1_task_proto_enumTypes,
		MessageInfos:      file_api_queue_v1_task_proto_msgTypes,
	}.Build()
	File_api_queue_v1_task_proto = out.File
	file_api_queue_v1_task_proto_rawDesc = nil
	file_api_queue_v1_task_proto_goTypes = nil
	file_api_queue_v1_task_proto_depIdxs = nil
}
