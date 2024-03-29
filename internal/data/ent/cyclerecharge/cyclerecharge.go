// Code generated by ent, DO NOT EDIT.

package cyclerecharge

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the cyclerecharge type in the database.
	Label = "cycle_recharge"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFkUserID holds the string denoting the fk_user_id field in the database.
	FieldFkUserID = "fk_user_id"
	// FieldOutTradeNo holds the string denoting the out_trade_no field in the database.
	FieldOutTradeNo = "out_trade_no"
	// FieldAlipayTradeNo holds the string denoting the alipay_trade_no field in the database.
	FieldAlipayTradeNo = "alipay_trade_no"
	// FieldRechargeChannel holds the string denoting the recharge_channel field in the database.
	FieldRechargeChannel = "recharge_channel"
	// FieldRedeemCode holds the string denoting the redeem_code field in the database.
	FieldRedeemCode = "redeem_code"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldPayAmount holds the string denoting the pay_amount field in the database.
	FieldPayAmount = "pay_amount"
	// FieldTotalAmount holds the string denoting the total_amount field in the database.
	FieldTotalAmount = "total_amount"
	// FieldBuyCycle holds the string denoting the buy_cycle field in the database.
	FieldBuyCycle = "buy_cycle"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the cyclerecharge in the database.
	Table = "cycle_recharges"
)

// Columns holds all SQL columns for cyclerecharge fields.
var Columns = []string{
	FieldID,
	FieldFkUserID,
	FieldOutTradeNo,
	FieldAlipayTradeNo,
	FieldRechargeChannel,
	FieldRedeemCode,
	FieldState,
	FieldPayAmount,
	FieldTotalAmount,
	FieldBuyCycle,
	FieldCreateTime,
	FieldUpdateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the CycleRecharge queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFkUserID orders the results by the fk_user_id field.
func ByFkUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFkUserID, opts...).ToFunc()
}

// ByOutTradeNo orders the results by the out_trade_no field.
func ByOutTradeNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOutTradeNo, opts...).ToFunc()
}

// ByAlipayTradeNo orders the results by the alipay_trade_no field.
func ByAlipayTradeNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlipayTradeNo, opts...).ToFunc()
}

// ByRechargeChannel orders the results by the recharge_channel field.
func ByRechargeChannel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRechargeChannel, opts...).ToFunc()
}

// ByRedeemCode orders the results by the redeem_code field.
func ByRedeemCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRedeemCode, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByPayAmount orders the results by the pay_amount field.
func ByPayAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPayAmount, opts...).ToFunc()
}

// ByTotalAmount orders the results by the total_amount field.
func ByTotalAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalAmount, opts...).ToFunc()
}

// ByBuyCycle orders the results by the buy_cycle field.
func ByBuyCycle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBuyCycle, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}
