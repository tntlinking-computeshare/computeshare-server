// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cycleorder"
)

// CycleOrder is the model entity for the CycleOrder schema.
type CycleOrder struct {
	config `json:"-"`
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
	CreateTime   time.Time `json:"create_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CycleOrder) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cycleorder.FieldCycle:
			values[i] = new(sql.NullFloat64)
		case cycleorder.FieldOrderNo, cycleorder.FieldProductName, cycleorder.FieldProductDesc, cycleorder.FieldSymbol:
			values[i] = new(sql.NullString)
		case cycleorder.FieldCreateTime:
			values[i] = new(sql.NullTime)
		case cycleorder.FieldID, cycleorder.FieldFkUserID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CycleOrder fields.
func (co *CycleOrder) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cycleorder.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				co.ID = *value
			}
		case cycleorder.FieldFkUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field fk_user_id", values[i])
			} else if value != nil {
				co.FkUserID = *value
			}
		case cycleorder.FieldOrderNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field order_no", values[i])
			} else if value.Valid {
				co.OrderNo = value.String
			}
		case cycleorder.FieldProductName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product_name", values[i])
			} else if value.Valid {
				co.ProductName = value.String
			}
		case cycleorder.FieldProductDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product_desc", values[i])
			} else if value.Valid {
				co.ProductDesc = value.String
			}
		case cycleorder.FieldSymbol:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field symbol", values[i])
			} else if value.Valid {
				co.Symbol = value.String
			}
		case cycleorder.FieldCycle:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field cycle", values[i])
			} else if value.Valid {
				co.Cycle = value.Float64
			}
		case cycleorder.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				co.CreateTime = value.Time
			}
		default:
			co.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CycleOrder.
// This includes values selected through modifiers, order, etc.
func (co *CycleOrder) Value(name string) (ent.Value, error) {
	return co.selectValues.Get(name)
}

// Update returns a builder for updating this CycleOrder.
// Note that you need to call CycleOrder.Unwrap() before calling this method if this CycleOrder
// was returned from a transaction, and the transaction was committed or rolled back.
func (co *CycleOrder) Update() *CycleOrderUpdateOne {
	return NewCycleOrderClient(co.config).UpdateOne(co)
}

// Unwrap unwraps the CycleOrder entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (co *CycleOrder) Unwrap() *CycleOrder {
	_tx, ok := co.config.driver.(*txDriver)
	if !ok {
		panic("ent: CycleOrder is not a transactional entity")
	}
	co.config.driver = _tx.drv
	return co
}

// String implements the fmt.Stringer.
func (co *CycleOrder) String() string {
	var builder strings.Builder
	builder.WriteString("CycleOrder(")
	builder.WriteString(fmt.Sprintf("id=%v, ", co.ID))
	builder.WriteString("fk_user_id=")
	builder.WriteString(fmt.Sprintf("%v", co.FkUserID))
	builder.WriteString(", ")
	builder.WriteString("order_no=")
	builder.WriteString(co.OrderNo)
	builder.WriteString(", ")
	builder.WriteString("product_name=")
	builder.WriteString(co.ProductName)
	builder.WriteString(", ")
	builder.WriteString("product_desc=")
	builder.WriteString(co.ProductDesc)
	builder.WriteString(", ")
	builder.WriteString("symbol=")
	builder.WriteString(co.Symbol)
	builder.WriteString(", ")
	builder.WriteString("cycle=")
	builder.WriteString(fmt.Sprintf("%v", co.Cycle))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(co.CreateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CycleOrders is a parsable slice of CycleOrder.
type CycleOrders []*CycleOrder
