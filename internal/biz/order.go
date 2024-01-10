package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/google/uuid"
	global2 "github.com/mohaijiang/computeshare-server/api/global"
	"github.com/mohaijiang/computeshare-server/internal/conf"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"github.com/mohaijiang/computeshare-server/internal/utils"
	"github.com/shopspring/decimal"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Cycle is a Cycle model.
type Cycle struct {
	ID         uuid.UUID       `json:"id,omitempty"`
	FkUserId   uuid.UUID       `json:"fkUserId,omitempty"`
	Cycle      decimal.Decimal `json:"cycle,omitempty"`
	CreateTime time.Time       `json:"createTime,omitempty"`
	UpdateTime time.Time       `json:"updateTime,omitempty"`
}

// CycleRepo is a Cycle repo.
type CycleRepo interface {
	FindByUserID(context.Context, uuid.UUID) (*Cycle, error)
	Update(ctx context.Context, cycle *Cycle) error
}

type CycleRechargeRepo interface {
	FindByOutTradeNo(context.Context, string) (*CycleRecharge, error)
	CreateCycleRecharge(context.Context, *CycleRecharge) (*CycleRecharge, error)
	UpdateCycleRecharge(context.Context, *CycleRecharge) error
}

type CycleOrderRepo interface {
	PageByUserId(ctx context.Context, userId uuid.UUID, page, size int) (*global2.Page[*CycleOrder], error)
	CheckOrderNoExists(ctx context.Context, orderNo string) bool
	Create(ctx context.Context, order *CycleOrder) (*CycleOrder, error)
}

type CycleRedeemCodeRepo interface {
	FindByRedeemCode(context.Context, string) (*CycleRedeemCode, error)
	Update(ctx context.Context, cycleRedeemCode *CycleRedeemCode) error
}

type AlipayOrderRollbackRepo interface {
	FindByOutTradeNo(context.Context, string) (*AlipayOrderRollback, error)
	SaveAlipayOrderRollback(context.Context, *AlipayOrderRollback) (*AlipayOrderRollback, error)
}

// OrderUseCase is a cycle UseCase.
type OrderUseCase struct {
	cycleRepo               CycleRepo
	orderRepo               CycleOrderRepo
	cycleRechargeRepo       CycleRechargeRepo
	alipayOrderRollbackRepo AlipayOrderRollbackRepo
	cycleTransactionRepo    CycleTransactionRepo
	cycleRedeemCodeRepo     CycleRedeemCodeRepo
	log                     *log.Helper
	dispose                 conf.Dispose
}

// NewOrderUseCase new a cycle UseCase.
func NewOrderUseCase(cycleRepo CycleRepo,
	orderRepo CycleOrderRepo,
	cycleRechargeRepo CycleRechargeRepo,
	alipayOrderRollbackRepo AlipayOrderRollbackRepo,
	cycleTransactionRepo CycleTransactionRepo,
	cycleRedeemCodeRepo CycleRedeemCodeRepo,
	logger log.Logger,
	confDispose *conf.Dispose) *OrderUseCase {
	return &OrderUseCase{
		cycleRepo:               cycleRepo,
		orderRepo:               orderRepo,
		cycleRechargeRepo:       cycleRechargeRepo,
		alipayOrderRollbackRepo: alipayOrderRollbackRepo,
		cycleTransactionRepo:    cycleTransactionRepo,
		cycleRedeemCodeRepo:     cycleRedeemCodeRepo,
		log:                     log.NewHelper(logger),
		dispose:                 *confDispose,
	}
}

