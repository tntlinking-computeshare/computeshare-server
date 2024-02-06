// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.2
// source: api/dashboard/v1/dashboard.proto

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

const OperationDashboardCyclesCount = "/api.server.dashboard.v1.Dashboard/CyclesCount"
const OperationDashboardGatewaysCount = "/api.server.dashboard.v1.Dashboard/GatewaysCount"
const OperationDashboardGatewaysList = "/api.server.dashboard.v1.Dashboard/GatewaysList"
const OperationDashboardLastComputeInstancesCount = "/api.server.dashboard.v1.Dashboard/LastComputeInstancesCount"
const OperationDashboardProvidersCount = "/api.server.dashboard.v1.Dashboard/ProvidersCount"
const OperationDashboardProvidersList = "/api.server.dashboard.v1.Dashboard/ProvidersList"
const OperationDashboardSandboxCount = "/api.server.dashboard.v1.Dashboard/SandboxCount"
const OperationDashboardStorageBucketsVolumeNumList = "/api.server.dashboard.v1.Dashboard/StorageBucketsVolumeNumList"
const OperationDashboardStorageS3KeyCallCount = "/api.server.dashboard.v1.Dashboard/StorageS3KeyCallCount"
const OperationDashboardStoragesCount = "/api.server.dashboard.v1.Dashboard/StoragesCount"
const OperationDashboardStoragesProvidersList = "/api.server.dashboard.v1.Dashboard/StoragesProvidersList"

type DashboardHTTPServer interface {
	// CyclesCount已发放积分总数 回收积分总数 发放代金券总数 已充值总数
	CyclesCount(context.Context, *CyclesCountRequest) (*CyclesCountReply, error)
	// GatewaysCountGateway总数
	GatewaysCount(context.Context, *GatewaysCountRequest) (*GatewaysCountReply, error)
	// GatewaysListGateway列表 总端口数 已用端口数内网 外网
	GatewaysList(context.Context, *GatewaysListRequest) (*GatewaysListReply, error)
	// LastComputeInstancesCount最新创建虚拟机
	LastComputeInstancesCount(context.Context, *LastComputeInstancesCountRequest) (*LastComputeInstancesCountReply, error)
	// ProvidersCountProvider总数
	ProvidersCount(context.Context, *ProvidersCountRequest) (*ProvidersCountReply, error)
	// ProvidersListProvider列表 类型，规格，是否存活
	ProvidersList(context.Context, *ProvidersListRequest) (*ProvidersListReply, error)
	// SandboxCount沙箱调用总数
	SandboxCount(context.Context, *SandboxCountRequest) (*SandboxCountReply, error)
	// StorageBucketsVolumeNumList存储桶VolumeNum列表
	StorageBucketsVolumeNumList(context.Context, *StorageBucketsVolumeNumListRequest) (*StorageBucketsVolumeNumListReply, error)
	// StorageS3KeyCallCountS3KeyCall
	StorageS3KeyCallCount(context.Context, *StorageS3KeyCallCountRequest) (*StorageS3KeyCallCountReply, error)
	// StoragesCount存储总数 已使用总数
	StoragesCount(context.Context, *StoragesCountRequest) (*StoragesCountReply, error)
	// StoragesProvidersList存储提供者列表
	StoragesProvidersList(context.Context, *StoragesProvidersListRequest) (*StoragesProvidersListReply, error)
}

func RegisterDashboardHTTPServer(s *http.Server, srv DashboardHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/dashboard/providers/count", _Dashboard_ProvidersCount0_HTTP_Handler(srv))
	r.GET("/v1/dashboard/gateways/count", _Dashboard_GatewaysCount0_HTTP_Handler(srv))
	r.GET("/v1/dashboard/storages/count", _Dashboard_StoragesCount0_HTTP_Handler(srv))
	r.GET("/v1/dashboard/providers/volumes/count", _Dashboard_StoragesProvidersList0_HTTP_Handler(srv))
	r.GET("/v1/dashboard/buckets/volumes/count", _Dashboard_StorageBucketsVolumeNumList0_HTTP_Handler(srv))
	r.GET("/v1/dashboard/s3_key/call/count", _Dashboard_StorageS3KeyCallCount0_HTTP_Handler(srv))
	r.GET("/v1/dashboard/providers/list", _Dashboard_ProvidersList0_HTTP_Handler(srv))
	r.GET("/v1/dashboard/gateways/list", _Dashboard_GatewaysList0_HTTP_Handler(srv))
	r.GET("/v1/dashboard/cycles/count", _Dashboard_CyclesCount0_HTTP_Handler(srv))
	r.GET("/v1/dashboard/sandbox/count", _Dashboard_SandboxCount0_HTTP_Handler(srv))
	r.GET("/v1/dashboard/instances/count", _Dashboard_LastComputeInstancesCount0_HTTP_Handler(srv))
}

