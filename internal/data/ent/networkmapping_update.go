// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/networkmapping"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// NetworkMappingUpdate is the builder for updating NetworkMapping entities.
type NetworkMappingUpdate struct {
	config
	hooks    []Hook
	mutation *NetworkMappingMutation
}

// Where appends a list predicates to the NetworkMappingUpdate builder.
func (nmu *NetworkMappingUpdate) Where(ps ...predicate.NetworkMapping) *NetworkMappingUpdate {
	nmu.mutation.Where(ps...)
	return nmu
}

// SetName sets the "name" field.
func (nmu *NetworkMappingUpdate) SetName(s string) *NetworkMappingUpdate {
	nmu.mutation.SetName(s)
	return nmu
}

// SetFkGatewayID sets the "fk_gateway_id" field.
func (nmu *NetworkMappingUpdate) SetFkGatewayID(u uuid.UUID) *NetworkMappingUpdate {
	nmu.mutation.SetFkGatewayID(u)
	return nmu
}

// SetFkComputerID sets the "fk_computer_id" field.
func (nmu *NetworkMappingUpdate) SetFkComputerID(u uuid.UUID) *NetworkMappingUpdate {
	nmu.mutation.SetFkComputerID(u)
	return nmu
}

// SetGatewayPort sets the "gateway_port" field.
func (nmu *NetworkMappingUpdate) SetGatewayPort(i int) *NetworkMappingUpdate {
	nmu.mutation.ResetGatewayPort()
	nmu.mutation.SetGatewayPort(i)
	return nmu
}

// AddGatewayPort adds i to the "gateway_port" field.
func (nmu *NetworkMappingUpdate) AddGatewayPort(i int) *NetworkMappingUpdate {
	nmu.mutation.AddGatewayPort(i)
	return nmu
}

// SetComputerPort sets the "computer_port" field.
func (nmu *NetworkMappingUpdate) SetComputerPort(i int) *NetworkMappingUpdate {
	nmu.mutation.ResetComputerPort()
	nmu.mutation.SetComputerPort(i)
	return nmu
}

// AddComputerPort adds i to the "computer_port" field.
func (nmu *NetworkMappingUpdate) AddComputerPort(i int) *NetworkMappingUpdate {
	nmu.mutation.AddComputerPort(i)
	return nmu
}

