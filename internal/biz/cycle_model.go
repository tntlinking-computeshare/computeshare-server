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
	FkUserID uuid.UUID `json:"fkUserId,omitempty"`
	// 订单编号
	OrderNo string `json:"orderNo,omitempty"`
	// 产品名字
	ProductName string `json:"productName,omitempty"`
	// 产品描述
	ProductDesc string `json:"productDesc,omitempty"`
	// symbol
	Symbol string `json:"symbol,omitempty"`
	// Cycle holds the value of the "cycle" field.
	Cycle float64 `json:"cycle,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"createTime,omitempty"`
}

type CycleTransaction struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// cycleId
	FkCycleID uuid.UUID `json:"fkCycleId,omitempty"`
	// 用户id
	FkUserID uuid.UUID `json:"fkUserId,omitempty"`
	// fk_cycle_order_id
	FkCycleOrderID uuid.UUID `json:"fkCycleOrderId,omitempty"`
	// fk_cycle_recharge_id
	FkCycleRechargeID uuid.UUID `json:"fkCycleRechargeId,omitempty"`
	// 操作
	Operation string `json:"operation,omitempty"`
	// symbol
	Symbol string `json:"symbol,omitempty"`
	// Cycle holds the value of the "cycle" field.
	Cycle float64 `json:"cycle,omitempty"`
	// 余额
	Balance float64 `json:"balance,omitempty"`
	// 操作时间
	OperationTime time.Time `json:"operationTime,omitempty"`
}

type CycleRenewal struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// 用户id
	FkUserID uuid.UUID `json:"fkUserId,omitempty"`
	// 资源ID
	ResourceID uuid.UUID `json:"resourceId,omitempty"`
	// 资源类型
	ResourceType int `json:"resourceType,omitempty"`
	// 产品名字
	ProductName string `json:"productName,omitempty"`
	// 产品描述
	ProductDesc string `json:"productDesc,omitempty"`
	// 状态
	State int8 `json:"state,omitempty"`
	// 延长时间
	ExtendDay int8 `json:"extendDay,omitempty"`
	// 额外的价格
	ExtendPrice float64 `json:"extendPrice,omitempty"`
	// 到期时间
	DueTime *time.Time `json:"dueTime,omitempty"`
	// 续费时间
	RenewalTime *time.Time `json:"renewalTime,omitempty"`
	// 自动续费
	AutoRenewal bool `json:"autoRenewal,omitempty"`
}

type CycleRenewalDetail struct {

	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// 用户id
	FkUserID uuid.UUID `json:"fkUserId,omitempty"`
	// 资源ID
	ResourceID uuid.UUID `json:"resourceId,omitempty"`
	// 资源类型
	ResourceType int `json:"resourceType,omitempty"`
	// 产品名字
	ProductName string `json:"productName,omitempty"`
	// 产品描述
	ProductDesc string `json:"productDesc,omitempty"`
	// 状态
	State int8 `json:"state,omitempty"`
	// 延长时间
	ExtendDay int8 `json:"extendDay,omitempty"`
	// 额外的价格
	ExtendPrice float64 `json:"extendPrice,omitempty"`
	// 到期时间
	DueTime *time.Time `json:"dueTime,omitempty"`
	// 续费时间
	RenewalTime *time.Time `json:"renewalTime,omitempty"`
	// 自动续费
	AutoRenewal bool `json:"autoRenewal,omitempty"`
	// 实例id
	InstanceId uuid.UUID `json:"instanceId"`
	// 实例名
	InstanceName string `json:"instanceName"`
	// 实例规格
	InstanceSpec string `json:"instanceSpec"`
	// 镜像
	Image string `json:"image"`
	// 余额
	Balance float32 `json:"balance"`
}

type CycleRecharge struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// 用户id
	FkUserID uuid.UUID `json:"fkUserId,omitempty"`
	// 商家订单号
	OutTradeNo string `json:"outTradeNo,omitempty"`
	// 支付宝订单号
	AlipayTradeNo string `json:"alipayTradeNo,omitempty"`
	// 充值渠道
	RechargeChannel int `json:"rechargeChannel,omitempty"`
	// 兑换码
	RedeemCode string `json:"redeemCode,omitempty"`
	// 状态
	State string `json:"state,omitempty"`
	// 支付的钱
	PayAmount decimal.Decimal `json:"payAmount,omitempty"`
	// 收到的钱
	TotalAmount decimal.Decimal `json:"totalAmount,omitempty"`
	// 购买的周期
	BuyCycle decimal.Decimal `json:"buyCycle,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"createTime,omitempty"`
	// 创建时间
	UpdateTime time.Time `json:"updateTime,omitempty"`
}
