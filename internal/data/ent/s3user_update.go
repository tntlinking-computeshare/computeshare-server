// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/s3user"
)

// S3UserUpdate is the builder for updating S3User entities.
type S3UserUpdate struct {
	config
	hooks    []Hook
	mutation *S3UserMutation
}

// Where appends a list predicates to the S3UserUpdate builder.
func (su *S3UserUpdate) Where(ps ...predicate.S3User) *S3UserUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetFkUserID sets the "fk_user_id" field.
func (su *S3UserUpdate) SetFkUserID(u uuid.UUID) *S3UserUpdate {
	su.mutation.SetFkUserID(u)
	return su
}

// SetType sets the "type" field.
func (su *S3UserUpdate) SetType(i int8) *S3UserUpdate {
	su.mutation.ResetType()
	su.mutation.SetType(i)
	return su
}

// AddType adds i to the "type" field.
func (su *S3UserUpdate) AddType(i int8) *S3UserUpdate {
	su.mutation.AddType(i)
	return su
}

// SetAccessKey sets the "access_key" field.
func (su *S3UserUpdate) SetAccessKey(s string) *S3UserUpdate {
	su.mutation.SetAccessKey(s)
	return su
}

// SetSecretKey sets the "secret_key" field.
func (su *S3UserUpdate) SetSecretKey(s string) *S3UserUpdate {
	su.mutation.SetSecretKey(s)
	return su
}

