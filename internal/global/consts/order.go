package consts

type PaymentMethod int32

const (
	Alipay PaymentMethod = iota + 1
	Wechat
	BankCard
	RedemptionCode
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

type Operation string

const (
	RentingFileStorage       Operation = "租用文件存储"
	RentingCloudServers      Operation = "租用云服务器"
	RedemptionCodeRedemption Operation = "兑换码兑换"
	AlipayRecharge           Operation = "支付宝充值"
)

type OrderSymbol string

const (
	Recharge    OrderSymbol = "+"
	Consumption OrderSymbol = "-"
)