func _Dashboard_ProvidersCount0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ProvidersCountRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardProvidersCount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ProvidersCount(ctx, req.(*ProvidersCountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProvidersCountReply)
		return ctx.Result(200, reply)
	}
}

func _Dashboard_GatewaysCount0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GatewaysCountRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardGatewaysCount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GatewaysCount(ctx, req.(*GatewaysCountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GatewaysCountReply)
		return ctx.Result(200, reply)
	}
}

func _Dashboard_StoragesCount0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StoragesCountRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardStoragesCount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.StoragesCount(ctx, req.(*StoragesCountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StoragesCountReply)
		return ctx.Result(200, reply)
	}
}

func _Dashboard_StoragesProvidersList0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StoragesProvidersListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardStoragesProvidersList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.StoragesProvidersList(ctx, req.(*StoragesProvidersListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StoragesProvidersListReply)
		return ctx.Result(200, reply)
	}
}

func _Dashboard_StorageBucketsVolumeNumList0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StorageBucketsVolumeNumListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardStorageBucketsVolumeNumList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.StorageBucketsVolumeNumList(ctx, req.(*StorageBucketsVolumeNumListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StorageBucketsVolumeNumListReply)
		return ctx.Result(200, reply)
	}
}

func _Dashboard_StorageS3KeyCallCount0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StorageS3KeyCallCountRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardStorageS3KeyCallCount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.StorageS3KeyCallCount(ctx, req.(*StorageS3KeyCallCountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StorageS3KeyCallCountReply)
		return ctx.Result(200, reply)
	}
}

func _Dashboard_ProvidersList0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ProvidersListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardProvidersList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ProvidersList(ctx, req.(*ProvidersListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProvidersListReply)
		return ctx.Result(200, reply)
	}
}

func _Dashboard_GatewaysList0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GatewaysListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardGatewaysList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GatewaysList(ctx, req.(*GatewaysListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GatewaysListReply)
		return ctx.Result(200, reply)
	}
}

func _Dashboard_CyclesCount0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CyclesCountRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardCyclesCount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CyclesCount(ctx, req.(*CyclesCountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CyclesCountReply)
		return ctx.Result(200, reply)
	}
}

func _Dashboard_SandboxCount0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SandboxCountRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardSandboxCount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SandboxCount(ctx, req.(*SandboxCountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SandboxCountReply)
		return ctx.Result(200, reply)
	}
}

func _Dashboard_LastComputeInstancesCount0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LastComputeInstancesCountRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardLastComputeInstancesCount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LastComputeInstancesCount(ctx, req.(*LastComputeInstancesCountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LastComputeInstancesCountReply)
		return ctx.Result(200, reply)
	}
}

type DashboardHTTPClient interface {
	CyclesCount(ctx context.Context, req *CyclesCountRequest, opts ...http.CallOption) (rsp *CyclesCountReply, err error)
	GatewaysCount(ctx context.Context, req *GatewaysCountRequest, opts ...http.CallOption) (rsp *GatewaysCountReply, err error)
	GatewaysList(ctx context.Context, req *GatewaysListRequest, opts ...http.CallOption) (rsp *GatewaysListReply, err error)
	LastComputeInstancesCount(ctx context.Context, req *LastComputeInstancesCountRequest, opts ...http.CallOption) (rsp *LastComputeInstancesCountReply, err error)
	ProvidersCount(ctx context.Context, req *ProvidersCountRequest, opts ...http.CallOption) (rsp *ProvidersCountReply, err error)
	ProvidersList(ctx context.Context, req *ProvidersListRequest, opts ...http.CallOption) (rsp *ProvidersListReply, err error)
	SandboxCount(ctx context.Context, req *SandboxCountRequest, opts ...http.CallOption) (rsp *SandboxCountReply, err error)
	StorageBucketsVolumeNumList(ctx context.Context, req *StorageBucketsVolumeNumListRequest, opts ...http.CallOption) (rsp *StorageBucketsVolumeNumListReply, err error)
	StorageS3KeyCallCount(ctx context.Context, req *StorageS3KeyCallCountRequest, opts ...http.CallOption) (rsp *StorageS3KeyCallCountReply, err error)
	StoragesCount(ctx context.Context, req *StoragesCountRequest, opts ...http.CallOption) (rsp *StoragesCountReply, err error)
	StoragesProvidersList(ctx context.Context, req *StoragesProvidersListRequest, opts ...http.CallOption) (rsp *StoragesProvidersListReply, err error)
}

