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
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAgentCreateAgent = "/api.agent.v1.Agent/CreateAgent"

type AgentHTTPServer interface {
	CreateAgent(context.Context, *CreateAgentRequest) (*CreateAgentReply, error)
}

func RegisterAgentHTTPServer(s *http.Server, srv AgentHTTPServer) {
	r := s.Route("/")
	r.POST("/agent", _Agent_CreateAgent0_HTTP_Handler(srv))
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

type AgentHTTPClient interface {
	CreateAgent(ctx context.Context, req *CreateAgentRequest, opts ...http.CallOption) (rsp *CreateAgentReply, err error)
}

type AgentHTTPClientImpl struct {
	cc *http.Client
}

func NewAgentHTTPClient(client *http.Client) AgentHTTPClient {
	return &AgentHTTPClientImpl{client}
}

func (c *AgentHTTPClientImpl) CreateAgent(ctx context.Context, in *CreateAgentRequest, opts ...http.CallOption) (*CreateAgentReply, error) {
	var out CreateAgentReply
	pattern := "/agent"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAgentCreateAgent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
