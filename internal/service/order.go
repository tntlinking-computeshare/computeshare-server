package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/mohaijiang/computeshare-server/api/order/v1"
)

type OrderService struct {
	pb.UnimplementedOrderServer
	log *log.Helper
}

func NewOrderService(logger log.Logger) *OrderService {
	return &OrderService{
		log: log.NewHelper(logger),
	}
}

func (s *OrderService) AlipayPayNotify(ctx context.Context, req *pb.AlipayPayNotifyRequest) (*pb.AlipayPayNotifyReply, error) {
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
