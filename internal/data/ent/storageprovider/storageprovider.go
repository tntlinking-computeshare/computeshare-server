// Code generated by ent, DO NOT EDIT.

package storageprovider

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
)

const (
	// Label holds the string label denoting the storageprovider type in the database.
	Label = "storage_provider"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAgentID holds the string denoting the agent_id field in the database.
	FieldAgentID = "agent_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldMasterServer holds the string denoting the master_server field in the database.
	FieldMasterServer = "master_server"
	// FieldPublicIP holds the string denoting the public_ip field in the database.
	FieldPublicIP = "public_ip"
	// FieldPublicPort holds the string denoting the public_port field in the database.
	FieldPublicPort = "public_port"
	// FieldGrpcPort holds the string denoting the grpc_port field in the database.
	FieldGrpcPort = "grpc_port"
	// FieldCreatedTime holds the string denoting the created_time field in the database.
	FieldCreatedTime = "created_time"
	// Table holds the table name of the storageprovider in the database.
	Table = "storage_providers"
)

// Columns holds all SQL columns for storageprovider fields.
var Columns = []string{
	FieldID,
	FieldAgentID,
	FieldStatus,
	FieldMasterServer,
	FieldPublicIP,
	FieldPublicPort,
	FieldGrpcPort,
	FieldCreatedTime,
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus consts.StorageProviderStatus
	// MasterServerValidator is a validator for the "master_server" field. It is called by the builders before save.
	MasterServerValidator func(string) error
	// PublicIPValidator is a validator for the "public_ip" field. It is called by the builders before save.
	PublicIPValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the StorageProvider queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAgentID orders the results by the agent_id field.
func ByAgentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAgentID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByMasterServer orders the results by the master_server field.
func ByMasterServer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMasterServer, opts...).ToFunc()
}

// ByPublicIP orders the results by the public_ip field.
func ByPublicIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublicIP, opts...).ToFunc()
}

// ByPublicPort orders the results by the public_port field.
func ByPublicPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublicPort, opts...).ToFunc()
}

// ByGrpcPort orders the results by the grpc_port field.
func ByGrpcPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGrpcPort, opts...).ToFunc()
}

// ByCreatedTime orders the results by the created_time field.
func ByCreatedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedTime, opts...).ToFunc()
}