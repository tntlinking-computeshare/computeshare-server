// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/computeimage"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// ComputeImageDelete is the builder for deleting a ComputeImage entity.
type ComputeImageDelete struct {
	config
	hooks    []Hook
	mutation *ComputeImageMutation
}

// Where appends a list predicates to the ComputeImageDelete builder.
func (cid *ComputeImageDelete) Where(ps ...predicate.ComputeImage) *ComputeImageDelete {
	cid.mutation.Where(ps...)
	return cid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cid *ComputeImageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cid.sqlExec, cid.mutation, cid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cid *ComputeImageDelete) ExecX(ctx context.Context) int {
	n, err := cid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cid *ComputeImageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(computeimage.Table, sqlgraph.NewFieldSpec(computeimage.FieldID, field.TypeInt32))
	if ps := cid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cid.mutation.done = true
	return affected, err
}

// ComputeImageDeleteOne is the builder for deleting a single ComputeImage entity.
type ComputeImageDeleteOne struct {
	cid *ComputeImageDelete
}

// Where appends a list predicates to the ComputeImageDelete builder.
func (cido *ComputeImageDeleteOne) Where(ps ...predicate.ComputeImage) *ComputeImageDeleteOne {
	cido.cid.mutation.Where(ps...)
	return cido
}

// Exec executes the deletion query.
func (cido *ComputeImageDeleteOne) Exec(ctx context.Context) error {
	n, err := cido.cid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{computeimage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cido *ComputeImageDeleteOne) ExecX(ctx context.Context) {
	if err := cido.Exec(ctx); err != nil {
		panic(err)
	}
}