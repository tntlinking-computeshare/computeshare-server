package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"github.com/samber/lo"

	pb "github.com/mohaijiang/computeshare-server/api/order/v1"
)

type OrderService struct {
	pb.UnimplementedOrderServer
	log                     *log.Helper
	orderUseCase            *biz.OrderUseCase
	cycleTransactionUseCase *biz.CycleTransactionUseCase
}

func NewOrderService(logger log.Logger,
	orderUseCase *biz.OrderUseCase,
	cycleTransactionUseCase *biz.CycleTransactionUseCase,
) *OrderService {
	return &OrderService{
		log:                     log.NewHelper(logger),
		orderUseCase:            orderUseCase,
		cycleTransactionUseCase: cycleTransactionUseCase,
	}
}

func (o *OrderService) AlipayPayNotify(ctx context.Context, req *pb.AlipayPayNotifyRequest) (*pb.AlipayPayNotifyReply, error) {
	if req.TradeStatus == "WAIT_BUYER_PAY" {
		log.Log(log.LevelInfo, "交易创建，等待买家付款。")
		log.Log(log.LevelInfo, req)
	} else if req.TradeStatus == "TRADE_CLOSED" {
		log.Log(log.LevelInfo, "未付款交易超时关闭，或支付完成后全额退款。")
		log.Log(log.LevelInfo, req)
	} else if req.TradeStatus == "TRADE_SUCCESS" {
		log.Log(log.LevelInfo, "交易支付成功。")
		log.Log(log.LevelInfo, req)
	} else if req.TradeStatus == "TRADE_FINISHED" {
		log.Log(log.LevelInfo, "交易结束，不可退款。")
		log.Log(log.LevelInfo, req)
	} else {
		log.Log(log.LevelInfo, "未知的交易状态"+req.TradeStatus)
		log.Log(log.LevelInfo, req)
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

	return &pb.RechargeCycleByRedeemCodeReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
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
