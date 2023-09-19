// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/computespec"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// ComputeSpecDelete is the builder for deleting a ComputeSpec entity.
type ComputeSpecDelete struct {
	config
	hooks    []Hook
	mutation *ComputeSpecMutation
}

// Where appends a list predicates to the ComputeSpecDelete builder.
func (csd *ComputeSpecDelete) Where(ps ...predicate.ComputeSpec) *ComputeSpecDelete {
	csd.mutation.Where(ps...)
	return csd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (csd *ComputeSpecDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, csd.sqlExec, csd.mutation, csd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (csd *ComputeSpecDelete) ExecX(ctx context.Context) int {
	n, err := csd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (csd *ComputeSpecDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(computespec.Table, sqlgraph.NewFieldSpec(computespec.FieldID, field.TypeInt32))
	if ps := csd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, csd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	csd.mutation.done = true
	return affected, err
}

// ComputeSpecDeleteOne is the builder for deleting a single ComputeSpec entity.
type ComputeSpecDeleteOne struct {
	csd *ComputeSpecDelete
}

// Where appends a list predicates to the ComputeSpecDelete builder.
func (csdo *ComputeSpecDeleteOne) Where(ps ...predicate.ComputeSpec) *ComputeSpecDeleteOne {
	csdo.csd.mutation.Where(ps...)
	return csdo
}

// Exec executes the deletion query.
func (csdo *ComputeSpecDeleteOne) Exec(ctx context.Context) error {
	n, err := csdo.csd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{computespec.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (csdo *ComputeSpecDeleteOne) ExecX(ctx context.Context) {
	if err := csdo.Exec(ctx); err != nil {
		panic(err)
	}
}