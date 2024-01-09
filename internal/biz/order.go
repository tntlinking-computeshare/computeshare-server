package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/google/uuid"
	global2 "github.com/mohaijiang/computeshare-server/api/global"
	"github.com/mohaijiang/computeshare-server/internal/conf"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"github.com/mohaijiang/computeshare-server/internal/utils"
	"github.com/shopspring/decimal"
	"io"
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
}

type CycleRechargeRepo interface {
	FindByOutTradeNo(context.Context, string) (*CycleRecharge, error)
	CreateCycleRecharge(context.Context, *CycleRecharge) (*CycleRecharge, error)
	UpdateCycleRecharge(context.Context, *CycleRecharge) error
}

type CycleOrderRepo interface {
	PageByUserId(ctx context.Context, userId uuid.UUID, page, size int) (*global2.Page[*CycleOrder], error)
}

// OrderUseCase is a cycle UseCase.
type OrderUseCase struct {
	cycleRepo         CycleRepo
	orderRepo         CycleOrderRepo
	cycleRechargeRepo CycleRechargeRepo
	log               *log.Helper
	dispose           conf.Dispose
}

// NewOrderUseCase new a cycle UseCase.
func NewOrderUseCase(cycleRepo CycleRepo, orderRepo CycleOrderRepo, cycleRechargeRepo CycleRechargeRepo, logger log.Logger, confDispose *conf.Dispose) *OrderUseCase {
	return &OrderUseCase{
		cycleRepo:         cycleRepo,
		orderRepo:         orderRepo,
		cycleRechargeRepo: cycleRechargeRepo,
		log:               log.NewHelper(logger),
		dispose:           *confDispose,
	}
}

func (c *OrderUseCase) RechargeCycleByAlipay(ctx context.Context, userId uuid.UUID, cycle, amount float64) (outTradeNo string, url string, err error) {
	outTradeNo = utils.GetOutTradeNo()
	cycleRecharge := CycleRecharge{
		FkUserID:        userId,
		OutTradeNo:      outTradeNo,
		RechargeChannel: int(consts.Alipay),
		PayAmount:       decimal.NewFromFloat(amount),
		BuyCycle:        decimal.NewFromFloat(cycle),
	}
	recharge, err := c.cycleRechargeRepo.CreateCycleRecharge(ctx, &cycleRecharge)
	if err != nil {
		return "", "", err
	}
	alipayPublicCert, _ := os.Open(c.dispose.Alipay.AlipayPublicCertPath)
	alipayRootCert, _ := os.Open(c.dispose.Alipay.AlipayRootCertPath)
	appPublicCert, _ := os.Open(c.dispose.Alipay.AppPublicCertPath)
	alipayPublicCertContent, _ := io.ReadAll(alipayPublicCert)
	alipayRootContent, _ := io.ReadAll(alipayRootCert)
	appPublicContent, _ := io.ReadAll(appPublicCert)
	client, err := alipay.NewClient(c.dispose.Alipay.AppId, c.dispose.Alipay.AppPrivateKey, false)
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
							SetReturnUrl(c.dispose.Alipay.PayReturnUrl). // 设置返回URL，付款结束后跳转的url
							SetNotifyUrl(c.dispose.Alipay.PayNotifyUrl)  // 设置异步通知URL

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
	bm.Set("subject", "共享算力Cycle购买"). // 标题
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

func (c *OrderUseCase) OrderList(ctx context.Context, page, size int32) (*global2.Page[*CycleOrder], error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New(400, "unauthorized", "unauthorized")
	}

	userId := claim.GetUserId()
	return c.orderRepo.PageByUserId(ctx, userId, int(page), int(size))

}