// SetStatus sets the "status" field.
func (nmu *NetworkMappingUpdate) SetStatus(i int) *NetworkMappingUpdate {
	nmu.mutation.ResetStatus()
	nmu.mutation.SetStatus(i)
	return nmu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (nmu *NetworkMappingUpdate) SetNillableStatus(i *int) *NetworkMappingUpdate {
	if i != nil {
		nmu.SetStatus(*i)
	}
	return nmu
}

// AddStatus adds i to the "status" field.
func (nmu *NetworkMappingUpdate) AddStatus(i int) *NetworkMappingUpdate {
	nmu.mutation.AddStatus(i)
	return nmu
}

// Mutation returns the NetworkMappingMutation object of the builder.
func (nmu *NetworkMappingUpdate) Mutation() *NetworkMappingMutation {
	return nmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nmu *NetworkMappingUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nmu.sqlSave, nmu.mutation, nmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nmu *NetworkMappingUpdate) SaveX(ctx context.Context) int {
	affected, err := nmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nmu *NetworkMappingUpdate) Exec(ctx context.Context) error {
	_, err := nmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nmu *NetworkMappingUpdate) ExecX(ctx context.Context) {
	if err := nmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nmu *NetworkMappingUpdate) check() error {
	if v, ok := nmu.mutation.Name(); ok {
		if err := networkmapping.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "NetworkMapping.name": %w`, err)}
		}
	}
	return nil
}

func (nmu *NetworkMappingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nmu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(networkmapping.Table, networkmapping.Columns, sqlgraph.NewFieldSpec(networkmapping.FieldID, field.TypeUUID))
	if ps := nmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nmu.mutation.Name(); ok {
		_spec.SetField(networkmapping.FieldName, field.TypeString, value)
	}
	if value, ok := nmu.mutation.FkGatewayID(); ok {
		_spec.SetField(networkmapping.FieldFkGatewayID, field.TypeUUID, value)
	}
	if value, ok := nmu.mutation.FkComputerID(); ok {
		_spec.SetField(networkmapping.FieldFkComputerID, field.TypeUUID, value)
	}
	if value, ok := nmu.mutation.GatewayPort(); ok {
		_spec.SetField(networkmapping.FieldGatewayPort, field.TypeInt, value)
	}
	if value, ok := nmu.mutation.AddedGatewayPort(); ok {
		_spec.AddField(networkmapping.FieldGatewayPort, field.TypeInt, value)
	}
	if value, ok := nmu.mutation.ComputerPort(); ok {
		_spec.SetField(networkmapping.FieldComputerPort, field.TypeInt, value)
	}
	if value, ok := nmu.mutation.AddedComputerPort(); ok {
		_spec.AddField(networkmapping.FieldComputerPort, field.TypeInt, value)
	}
	if value, ok := nmu.mutation.Status(); ok {
		_spec.SetField(networkmapping.FieldStatus, field.TypeInt, value)
	}
	if value, ok := nmu.mutation.AddedStatus(); ok {
		_spec.AddField(networkmapping.FieldStatus, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{networkmapping.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nmu.mutation.done = true
	return n, nil
}

// NetworkMappingUpdateOne is the builder for updating a single NetworkMapping entity.
type NetworkMappingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NetworkMappingMutation
}

// SetName sets the "name" field.
func (nmuo *NetworkMappingUpdateOne) SetName(s string) *NetworkMappingUpdateOne {
	nmuo.mutation.SetName(s)
	return nmuo
}

// SetFkGatewayID sets the "fk_gateway_id" field.
func (nmuo *NetworkMappingUpdateOne) SetFkGatewayID(u uuid.UUID) *NetworkMappingUpdateOne {
	nmuo.mutation.SetFkGatewayID(u)
	return nmuo
}

// SetFkComputerID sets the "fk_computer_id" field.
func (nmuo *NetworkMappingUpdateOne) SetFkComputerID(u uuid.UUID) *NetworkMappingUpdateOne {
	nmuo.mutation.SetFkComputerID(u)
	return nmuo
}

// SetGatewayPort sets the "gateway_port" field.
func (nmuo *NetworkMappingUpdateOne) SetGatewayPort(i int) *NetworkMappingUpdateOne {
	nmuo.mutation.ResetGatewayPort()
	nmuo.mutation.SetGatewayPort(i)
	return nmuo
}

// AddGatewayPort adds i to the "gateway_port" field.
func (nmuo *NetworkMappingUpdateOne) AddGatewayPort(i int) *NetworkMappingUpdateOne {
	nmuo.mutation.AddGatewayPort(i)
	return nmuo
}

// SetComputerPort sets the "computer_port" field.
func (nmuo *NetworkMappingUpdateOne) SetComputerPort(i int) *NetworkMappingUpdateOne {
	nmuo.mutation.ResetComputerPort()
	nmuo.mutation.SetComputerPort(i)
	return nmuo
}

// AddComputerPort adds i to the "computer_port" field.
func (nmuo *NetworkMappingUpdateOne) AddComputerPort(i int) *NetworkMappingUpdateOne {
	nmuo.mutation.AddComputerPort(i)
	return nmuo
}

// SetStatus sets the "status" field.
func (nmuo *NetworkMappingUpdateOne) SetStatus(i int) *NetworkMappingUpdateOne {
	nmuo.mutation.ResetStatus()
	nmuo.mutation.SetStatus(i)
	return nmuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (nmuo *NetworkMappingUpdateOne) SetNillableStatus(i *int) *NetworkMappingUpdateOne {
	if i != nil {
		nmuo.SetStatus(*i)
	}
	return nmuo
}

// AddStatus adds i to the "status" field.
func (nmuo *NetworkMappingUpdateOne) AddStatus(i int) *NetworkMappingUpdateOne {
	nmuo.mutation.AddStatus(i)
	return nmuo
}

// Mutation returns the NetworkMappingMutation object of the builder.
func (nmuo *NetworkMappingUpdateOne) Mutation() *NetworkMappingMutation {
	return nmuo.mutation
}

// Where appends a list predicates to the NetworkMappingUpdate builder.
func (nmuo *NetworkMappingUpdateOne) Where(ps ...predicate.NetworkMapping) *NetworkMappingUpdateOne {
	nmuo.mutation.Where(ps...)
	return nmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nmuo *NetworkMappingUpdateOne) Select(field string, fields ...string) *NetworkMappingUpdateOne {
	nmuo.fields = append([]string{field}, fields...)
	return nmuo
}

// Save executes the query and returns the updated NetworkMapping entity.
func (nmuo *NetworkMappingUpdateOne) Save(ctx context.Context) (*NetworkMapping, error) {
	return withHooks(ctx, nmuo.sqlSave, nmuo.mutation, nmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nmuo *NetworkMappingUpdateOne) SaveX(ctx context.Context) *NetworkMapping {
	node, err := nmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nmuo *NetworkMappingUpdateOne) Exec(ctx context.Context) error {
	_, err := nmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nmuo *NetworkMappingUpdateOne) ExecX(ctx context.Context) {
	if err := nmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nmuo *NetworkMappingUpdateOne) check() error {
	if v, ok := nmuo.mutation.Name(); ok {
		if err := networkmapping.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "NetworkMapping.name": %w`, err)}
		}
	}
	return nil
}