// SetCreateTime sets the "create_time" field.
func (su *S3UserUpdate) SetCreateTime(t time.Time) *S3UserUpdate {
	su.mutation.SetCreateTime(t)
	return su
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (su *S3UserUpdate) SetNillableCreateTime(t *time.Time) *S3UserUpdate {
	if t != nil {
		su.SetCreateTime(*t)
	}
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *S3UserUpdate) SetUpdateTime(t time.Time) *S3UserUpdate {
	su.mutation.SetUpdateTime(t)
	return su
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (su *S3UserUpdate) SetNillableUpdateTime(t *time.Time) *S3UserUpdate {
	if t != nil {
		su.SetUpdateTime(*t)
	}
	return su
}

// Mutation returns the S3UserMutation object of the builder.
func (su *S3UserUpdate) Mutation() *S3UserMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *S3UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *S3UserUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *S3UserUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *S3UserUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *S3UserUpdate) check() error {
	if v, ok := su.mutation.AccessKey(); ok {
		if err := s3user.AccessKeyValidator(v); err != nil {
			return &ValidationError{Name: "access_key", err: fmt.Errorf(`ent: validator failed for field "S3User.access_key": %w`, err)}
		}
	}
	if v, ok := su.mutation.SecretKey(); ok {
		if err := s3user.SecretKeyValidator(v); err != nil {
			return &ValidationError{Name: "secret_key", err: fmt.Errorf(`ent: validator failed for field "S3User.secret_key": %w`, err)}
		}
	}
	return nil
}

func (su *S3UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(s3user.Table, s3user.Columns, sqlgraph.NewFieldSpec(s3user.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.FkUserID(); ok {
		_spec.SetField(s3user.FieldFkUserID, field.TypeUUID, value)
	}
	if value, ok := su.mutation.GetType(); ok {
		_spec.SetField(s3user.FieldType, field.TypeInt8, value)
	}
	if value, ok := su.mutation.AddedType(); ok {
		_spec.AddField(s3user.FieldType, field.TypeInt8, value)
	}
	if value, ok := su.mutation.AccessKey(); ok {
		_spec.SetField(s3user.FieldAccessKey, field.TypeString, value)
	}
	if value, ok := su.mutation.SecretKey(); ok {
		_spec.SetField(s3user.FieldSecretKey, field.TypeString, value)
	}
	if value, ok := su.mutation.CreateTime(); ok {
		_spec.SetField(s3user.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.SetField(s3user.FieldUpdateTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{s3user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// S3UserUpdateOne is the builder for updating a single S3User entity.
type S3UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *S3UserMutation
}

// SetFkUserID sets the "fk_user_id" field.
func (suo *S3UserUpdateOne) SetFkUserID(u uuid.UUID) *S3UserUpdateOne {
	suo.mutation.SetFkUserID(u)
	return suo
}

// SetType sets the "type" field.
func (suo *S3UserUpdateOne) SetType(i int8) *S3UserUpdateOne {
	suo.mutation.ResetType()
	suo.mutation.SetType(i)
	return suo
}

// AddType adds i to the "type" field.
func (suo *S3UserUpdateOne) AddType(i int8) *S3UserUpdateOne {
	suo.mutation.AddType(i)
	return suo
}

// SetAccessKey sets the "access_key" field.
func (suo *S3UserUpdateOne) SetAccessKey(s string) *S3UserUpdateOne {
	suo.mutation.SetAccessKey(s)
	return suo
}

// SetSecretKey sets the "secret_key" field.
func (suo *S3UserUpdateOne) SetSecretKey(s string) *S3UserUpdateOne {
	suo.mutation.SetSecretKey(s)
	return suo
}

// SetCreateTime sets the "create_time" field.
func (suo *S3UserUpdateOne) SetCreateTime(t time.Time) *S3UserUpdateOne {
	suo.mutation.SetCreateTime(t)
	return suo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (suo *S3UserUpdateOne) SetNillableCreateTime(t *time.Time) *S3UserUpdateOne {
	if t != nil {
		suo.SetCreateTime(*t)
	}
	return suo
}

// SetUpdateTime sets the "update_time" field.
func (suo *S3UserUpdateOne) SetUpdateTime(t time.Time) *S3UserUpdateOne {
	suo.mutation.SetUpdateTime(t)
	return suo
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (suo *S3UserUpdateOne) SetNillableUpdateTime(t *time.Time) *S3UserUpdateOne {
	if t != nil {
		suo.SetUpdateTime(*t)
	}
	return suo
}

// Mutation returns the S3UserMutation object of the builder.
func (suo *S3UserUpdateOne) Mutation() *S3UserMutation {
	return suo.mutation
}

// Where appends a list predicates to the S3UserUpdate builder.
func (suo *S3UserUpdateOne) Where(ps ...predicate.S3User) *S3UserUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *S3UserUpdateOne) Select(field string, fields ...string) *S3UserUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated S3User entity.
func (suo *S3UserUpdateOne) Save(ctx context.Context) (*S3User, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *S3UserUpdateOne) SaveX(ctx context.Context) *S3User {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *S3UserUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *S3UserUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *S3UserUpdateOne) check() error {
	if v, ok := suo.mutation.AccessKey(); ok {
		if err := s3user.AccessKeyValidator(v); err != nil {
			return &ValidationError{Name: "access_key", err: fmt.Errorf(`ent: validator failed for field "S3User.access_key": %w`, err)}
		}
	}
	if v, ok := suo.mutation.SecretKey(); ok {
		if err := s3user.SecretKeyValidator(v); err != nil {
			return &ValidationError{Name: "secret_key", err: fmt.Errorf(`ent: validator failed for field "S3User.secret_key": %w`, err)}
		}
	}
	return nil
}

func (suo *S3UserUpdateOne) sqlSave(ctx context.Context) (_node *S3User, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(s3user.Table, s3user.Columns, sqlgraph.NewFieldSpec(s3user.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "S3User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, s3user.FieldID)
		for _, f := range fields {
			if !s3user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != s3user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.FkUserID(); ok {
		_spec.SetField(s3user.FieldFkUserID, field.TypeUUID, value)
	}
	if value, ok := suo.mutation.GetType(); ok {
		_spec.SetField(s3user.FieldType, field.TypeInt8, value)
	}
	if value, ok := suo.mutation.AddedType(); ok {
		_spec.AddField(s3user.FieldType, field.TypeInt8, value)
	}
	if value, ok := suo.mutation.AccessKey(); ok {
		_spec.SetField(s3user.FieldAccessKey, field.TypeString, value)
	}
	if value, ok := suo.mutation.SecretKey(); ok {
		_spec.SetField(s3user.FieldSecretKey, field.TypeString, value)
	}
	if value, ok := suo.mutation.CreateTime(); ok {
		_spec.SetField(s3user.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.SetField(s3user.FieldUpdateTime, field.TypeTime, value)
	}
	_node = &S3User{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{s3user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
