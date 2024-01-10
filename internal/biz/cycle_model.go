package biz

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type CycleOrder struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// 用户id
	FkUserID uuid.UUID `json:"fk_user_id,omitempty"`
	// 订单编号
	OrderNo string `json:"order_no,omitempty"`
	// 产品名字
	ProductName string `json:"product_name,omitempty"`
	// 产品描述
	ProductDesc string `json:"product_desc,omitempty"`
	// symbol
	Symbol string `json:"symbol,omitempty"`
	// Cycle holds the value of the "cycle" field.
	Cycle float64 `json:"cycle,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
}

type CycleTransaction struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// cycleId
	FkCycleID uuid.UUID `json:"fk_cycle_id,omitempty"`
	// 用户id
	FkUserID uuid.UUID `json:"fk_user_id,omitempty"`
	// fk_cycle_order_id
	FkCycleOrderID uuid.UUID `json:"fk_cycle_order_id,omitempty"`
	// fk_cycle_recharge_id
	FkCycleRechargeID uuid.UUID `json:"fk_cycle_recharge_id,omitempty"`
	// 操作
	Operation string `json:"operation,omitempty"`
	// symbol
	Symbol string `json:"symbol,omitempty"`
	// Cycle holds the value of the "cycle" field.
	Cycle float64 `json:"cycle,omitempty"`
	// 余额
	Balance float64 `json:"balance,omitempty"`
	// 操作时间
	OperationTime time.Time `json:"operation_time,omitempty"`
}

type CycleRenewal struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// 用户id
	FkUserID uuid.UUID `json:"fk_user_id,omitempty"`
	// 资源ID
	ResourceID uuid.UUID `json:"resource_id,omitempty"`
	// 资源类型
	ResourceType int `json:"resource_type,omitempty"`
	// 产品名字
	ProductName string `json:"product_name,omitempty"`
	// 产品描述
	ProductDesc string `json:"product_desc,omitempty"`
	// 状态
	State int8 `json:"state,omitempty"`
	// 延长时间
	ExtendDay int8 `json:"extend_day,omitempty"`
	// 额外的价格
	ExtendPrice float64 `json:"extend_price,omitempty"`
	// 到期时间
	DueTime *time.Time `json:"due_time,omitempty"`
	// 续费时间
	RenewalTime *time.Time `json:"renewal_time,omitempty"`
	// 自动续费
	AutoRenewal bool `json:"auto_renewal,omitempty"`
}

type CycleRecharge struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// 用户id
	FkUserID uuid.UUID `json:"fk_user_id,omitempty"`
	// 商家订单号
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// 支付宝订单号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty"`
	// 充值渠道
	RechargeChannel int `json:"recharge_channel,omitempty"`
	// 兑换码
	RedeemCode string `json:"redeem_code,omitempty"`
	// 状态
	State string `json:"state,omitempty"`
	// 支付的钱
	PayAmount decimal.Decimal `json:"pay_amount,omitempty"`
	// 收到的钱
	TotalAmount decimal.Decimal `json:"total_amount,omitempty"`
	// 购买的周期
	BuyCycle decimal.Decimal `json:"buy_cycle,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 创建时间
	UpdateTime time.Time `json:"update_time,omitempty"`
}

type CycleRedeemCode struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// 用户id
	FkUserID uuid.UUID `json:"fk_user_id,omitempty"`
	// 兑换码
	RedeemCode string `json:"redeem_code,omitempty"`
	// 兑换码对应的周期
	Cycle decimal.Decimal `json:"cycle,omitempty"`
	// 状态
	State bool `json:"state,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 使用
	UseTime time.Time `json:"use_time,omitempty"`
}

type AlipayOrderRollback struct {
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 通知校验ID
	NotifyID string `json:"notify_id,omitempty"`
	// 通知类型
	NotifyType string `json:"notify_type,omitempty"`
	// 通知时间
	NotifyTime string `json:"notify_time,omitempty"`
	// 编码格式，如 utf-8、gbk、gb2312 等
	Charset string `json:"charset,omitempty"`
	// 调用的接口版本，固定为：1.0
	Version string `json:"version,omitempty"`
	// 签名类型
	SignType string `json:"sign_type,omitempty"`
	// 签名
	Sign string `json:"sign,omitempty"`
	// 支付成功的各个渠道金额信息。详情可查看 资金明细信息说明
	FundBillList string `json:"fund_bill_list,omitempty"`
	// 实收金额
	ReceiptAmount string `json:"receipt_amount,omitempty"`
	// 用户在交易中支付的可开发票的金额
	InvoiceAmount string `json:"invoice_amount,omitempty"`
	// 付款金额
	BuyerPayAmount string `json:"buyer_pay_amount,omitempty"`
	// 集分宝金额
	PointAmount string `json:"point_amount,omitempty"`
	// 本交易支付时所有优惠券信息，详情可查看 优惠券信息说明
	VoucherDetailList string `json:"voucher_detail_list,omitempty"`
	// 公共回传参数，如果请求时传递了该参数，则返回给商家时会在异步通知时将该参数原样返回。本参数必须进行 UrlEncode 之后才可以发送给支付宝。
	PassbackParams string `json:"passback_params,omitempty"`
	// 支付宝交易号
	TradeNo string `json:"trade_no,omitempty"`
	// 开发者id
	AppID string `json:"app_id,omitempty"`
	// 商户订单号
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// 商户业务号
	OutBizNo string `json:"out_biz_no,omitempty"`
	// 买家支付宝ID
	BuyerID string `json:"buyer_id,omitempty"`
	// 卖家支付宝id
	SellerID string `json:"seller_id,omitempty"`
	// 交易状态
	TradeStatus string `json:"trade_status,omitempty"`
	// 订单金额
	TotalAmount string `json:"total_amount,omitempty"`
	// 总退款金额
	RefundFee string `json:"refund_fee,omitempty"`
	// 订单标题
	Subject string `json:"subject,omitempty"`
	// 订单的备注、描述、明细等。对应请求时的 body 参数，原样通知回来
	Body string `json:"body,omitempty"`
	// 交易创建时间
	GmtCreate string `json:"gmt_create,omitempty"`
	// 交易付款时间
	GmtPayment string `json:"gmt_payment,omitempty"`
	// 交易关闭时间
	GmtClose string `json:"gmt_close,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 创建时间
	UpdateTime time.Time `json:"update_time,omitempty"`
}
