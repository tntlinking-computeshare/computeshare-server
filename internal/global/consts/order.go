package consts

type PaymentMethod int32

const (
	Alipay PaymentMethod = iota + 1
	Wechat
	BankCard
)

type PayOrderState string

const (
	WaitBuyerPay  PayOrderState = "WAIT_BUYER_PAY"
	TradeClosed   PayOrderState = "TRADE_CLOSED"
	TradeSuccess  PayOrderState = "TRADE_SUCCESS"
	TradeFinished PayOrderState = "TRADE_FINISHED"
)

type OrderSubject string

const (
	PayCycles OrderSubject = "购买Cycles"
)

type OrderPayTime string

const (
	AlipayPayTime OrderPayTime = "5m"
)

// RenewalResourceType 续费资源类型
type RenewalResourceType int

const (
	// RenewalResourceType_Resource  资源
	RenewalResourceType_Resource RenewalResourceType = iota
	// RenewalResourceType_Storage 存储
	RenewalResourceType_Storage
)

// RenewalState 续费状态
type RenewalState int

const (
	// RenewalState_IN_SERVICE 服务中
	RenewalState_IN_SERVICE RenewalState = iota
	// RenewalState_STOP 已停止
	RenewalState_STOP
)
