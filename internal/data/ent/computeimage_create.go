// Code generated by ent, DO NOT EDIT.

package ent

import (
	"computeshare-server/internal/data/ent/computeimage"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ComputeImageCreate is the builder for creating a ComputeImage entity.
type ComputeImageCreate struct {
	config
	mutation *ComputeImageMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cic *ComputeImageCreate) SetName(s string) *ComputeImageCreate {
	cic.mutation.SetName(s)
	return cic
}

// SetImage sets the "image" field.
func (cic *ComputeImageCreate) SetImage(s string) *ComputeImageCreate {
	cic.mutation.SetImage(s)
	return cic
}

// SetTag sets the "tag" field.
func (cic *ComputeImageCreate) SetTag(s string) *ComputeImageCreate {
	cic.mutation.SetTag(s)
	return cic
}

// SetPort sets the "port" field.
func (cic *ComputeImageCreate) SetPort(i int32) *ComputeImageCreate {
	cic.mutation.SetPort(i)
	return cic
}

// SetID sets the "id" field.
func (cic *ComputeImageCreate) SetID(i int32) *ComputeImageCreate {
	cic.mutation.SetID(i)
	return cic
}

// Mutation returns the ComputeImageMutation object of the builder.
func (cic *ComputeImageCreate) Mutation() *ComputeImageMutation {
	return cic.mutation
}

// Save creates the ComputeImage in the database.
func (cic *ComputeImageCreate) Save(ctx context.Context) (*ComputeImage, error) {
	return withHooks(ctx, cic.sqlSave, cic.mutation, cic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cic *ComputeImageCreate) SaveX(ctx context.Context) *ComputeImage {
	v, err := cic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cic *ComputeImageCreate) Exec(ctx context.Context) error {
	_, err := cic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cic *ComputeImageCreate) ExecX(ctx context.Context) {
	if err := cic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cic *ComputeImageCreate) check() error {
	if _, ok := cic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ComputeImage.name"`)}
	}
	if v, ok := cic.mutation.Name(); ok {
		if err := computeimage.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ComputeImage.name": %w`, err)}
		}
	}
	if _, ok := cic.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "ComputeImage.image"`)}
	}
	if v, ok := cic.mutation.Image(); ok {
		if err := computeimage.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "ComputeImage.image": %w`, err)}
		}
	}
	if _, ok := cic.mutation.Tag(); !ok {
		return &ValidationError{Name: "tag", err: errors.New(`ent: missing required field "ComputeImage.tag"`)}
	}
	if v, ok := cic.mutation.Tag(); ok {
		if err := computeimage.TagValidator(v); err != nil {
			return &ValidationError{Name: "tag", err: fmt.Errorf(`ent: validator failed for field "ComputeImage.tag": %w`, err)}
		}
	}
	if _, ok := cic.mutation.Port(); !ok {
		return &ValidationError{Name: "port", err: errors.New(`ent: missing required field "ComputeImage.port"`)}
	}
	return nil
}

func (cic *ComputeImageCreate) sqlSave(ctx context.Context) (*ComputeImage, error) {
	if err := cic.check(); err != nil {
		return nil, err
	}
	_node, _spec := cic.createSpec()
	if err := sqlgraph.CreateNode(ctx, cic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	cic.mutation.id = &_node.ID
	cic.mutation.done = true
	return _node, nil
}

func (cic *ComputeImageCreate) createSpec() (*ComputeImage, *sqlgraph.CreateSpec) {
	var (
		_node = &ComputeImage{config: cic.config}
		_spec = sqlgraph.NewCreateSpec(computeimage.Table, sqlgraph.NewFieldSpec(computeimage.FieldID, field.TypeInt32))
	)
	if id, ok := cic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cic.mutation.Name(); ok {
		_spec.SetField(computeimage.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cic.mutation.Image(); ok {
		_spec.SetField(computeimage.FieldImage, field.TypeString, value)
		_node.Image = value
	}
	if value, ok := cic.mutation.Tag(); ok {
		_spec.SetField(computeimage.FieldTag, field.TypeString, value)
		_node.Tag = value
	}
	if value, ok := cic.mutation.Port(); ok {
		_spec.SetField(computeimage.FieldPort, field.TypeInt32, value)
		_node.Port = value
	}
	return _node, _spec
}

// ComputeImageCreateBulk is the builder for creating many ComputeImage entities in bulk.
type ComputeImageCreateBulk struct {
	config
	builders []*ComputeImageCreate
}

// Save creates the ComputeImage entities in the database.
func (cicb *ComputeImageCreateBulk) Save(ctx context.Context) ([]*ComputeImage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cicb.builders))
	nodes := make([]*ComputeImage, len(cicb.builders))
	mutators := make([]Mutator, len(cicb.builders))
	for i := range cicb.builders {
		func(i int, root context.Context) {
			builder := cicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ComputeImageMutation)
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
					_, err = mutators[i+1].Mutate(root, cicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, cicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cicb *ComputeImageCreateBulk) SaveX(ctx context.Context) []*ComputeImage {
	v, err := cicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cicb *ComputeImageCreateBulk) Exec(ctx context.Context) error {
	_, err := cicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cicb *ComputeImageCreateBulk) ExecX(ctx context.Context) {
	if err := cicb.Exec(ctx); err != nil {
		panic(err)
	}
}
