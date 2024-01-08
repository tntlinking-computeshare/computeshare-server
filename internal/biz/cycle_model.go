package biz

import (
	"github.com/google/uuid"
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
