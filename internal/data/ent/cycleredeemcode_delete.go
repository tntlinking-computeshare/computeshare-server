// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cycleredeemcode"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// CycleRedeemCodeDelete is the builder for deleting a CycleRedeemCode entity.
type CycleRedeemCodeDelete struct {
	config
	hooks    []Hook
	mutation *CycleRedeemCodeMutation
}

// Where appends a list predicates to the CycleRedeemCodeDelete builder.
func (crcd *CycleRedeemCodeDelete) Where(ps ...predicate.CycleRedeemCode) *CycleRedeemCodeDelete {
	crcd.mutation.Where(ps...)
	return crcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (crcd *CycleRedeemCodeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, crcd.sqlExec, crcd.mutation, crcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (crcd *CycleRedeemCodeDelete) ExecX(ctx context.Context) int {
	n, err := crcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (crcd *CycleRedeemCodeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(cycleredeemcode.Table, sqlgraph.NewFieldSpec(cycleredeemcode.FieldID, field.TypeUUID))
	if ps := crcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, crcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	crcd.mutation.done = true
	return affected, err
}

// CycleRedeemCodeDeleteOne is the builder for deleting a single CycleRedeemCode entity.
type CycleRedeemCodeDeleteOne struct {
	crcd *CycleRedeemCodeDelete
}

// Where appends a list predicates to the CycleRedeemCodeDelete builder.
func (crcdo *CycleRedeemCodeDeleteOne) Where(ps ...predicate.CycleRedeemCode) *CycleRedeemCodeDeleteOne {
	crcdo.crcd.mutation.Where(ps...)
	return crcdo
}

// Exec executes the deletion query.
func (crcdo *CycleRedeemCodeDeleteOne) Exec(ctx context.Context) error {
	n, err := crcdo.crcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cycleredeemcode.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (crcdo *CycleRedeemCodeDeleteOne) ExecX(ctx context.Context) {
	if err := crcdo.Exec(ctx); err != nil {
		panic(err)
	}
}