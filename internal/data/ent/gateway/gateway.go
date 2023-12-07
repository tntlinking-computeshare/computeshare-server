// Code generated by ent, DO NOT EDIT.

package gateway

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the gateway type in the database.
	Label = "gateway"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// Table holds the table name of the gateway in the database.
	Table = "gateways"
)

// Columns holds all SQL columns for gateway fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldIP,
	FieldPort,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Gateway queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByIP orders the results by the ip field.
func ByIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIP, opts...).ToFunc()
}

// ByPort orders the results by the port field.
func ByPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPort, opts...).ToFunc()
}