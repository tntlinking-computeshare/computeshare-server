// Code generated by ent, DO NOT EDIT.

package cycleorder

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the cycleorder type in the database.
	Label = "cycle_order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFkUserID holds the string denoting the fk_user_id field in the database.
	FieldFkUserID = "fk_user_id"
	// FieldOrderNo holds the string denoting the order_no field in the database.
	FieldOrderNo = "order_no"
	// FieldProductName holds the string denoting the product_name field in the database.
	FieldProductName = "product_name"
	// FieldProductDesc holds the string denoting the product_desc field in the database.
	FieldProductDesc = "product_desc"
	// FieldSymbol holds the string denoting the symbol field in the database.
	FieldSymbol = "symbol"
	// FieldCycle holds the string denoting the cycle field in the database.
	FieldCycle = "cycle"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// Table holds the table name of the cycleorder in the database.
	Table = "cycle_orders"
)

// Columns holds all SQL columns for cycleorder fields.
var Columns = []string{
	FieldID,
	FieldFkUserID,
	FieldOrderNo,
	FieldProductName,
	FieldProductDesc,
	FieldSymbol,
	FieldCycle,
	FieldCreateTime,
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
	// OrderNoValidator is a validator for the "order_no" field. It is called by the builders before save.
	OrderNoValidator func(string) error
	// ProductNameValidator is a validator for the "product_name" field. It is called by the builders before save.
	ProductNameValidator func(string) error
	// ProductDescValidator is a validator for the "product_desc" field. It is called by the builders before save.
	ProductDescValidator func(string) error
	// SymbolValidator is a validator for the "symbol" field. It is called by the builders before save.
	SymbolValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the CycleOrder queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFkUserID orders the results by the fk_user_id field.
func ByFkUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFkUserID, opts...).ToFunc()
}

// ByOrderNo orders the results by the order_no field.
func ByOrderNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderNo, opts...).ToFunc()
}

// ByProductName orders the results by the product_name field.
func ByProductName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductName, opts...).ToFunc()
}

// ByProductDesc orders the results by the product_desc field.
func ByProductDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductDesc, opts...).ToFunc()
}

// BySymbol orders the results by the symbol field.
func BySymbol(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSymbol, opts...).ToFunc()
}

// ByCycle orders the results by the cycle field.
func ByCycle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCycle, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}
