// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.2
// source: api/network_mapping/v1/network_mapping.proto

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

const OperationNetworkMappingCreateNetworkMapping = "/api.server.network_mapping.v1.NetworkMapping/CreateNetworkMapping"
const OperationNetworkMappingDeleteNetworkMapping = "/api.server.network_mapping.v1.NetworkMapping/DeleteNetworkMapping"
const OperationNetworkMappingGetNetworkMapping = "/api.server.network_mapping.v1.NetworkMapping/GetNetworkMapping"
const OperationNetworkMappingNextNetworkMapping = "/api.server.network_mapping.v1.NetworkMapping/NextNetworkMapping"
const OperationNetworkMappingPageNetworkMapping = "/api.server.network_mapping.v1.NetworkMapping/PageNetworkMapping"
const OperationNetworkMappingUpdateNetworkMapping = "/api.server.network_mapping.v1.NetworkMapping/UpdateNetworkMapping"

type NetworkMappingHTTPServer interface {
	CreateNetworkMapping(context.Context, *CreateNetworkMappingRequest) (*CreateNetworkMappingReply, error)
	DeleteNetworkMapping(context.Context, *DeleteNetworkMappingRequest) (*DeleteNetworkMappingReply, error)
	GetNetworkMapping(context.Context, *GetNetworkMappingRequest) (*GetNetworkMappingReply, error)
	NextNetworkMapping(context.Context, *NextNetworkMappingRequest) (*NextNetworkMappingReply, error)
	PageNetworkMapping(context.Context, *PageNetworkMappingRequest) (*PageNetworkMappingReply, error)
	UpdateNetworkMapping(context.Context, *UpdateNetworkMappingRequest) (*UpdateNetworkMappingReply, error)
}

func RegisterNetworkMappingHTTPServer(s *http.Server, srv NetworkMappingHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/network-mappings", _NetworkMapping_CreateNetworkMapping0_HTTP_Handler(srv))
	r.GET("/v1/network-mappings/page", _NetworkMapping_PageNetworkMapping0_HTTP_Handler(srv))
	r.GET("/v1/network-mappings/next", _NetworkMapping_NextNetworkMapping0_HTTP_Handler(srv))
	r.GET("/v1/network-mappings/{id}", _NetworkMapping_GetNetworkMapping0_HTTP_Handler(srv))
	r.DELETE("/v1/network-mappings/{id}", _NetworkMapping_DeleteNetworkMapping0_HTTP_Handler(srv))
	r.PUT("/v1/network-mappings/{id}", _NetworkMapping_UpdateNetworkMapping0_HTTP_Handler(srv))
}

func _NetworkMapping_CreateNetworkMapping0_HTTP_Handler(srv NetworkMappingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateNetworkMappingRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNetworkMappingCreateNetworkMapping)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateNetworkMapping(ctx, req.(*CreateNetworkMappingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateNetworkMappingReply)
		return ctx.Result(200, reply)
	}
}

func _NetworkMapping_PageNetworkMapping0_HTTP_Handler(srv NetworkMappingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PageNetworkMappingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNetworkMappingPageNetworkMapping)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PageNetworkMapping(ctx, req.(*PageNetworkMappingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PageNetworkMappingReply)
		return ctx.Result(200, reply)
	}
}

func _NetworkMapping_NextNetworkMapping0_HTTP_Handler(srv NetworkMappingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NextNetworkMappingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNetworkMappingNextNetworkMapping)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NextNetworkMapping(ctx, req.(*NextNetworkMappingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NextNetworkMappingReply)
		return ctx.Result(200, reply)
	}
}

func _NetworkMapping_GetNetworkMapping0_HTTP_Handler(srv NetworkMappingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetNetworkMappingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNetworkMappingGetNetworkMapping)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNetworkMapping(ctx, req.(*GetNetworkMappingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetNetworkMappingReply)
		return ctx.Result(200, reply)
	}
}

