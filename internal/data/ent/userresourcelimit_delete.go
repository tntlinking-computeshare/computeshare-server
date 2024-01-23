// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/userresourcelimit"
)

// UserResourceLimitDelete is the builder for deleting a UserResourceLimit entity.
type UserResourceLimitDelete struct {
	config
	hooks    []Hook
	mutation *UserResourceLimitMutation
}

// Where appends a list predicates to the UserResourceLimitDelete builder.
func (urld *UserResourceLimitDelete) Where(ps ...predicate.UserResourceLimit) *UserResourceLimitDelete {
	urld.mutation.Where(ps...)
	return urld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (urld *UserResourceLimitDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, urld.sqlExec, urld.mutation, urld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (urld *UserResourceLimitDelete) ExecX(ctx context.Context) int {
	n, err := urld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (urld *UserResourceLimitDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userresourcelimit.Table, sqlgraph.NewFieldSpec(userresourcelimit.FieldID, field.TypeUUID))
	if ps := urld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, urld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	urld.mutation.done = true
	return affected, err
}

// UserResourceLimitDeleteOne is the builder for deleting a single UserResourceLimit entity.
type UserResourceLimitDeleteOne struct {
	urld *UserResourceLimitDelete
}

// Where appends a list predicates to the UserResourceLimitDelete builder.
func (urldo *UserResourceLimitDeleteOne) Where(ps ...predicate.UserResourceLimit) *UserResourceLimitDeleteOne {
	urldo.urld.mutation.Where(ps...)
	return urldo
}

// Exec executes the deletion query.
func (urldo *UserResourceLimitDeleteOne) Exec(ctx context.Context) error {
	n, err := urldo.urld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userresourcelimit.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (urldo *UserResourceLimitDeleteOne) ExecX(ctx context.Context) {
	if err := urldo.Exec(ctx); err != nil {
		panic(err)
	}
}
