// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: api/agent/v1/agent.proto

package v1

import (
	context "context"
	v1 "github.com/mohaijiang/computeshare-server/api/compute/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Agent_CreateAgent_FullMethodName          = "/api.server.agent.v1.Agent/CreateAgent"
	Agent_UpdateAgent_FullMethodName          = "/api.server.agent.v1.Agent/UpdateAgent"
	Agent_DeleteAgent_FullMethodName          = "/api.server.agent.v1.Agent/DeleteAgent"
	Agent_GetAgent_FullMethodName             = "/api.server.agent.v1.Agent/GetAgent"
	Agent_ListAgent_FullMethodName            = "/api.server.agent.v1.Agent/ListAgent"
	Agent_ListAgentInstance_FullMethodName    = "/api.server.agent.v1.Agent/ListAgentInstance"
	Agent_ReportInstanceStatus_FullMethodName = "/api.server.agent.v1.Agent/ReportInstanceStatus"
)

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	CreateAgent(ctx context.Context, in *CreateAgentRequest, opts ...grpc.CallOption) (*CreateAgentReply, error)
	UpdateAgent(ctx context.Context, in *UpdateAgentRequest, opts ...grpc.CallOption) (*UpdateAgentReply, error)
	DeleteAgent(ctx context.Context, in *DeleteAgentRequest, opts ...grpc.CallOption) (*DeleteAgentReply, error)
	GetAgent(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*GetAgentReply, error)
	ListAgent(ctx context.Context, in *ListAgentRequest, opts ...grpc.CallOption) (*ListAgentReply, error)
	ListAgentInstance(ctx context.Context, in *ListAgentInstanceReq, opts ...grpc.CallOption) (*v1.ListInstanceReply, error)
	ReportInstanceStatus(ctx context.Context, in *v1.Instance, opts ...grpc.CallOption) (*ReportInstanceStatusReply, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) CreateAgent(ctx context.Context, in *CreateAgentRequest, opts ...grpc.CallOption) (*CreateAgentReply, error) {
	out := new(CreateAgentReply)
	err := c.cc.Invoke(ctx, Agent_CreateAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) UpdateAgent(ctx context.Context, in *UpdateAgentRequest, opts ...grpc.CallOption) (*UpdateAgentReply, error) {
	out := new(UpdateAgentReply)
	err := c.cc.Invoke(ctx, Agent_UpdateAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) DeleteAgent(ctx context.Context, in *DeleteAgentRequest, opts ...grpc.CallOption) (*DeleteAgentReply, error) {
	out := new(DeleteAgentReply)
	err := c.cc.Invoke(ctx, Agent_DeleteAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetAgent(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*GetAgentReply, error) {
	out := new(GetAgentReply)
	err := c.cc.Invoke(ctx, Agent_GetAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ListAgent(ctx context.Context, in *ListAgentRequest, opts ...grpc.CallOption) (*ListAgentReply, error) {
	out := new(ListAgentReply)
	err := c.cc.Invoke(ctx, Agent_ListAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ListAgentInstance(ctx context.Context, in *ListAgentInstanceReq, opts ...grpc.CallOption) (*v1.ListInstanceReply, error) {
	out := new(v1.ListInstanceReply)
	err := c.cc.Invoke(ctx, Agent_ListAgentInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ReportInstanceStatus(ctx context.Context, in *v1.Instance, opts ...grpc.CallOption) (*ReportInstanceStatusReply, error) {
	out := new(ReportInstanceStatusReply)
	err := c.cc.Invoke(ctx, Agent_ReportInstanceStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	CreateAgent(context.Context, *CreateAgentRequest) (*CreateAgentReply, error)
	UpdateAgent(context.Context, *UpdateAgentRequest) (*UpdateAgentReply, error)
	DeleteAgent(context.Context, *DeleteAgentRequest) (*DeleteAgentReply, error)
	GetAgent(context.Context, *GetAgentRequest) (*GetAgentReply, error)
	ListAgent(context.Context, *ListAgentRequest) (*ListAgentReply, error)
	ListAgentInstance(context.Context, *ListAgentInstanceReq) (*v1.ListInstanceReply, error)
	ReportInstanceStatus(context.Context, *v1.Instance) (*ReportInstanceStatusReply, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) CreateAgent(context.Context, *CreateAgentRequest) (*CreateAgentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgent not implemented")
}
func (UnimplementedAgentServer) UpdateAgent(context.Context, *UpdateAgentRequest) (*UpdateAgentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAgent not implemented")
}
func (UnimplementedAgentServer) DeleteAgent(context.Context, *DeleteAgentRequest) (*DeleteAgentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAgent not implemented")
}
func (UnimplementedAgentServer) GetAgent(context.Context, *GetAgentRequest) (*GetAgentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgent not implemented")
}
func (UnimplementedAgentServer) ListAgent(context.Context, *ListAgentRequest) (*ListAgentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgent not implemented")
}
func (UnimplementedAgentServer) ListAgentInstance(context.Context, *ListAgentInstanceReq) (*v1.ListInstanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgentInstance not implemented")
}
func (UnimplementedAgentServer) ReportInstanceStatus(context.Context, *v1.Instance) (*ReportInstanceStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportInstanceStatus not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_CreateAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CreateAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_CreateAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CreateAgent(ctx, req.(*CreateAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_UpdateAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).UpdateAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_UpdateAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).UpdateAgent(ctx, req.(*UpdateAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_DeleteAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).DeleteAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_DeleteAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).DeleteAgent(ctx, req.(*DeleteAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetAgent(ctx, req.(*GetAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ListAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ListAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ListAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ListAgent(ctx, req.(*ListAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ListAgentInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAgentInstanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ListAgentInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ListAgentInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ListAgentInstance(ctx, req.(*ListAgentInstanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ReportInstanceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Instance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ReportInstanceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ReportInstanceStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ReportInstanceStatus(ctx, req.(*v1.Instance))
	}
	return interceptor(ctx, in, info, handler)
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.server.agent.v1.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAgent",
			Handler:    _Agent_CreateAgent_Handler,
		},
		{
			MethodName: "UpdateAgent",
			Handler:    _Agent_UpdateAgent_Handler,
		},
		{
			MethodName: "DeleteAgent",
			Handler:    _Agent_DeleteAgent_Handler,
		},
		{
			MethodName: "GetAgent",
			Handler:    _Agent_GetAgent_Handler,
		},
		{
			MethodName: "ListAgent",
			Handler:    _Agent_ListAgent_Handler,
		},
		{
			MethodName: "ListAgentInstance",
			Handler:    _Agent_ListAgentInstance_Handler,
		},
		{
			MethodName: "ReportInstanceStatus",
			Handler:    _Agent_ReportInstanceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/agent/v1/agent.proto",
}
