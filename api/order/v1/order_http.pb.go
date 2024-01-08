// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.2
// source: api/order/v1/order.proto

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

const OperationOrderAlipayPayNotify = "/api.server.order.v1.Order/AlipayPayNotify"
const OperationOrderCycleRenewalList = "/api.server.order.v1.Order/CycleRenewalList"
const OperationOrderCycleTransactionList = "/api.server.order.v1.Order/CycleTransactionList"
const OperationOrderOrderList = "/api.server.order.v1.Order/OrderList"
const OperationOrderRechargeCycleByAlipay = "/api.server.order.v1.Order/RechargeCycleByAlipay"
const OperationOrderRechargeCycleByRedeemCode = "/api.server.order.v1.Order/RechargeCycleByRedeemCode"

type OrderHTTPServer interface {
	AlipayPayNotify(context.Context, *AlipayPayNotifyRequest) (*AlipayPayNotifyReply, error)
	CycleRenewalList(context.Context, *CycleRenewalListRequest) (*CycleRenewalListReply, error)
	CycleTransactionList(context.Context, *CycleTransactionListRequest) (*CycleTransactionListReply, error)
	OrderList(context.Context, *OrderListRequest) (*OrderListReply, error)
	RechargeCycleByAlipay(context.Context, *RechargeCycleByAlipayRequest) (*RechargeCycleByAlipayReply, error)
	RechargeCycleByRedeemCode(context.Context, *RechargeCycleByRedeemCodeRequest) (*RechargeCycleByRedeemCodeReply, error)
}

func RegisterOrderHTTPServer(s *http.Server, srv OrderHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/alipay/pay/notify", _Order_AlipayPayNotify0_HTTP_Handler(srv))
	r.POST("/v1/cycle/recharge", _Order_RechargeCycleByAlipay0_HTTP_Handler(srv))
	r.POST("/v1/cycle/redeem", _Order_RechargeCycleByRedeemCode0_HTTP_Handler(srv))
	r.GET("/v1/order", _Order_OrderList0_HTTP_Handler(srv))
	r.GET("/v1/cycle/transaction", _Order_CycleTransactionList0_HTTP_Handler(srv))
	r.GET("/v1/cycle/renew", _Order_CycleRenewalList0_HTTP_Handler(srv))
}

func _Order_AlipayPayNotify0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AlipayPayNotifyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderAlipayPayNotify)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AlipayPayNotify(ctx, req.(*AlipayPayNotifyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AlipayPayNotifyReply)
		return ctx.Result(200, reply)
	}
}

func _Order_RechargeCycleByAlipay0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RechargeCycleByAlipayRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderRechargeCycleByAlipay)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RechargeCycleByAlipay(ctx, req.(*RechargeCycleByAlipayRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RechargeCycleByAlipayReply)
		return ctx.Result(200, reply)
	}
}

func _Order_RechargeCycleByRedeemCode0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RechargeCycleByRedeemCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderRechargeCycleByRedeemCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RechargeCycleByRedeemCode(ctx, req.(*RechargeCycleByRedeemCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RechargeCycleByRedeemCodeReply)
		return ctx.Result(200, reply)
	}
}

func _Order_OrderList0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OrderListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderOrderList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OrderList(ctx, req.(*OrderListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OrderListReply)
		return ctx.Result(200, reply)
	}
}

func _Order_CycleTransactionList0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CycleTransactionListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderCycleTransactionList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CycleTransactionList(ctx, req.(*CycleTransactionListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CycleTransactionListReply)
		return ctx.Result(200, reply)
	}
}

func _Order_CycleRenewalList0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CycleRenewalListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderCycleRenewalList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CycleRenewalList(ctx, req.(*CycleRenewalListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CycleRenewalListReply)
		return ctx.Result(200, reply)
	}
}

type OrderHTTPClient interface {
	AlipayPayNotify(ctx context.Context, req *AlipayPayNotifyRequest, opts ...http.CallOption) (rsp *AlipayPayNotifyReply, err error)
	CycleRenewalList(ctx context.Context, req *CycleRenewalListRequest, opts ...http.CallOption) (rsp *CycleRenewalListReply, err error)
	CycleTransactionList(ctx context.Context, req *CycleTransactionListRequest, opts ...http.CallOption) (rsp *CycleTransactionListReply, err error)
	OrderList(ctx context.Context, req *OrderListRequest, opts ...http.CallOption) (rsp *OrderListReply, err error)
	RechargeCycleByAlipay(ctx context.Context, req *RechargeCycleByAlipayRequest, opts ...http.CallOption) (rsp *RechargeCycleByAlipayReply, err error)
	RechargeCycleByRedeemCode(ctx context.Context, req *RechargeCycleByRedeemCodeRequest, opts ...http.CallOption) (rsp *RechargeCycleByRedeemCodeReply, err error)
}

type OrderHTTPClientImpl struct {
	cc *http.Client
}

func NewOrderHTTPClient(client *http.Client) OrderHTTPClient {
	return &OrderHTTPClientImpl{client}
}

func (c *OrderHTTPClientImpl) AlipayPayNotify(ctx context.Context, in *AlipayPayNotifyRequest, opts ...http.CallOption) (*AlipayPayNotifyReply, error) {
	var out AlipayPayNotifyReply
	pattern := "/v1/alipay/pay/notify"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOrderAlipayPayNotify))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) CycleRenewalList(ctx context.Context, in *CycleRenewalListRequest, opts ...http.CallOption) (*CycleRenewalListReply, error) {
	var out CycleRenewalListReply
	pattern := "/v1/cycle/renew"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOrderCycleRenewalList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) CycleTransactionList(ctx context.Context, in *CycleTransactionListRequest, opts ...http.CallOption) (*CycleTransactionListReply, error) {
	var out CycleTransactionListReply
	pattern := "/v1/cycle/transaction"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOrderCycleTransactionList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) OrderList(ctx context.Context, in *OrderListRequest, opts ...http.CallOption) (*OrderListReply, error) {
	var out OrderListReply
	pattern := "/v1/order"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOrderOrderList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) RechargeCycleByAlipay(ctx context.Context, in *RechargeCycleByAlipayRequest, opts ...http.CallOption) (*RechargeCycleByAlipayReply, error) {
	var out RechargeCycleByAlipayReply
	pattern := "/v1/cycle/recharge"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOrderRechargeCycleByAlipay))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) RechargeCycleByRedeemCode(ctx context.Context, in *RechargeCycleByRedeemCodeRequest, opts ...http.CallOption) (*RechargeCycleByRedeemCodeReply, error) {
	var out RechargeCycleByRedeemCodeReply
	pattern := "/v1/cycle/redeem"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOrderRechargeCycleByRedeemCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
