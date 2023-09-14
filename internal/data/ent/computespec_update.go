// Code generated by ent, DO NOT EDIT.

package ent

import (
	"computeshare-server/internal/data/ent/computespec"
	"computeshare-server/internal/data/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ComputeSpecUpdate is the builder for updating ComputeSpec entities.
type ComputeSpecUpdate struct {
	config
	hooks    []Hook
	mutation *ComputeSpecMutation
}

// Where appends a list predicates to the ComputeSpecUpdate builder.
func (csu *ComputeSpecUpdate) Where(ps ...predicate.ComputeSpec) *ComputeSpecUpdate {
	csu.mutation.Where(ps...)
	return csu
}

// SetCore sets the "core" field.
func (csu *ComputeSpecUpdate) SetCore(s string) *ComputeSpecUpdate {
	csu.mutation.SetCore(s)
	return csu
}

// SetMemory sets the "memory" field.
func (csu *ComputeSpecUpdate) SetMemory(s string) *ComputeSpecUpdate {
	csu.mutation.SetMemory(s)
	return csu
}

// Mutation returns the ComputeSpecMutation object of the builder.
func (csu *ComputeSpecUpdate) Mutation() *ComputeSpecMutation {
	return csu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (csu *ComputeSpecUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, csu.sqlSave, csu.mutation, csu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (csu *ComputeSpecUpdate) SaveX(ctx context.Context) int {
	affected, err := csu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (csu *ComputeSpecUpdate) Exec(ctx context.Context) error {
	_, err := csu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csu *ComputeSpecUpdate) ExecX(ctx context.Context) {
	if err := csu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csu *ComputeSpecUpdate) check() error {
	if v, ok := csu.mutation.Core(); ok {
		if err := computespec.CoreValidator(v); err != nil {
			return &ValidationError{Name: "core", err: fmt.Errorf(`ent: validator failed for field "ComputeSpec.core": %w`, err)}
		}
	}
	if v, ok := csu.mutation.Memory(); ok {
		if err := computespec.MemoryValidator(v); err != nil {
			return &ValidationError{Name: "memory", err: fmt.Errorf(`ent: validator failed for field "ComputeSpec.memory": %w`, err)}
		}
	}
	return nil
}

func (csu *ComputeSpecUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := csu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(computespec.Table, computespec.Columns, sqlgraph.NewFieldSpec(computespec.FieldID, field.TypeInt32))
	if ps := csu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csu.mutation.Core(); ok {
		_spec.SetField(computespec.FieldCore, field.TypeString, value)
	}
	if value, ok := csu.mutation.Memory(); ok {
		_spec.SetField(computespec.FieldMemory, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, csu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{computespec.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	csu.mutation.done = true
	return n, nil
}

// ComputeSpecUpdateOne is the builder for updating a single ComputeSpec entity.
type ComputeSpecUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ComputeSpecMutation
}

// SetCore sets the "core" field.
func (csuo *ComputeSpecUpdateOne) SetCore(s string) *ComputeSpecUpdateOne {
	csuo.mutation.SetCore(s)
	return csuo
}

// SetMemory sets the "memory" field.
func (csuo *ComputeSpecUpdateOne) SetMemory(s string) *ComputeSpecUpdateOne {
	csuo.mutation.SetMemory(s)
	return csuo
}

// Mutation returns the ComputeSpecMutation object of the builder.
func (csuo *ComputeSpecUpdateOne) Mutation() *ComputeSpecMutation {
	return csuo.mutation
}

// Where appends a list predicates to the ComputeSpecUpdate builder.
func (csuo *ComputeSpecUpdateOne) Where(ps ...predicate.ComputeSpec) *ComputeSpecUpdateOne {
	csuo.mutation.Where(ps...)
	return csuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (csuo *ComputeSpecUpdateOne) Select(field string, fields ...string) *ComputeSpecUpdateOne {
	csuo.fields = append([]string{field}, fields...)
	return csuo
}

// Save executes the query and returns the updated ComputeSpec entity.
func (csuo *ComputeSpecUpdateOne) Save(ctx context.Context) (*ComputeSpec, error) {
	return withHooks(ctx, csuo.sqlSave, csuo.mutation, csuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (csuo *ComputeSpecUpdateOne) SaveX(ctx context.Context) *ComputeSpec {
	node, err := csuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (csuo *ComputeSpecUpdateOne) Exec(ctx context.Context) error {
	_, err := csuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csuo *ComputeSpecUpdateOne) ExecX(ctx context.Context) {
	if err := csuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csuo *ComputeSpecUpdateOne) check() error {
	if v, ok := csuo.mutation.Core(); ok {
		if err := computespec.CoreValidator(v); err != nil {
			return &ValidationError{Name: "core", err: fmt.Errorf(`ent: validator failed for field "ComputeSpec.core": %w`, err)}
		}
	}
	if v, ok := csuo.mutation.Memory(); ok {
		if err := computespec.MemoryValidator(v); err != nil {
			return &ValidationError{Name: "memory", err: fmt.Errorf(`ent: validator failed for field "ComputeSpec.memory": %w`, err)}
		}
	}
	return nil
}

func (csuo *ComputeSpecUpdateOne) sqlSave(ctx context.Context) (_node *ComputeSpec, err error) {
	if err := csuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(computespec.Table, computespec.Columns, sqlgraph.NewFieldSpec(computespec.FieldID, field.TypeInt32))
	id, ok := csuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ComputeSpec.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := csuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, computespec.FieldID)
		for _, f := range fields {
			if !computespec.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != computespec.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := csuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csuo.mutation.Core(); ok {
		_spec.SetField(computespec.FieldCore, field.TypeString, value)
	}
	if value, ok := csuo.mutation.Memory(); ok {
		_spec.SetField(computespec.FieldMemory, field.TypeString, value)
	}
	_node = &ComputeSpec{config: csuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, csuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{computespec.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	csuo.mutation.done = true
	return _node, nil
}
