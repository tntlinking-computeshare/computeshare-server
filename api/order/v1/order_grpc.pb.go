// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: api/order/v1/order.proto

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
	Order_AlipayPayNotify_FullMethodName           = "/api.server.order.v1.Order/AlipayPayNotify"
	Order_RechargeCycleByAlipay_FullMethodName     = "/api.server.order.v1.Order/RechargeCycleByAlipay"
	Order_GetRechargeState_FullMethodName          = "/api.server.order.v1.Order/GetRechargeState"
	Order_RechargeCycleByRedeemCode_FullMethodName = "/api.server.order.v1.Order/RechargeCycleByRedeemCode"
	Order_GetCycleBalance_FullMethodName           = "/api.server.order.v1.Order/GetCycleBalance"
	Order_OrderList_FullMethodName                 = "/api.server.order.v1.Order/OrderList"
	Order_CycleTransactionList_FullMethodName      = "/api.server.order.v1.Order/CycleTransactionList"
	Order_CycleRenewalDetail_FullMethodName        = "/api.server.order.v1.Order/CycleRenewalDetail"
	Order_CycleRenewalList_FullMethodName          = "/api.server.order.v1.Order/CycleRenewalList"
	Order_CycleRenewalOpen_FullMethodName          = "/api.server.order.v1.Order/CycleRenewalOpen"
	Order_CycleRenewalClose_FullMethodName         = "/api.server.order.v1.Order/CycleRenewalClose"
	Order_ManualRenew_FullMethodName               = "/api.server.order.v1.Order/ManualRenew"
	Order_RenewDailyCheck_FullMethodName           = "/api.server.order.v1.Order/RenewDailyCheck"
)

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	AlipayPayNotify(ctx context.Context, in *AlipayPayNotifyRequest, opts ...grpc.CallOption) (*AlipayPayNotifyReply, error)
	RechargeCycleByAlipay(ctx context.Context, in *RechargeCycleByAlipayRequest, opts ...grpc.CallOption) (*RechargeCycleByAlipayReply, error)
	GetRechargeState(ctx context.Context, in *GetRechargeStateRequest, opts ...grpc.CallOption) (*GetRechargeStateReply, error)
	RechargeCycleByRedeemCode(ctx context.Context, in *RechargeCycleByRedeemCodeRequest, opts ...grpc.CallOption) (*RechargeCycleByRedeemCodeReply, error)
	GetCycleBalance(ctx context.Context, in *GetCycleBalanceRequest, opts ...grpc.CallOption) (*GetCycleBalanceReply, error)
	OrderList(ctx context.Context, in *OrderListRequest, opts ...grpc.CallOption) (*OrderListReply, error)
	CycleTransactionList(ctx context.Context, in *CycleTransactionListRequest, opts ...grpc.CallOption) (*CycleTransactionListReply, error)
	CycleRenewalDetail(ctx context.Context, in *CycleRenewalGetRequest, opts ...grpc.CallOption) (*CycleRenewalGetReply, error)
	CycleRenewalList(ctx context.Context, in *CycleRenewalListRequest, opts ...grpc.CallOption) (*CycleRenewalListReply, error)
	CycleRenewalOpen(ctx context.Context, in *CycleRenewalGetRequest, opts ...grpc.CallOption) (*CycleRenewalBaseReply, error)
	CycleRenewalClose(ctx context.Context, in *CycleRenewalGetRequest, opts ...grpc.CallOption) (*CycleRenewalBaseReply, error)
	ManualRenew(ctx context.Context, in *ManualRenewRequest, opts ...grpc.CallOption) (*ManualRenewReply, error)
	RenewDailyCheck(ctx context.Context, in *DailyCheckRequest, opts ...grpc.CallOption) (*DailyCheckReply, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) AlipayPayNotify(ctx context.Context, in *AlipayPayNotifyRequest, opts ...grpc.CallOption) (*AlipayPayNotifyReply, error) {
	out := new(AlipayPayNotifyReply)
	err := c.cc.Invoke(ctx, Order_AlipayPayNotify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) RechargeCycleByAlipay(ctx context.Context, in *RechargeCycleByAlipayRequest, opts ...grpc.CallOption) (*RechargeCycleByAlipayReply, error) {
	out := new(RechargeCycleByAlipayReply)
	err := c.cc.Invoke(ctx, Order_RechargeCycleByAlipay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetRechargeState(ctx context.Context, in *GetRechargeStateRequest, opts ...grpc.CallOption) (*GetRechargeStateReply, error) {
	out := new(GetRechargeStateReply)
	err := c.cc.Invoke(ctx, Order_GetRechargeState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) RechargeCycleByRedeemCode(ctx context.Context, in *RechargeCycleByRedeemCodeRequest, opts ...grpc.CallOption) (*RechargeCycleByRedeemCodeReply, error) {
	out := new(RechargeCycleByRedeemCodeReply)
	err := c.cc.Invoke(ctx, Order_RechargeCycleByRedeemCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetCycleBalance(ctx context.Context, in *GetCycleBalanceRequest, opts ...grpc.CallOption) (*GetCycleBalanceReply, error) {
	out := new(GetCycleBalanceReply)
	err := c.cc.Invoke(ctx, Order_GetCycleBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderList(ctx context.Context, in *OrderListRequest, opts ...grpc.CallOption) (*OrderListReply, error) {
	out := new(OrderListReply)
	err := c.cc.Invoke(ctx, Order_OrderList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CycleTransactionList(ctx context.Context, in *CycleTransactionListRequest, opts ...grpc.CallOption) (*CycleTransactionListReply, error) {
	out := new(CycleTransactionListReply)
	err := c.cc.Invoke(ctx, Order_CycleTransactionList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CycleRenewalDetail(ctx context.Context, in *CycleRenewalGetRequest, opts ...grpc.CallOption) (*CycleRenewalGetReply, error) {
	out := new(CycleRenewalGetReply)
	err := c.cc.Invoke(ctx, Order_CycleRenewalDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CycleRenewalList(ctx context.Context, in *CycleRenewalListRequest, opts ...grpc.CallOption) (*CycleRenewalListReply, error) {
	out := new(CycleRenewalListReply)
	err := c.cc.Invoke(ctx, Order_CycleRenewalList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CycleRenewalOpen(ctx context.Context, in *CycleRenewalGetRequest, opts ...grpc.CallOption) (*CycleRenewalBaseReply, error) {
	out := new(CycleRenewalBaseReply)
	err := c.cc.Invoke(ctx, Order_CycleRenewalOpen_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CycleRenewalClose(ctx context.Context, in *CycleRenewalGetRequest, opts ...grpc.CallOption) (*CycleRenewalBaseReply, error) {
	out := new(CycleRenewalBaseReply)
	err := c.cc.Invoke(ctx, Order_CycleRenewalClose_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) ManualRenew(ctx context.Context, in *ManualRenewRequest, opts ...grpc.CallOption) (*ManualRenewReply, error) {
	out := new(ManualRenewReply)
	err := c.cc.Invoke(ctx, Order_ManualRenew_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) RenewDailyCheck(ctx context.Context, in *DailyCheckRequest, opts ...grpc.CallOption) (*DailyCheckReply, error) {
	out := new(DailyCheckReply)
	err := c.cc.Invoke(ctx, Order_RenewDailyCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	AlipayPayNotify(context.Context, *AlipayPayNotifyRequest) (*AlipayPayNotifyReply, error)
	RechargeCycleByAlipay(context.Context, *RechargeCycleByAlipayRequest) (*RechargeCycleByAlipayReply, error)
	GetRechargeState(context.Context, *GetRechargeStateRequest) (*GetRechargeStateReply, error)
	RechargeCycleByRedeemCode(context.Context, *RechargeCycleByRedeemCodeRequest) (*RechargeCycleByRedeemCodeReply, error)
	GetCycleBalance(context.Context, *GetCycleBalanceRequest) (*GetCycleBalanceReply, error)
	OrderList(context.Context, *OrderListRequest) (*OrderListReply, error)
	CycleTransactionList(context.Context, *CycleTransactionListRequest) (*CycleTransactionListReply, error)
	CycleRenewalDetail(context.Context, *CycleRenewalGetRequest) (*CycleRenewalGetReply, error)
	CycleRenewalList(context.Context, *CycleRenewalListRequest) (*CycleRenewalListReply, error)
	CycleRenewalOpen(context.Context, *CycleRenewalGetRequest) (*CycleRenewalBaseReply, error)
	CycleRenewalClose(context.Context, *CycleRenewalGetRequest) (*CycleRenewalBaseReply, error)
	ManualRenew(context.Context, *ManualRenewRequest) (*ManualRenewReply, error)
	RenewDailyCheck(context.Context, *DailyCheckRequest) (*DailyCheckReply, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) AlipayPayNotify(context.Context, *AlipayPayNotifyRequest) (*AlipayPayNotifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlipayPayNotify not implemented")
}
func (UnimplementedOrderServer) RechargeCycleByAlipay(context.Context, *RechargeCycleByAlipayRequest) (*RechargeCycleByAlipayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RechargeCycleByAlipay not implemented")
}
func (UnimplementedOrderServer) GetRechargeState(context.Context, *GetRechargeStateRequest) (*GetRechargeStateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRechargeState not implemented")
}
func (UnimplementedOrderServer) RechargeCycleByRedeemCode(context.Context, *RechargeCycleByRedeemCodeRequest) (*RechargeCycleByRedeemCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RechargeCycleByRedeemCode not implemented")
}
func (UnimplementedOrderServer) GetCycleBalance(context.Context, *GetCycleBalanceRequest) (*GetCycleBalanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCycleBalance not implemented")
}
func (UnimplementedOrderServer) OrderList(context.Context, *OrderListRequest) (*OrderListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderList not implemented")
}
func (UnimplementedOrderServer) CycleTransactionList(context.Context, *CycleTransactionListRequest) (*CycleTransactionListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CycleTransactionList not implemented")
}
func (UnimplementedOrderServer) CycleRenewalDetail(context.Context, *CycleRenewalGetRequest) (*CycleRenewalGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CycleRenewalDetail not implemented")
}
func (UnimplementedOrderServer) CycleRenewalList(context.Context, *CycleRenewalListRequest) (*CycleRenewalListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CycleRenewalList not implemented")
}
func (UnimplementedOrderServer) CycleRenewalOpen(context.Context, *CycleRenewalGetRequest) (*CycleRenewalBaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CycleRenewalOpen not implemented")
}
func (UnimplementedOrderServer) CycleRenewalClose(context.Context, *CycleRenewalGetRequest) (*CycleRenewalBaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CycleRenewalClose not implemented")
}
func (UnimplementedOrderServer) ManualRenew(context.Context, *ManualRenewRequest) (*ManualRenewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManualRenew not implemented")
}
func (UnimplementedOrderServer) RenewDailyCheck(context.Context, *DailyCheckRequest) (*DailyCheckReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewDailyCheck not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_AlipayPayNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlipayPayNotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).AlipayPayNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_AlipayPayNotify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).AlipayPayNotify(ctx, req.(*AlipayPayNotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_RechargeCycleByAlipay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RechargeCycleByAlipayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).RechargeCycleByAlipay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_RechargeCycleByAlipay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).RechargeCycleByAlipay(ctx, req.(*RechargeCycleByAlipayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetRechargeState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRechargeStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetRechargeState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_GetRechargeState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetRechargeState(ctx, req.(*GetRechargeStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_RechargeCycleByRedeemCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RechargeCycleByRedeemCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).RechargeCycleByRedeemCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_RechargeCycleByRedeemCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).RechargeCycleByRedeemCode(ctx, req.(*RechargeCycleByRedeemCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetCycleBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCycleBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetCycleBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_GetCycleBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetCycleBalance(ctx, req.(*GetCycleBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_OrderList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderList(ctx, req.(*OrderListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CycleTransactionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CycleTransactionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CycleTransactionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CycleTransactionList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CycleTransactionList(ctx, req.(*CycleTransactionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CycleRenewalDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CycleRenewalGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CycleRenewalDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CycleRenewalDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CycleRenewalDetail(ctx, req.(*CycleRenewalGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CycleRenewalList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CycleRenewalListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CycleRenewalList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CycleRenewalList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CycleRenewalList(ctx, req.(*CycleRenewalListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CycleRenewalOpen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CycleRenewalGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CycleRenewalOpen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CycleRenewalOpen_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CycleRenewalOpen(ctx, req.(*CycleRenewalGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CycleRenewalClose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CycleRenewalGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CycleRenewalClose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CycleRenewalClose_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CycleRenewalClose(ctx, req.(*CycleRenewalGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_ManualRenew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManualRenewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).ManualRenew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_ManualRenew_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).ManualRenew(ctx, req.(*ManualRenewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_RenewDailyCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DailyCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).RenewDailyCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_RenewDailyCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).RenewDailyCheck(ctx, req.(*DailyCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.server.order.v1.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AlipayPayNotify",
			Handler:    _Order_AlipayPayNotify_Handler,
		},
		{
			MethodName: "RechargeCycleByAlipay",
			Handler:    _Order_RechargeCycleByAlipay_Handler,
		},
		{
			MethodName: "GetRechargeState",
			Handler:    _Order_GetRechargeState_Handler,
		},
		{
			MethodName: "RechargeCycleByRedeemCode",
			Handler:    _Order_RechargeCycleByRedeemCode_Handler,
		},
		{
			MethodName: "GetCycleBalance",
			Handler:    _Order_GetCycleBalance_Handler,
		},
		{
			MethodName: "OrderList",
			Handler:    _Order_OrderList_Handler,
		},
		{
			MethodName: "CycleTransactionList",
			Handler:    _Order_CycleTransactionList_Handler,
		},
		{
			MethodName: "CycleRenewalDetail",
			Handler:    _Order_CycleRenewalDetail_Handler,
		},
		{
			MethodName: "CycleRenewalList",
			Handler:    _Order_CycleRenewalList_Handler,
		},
		{
			MethodName: "CycleRenewalOpen",
			Handler:    _Order_CycleRenewalOpen_Handler,
		},
		{
			MethodName: "CycleRenewalClose",
			Handler:    _Order_CycleRenewalClose_Handler,
		},
		{
			MethodName: "ManualRenew",
			Handler:    _Order_ManualRenew_Handler,
		},
		{
			MethodName: "RenewDailyCheck",
			Handler:    _Order_RenewDailyCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/order/v1/order.proto",
}
