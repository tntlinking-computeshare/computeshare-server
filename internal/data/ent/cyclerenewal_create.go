// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cyclerenewal"
)

// CycleRenewalCreate is the builder for creating a CycleRenewal entity.
type CycleRenewalCreate struct {
	config
	mutation *CycleRenewalMutation
	hooks    []Hook
}

// SetFkUserID sets the "fk_user_id" field.
func (crc *CycleRenewalCreate) SetFkUserID(u uuid.UUID) *CycleRenewalCreate {
	crc.mutation.SetFkUserID(u)
	return crc
}

// SetResourceID sets the "resource_id" field.
func (crc *CycleRenewalCreate) SetResourceID(u uuid.UUID) *CycleRenewalCreate {
	crc.mutation.SetResourceID(u)
	return crc
}

// SetResourceType sets the "resource_type" field.
func (crc *CycleRenewalCreate) SetResourceType(i int) *CycleRenewalCreate {
	crc.mutation.SetResourceType(i)
	return crc
}

// SetProductName sets the "product_name" field.
func (crc *CycleRenewalCreate) SetProductName(s string) *CycleRenewalCreate {
	crc.mutation.SetProductName(s)
	return crc
}

// SetProductDesc sets the "product_desc" field.
func (crc *CycleRenewalCreate) SetProductDesc(s string) *CycleRenewalCreate {
	crc.mutation.SetProductDesc(s)
	return crc
}

// SetState sets the "state" field.
func (crc *CycleRenewalCreate) SetState(i int8) *CycleRenewalCreate {
	crc.mutation.SetState(i)
	return crc
}

// SetExtendDay sets the "extend_day" field.
func (crc *CycleRenewalCreate) SetExtendDay(i int8) *CycleRenewalCreate {
	crc.mutation.SetExtendDay(i)
	return crc
}

// SetExtendPrice sets the "extend_price" field.
func (crc *CycleRenewalCreate) SetExtendPrice(f float64) *CycleRenewalCreate {
	crc.mutation.SetExtendPrice(f)
	return crc
}

// SetDueTime sets the "due_time" field.
func (crc *CycleRenewalCreate) SetDueTime(t time.Time) *CycleRenewalCreate {
	crc.mutation.SetDueTime(t)
	return crc
}

// SetNillableDueTime sets the "due_time" field if the given value is not nil.
func (crc *CycleRenewalCreate) SetNillableDueTime(t *time.Time) *CycleRenewalCreate {
	if t != nil {
		crc.SetDueTime(*t)
	}
	return crc
}

// SetRenewalTime sets the "renewal_time" field.
func (crc *CycleRenewalCreate) SetRenewalTime(t time.Time) *CycleRenewalCreate {
	crc.mutation.SetRenewalTime(t)
	return crc
}

// SetNillableRenewalTime sets the "renewal_time" field if the given value is not nil.
func (crc *CycleRenewalCreate) SetNillableRenewalTime(t *time.Time) *CycleRenewalCreate {
	if t != nil {
		crc.SetRenewalTime(*t)
	}
	return crc
}

// SetAutoRenewal sets the "auto_renewal" field.
func (crc *CycleRenewalCreate) SetAutoRenewal(b bool) *CycleRenewalCreate {
	crc.mutation.SetAutoRenewal(b)
	return crc
}

