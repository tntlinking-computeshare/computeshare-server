// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cyclerenewal"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// CycleRenewalDelete is the builder for deleting a CycleRenewal entity.
type CycleRenewalDelete struct {
	config
	hooks    []Hook
	mutation *CycleRenewalMutation
}

// Where appends a list predicates to the CycleRenewalDelete builder.
func (crd *CycleRenewalDelete) Where(ps ...predicate.CycleRenewal) *CycleRenewalDelete {
	crd.mutation.Where(ps...)
	return crd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (crd *CycleRenewalDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, crd.sqlExec, crd.mutation, crd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (crd *CycleRenewalDelete) ExecX(ctx context.Context) int {
	n, err := crd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (crd *CycleRenewalDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(cyclerenewal.Table, sqlgraph.NewFieldSpec(cyclerenewal.FieldID, field.TypeUUID))
	if ps := crd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, crd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	crd.mutation.done = true
	return affected, err
}

// CycleRenewalDeleteOne is the builder for deleting a single CycleRenewal entity.
type CycleRenewalDeleteOne struct {
	crd *CycleRenewalDelete
}

// Where appends a list predicates to the CycleRenewalDelete builder.
func (crdo *CycleRenewalDeleteOne) Where(ps ...predicate.CycleRenewal) *CycleRenewalDeleteOne {
	crdo.crd.mutation.Where(ps...)
	return crdo
}

// Exec executes the deletion query.
func (crdo *CycleRenewalDeleteOne) Exec(ctx context.Context) error {
	n, err := crdo.crd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cyclerenewal.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (crdo *CycleRenewalDeleteOne) ExecX(ctx context.Context) {
	if err := crdo.Exec(ctx); err != nil {
		panic(err)
	}
}
