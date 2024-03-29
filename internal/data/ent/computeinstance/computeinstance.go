// Code generated by ent, DO NOT EDIT.

package computeinstance

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the computeinstance type in the database.
	Label = "compute_instance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOwner holds the string denoting the owner field in the database.
	FieldOwner = "owner"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCore holds the string denoting the core field in the database.
	FieldCore = "core"
	// FieldMemory holds the string denoting the memory field in the database.
	FieldMemory = "memory"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldImageID holds the string denoting the image_id field in the database.
	FieldImageID = "image_id"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// FieldExpirationTime holds the string denoting the expiration_time field in the database.
	FieldExpirationTime = "expiration_time"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldContainerID holds the string denoting the container_id field in the database.
	FieldContainerID = "container_id"
	// FieldAgentID holds the string denoting the agent_id field in the database.
	FieldAgentID = "agent_id"
	// FieldVncIP holds the string denoting the vnc_ip field in the database.
	FieldVncIP = "vnc_ip"
	// FieldVncPort holds the string denoting the vnc_port field in the database.
	FieldVncPort = "vnc_port"
	// FieldDockerCompose holds the string denoting the docker_compose field in the database.
	FieldDockerCompose = "docker_compose"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// Table holds the table name of the computeinstance in the database.
	Table = "compute_instances"
)

// Columns holds all SQL columns for computeinstance fields.
var Columns = []string{
	FieldID,
	FieldOwner,
	FieldName,
	FieldCore,
	FieldMemory,
	FieldImage,
	FieldImageID,
	FieldPort,
	FieldExpirationTime,
	FieldStatus,
	FieldContainerID,
	FieldAgentID,
	FieldVncIP,
	FieldVncPort,
	FieldDockerCompose,
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
	// OwnerValidator is a validator for the "owner" field. It is called by the builders before save.
	OwnerValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// ImageValidator is a validator for the "image" field. It is called by the builders before save.
	ImageValidator func(string) error
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the ComputeInstance queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOwner orders the results by the owner field.
func ByOwner(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwner, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCore orders the results by the core field.
func ByCore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCore, opts...).ToFunc()
}

// ByMemory orders the results by the memory field.
func ByMemory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemory, opts...).ToFunc()
}

// ByImage orders the results by the image field.
func ByImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImage, opts...).ToFunc()
}

// ByImageID orders the results by the image_id field.
func ByImageID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImageID, opts...).ToFunc()
}

// ByPort orders the results by the port field.
func ByPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPort, opts...).ToFunc()
}

// ByExpirationTime orders the results by the expiration_time field.
func ByExpirationTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpirationTime, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByContainerID orders the results by the container_id field.
func ByContainerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContainerID, opts...).ToFunc()
}

// ByAgentID orders the results by the agent_id field.
func ByAgentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAgentID, opts...).ToFunc()
}

// ByVncIP orders the results by the vnc_ip field.
func ByVncIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVncIP, opts...).ToFunc()
}

// ByVncPort orders the results by the vnc_port field.
func ByVncPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVncPort, opts...).ToFunc()
}

// ByDockerCompose orders the results by the docker_compose field.
func ByDockerCompose(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDockerCompose, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}