func (o *OrderUseCase) RechargeCycleByAlipay(ctx context.Context, userId uuid.UUID, cycle, amount float64) (outTradeNo string, url string, err error) {
	amountDecimal := decimal.NewFromFloat(amount).Round(2)
	cycleDecimal := decimal.NewFromFloat(cycle).Round(2)
	if !amountDecimal.Mul(decimal.NewFromFloat(1000.00)).Round(2).Equal(cycleDecimal) {
		return "", "", errors.New(400, "Amount_Error", "充值比例不正确")
	}
	outTradeNo = utils.GetOutTradeNo()
	cycleRecharge := CycleRecharge{
		FkUserID:        userId,
		OutTradeNo:      outTradeNo,
		RechargeChannel: int(consts.Alipay),
		PayAmount:       amountDecimal,
		BuyCycle:        cycleDecimal,
	}
	recharge, err := o.cycleRechargeRepo.CreateCycleRecharge(ctx, &cycleRecharge)
	if err != nil {
		return "", "", err
	}
	alipayPublicCert, _ := os.Open(o.dispose.Alipay.AlipayPublicCertPath)
	alipayRootCert, _ := os.Open(o.dispose.Alipay.AlipayRootCertPath)
	appPublicCert, _ := os.Open(o.dispose.Alipay.AppPublicCertPath)
	alipayPublicCertContent, _ := io.ReadAll(alipayPublicCert)
	alipayRootContent, _ := io.ReadAll(alipayRootCert)
	appPublicContent, _ := io.ReadAll(appPublicCert)
	client, err := alipay.NewClient(o.dispose.Alipay.AppId, o.dispose.Alipay.AppPrivateKey, false)
	if err != nil {
		log.Log(log.LevelError, "alipay.NewClient", err)
		return "", "", err
	}
	client.DebugSwitch = gopay.DebugOn
	// 设置支付宝请求 公共参数
	// 注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).                     // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2).                    // 设置签名类型，不设置默认 RSA2
							SetReturnUrl(o.dispose.Alipay.PayReturnUrl). // 设置返回URL，付款结束后跳转的url
							SetNotifyUrl(o.dispose.Alipay.PayNotifyUrl)  // 设置异步通知URL

	// 自动同步验签（只支持证书模式）
	// 传入 alipayCertPublicKey_RSA2.crt 内容
	client.AutoVerifySign(alipayPublicCertContent)

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	//err := client.SetCertSnByPath("appCertPublicKey.crt", "alipayRootCert.crt", "alipayCertPublicKey_RSA2.crt")
	// 证书内容
	err = client.SetCertSnByContent(appPublicContent, alipayRootContent, alipayPublicCertContent)
	if err != nil {
		log.Log(log.LevelError, "SetCertSnByContent", err)
		return "", "", err
	}
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("subject", consts.PayCycles). // 标题
						Set("out_trade_no", recharge.OutTradeNo).     // 订单号，支付成功后会返回
						Set("total_amount", recharge.PayAmount).      // 订单金额
						Set("timeout_express", consts.AlipayPayTime). // 支付超时时间
						Set("product_code", "FAST_INSTANT_TRADE_PAY") // 必填 具体参考文档

	alipayUrl, err := client.TradePagePay(context.Background(), bm)
	if err != nil {
		log.Log(log.LevelError, "TradePagePay", err)
		return "", "", err
	}
	return outTradeNo, alipayUrl, nil

}

func (o *OrderUseCase) AlipayPayNotify(ctx context.Context, alipayOrderRollback AlipayOrderRollback) (err error) {
	rollback, err := o.alipayOrderRollbackRepo.SaveAlipayOrderRollback(ctx, &alipayOrderRollback)
	if err != nil {
		return err
	}
	if rollback.TradeStatus == string(consts.WaitBuyerPay) {
		log.Log(log.LevelInfo, "AlipayPayNotifyStatus", "交易创建，等待买家付款。")
	} else if rollback.TradeStatus == string(consts.TradeClosed) {
		log.Log(log.LevelInfo, "AlipayPayNotifyStatus", "未付款交易超时关闭，或支付完成后全额退款。")
	} else if rollback.TradeStatus == string(consts.TradeSuccess) {
		log.Log(log.LevelInfo, "AlipayPayNotifyStatus", "交易支付成功。")
		err = o.AlipayPayNotifyFollowUp(ctx, *rollback)
	} else if rollback.TradeStatus == string(consts.TradeFinished) {
		log.Log(log.LevelInfo, "AlipayPayNotifyStatus", "交易结束，不可退款。")
	} else {
		log.Log(log.LevelInfo, "AlipayPayNotifyStatus", "未知的交易状态"+rollback.TradeStatus)
	}
	return err
}

func (o *OrderUseCase) AlipayPayNotifyFollowUp(ctx context.Context, rollback AlipayOrderRollback) (err error) {
	cycleRecharge, err := o.cycleRechargeRepo.FindByOutTradeNo(ctx, rollback.OutTradeNo)
	if err != nil {
		return err
	}
	cycle, err := o.cycleRepo.FindByUserID(ctx, cycleRecharge.FkUserID)
	if err != nil {
		return err
	}
	buyCycle, _ := cycleRecharge.BuyCycle.Round(2).Float64()
	balanceDecimal := cycle.Cycle.Add(cycleRecharge.BuyCycle)
	balance, _ := balanceDecimal.Round(2).Float64()
	cycleTransaction := CycleTransaction{
		FkCycleID:         cycle.ID,
		FkUserID:          cycle.FkUserId,
		FkCycleOrderID:    uuid.Nil,
		FkCycleRechargeID: cycleRecharge.ID,
		Operation:         string(consts.AlipayRecharge),
		Symbol:            string(consts.Recharge),
		Cycle:             buyCycle,
		Balance:           balance,
		OperationTime:     time.Now(),
	}
	_, err = o.cycleTransactionRepo.Create(ctx, &cycleTransaction)
	if err != nil {
		log.Log(log.LevelInfo, "o.cycleTransactionRepo.Create", err)
		return err
	}
	cycle.Cycle = balanceDecimal
	err = o.cycleRepo.Update(ctx, cycle)
	if err != nil {
		log.Log(log.LevelInfo, "o.cycleRepo.Update", err)
		return err
	}
	totalAmount, err := decimal.NewFromString(rollback.TotalAmount)
	if err != nil {
		log.Log(log.LevelInfo, "decimal.NewFromString", err)
		return err
	}
	cycleRecharge.AlipayTradeNo = rollback.TradeNo
	cycleRecharge.State = string(consts.TradeSuccess)
	cycleRecharge.TotalAmount = totalAmount
	err = o.cycleRechargeRepo.UpdateCycleRecharge(ctx, cycleRecharge)
	if err != nil {
		log.Log(log.LevelInfo, "o.cycleRechargeRepo.UpdateCycleRecharge", err)
		return err
	}
	return nil
}

