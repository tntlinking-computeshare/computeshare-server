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
