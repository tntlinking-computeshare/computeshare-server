package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"github.com/samber/lo"
	"strconv"

	pb "github.com/mohaijiang/computeshare-server/api/order/v1"
)

type OrderService struct {
	pb.UnimplementedOrderServer
	log                     *log.Helper
	orderUseCase            *biz.OrderUseCase
	cycleTransactionUseCase *biz.CycleTransactionUseCase
	cycleRenewalUseCase     *biz.CycleRenewalUseCase
}

func NewOrderService(logger log.Logger,
	orderUseCase *biz.OrderUseCase,
	cycleTransactionUseCase *biz.CycleTransactionUseCase,
	cycleRenewalUseCase *biz.CycleRenewalUseCase,
) *OrderService {
	return &OrderService{
		log:                     log.NewHelper(logger),
		orderUseCase:            orderUseCase,
		cycleTransactionUseCase: cycleTransactionUseCase,
		cycleRenewalUseCase:     cycleRenewalUseCase,
	}
}

func (o *OrderService) AlipayPayNotify(ctx context.Context, req *pb.AlipayPayNotifyRequest) (*pb.AlipayPayNotifyReply, error) {
	var alipayOrderRollback biz.AlipayOrderRollback
	copier.Copy(alipayOrderRollback, req)
	err := o.orderUseCase.AlipayPayNotify(ctx, alipayOrderRollback)
	if err != nil {
		return nil, err
	}
	return &pb.AlipayPayNotifyReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}

func (o *OrderService) RechargeCycleByAlipay(ctx context.Context, req *pb.RechargeCycleByAlipayRequest) (*pb.RechargeCycleByAlipayReply, error) {
	var outTradeNo string
	var url string
	var err error
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	if req.RechargeChannel == int32(consts.Alipay) {
		outTradeNo, url, err = o.orderUseCase.RechargeCycleByAlipay(ctx, userId, float64(req.Cycle), float64(req.Amount))
	} else {
		return nil, errors.New("不支持的支付方式")
	}
	if err != nil {
		return nil, err
	}
	return &pb.RechargeCycleByAlipayReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.RechargeCycleByAlipayReply_Pay{
			OutTradeNo: outTradeNo,
			Url:        url,
		},
	}, nil
}

func (o *OrderService) RechargeCycleByRedeemCode(ctx context.Context, req *pb.RechargeCycleByRedeemCodeRequest) (*pb.RechargeCycleByRedeemCodeReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	redeemCycle, err := o.orderUseCase.RechargeCycleByRedeemCode(ctx, userId, req.RedeemCode)
	return &pb.RechargeCycleByRedeemCodeReply{
		Code:    200,
		Message: SUCCESS,
		Data:    redeemCycle,
	}, err
}

func (o *OrderService) OrderList(ctx context.Context, req *pb.OrderListRequest) (*pb.OrderListReply, error) {
	pageData, err := o.orderUseCase.OrderList(ctx, req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	return &pb.OrderListReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.OrderPage{
			Total: pageData.Total,
			Page:  pageData.Page,
			Size:  pageData.Size,
			Data:  lo.Map(pageData.Data, o.toBiz),
		},
	}, err
}

func (o *OrderService) toBiz(item *biz.CycleOrder, _ int) *pb.OrderInfo {
	if item == nil {
		return nil
	}

	return &pb.OrderInfo{
		Id:          item.ID.String(),
		OrderNo:     item.OrderNo,
		ProductName: item.ProductName,
		ProductDesc: item.ProductDesc,
		Symbol:      item.Symbol,
		Cycle:       float32(item.Cycle),
		CreateTime:  item.CreateTime.UnixMilli(),
	}
}

func (o *OrderService) CycleTransactionList(ctx context.Context, req *pb.CycleTransactionListRequest) (*pb.CycleTransactionListReply, error) {
	pageData, err := o.cycleTransactionUseCase.PageByUser(ctx, req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	return &pb.CycleTransactionListReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.CycleTransactionPage{
			Total: pageData.Total,
			Page:  pageData.Page,
			Size:  pageData.Size,
			Data:  lo.Map(pageData.Data, o.toCycleTransactionBiz),
		},
	}, err
}

func (o *OrderService) toCycleTransactionBiz(item *biz.CycleTransaction, _ int) *pb.CycleTransactionInfo {

	if item == nil {
		return nil
	}

	return &pb.CycleTransactionInfo{
		Id:            item.ID.String(),
		Operation:     item.Operation,
		Symbol:        item.Symbol,
		Cycle:         float32(item.Cycle),
		OperationTime: item.OperationTime.UnixMilli(),
	}
}

func (o *OrderService) CycleRenewalList(ctx context.Context, req *pb.CycleRenewalListRequest) (*pb.CycleRenewalListReply, error) {
	pageData, err := o.cycleRenewalUseCase.PageByUser(ctx, req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	return &pb.CycleRenewalListReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.CycleRenewalPage{
			Total: pageData.Total,
			Page:  pageData.Page,
			Size:  pageData.Size,
			Data:  lo.Map(pageData.Data, o.toCycleRenewalBiz),
		},
	}, err
}

func (o *OrderService) toCycleRenewalBiz(item *biz.CycleRenewal, _ int) *pb.CycleRenewalInfo {
	if item == nil {
		return nil
	}

	dueTime := ""
	if item.DueTime != nil {
		dueTime = strconv.Itoa(int(item.DueTime.UnixMilli()))
	}
	renewTime := ""
	if item.RenewalTime != nil {
		renewTime = strconv.Itoa(int(item.RenewalTime.UnixMilli()))
	}
	return &pb.CycleRenewalInfo{
		Id:          item.ID.String(),
		ProductName: item.ProductName,
		ProductDesc: item.ProductDesc,
		State:       int32(item.State),
		DueTime:     dueTime,
		RenewalTime: renewTime,
	}
}

func (o *OrderService) CycleRenewalOpen(ctx context.Context, req *pb.CycleRenewalGetRequest) (*pb.CycleRenewalBaseReply, error) {
	renewalId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	err = o.cycleRenewalUseCase.OpenRenewal(ctx, renewalId)
	return &pb.CycleRenewalBaseReply{
		Code:    200,
		Message: SUCCESS,
	}, err
}

func (o *OrderService) CycleRenewalClose(ctx context.Context, req *pb.CycleRenewalGetRequest) (*pb.CycleRenewalBaseReply, error) {

	renewalId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	err = o.cycleRenewalUseCase.CloseRenewal(ctx, renewalId)
	return &pb.CycleRenewalBaseReply{
		Code:    200,
		Message: SUCCESS,
	}, err
}

func (o *OrderService) CycleRenewalInfo(ctx context.Context, req *pb.CycleRenewalGetRequest) (*pb.CycleRenewalGetReply, error) {
	renewalId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	renewal, err := o.cycleRenewalUseCase.Get(ctx, renewalId)
	return &pb.CycleRenewalGetReply{
		Code:    200,
		Message: SUCCESS,
		Data:    o.toCycleRenewalBiz(renewal, 0),
	}, err
}

func (o *OrderService) ManualRenew(ctx context.Context, req *pb.ManualRenewRequest) (*pb.ManualRenewReply, error) {
	renewalId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	err = o.cycleRenewalUseCase.ManualRenew(ctx, renewalId)
	return &pb.ManualRenewReply{
		Code:    200,
		Message: SUCCESS,
	}, err
}