func (o *OrderUseCase) RechargeCycleByRedeemCode(ctx context.Context, userId uuid.UUID, redeemCode string) (redeemCycle string, err error) {
	cycleRedeemCode, err := o.cycleRedeemCodeRepo.FindByRedeemCode(ctx, redeemCode)
	if err != nil && ent.IsNotFound(err) {
		return "", errors.New(400, "Redeem_Code_Invalid", "兑换码不正确")
	} else if err != nil {
		return "", err
	}
	if cycleRedeemCode.State {
		return "", errors.New(400, "Redeem_Code_Invalid", "兑换码已经被使用")
	}
	cycle, err := o.cycleRepo.FindByUserID(ctx, userId)
	if err != nil {
		return "", err
	}
	outTradeNo := utils.GetOutTradeNo()
	//生成充值数据
	cycleRecharge := CycleRecharge{
		FkUserID:        userId,
		OutTradeNo:      outTradeNo,
		RechargeChannel: int(consts.RedemptionCode),
		RedeemCode:      redeemCode,
		State:           string(consts.TradeSuccess),
		PayAmount:       decimal.Decimal{},
		TotalAmount:     decimal.Decimal{},
		BuyCycle:        cycleRedeemCode.Cycle,
	}
	createCycleRecharge, err := o.cycleRechargeRepo.CreateCycleRecharge(ctx, &cycleRecharge)
	if err != nil {
		return "", err
	}
	//生成交易数据
	buyCycle, _ := createCycleRecharge.BuyCycle.Round(2).Float64()
	balanceDecimal := cycle.Cycle.Add(createCycleRecharge.BuyCycle)
	balance, _ := balanceDecimal.Round(2).Float64()
	cycleTransaction := CycleTransaction{
		FkCycleID:         cycle.ID,
		FkUserID:          cycle.FkUserId,
		FkCycleOrderID:    uuid.Nil,
		FkCycleRechargeID: createCycleRecharge.ID,
		Operation:         string(consts.RedemptionCodeRedemption),
		Symbol:            string(consts.Recharge),
		Cycle:             buyCycle,
		Balance:           balance,
		OperationTime:     time.Now(),
	}
	_, err = o.cycleTransactionRepo.Create(ctx, &cycleTransaction)
	if err != nil {
		log.Log(log.LevelInfo, "o.cycleTransactionRepo.Create", err)
		return "", err
	}
	cycle.Cycle = balanceDecimal
	err = o.cycleRepo.Update(ctx, cycle)
	if err != nil {
		log.Log(log.LevelInfo, "o.cycleRepo.Update", err)
		return "", err
	}
	cycleRedeemCode.State = true
	cycleRedeemCode.FkUserID = userId
	cycleRedeemCode.UseTime = time.Now()
	err = o.cycleRedeemCodeRepo.Update(ctx, cycleRedeemCode)
	if err != nil {
		log.Log(log.LevelInfo, "o.cycleRedeemCodeRepo.Update", err)
		return "", err
	}
	return cycleRedeemCode.Cycle.StringFixed(2), nil
}

func (o *OrderUseCase) GetCycleBalance(ctx context.Context, userId uuid.UUID) (redeemCycle string, err error) {
	cycle, err := o.cycleRepo.FindByUserID(ctx, userId)
	if err != nil {
		return "", err
	}
	return cycle.Cycle.StringFixed(2), nil
}

func (o *OrderUseCase) OrderList(ctx context.Context, page, size int32) (*global2.Page[*CycleOrder], error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New(400, "unauthorized", "unauthorized")
	}

	userId := claim.GetUserId()
	return o.orderRepo.PageByUserId(ctx, userId, int(page), int(size))

}

func NewOrderNo() string {
	// 设置随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成0到9999之间的随机数
	randomNumber := r.Intn(10000)
	// 格式化为字符串，并补足到4位长度
	formattedNumber := fmt.Sprintf("%04d", randomNumber)
	return fmt.Sprintf("%s0000%s", time.Now().Format("20060102"), formattedNumber)
}