func _NetworkMapping_DeleteNetworkMapping0_HTTP_Handler(srv NetworkMappingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteNetworkMappingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNetworkMappingDeleteNetworkMapping)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteNetworkMapping(ctx, req.(*DeleteNetworkMappingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteNetworkMappingReply)
		return ctx.Result(200, reply)
	}
}

func _NetworkMapping_UpdateNetworkMapping0_HTTP_Handler(srv NetworkMappingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateNetworkMappingRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNetworkMappingUpdateNetworkMapping)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateNetworkMapping(ctx, req.(*UpdateNetworkMappingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateNetworkMappingReply)
		return ctx.Result(200, reply)
	}
}

type NetworkMappingHTTPClient interface {
	CreateNetworkMapping(ctx context.Context, req *CreateNetworkMappingRequest, opts ...http.CallOption) (rsp *CreateNetworkMappingReply, err error)
	DeleteNetworkMapping(ctx context.Context, req *DeleteNetworkMappingRequest, opts ...http.CallOption) (rsp *DeleteNetworkMappingReply, err error)
	GetNetworkMapping(ctx context.Context, req *GetNetworkMappingRequest, opts ...http.CallOption) (rsp *GetNetworkMappingReply, err error)
	NextNetworkMapping(ctx context.Context, req *NextNetworkMappingRequest, opts ...http.CallOption) (rsp *NextNetworkMappingReply, err error)
	PageNetworkMapping(ctx context.Context, req *PageNetworkMappingRequest, opts ...http.CallOption) (rsp *PageNetworkMappingReply, err error)
	UpdateNetworkMapping(ctx context.Context, req *UpdateNetworkMappingRequest, opts ...http.CallOption) (rsp *UpdateNetworkMappingReply, err error)
}

type NetworkMappingHTTPClientImpl struct {
	cc *http.Client
}

func NewNetworkMappingHTTPClient(client *http.Client) NetworkMappingHTTPClient {
	return &NetworkMappingHTTPClientImpl{client}
}

func (c *NetworkMappingHTTPClientImpl) CreateNetworkMapping(ctx context.Context, in *CreateNetworkMappingRequest, opts ...http.CallOption) (*CreateNetworkMappingReply, error) {
	var out CreateNetworkMappingReply
	pattern := "/v1/network-mappings"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationNetworkMappingCreateNetworkMapping))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NetworkMappingHTTPClientImpl) DeleteNetworkMapping(ctx context.Context, in *DeleteNetworkMappingRequest, opts ...http.CallOption) (*DeleteNetworkMappingReply, error) {
	var out DeleteNetworkMappingReply
	pattern := "/v1/network-mappings/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNetworkMappingDeleteNetworkMapping))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NetworkMappingHTTPClientImpl) GetNetworkMapping(ctx context.Context, in *GetNetworkMappingRequest, opts ...http.CallOption) (*GetNetworkMappingReply, error) {
	var out GetNetworkMappingReply
	pattern := "/v1/network-mappings/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNetworkMappingGetNetworkMapping))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NetworkMappingHTTPClientImpl) NextNetworkMapping(ctx context.Context, in *NextNetworkMappingRequest, opts ...http.CallOption) (*NextNetworkMappingReply, error) {
	var out NextNetworkMappingReply
	pattern := "/v1/network-mappings/next"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNetworkMappingNextNetworkMapping))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NetworkMappingHTTPClientImpl) PageNetworkMapping(ctx context.Context, in *PageNetworkMappingRequest, opts ...http.CallOption) (*PageNetworkMappingReply, error) {
	var out PageNetworkMappingReply
	pattern := "/v1/network-mappings/page"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNetworkMappingPageNetworkMapping))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NetworkMappingHTTPClientImpl) UpdateNetworkMapping(ctx context.Context, in *UpdateNetworkMappingRequest, opts ...http.CallOption) (*UpdateNetworkMappingReply, error) {
	var out UpdateNetworkMappingReply
	pattern := "/v1/network-mappings/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationNetworkMappingUpdateNetworkMapping))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
