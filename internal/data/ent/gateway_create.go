// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/gateway"
)

// GatewayCreate is the builder for creating a Gateway entity.
type GatewayCreate struct {
	config
	mutation *GatewayMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (gc *GatewayCreate) SetName(s string) *GatewayCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetIP sets the "ip" field.
func (gc *GatewayCreate) SetIP(s string) *GatewayCreate {
	gc.mutation.SetIP(s)
	return gc
}

// SetPort sets the "port" field.
func (gc *GatewayCreate) SetPort(i int32) *GatewayCreate {
	gc.mutation.SetPort(i)
	return gc
}

// SetInternalIP sets the "internal_ip" field.
func (gc *GatewayCreate) SetInternalIP(s string) *GatewayCreate {
	gc.mutation.SetInternalIP(s)
	return gc
}

// SetID sets the "id" field.
func (gc *GatewayCreate) SetID(u uuid.UUID) *GatewayCreate {
	gc.mutation.SetID(u)
	return gc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gc *GatewayCreate) SetNillableID(u *uuid.UUID) *GatewayCreate {
	if u != nil {
		gc.SetID(*u)
	}
	return gc
}

// Mutation returns the GatewayMutation object of the builder.
func (gc *GatewayCreate) Mutation() *GatewayMutation {
	return gc.mutation
}

// Save creates the Gateway in the database.
func (gc *GatewayCreate) Save(ctx context.Context) (*Gateway, error) {
	gc.defaults()
	return withHooks(ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GatewayCreate) SaveX(ctx context.Context) *Gateway {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GatewayCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GatewayCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GatewayCreate) defaults() {
	if _, ok := gc.mutation.ID(); !ok {
		v := gateway.DefaultID()
		gc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GatewayCreate) check() error {
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Gateway.name"`)}
	}
	if v, ok := gc.mutation.Name(); ok {
		if err := gateway.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Gateway.name": %w`, err)}
		}
	}
	if _, ok := gc.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New(`ent: missing required field "Gateway.ip"`)}
	}
	if _, ok := gc.mutation.Port(); !ok {
		return &ValidationError{Name: "port", err: errors.New(`ent: missing required field "Gateway.port"`)}
	}
	if _, ok := gc.mutation.InternalIP(); !ok {
		return &ValidationError{Name: "internal_ip", err: errors.New(`ent: missing required field "Gateway.internal_ip"`)}
	}
	return nil
}

func (gc *GatewayCreate) sqlSave(ctx context.Context) (*Gateway, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
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
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GatewayCreate) createSpec() (*Gateway, *sqlgraph.CreateSpec) {
	var (
		_node = &Gateway{config: gc.config}
		_spec = sqlgraph.NewCreateSpec(gateway.Table, sqlgraph.NewFieldSpec(gateway.FieldID, field.TypeUUID))
	)
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := gc.mutation.Name(); ok {
		_spec.SetField(gateway.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := gc.mutation.IP(); ok {
		_spec.SetField(gateway.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if value, ok := gc.mutation.Port(); ok {
		_spec.SetField(gateway.FieldPort, field.TypeInt32, value)
		_node.Port = value
	}
	if value, ok := gc.mutation.InternalIP(); ok {
		_spec.SetField(gateway.FieldInternalIP, field.TypeString, value)
		_node.InternalIP = value
	}
	return _node, _spec
}

// GatewayCreateBulk is the builder for creating many Gateway entities in bulk.
type GatewayCreateBulk struct {
	config
	builders []*GatewayCreate
}

// Save creates the Gateway entities in the database.
func (gcb *GatewayCreateBulk) Save(ctx context.Context) ([]*Gateway, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Gateway, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GatewayMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GatewayCreateBulk) SaveX(ctx context.Context) []*Gateway {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GatewayCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GatewayCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}
