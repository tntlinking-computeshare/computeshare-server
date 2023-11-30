// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/computeinstance"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
)

// ComputeInstance is the model entity for the ComputeInstance schema.
type ComputeInstance struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner string `json:"owner,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Core holds the value of the "core" field.
	Core string `json:"core,omitempty"`
	// Memory holds the value of the "memory" field.
	Memory string `json:"memory,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// 容器端口
	Port string `json:"port,omitempty"`
	// ExpirationTime holds the value of the "expiration_time" field.
	ExpirationTime time.Time `json:"expiration_time,omitempty"`
	// 0: 启动中,1:运行中,2:连接中断, 3:过期
	Status consts.InstanceStatus `json:"status,omitempty"`
	// 容器id
	ContainerID string `json:"container_id,omitempty"`
	// p2p agent Id
	AgentID string `json:"agent_id,omitempty"`
	// 容器启动命令
	Command string `json:"command,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ComputeInstanceQuery when eager-loading is set.
	Edges        ComputeInstanceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ComputeInstanceEdges holds the relations/edges for other nodes in the graph.
type ComputeInstanceEdges struct {
	// NetworkMappings holds the value of the networkMappings edge.
	NetworkMappings []*NetworkMapping `json:"networkMappings,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// NetworkMappingsOrErr returns the NetworkMappings value or an error if the edge
// was not loaded in eager-loading.
func (e ComputeInstanceEdges) NetworkMappingsOrErr() ([]*NetworkMapping, error) {
	if e.loadedTypes[0] {
		return e.NetworkMappings, nil
	}
	return nil, &NotLoadedError{edge: "networkMappings"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ComputeInstance) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case computeinstance.FieldStatus:
			values[i] = new(sql.NullInt64)
		case computeinstance.FieldOwner, computeinstance.FieldName, computeinstance.FieldCore, computeinstance.FieldMemory, computeinstance.FieldImage, computeinstance.FieldPort, computeinstance.FieldContainerID, computeinstance.FieldAgentID, computeinstance.FieldCommand:
			values[i] = new(sql.NullString)
		case computeinstance.FieldExpirationTime:
			values[i] = new(sql.NullTime)
		case computeinstance.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ComputeInstance fields.
func (ci *ComputeInstance) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case computeinstance.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ci.ID = *value
			}
		case computeinstance.FieldOwner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner", values[i])
			} else if value.Valid {
				ci.Owner = value.String
			}
		case computeinstance.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ci.Name = value.String
			}
		case computeinstance.FieldCore:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field core", values[i])
			} else if value.Valid {
				ci.Core = value.String
			}
		case computeinstance.FieldMemory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memory", values[i])
			} else if value.Valid {
				ci.Memory = value.String
			}
		case computeinstance.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				ci.Image = value.String
			}
		case computeinstance.FieldPort:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field port", values[i])
			} else if value.Valid {
				ci.Port = value.String
			}
		case computeinstance.FieldExpirationTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiration_time", values[i])
			} else if value.Valid {
				ci.ExpirationTime = value.Time
			}
		case computeinstance.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ci.Status = consts.InstanceStatus(value.Int64)
			}
		case computeinstance.FieldContainerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field container_id", values[i])
			} else if value.Valid {
				ci.ContainerID = value.String
			}
		case computeinstance.FieldAgentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agent_id", values[i])
			} else if value.Valid {
				ci.AgentID = value.String
			}
		case computeinstance.FieldCommand:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field command", values[i])
			} else if value.Valid {
				ci.Command = value.String
			}
		default:
			ci.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ComputeInstance.
// This includes values selected through modifiers, order, etc.
func (ci *ComputeInstance) Value(name string) (ent.Value, error) {
	return ci.selectValues.Get(name)
}

// QueryNetworkMappings queries the "networkMappings" edge of the ComputeInstance entity.
func (ci *ComputeInstance) QueryNetworkMappings() *NetworkMappingQuery {
	return NewComputeInstanceClient(ci.config).QueryNetworkMappings(ci)
}

// Update returns a builder for updating this ComputeInstance.
// Note that you need to call ComputeInstance.Unwrap() before calling this method if this ComputeInstance
// was returned from a transaction, and the transaction was committed or rolled back.
func (ci *ComputeInstance) Update() *ComputeInstanceUpdateOne {
	return NewComputeInstanceClient(ci.config).UpdateOne(ci)
}

// Unwrap unwraps the ComputeInstance entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ci *ComputeInstance) Unwrap() *ComputeInstance {
	_tx, ok := ci.config.driver.(*txDriver)
	if !ok {
		panic("ent: ComputeInstance is not a transactional entity")
	}
	ci.config.driver = _tx.drv
	return ci
}

// String implements the fmt.Stringer.
func (ci *ComputeInstance) String() string {
	var builder strings.Builder
	builder.WriteString("ComputeInstance(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ci.ID))
	builder.WriteString("owner=")
	builder.WriteString(ci.Owner)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ci.Name)
	builder.WriteString(", ")
	builder.WriteString("core=")
	builder.WriteString(ci.Core)
	builder.WriteString(", ")
	builder.WriteString("memory=")
	builder.WriteString(ci.Memory)
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(ci.Image)
	builder.WriteString(", ")
	builder.WriteString("port=")
	builder.WriteString(ci.Port)
	builder.WriteString(", ")
	builder.WriteString("expiration_time=")
	builder.WriteString(ci.ExpirationTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ci.Status))
	builder.WriteString(", ")
	builder.WriteString("container_id=")
	builder.WriteString(ci.ContainerID)
	builder.WriteString(", ")
	builder.WriteString("agent_id=")
	builder.WriteString(ci.AgentID)
	builder.WriteString(", ")
	builder.WriteString("command=")
	builder.WriteString(ci.Command)
	builder.WriteByte(')')
	return builder.String()
}

// ComputeInstances is a parsable slice of ComputeInstance.
type ComputeInstances []*ComputeInstance