type DashboardHTTPClientImpl struct {
	cc *http.Client
}

func NewDashboardHTTPClient(client *http.Client) DashboardHTTPClient {
	return &DashboardHTTPClientImpl{client}
}

func (c *DashboardHTTPClientImpl) CyclesCount(ctx context.Context, in *CyclesCountRequest, opts ...http.CallOption) (*CyclesCountReply, error) {
	var out CyclesCountReply
	pattern := "/v1/dashboard/cycles/count"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDashboardCyclesCount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DashboardHTTPClientImpl) GatewaysCount(ctx context.Context, in *GatewaysCountRequest, opts ...http.CallOption) (*GatewaysCountReply, error) {
	var out GatewaysCountReply
	pattern := "/v1/dashboard/gateways/count"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDashboardGatewaysCount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DashboardHTTPClientImpl) GatewaysList(ctx context.Context, in *GatewaysListRequest, opts ...http.CallOption) (*GatewaysListReply, error) {
	var out GatewaysListReply
	pattern := "/v1/dashboard/gateways/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDashboardGatewaysList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DashboardHTTPClientImpl) LastComputeInstancesCount(ctx context.Context, in *LastComputeInstancesCountRequest, opts ...http.CallOption) (*LastComputeInstancesCountReply, error) {
	var out LastComputeInstancesCountReply
	pattern := "/v1/dashboard/instances/count"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDashboardLastComputeInstancesCount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DashboardHTTPClientImpl) ProvidersCount(ctx context.Context, in *ProvidersCountRequest, opts ...http.CallOption) (*ProvidersCountReply, error) {
	var out ProvidersCountReply
	pattern := "/v1/dashboard/providers/count"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDashboardProvidersCount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DashboardHTTPClientImpl) ProvidersList(ctx context.Context, in *ProvidersListRequest, opts ...http.CallOption) (*ProvidersListReply, error) {
	var out ProvidersListReply
	pattern := "/v1/dashboard/providers/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDashboardProvidersList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DashboardHTTPClientImpl) SandboxCount(ctx context.Context, in *SandboxCountRequest, opts ...http.CallOption) (*SandboxCountReply, error) {
	var out SandboxCountReply
	pattern := "/v1/dashboard/sandbox/count"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDashboardSandboxCount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DashboardHTTPClientImpl) StorageBucketsVolumeNumList(ctx context.Context, in *StorageBucketsVolumeNumListRequest, opts ...http.CallOption) (*StorageBucketsVolumeNumListReply, error) {
	var out StorageBucketsVolumeNumListReply
	pattern := "/v1/dashboard/buckets/volumes/count"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDashboardStorageBucketsVolumeNumList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DashboardHTTPClientImpl) StorageS3KeyCallCount(ctx context.Context, in *StorageS3KeyCallCountRequest, opts ...http.CallOption) (*StorageS3KeyCallCountReply, error) {
	var out StorageS3KeyCallCountReply
	pattern := "/v1/dashboard/s3_key/call/count"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDashboardStorageS3KeyCallCount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DashboardHTTPClientImpl) StoragesCount(ctx context.Context, in *StoragesCountRequest, opts ...http.CallOption) (*StoragesCountReply, error) {
	var out StoragesCountReply
	pattern := "/v1/dashboard/storages/count"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDashboardStoragesCount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DashboardHTTPClientImpl) StoragesProvidersList(ctx context.Context, in *StoragesProvidersListRequest, opts ...http.CallOption) (*StoragesProvidersListReply, error) {
	var out StoragesProvidersListReply
	pattern := "/v1/dashboard/providers/volumes/count"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDashboardStoragesProvidersList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