func (nmuo *NetworkMappingUpdateOne) sqlSave(ctx context.Context) (_node *NetworkMapping, err error) {
	if err := nmuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(networkmapping.Table, networkmapping.Columns, sqlgraph.NewFieldSpec(networkmapping.FieldID, field.TypeUUID))
	id, ok := nmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NetworkMapping.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, networkmapping.FieldID)
		for _, f := range fields {
			if !networkmapping.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != networkmapping.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nmuo.mutation.Name(); ok {
		_spec.SetField(networkmapping.FieldName, field.TypeString, value)
	}
	if value, ok := nmuo.mutation.FkGatewayID(); ok {
		_spec.SetField(networkmapping.FieldFkGatewayID, field.TypeUUID, value)
	}
	if value, ok := nmuo.mutation.FkComputerID(); ok {
		_spec.SetField(networkmapping.FieldFkComputerID, field.TypeUUID, value)
	}
	if value, ok := nmuo.mutation.GatewayPort(); ok {
		_spec.SetField(networkmapping.FieldGatewayPort, field.TypeInt, value)
	}
	if value, ok := nmuo.mutation.AddedGatewayPort(); ok {
		_spec.AddField(networkmapping.FieldGatewayPort, field.TypeInt, value)
	}
	if value, ok := nmuo.mutation.ComputerPort(); ok {
		_spec.SetField(networkmapping.FieldComputerPort, field.TypeInt, value)
	}
	if value, ok := nmuo.mutation.AddedComputerPort(); ok {
		_spec.AddField(networkmapping.FieldComputerPort, field.TypeInt, value)
	}
	if value, ok := nmuo.mutation.Status(); ok {
		_spec.SetField(networkmapping.FieldStatus, field.TypeInt, value)
	}
	if value, ok := nmuo.mutation.AddedStatus(); ok {
		_spec.AddField(networkmapping.FieldStatus, field.TypeInt, value)
	}
	_node = &NetworkMapping{config: nmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{networkmapping.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nmuo.mutation.done = true
	return _node, nil
}
