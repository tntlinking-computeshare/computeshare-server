// Code generated by ent, DO NOT EDIT.

package computeimage

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the computeimage type in the database.
	Label = "compute_image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldTag holds the string denoting the tag field in the database.
	FieldTag = "tag"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// Table holds the table name of the computeimage in the database.
	Table = "compute_images"
)

// Columns holds all SQL columns for computeimage fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldImage,
	FieldTag,
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
	// ImageValidator is a validator for the "image" field. It is called by the builders before save.
	ImageValidator func(string) error
	// TagValidator is a validator for the "tag" field. It is called by the builders before save.
	TagValidator func(string) error
)

// OrderOption defines the ordering options for the ComputeImage queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByImage orders the results by the image field.
func ByImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImage, opts...).ToFunc()
}

// ByTag orders the results by the tag field.
func ByTag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTag, opts...).ToFunc()
}

// ByPort orders the results by the port field.
func ByPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPort, opts...).ToFunc()
}
