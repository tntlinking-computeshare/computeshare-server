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
	"github.com/mohaijiang/computeshare-server/internal/data/ent/domainbinding"
)

// DomainBindingCreate is the builder for creating a DomainBinding entity.
type DomainBindingCreate struct {
	config
	mutation *DomainBindingMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (dbc *DomainBindingCreate) SetUserID(u uuid.UUID) *DomainBindingCreate {
	dbc.mutation.SetUserID(u)
	return dbc
}

// SetFkComputeInstanceID sets the "fk_compute_instance_id" field.
func (dbc *DomainBindingCreate) SetFkComputeInstanceID(u uuid.UUID) *DomainBindingCreate {
	dbc.mutation.SetFkComputeInstanceID(u)
	return dbc
}

// SetFkNetworkMappingID sets the "fk_network_mapping_id" field.
func (dbc *DomainBindingCreate) SetFkNetworkMappingID(u uuid.UUID) *DomainBindingCreate {
	dbc.mutation.SetFkNetworkMappingID(u)
	return dbc
}

// SetName sets the "name" field.
func (dbc *DomainBindingCreate) SetName(s string) *DomainBindingCreate {
	dbc.mutation.SetName(s)
	return dbc
}

// SetDomain sets the "domain" field.
func (dbc *DomainBindingCreate) SetDomain(s string) *DomainBindingCreate {
	dbc.mutation.SetDomain(s)
	return dbc
}

// SetGatewayPort sets the "gateway_port" field.
func (dbc *DomainBindingCreate) SetGatewayPort(i int32) *DomainBindingCreate {
	dbc.mutation.SetGatewayPort(i)
	return dbc
}

// SetCreateTime sets the "create_time" field.
func (dbc *DomainBindingCreate) SetCreateTime(t time.Time) *DomainBindingCreate {
	dbc.mutation.SetCreateTime(t)
	return dbc
}

// SetID sets the "id" field.
func (dbc *DomainBindingCreate) SetID(u uuid.UUID) *DomainBindingCreate {
	dbc.mutation.SetID(u)
	return dbc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dbc *DomainBindingCreate) SetNillableID(u *uuid.UUID) *DomainBindingCreate {
	if u != nil {
		dbc.SetID(*u)
	}
	return dbc
}

// Mutation returns the DomainBindingMutation object of the builder.
func (dbc *DomainBindingCreate) Mutation() *DomainBindingMutation {
	return dbc.mutation
}

// Save creates the DomainBinding in the database.
func (dbc *DomainBindingCreate) Save(ctx context.Context) (*DomainBinding, error) {
	dbc.defaults()
	return withHooks(ctx, dbc.sqlSave, dbc.mutation, dbc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dbc *DomainBindingCreate) SaveX(ctx context.Context) *DomainBinding {
	v, err := dbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dbc *DomainBindingCreate) Exec(ctx context.Context) error {
	_, err := dbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dbc *DomainBindingCreate) ExecX(ctx context.Context) {
	if err := dbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dbc *DomainBindingCreate) defaults() {
	if _, ok := dbc.mutation.ID(); !ok {
		v := domainbinding.DefaultID()
		dbc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dbc *DomainBindingCreate) check() error {
	if _, ok := dbc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "DomainBinding.user_id"`)}
	}
	if _, ok := dbc.mutation.FkComputeInstanceID(); !ok {
		return &ValidationError{Name: "fk_compute_instance_id", err: errors.New(`ent: missing required field "DomainBinding.fk_compute_instance_id"`)}
	}
	if _, ok := dbc.mutation.FkNetworkMappingID(); !ok {
		return &ValidationError{Name: "fk_network_mapping_id", err: errors.New(`ent: missing required field "DomainBinding.fk_network_mapping_id"`)}
	}
	if _, ok := dbc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "DomainBinding.name"`)}
	}
	if v, ok := dbc.mutation.Name(); ok {
		if err := domainbinding.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "DomainBinding.name": %w`, err)}
		}
	}
	if _, ok := dbc.mutation.Domain(); !ok {
		return &ValidationError{Name: "domain", err: errors.New(`ent: missing required field "DomainBinding.domain"`)}
	}
	if v, ok := dbc.mutation.Domain(); ok {
		if err := domainbinding.DomainValidator(v); err != nil {
			return &ValidationError{Name: "domain", err: fmt.Errorf(`ent: validator failed for field "DomainBinding.domain": %w`, err)}
		}
	}
	if _, ok := dbc.mutation.GatewayPort(); !ok {
		return &ValidationError{Name: "gateway_port", err: errors.New(`ent: missing required field "DomainBinding.gateway_port"`)}
	}
	if _, ok := dbc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "DomainBinding.create_time"`)}
	}
	return nil
}

func (dbc *DomainBindingCreate) sqlSave(ctx context.Context) (*DomainBinding, error) {
	if err := dbc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dbc.driver, _spec); err != nil {
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
	dbc.mutation.id = &_node.ID
	dbc.mutation.done = true
	return _node, nil
}

func (dbc *DomainBindingCreate) createSpec() (*DomainBinding, *sqlgraph.CreateSpec) {
	var (
		_node = &DomainBinding{config: dbc.config}
		_spec = sqlgraph.NewCreateSpec(domainbinding.Table, sqlgraph.NewFieldSpec(domainbinding.FieldID, field.TypeUUID))
	)
	if id, ok := dbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dbc.mutation.UserID(); ok {
		_spec.SetField(domainbinding.FieldUserID, field.TypeUUID, value)
		_node.UserID = value
	}
	if value, ok := dbc.mutation.FkComputeInstanceID(); ok {
		_spec.SetField(domainbinding.FieldFkComputeInstanceID, field.TypeUUID, value)
		_node.FkComputeInstanceID = value
	}
	if value, ok := dbc.mutation.FkNetworkMappingID(); ok {
		_spec.SetField(domainbinding.FieldFkNetworkMappingID, field.TypeUUID, value)
		_node.FkNetworkMappingID = value
	}
	if value, ok := dbc.mutation.Name(); ok {
		_spec.SetField(domainbinding.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dbc.mutation.Domain(); ok {
		_spec.SetField(domainbinding.FieldDomain, field.TypeString, value)
		_node.Domain = value
	}
	if value, ok := dbc.mutation.GatewayPort(); ok {
		_spec.SetField(domainbinding.FieldGatewayPort, field.TypeInt32, value)
		_node.GatewayPort = value
	}
	if value, ok := dbc.mutation.CreateTime(); ok {
		_spec.SetField(domainbinding.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	return _node, _spec
}

// DomainBindingCreateBulk is the builder for creating many DomainBinding entities in bulk.
type DomainBindingCreateBulk struct {
	config
	builders []*DomainBindingCreate
}

// Save creates the DomainBinding entities in the database.
func (dbcb *DomainBindingCreateBulk) Save(ctx context.Context) ([]*DomainBinding, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dbcb.builders))
	nodes := make([]*DomainBinding, len(dbcb.builders))
	mutators := make([]Mutator, len(dbcb.builders))
	for i := range dbcb.builders {
		func(i int, root context.Context) {
			builder := dbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DomainBindingMutation)
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
					_, err = mutators[i+1].Mutate(root, dbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dbcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dbcb *DomainBindingCreateBulk) SaveX(ctx context.Context) []*DomainBinding {
	v, err := dbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dbcb *DomainBindingCreateBulk) Exec(ctx context.Context) error {
	_, err := dbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dbcb *DomainBindingCreateBulk) ExecX(ctx context.Context) {
	if err := dbcb.Exec(ctx); err != nil {
		panic(err)
	}
}