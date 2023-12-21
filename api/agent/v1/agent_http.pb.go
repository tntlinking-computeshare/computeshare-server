// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.2
// source: api/agent/v1/agent.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	v1 "github.com/mohaijiang/computeshare-server/api/compute/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAgentCreateAgent = "/api.server.agent.v1.Agent/CreateAgent"
const OperationAgentDeleteAgent = "/api.server.agent.v1.Agent/DeleteAgent"
const OperationAgentGetAgent = "/api.server.agent.v1.Agent/GetAgent"
const OperationAgentListAgent = "/api.server.agent.v1.Agent/ListAgent"
const OperationAgentListAgentInstance = "/api.server.agent.v1.Agent/ListAgentInstance"
const OperationAgentReportInstanceStatus = "/api.server.agent.v1.Agent/ReportInstanceStatus"
const OperationAgentUpdateAgent = "/api.server.agent.v1.Agent/UpdateAgent"

type AgentHTTPServer interface {
	CreateAgent(context.Context, *CreateAgentRequest) (*CreateAgentReply, error)
	DeleteAgent(context.Context, *DeleteAgentRequest) (*DeleteAgentReply, error)
	GetAgent(context.Context, *GetAgentRequest) (*GetAgentReply, error)
	ListAgent(context.Context, *ListAgentRequest) (*ListAgentReply, error)
	ListAgentInstance(context.Context, *ListAgentInstanceReq) (*v1.ListInstanceReply, error)
	ReportInstanceStatus(context.Context, *v1.Instance) (*ReportInstanceStatusReply, error)
	UpdateAgent(context.Context, *UpdateAgentRequest) (*UpdateAgentReply, error)
}

func RegisterAgentHTTPServer(s *http.Server, srv AgentHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/agent", _Agent_CreateAgent0_HTTP_Handler(srv))
	r.PUT("/v1/agent/{id}", _Agent_UpdateAgent0_HTTP_Handler(srv))
	r.DELETE("/v1/agent/{id}", _Agent_DeleteAgent0_HTTP_Handler(srv))
	r.GET("/v1/agent/{id}", _Agent_GetAgent0_HTTP_Handler(srv))
	r.GET("/v1/agent", _Agent_ListAgent0_HTTP_Handler(srv))
	r.GET("/v1/agent/instance/{mac}", _Agent_ListAgentInstance0_HTTP_Handler(srv))
	r.PUT("/v1/agent/instance/report", _Agent_ReportInstanceStatus0_HTTP_Handler(srv))
}

func _Agent_CreateAgent0_HTTP_Handler(srv AgentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateAgentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAgentCreateAgent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAgent(ctx, req.(*CreateAgentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateAgentReply)
		return ctx.Result(200, reply)
	}
}

func _Agent_UpdateAgent0_HTTP_Handler(srv AgentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateAgentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAgentUpdateAgent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateAgent(ctx, req.(*UpdateAgentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateAgentReply)
		return ctx.Result(200, reply)
	}
}

func _Agent_DeleteAgent0_HTTP_Handler(srv AgentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteAgentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAgentDeleteAgent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteAgent(ctx, req.(*DeleteAgentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteAgentReply)
		return ctx.Result(200, reply)
	}
}

func _Agent_GetAgent0_HTTP_Handler(srv AgentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAgentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAgentGetAgent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAgent(ctx, req.(*GetAgentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAgentReply)
		return ctx.Result(200, reply)
	}
}

func _Agent_ListAgent0_HTTP_Handler(srv AgentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListAgentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAgentListAgent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListAgent(ctx, req.(*ListAgentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAgentReply)
		return ctx.Result(200, reply)
	}
}

func _Agent_ListAgentInstance0_HTTP_Handler(srv AgentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListAgentInstanceReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAgentListAgentInstance)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListAgentInstance(ctx, req.(*ListAgentInstanceReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.ListInstanceReply)
		return ctx.Result(200, reply)
	}
}

func _Agent_ReportInstanceStatus0_HTTP_Handler(srv AgentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.Instance
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAgentReportInstanceStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ReportInstanceStatus(ctx, req.(*v1.Instance))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReportInstanceStatusReply)
		return ctx.Result(200, reply)
	}
}

type AgentHTTPClient interface {
	CreateAgent(ctx context.Context, req *CreateAgentRequest, opts ...http.CallOption) (rsp *CreateAgentReply, err error)
	DeleteAgent(ctx context.Context, req *DeleteAgentRequest, opts ...http.CallOption) (rsp *DeleteAgentReply, err error)
	GetAgent(ctx context.Context, req *GetAgentRequest, opts ...http.CallOption) (rsp *GetAgentReply, err error)
	ListAgent(ctx context.Context, req *ListAgentRequest, opts ...http.CallOption) (rsp *ListAgentReply, err error)
	ListAgentInstance(ctx context.Context, req *ListAgentInstanceReq, opts ...http.CallOption) (rsp *v1.ListInstanceReply, err error)
	ReportInstanceStatus(ctx context.Context, req *v1.Instance, opts ...http.CallOption) (rsp *ReportInstanceStatusReply, err error)
	UpdateAgent(ctx context.Context, req *UpdateAgentRequest, opts ...http.CallOption) (rsp *UpdateAgentReply, err error)
}

type AgentHTTPClientImpl struct {
	cc *http.Client
}

func NewAgentHTTPClient(client *http.Client) AgentHTTPClient {
	return &AgentHTTPClientImpl{client}
}

func (c *AgentHTTPClientImpl) CreateAgent(ctx context.Context, in *CreateAgentRequest, opts ...http.CallOption) (*CreateAgentReply, error) {
	var out CreateAgentReply
	pattern := "/v1/agent"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAgentCreateAgent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AgentHTTPClientImpl) DeleteAgent(ctx context.Context, in *DeleteAgentRequest, opts ...http.CallOption) (*DeleteAgentReply, error) {
	var out DeleteAgentReply
	pattern := "/v1/agent/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAgentDeleteAgent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AgentHTTPClientImpl) GetAgent(ctx context.Context, in *GetAgentRequest, opts ...http.CallOption) (*GetAgentReply, error) {
	var out GetAgentReply
	pattern := "/v1/agent/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAgentGetAgent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AgentHTTPClientImpl) ListAgent(ctx context.Context, in *ListAgentRequest, opts ...http.CallOption) (*ListAgentReply, error) {
	var out ListAgentReply
	pattern := "/v1/agent"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAgentListAgent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AgentHTTPClientImpl) ListAgentInstance(ctx context.Context, in *ListAgentInstanceReq, opts ...http.CallOption) (*v1.ListInstanceReply, error) {
	var out v1.ListInstanceReply
	pattern := "/v1/agent/instance/{mac}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAgentListAgentInstance))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AgentHTTPClientImpl) ReportInstanceStatus(ctx context.Context, in *v1.Instance, opts ...http.CallOption) (*ReportInstanceStatusReply, error) {
	var out ReportInstanceStatusReply
	pattern := "/v1/agent/instance/report"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAgentReportInstanceStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AgentHTTPClientImpl) UpdateAgent(ctx context.Context, in *UpdateAgentRequest, opts ...http.CallOption) (*UpdateAgentReply, error) {
	var out UpdateAgentReply
	pattern := "/v1/agent/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAgentUpdateAgent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