// SetID sets the "id" field.
func (crc *CycleRenewalCreate) SetID(u uuid.UUID) *CycleRenewalCreate {
	crc.mutation.SetID(u)
	return crc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (crc *CycleRenewalCreate) SetNillableID(u *uuid.UUID) *CycleRenewalCreate {
	if u != nil {
		crc.SetID(*u)
	}
	return crc
}

// Mutation returns the CycleRenewalMutation object of the builder.
func (crc *CycleRenewalCreate) Mutation() *CycleRenewalMutation {
	return crc.mutation
}

// Save creates the CycleRenewal in the database.
func (crc *CycleRenewalCreate) Save(ctx context.Context) (*CycleRenewal, error) {
	crc.defaults()
	return withHooks(ctx, crc.sqlSave, crc.mutation, crc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (crc *CycleRenewalCreate) SaveX(ctx context.Context) *CycleRenewal {
	v, err := crc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (crc *CycleRenewalCreate) Exec(ctx context.Context) error {
	_, err := crc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crc *CycleRenewalCreate) ExecX(ctx context.Context) {
	if err := crc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (crc *CycleRenewalCreate) defaults() {
	if _, ok := crc.mutation.ID(); !ok {
		v := cyclerenewal.DefaultID()
		crc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (crc *CycleRenewalCreate) check() error {
	if _, ok := crc.mutation.FkUserID(); !ok {
		return &ValidationError{Name: "fk_user_id", err: errors.New(`ent: missing required field "CycleRenewal.fk_user_id"`)}
	}
	if _, ok := crc.mutation.ResourceID(); !ok {
		return &ValidationError{Name: "resource_id", err: errors.New(`ent: missing required field "CycleRenewal.resource_id"`)}
	}
	if _, ok := crc.mutation.ResourceType(); !ok {
		return &ValidationError{Name: "resource_type", err: errors.New(`ent: missing required field "CycleRenewal.resource_type"`)}
	}
	if v, ok := crc.mutation.ResourceType(); ok {
		if err := cyclerenewal.ResourceTypeValidator(v); err != nil {
			return &ValidationError{Name: "resource_type", err: fmt.Errorf(`ent: validator failed for field "CycleRenewal.resource_type": %w`, err)}
		}
	}
	if _, ok := crc.mutation.ProductName(); !ok {
		return &ValidationError{Name: "product_name", err: errors.New(`ent: missing required field "CycleRenewal.product_name"`)}
	}
	if v, ok := crc.mutation.ProductName(); ok {
		if err := cyclerenewal.ProductNameValidator(v); err != nil {
			return &ValidationError{Name: "product_name", err: fmt.Errorf(`ent: validator failed for field "CycleRenewal.product_name": %w`, err)}
		}
	}
	if _, ok := crc.mutation.ProductDesc(); !ok {
		return &ValidationError{Name: "product_desc", err: errors.New(`ent: missing required field "CycleRenewal.product_desc"`)}
	}
	if v, ok := crc.mutation.ProductDesc(); ok {
		if err := cyclerenewal.ProductDescValidator(v); err != nil {
			return &ValidationError{Name: "product_desc", err: fmt.Errorf(`ent: validator failed for field "CycleRenewal.product_desc": %w`, err)}
		}
	}
	if _, ok := crc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "CycleRenewal.state"`)}
	}
	if _, ok := crc.mutation.ExtendDay(); !ok {
		return &ValidationError{Name: "extend_day", err: errors.New(`ent: missing required field "CycleRenewal.extend_day"`)}
	}
	if _, ok := crc.mutation.ExtendPrice(); !ok {
		return &ValidationError{Name: "extend_price", err: errors.New(`ent: missing required field "CycleRenewal.extend_price"`)}
	}
	if _, ok := crc.mutation.AutoRenewal(); !ok {
		return &ValidationError{Name: "auto_renewal", err: errors.New(`ent: missing required field "CycleRenewal.auto_renewal"`)}
	}
	return nil
}

func (crc *CycleRenewalCreate) sqlSave(ctx context.Context) (*CycleRenewal, error) {
	if err := crc.check(); err != nil {
		return nil, err
	}
	_node, _spec := crc.createSpec()
	if err := sqlgraph.CreateNode(ctx, crc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	crc.mutation.id = &_node.ID
	crc.mutation.done = true
	return _node, nil
}

func (crc *CycleRenewalCreate) createSpec() (*CycleRenewal, *sqlgraph.CreateSpec) {
	var (
		_node = &CycleRenewal{config: crc.config}
		_spec = sqlgraph.NewCreateSpec(cyclerenewal.Table, sqlgraph.NewFieldSpec(cyclerenewal.FieldID, field.TypeUUID))
	)
	if id, ok := crc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := crc.mutation.FkUserID(); ok {
		_spec.SetField(cyclerenewal.FieldFkUserID, field.TypeUUID, value)
		_node.FkUserID = value
	}
	if value, ok := crc.mutation.ResourceID(); ok {
		_spec.SetField(cyclerenewal.FieldResourceID, field.TypeUUID, value)
		_node.ResourceID = value
	}
	if value, ok := crc.mutation.ResourceType(); ok {
		_spec.SetField(cyclerenewal.FieldResourceType, field.TypeInt, value)
		_node.ResourceType = value
	}
	if value, ok := crc.mutation.ProductName(); ok {
		_spec.SetField(cyclerenewal.FieldProductName, field.TypeString, value)
		_node.ProductName = value
	}
	if value, ok := crc.mutation.ProductDesc(); ok {
		_spec.SetField(cyclerenewal.FieldProductDesc, field.TypeString, value)
		_node.ProductDesc = value
	}
	if value, ok := crc.mutation.State(); ok {
		_spec.SetField(cyclerenewal.FieldState, field.TypeInt8, value)
		_node.State = value
	}
	if value, ok := crc.mutation.ExtendDay(); ok {
		_spec.SetField(cyclerenewal.FieldExtendDay, field.TypeInt8, value)
		_node.ExtendDay = value
	}
	if value, ok := crc.mutation.ExtendPrice(); ok {
		_spec.SetField(cyclerenewal.FieldExtendPrice, field.TypeFloat64, value)
		_node.ExtendPrice = value
	}
	if value, ok := crc.mutation.DueTime(); ok {
		_spec.SetField(cyclerenewal.FieldDueTime, field.TypeTime, value)
		_node.DueTime = &value
	}
	if value, ok := crc.mutation.RenewalTime(); ok {
		_spec.SetField(cyclerenewal.FieldRenewalTime, field.TypeTime, value)
		_node.RenewalTime = &value
	}
	if value, ok := crc.mutation.AutoRenewal(); ok {
		_spec.SetField(cyclerenewal.FieldAutoRenewal, field.TypeBool, value)
		_node.AutoRenewal = value
	}
	return _node, _spec
}

// CycleRenewalCreateBulk is the builder for creating many CycleRenewal entities in bulk.
type CycleRenewalCreateBulk struct {
	config
	builders []*CycleRenewalCreate
}

// Save creates the CycleRenewal entities in the database.
func (crcb *CycleRenewalCreateBulk) Save(ctx context.Context) ([]*CycleRenewal, error) {
	specs := make([]*sqlgraph.CreateSpec, len(crcb.builders))
	nodes := make([]*CycleRenewal, len(crcb.builders))
	mutators := make([]Mutator, len(crcb.builders))
	for i := range crcb.builders {
		func(i int, root context.Context) {
			builder := crcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CycleRenewalMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, crcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, crcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, crcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (crcb *CycleRenewalCreateBulk) SaveX(ctx context.Context) []*CycleRenewal {
	v, err := crcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (crcb *CycleRenewalCreateBulk) Exec(ctx context.Context) error {
	_, err := crcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crcb *CycleRenewalCreateBulk) ExecX(ctx context.Context) {
	if err := crcb.Exec(ctx); err != nil {
		panic(err)
	}
}
