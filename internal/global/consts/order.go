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
