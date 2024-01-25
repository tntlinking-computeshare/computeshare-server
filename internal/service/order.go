package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
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
	computeInstanceUseCase  *biz.ComputeInstanceUsercase
	db                      *ent.Client
}

func NewOrderService(logger log.Logger,
	orderUseCase *biz.OrderUseCase,
	cycleTransactionUseCase *biz.CycleTransactionUseCase,
	cycleRenewalUseCase *biz.CycleRenewalUseCase,
	computeInstanceUseCase *biz.ComputeInstanceUsercase,
	db *ent.Client,
) *OrderService {
	return &OrderService{
		log:                     log.NewHelper(logger),
		orderUseCase:            orderUseCase,
		cycleTransactionUseCase: cycleTransactionUseCase,
		cycleRenewalUseCase:     cycleRenewalUseCase,
		computeInstanceUseCase:  computeInstanceUseCase,
		db:                      db,
	}
}

func (o *OrderService) AlipayPayNotify(ctx context.Context, req *pb.AlipayPayNotifyRequest) (*pb.AlipayPayNotifyReply, error) {
	var alipayOrderRollback biz.AlipayOrderRollback
	copier.Copy(&alipayOrderRollback, &req)
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
		return nil, errors.New(400, "unauthorized", "未认证")
	}
	userId := claim.GetUserId()
	if req.RechargeChannel == int32(consts.Alipay) {
		outTradeNo, url, err = o.orderUseCase.RechargeCycleByAlipay(ctx, userId, float64(req.Cycle), float64(req.Amount))
	} else {
		return nil, errors.New(400, "un_support_pay_method", "不支持的支付方式")
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
		return nil, errors.New(400, "unauthorized", "未认证")
	}
	userId := claim.GetUserId()
	redeemCycle, err := o.orderUseCase.RechargeCycleByRedeemCode(ctx, userId, req.RedeemCode)
	return &pb.RechargeCycleByRedeemCodeReply{
		Code:    200,
		Message: SUCCESS,
		Data:    redeemCycle,
	}, err
}

func (o *OrderService) GetCycleBalance(ctx context.Context, req *pb.GetCycleBalanceRequest) (*pb.GetCycleBalanceReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New(400, "unauthorized", "未认证")
	}
	userId := claim.GetUserId()
	redeemCycle, err := o.orderUseCase.GetCycleBalance(ctx, userId)
	return &pb.GetCycleBalanceReply{
		Code:    200,
		Message: SUCCESS,
		Data:    redeemCycle,
	}, err
}

func (o *OrderService) GetRechargeState(ctx context.Context, req *pb.GetRechargeStateRequest) (*pb.GetRechargeStateReply, error) {
	_, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New(400, "unauthorized", "未认证")
	}
	state, err := o.orderUseCase.GetRechargeState(ctx, req.GetOutTradeNo())
	return &pb.GetRechargeStateReply{
		Code:    200,
		Message: SUCCESS,
		Data:    state,
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
		AutoRenew:   item.AutoRenewal,
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

func (o *OrderService) CycleRenewalDetail(ctx context.Context, req *pb.CycleRenewalGetRequest) (*pb.CycleRenewalGetReply, error) {
	renewalId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	detail, err := o.cycleRenewalUseCase.Detail(ctx, renewalId)
	dueTime := ""
	if detail.DueTime != nil {
		dueTime = strconv.Itoa(int(detail.DueTime.UnixMilli()))
	}
	renewalTime := ""
	if detail.RenewalTime != nil {
		renewalTime = strconv.Itoa(int(detail.RenewalTime.UnixMilli()))
	}
	return &pb.CycleRenewalGetReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.CycleRenewalDetailInfo{
			Id:           detail.ID.String(),
			ProductName:  detail.ProductName,
			ProductDesc:  detail.ProductDesc,
			State:        int32(detail.State),
			DueTime:      dueTime,
			RenewalTime:  renewalTime,
			InstanceId:   detail.InstanceId.String(),
			InstanceName: detail.InstanceName,
			InstanceSpec: detail.InstanceSpec,
			Image:        detail.Image,
			ExtendPrice:  float32(detail.ExtendPrice),
			ExtendDay:    int64(detail.ExtendDay),
			Balance:      detail.Balance,
		},
	}, err
}

func (o *OrderService) ManualRenew(ctx context.Context, req *pb.ManualRenewRequest) (*pb.ManualRenewReply, error) {
	renewalId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New(400, "unauthorized", "unauthorized")
	}

	userId := claim.GetUserId()

	err = o.cycleRenewalUseCase.ManualRenew(ctx, renewalId, userId, false)
	return &pb.ManualRenewReply{
		Code:    200,
		Message: SUCCESS,
	}, err
}

func (o *OrderService) RenewDailyCheck(_ context.Context, _ *pb.DailyCheckRequest) (*pb.DailyCheckReply, error) {
	o.cycleRenewalUseCase.DailyCheck(o.db)

	o.computeInstanceUseCase.NotificationOverDue(context.Background())

	return &pb.DailyCheckReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}
