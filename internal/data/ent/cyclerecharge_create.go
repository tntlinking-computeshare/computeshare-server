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
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cyclerecharge"
)

// CycleRechargeCreate is the builder for creating a CycleRecharge entity.
type CycleRechargeCreate struct {
	config
	mutation *CycleRechargeMutation
	hooks    []Hook
}

// SetFkUserID sets the "fk_user_id" field.
func (crc *CycleRechargeCreate) SetFkUserID(u uuid.UUID) *CycleRechargeCreate {
	crc.mutation.SetFkUserID(u)
	return crc
}

// SetOutTradeNo sets the "out_trade_no" field.
func (crc *CycleRechargeCreate) SetOutTradeNo(s string) *CycleRechargeCreate {
	crc.mutation.SetOutTradeNo(s)
	return crc
}

// SetAlipayTradeNo sets the "alipay_trade_no" field.
func (crc *CycleRechargeCreate) SetAlipayTradeNo(s string) *CycleRechargeCreate {
	crc.mutation.SetAlipayTradeNo(s)
	return crc
}

// SetRechargeChannel sets the "recharge_channel" field.
func (crc *CycleRechargeCreate) SetRechargeChannel(i int) *CycleRechargeCreate {
	crc.mutation.SetRechargeChannel(i)
	return crc
}

// SetRedeemCode sets the "redeem_code" field.
func (crc *CycleRechargeCreate) SetRedeemCode(s string) *CycleRechargeCreate {
	crc.mutation.SetRedeemCode(s)
	return crc
}

// SetState sets the "state" field.
func (crc *CycleRechargeCreate) SetState(s string) *CycleRechargeCreate {
	crc.mutation.SetState(s)
	return crc
}

// SetPayAmount sets the "pay_amount" field.
func (crc *CycleRechargeCreate) SetPayAmount(f float64) *CycleRechargeCreate {
	crc.mutation.SetPayAmount(f)
	return crc
}

// SetTotalAmount sets the "total_amount" field.
func (crc *CycleRechargeCreate) SetTotalAmount(f float64) *CycleRechargeCreate {
	crc.mutation.SetTotalAmount(f)
	return crc
}

// SetBuyCycle sets the "buy_cycle" field.
func (crc *CycleRechargeCreate) SetBuyCycle(f float64) *CycleRechargeCreate {
	crc.mutation.SetBuyCycle(f)
	return crc
}

// SetCreateTime sets the "create_time" field.
func (crc *CycleRechargeCreate) SetCreateTime(t time.Time) *CycleRechargeCreate {
	crc.mutation.SetCreateTime(t)
	return crc
}

// SetUpdateTime sets the "update_time" field.
func (crc *CycleRechargeCreate) SetUpdateTime(t time.Time) *CycleRechargeCreate {
	crc.mutation.SetUpdateTime(t)
	return crc
}

// SetID sets the "id" field.
func (crc *CycleRechargeCreate) SetID(u uuid.UUID) *CycleRechargeCreate {
	crc.mutation.SetID(u)
	return crc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (crc *CycleRechargeCreate) SetNillableID(u *uuid.UUID) *CycleRechargeCreate {
	if u != nil {
		crc.SetID(*u)
	}
	return crc
}

// Mutation returns the CycleRechargeMutation object of the builder.
func (crc *CycleRechargeCreate) Mutation() *CycleRechargeMutation {
	return crc.mutation
}

// Save creates the CycleRecharge in the database.
func (crc *CycleRechargeCreate) Save(ctx context.Context) (*CycleRecharge, error) {
	crc.defaults()
	return withHooks(ctx, crc.sqlSave, crc.mutation, crc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (crc *CycleRechargeCreate) SaveX(ctx context.Context) *CycleRecharge {
	v, err := crc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (crc *CycleRechargeCreate) Exec(ctx context.Context) error {
	_, err := crc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crc *CycleRechargeCreate) ExecX(ctx context.Context) {
	if err := crc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (crc *CycleRechargeCreate) defaults() {
	if _, ok := crc.mutation.ID(); !ok {
		v := cyclerecharge.DefaultID()
		crc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (crc *CycleRechargeCreate) check() error {
	if _, ok := crc.mutation.FkUserID(); !ok {
		return &ValidationError{Name: "fk_user_id", err: errors.New(`ent: missing required field "CycleRecharge.fk_user_id"`)}
	}
	if _, ok := crc.mutation.OutTradeNo(); !ok {
		return &ValidationError{Name: "out_trade_no", err: errors.New(`ent: missing required field "CycleRecharge.out_trade_no"`)}
	}
	if _, ok := crc.mutation.AlipayTradeNo(); !ok {
		return &ValidationError{Name: "alipay_trade_no", err: errors.New(`ent: missing required field "CycleRecharge.alipay_trade_no"`)}
	}
	if _, ok := crc.mutation.RechargeChannel(); !ok {
		return &ValidationError{Name: "recharge_channel", err: errors.New(`ent: missing required field "CycleRecharge.recharge_channel"`)}
	}
	if _, ok := crc.mutation.RedeemCode(); !ok {
		return &ValidationError{Name: "redeem_code", err: errors.New(`ent: missing required field "CycleRecharge.redeem_code"`)}
	}
	if _, ok := crc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "CycleRecharge.state"`)}
	}
	if _, ok := crc.mutation.PayAmount(); !ok {
		return &ValidationError{Name: "pay_amount", err: errors.New(`ent: missing required field "CycleRecharge.pay_amount"`)}
	}
	if _, ok := crc.mutation.TotalAmount(); !ok {
		return &ValidationError{Name: "total_amount", err: errors.New(`ent: missing required field "CycleRecharge.total_amount"`)}
	}
	if _, ok := crc.mutation.BuyCycle(); !ok {
		return &ValidationError{Name: "buy_cycle", err: errors.New(`ent: missing required field "CycleRecharge.buy_cycle"`)}
	}
	if _, ok := crc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "CycleRecharge.create_time"`)}
	}
	if _, ok := crc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "CycleRecharge.update_time"`)}
	}
	return nil
}

func (crc *CycleRechargeCreate) sqlSave(ctx context.Context) (*CycleRecharge, error) {
	if err := crc.check(); err != nil {
		return nil, err
	}
	_node, _spec := crc.createSpec()
	if err := sqlgraph.CreateNode(ctx, crc.driver, _spec); err != nil {
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
	crc.mutation.id = &_node.ID
	crc.mutation.done = true
	return _node, nil
}

func (crc *CycleRechargeCreate) createSpec() (*CycleRecharge, *sqlgraph.CreateSpec) {
	var (
		_node = &CycleRecharge{config: crc.config}
		_spec = sqlgraph.NewCreateSpec(cyclerecharge.Table, sqlgraph.NewFieldSpec(cyclerecharge.FieldID, field.TypeUUID))
	)
	if id, ok := crc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := crc.mutation.FkUserID(); ok {
		_spec.SetField(cyclerecharge.FieldFkUserID, field.TypeUUID, value)
		_node.FkUserID = value
	}
	if value, ok := crc.mutation.OutTradeNo(); ok {
		_spec.SetField(cyclerecharge.FieldOutTradeNo, field.TypeString, value)
		_node.OutTradeNo = value
	}
	if value, ok := crc.mutation.AlipayTradeNo(); ok {
		_spec.SetField(cyclerecharge.FieldAlipayTradeNo, field.TypeString, value)
		_node.AlipayTradeNo = value
	}
	if value, ok := crc.mutation.RechargeChannel(); ok {
		_spec.SetField(cyclerecharge.FieldRechargeChannel, field.TypeInt, value)
		_node.RechargeChannel = value
	}
	if value, ok := crc.mutation.RedeemCode(); ok {
		_spec.SetField(cyclerecharge.FieldRedeemCode, field.TypeString, value)
		_node.RedeemCode = value
	}
	if value, ok := crc.mutation.State(); ok {
		_spec.SetField(cyclerecharge.FieldState, field.TypeString, value)
		_node.State = value
	}
	if value, ok := crc.mutation.PayAmount(); ok {
		_spec.SetField(cyclerecharge.FieldPayAmount, field.TypeFloat64, value)
		_node.PayAmount = value
	}
	if value, ok := crc.mutation.TotalAmount(); ok {
		_spec.SetField(cyclerecharge.FieldTotalAmount, field.TypeFloat64, value)
		_node.TotalAmount = value
	}
	if value, ok := crc.mutation.BuyCycle(); ok {
		_spec.SetField(cyclerecharge.FieldBuyCycle, field.TypeFloat64, value)
		_node.BuyCycle = value
	}
	if value, ok := crc.mutation.CreateTime(); ok {
		_spec.SetField(cyclerecharge.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := crc.mutation.UpdateTime(); ok {
		_spec.SetField(cyclerecharge.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// CycleRechargeCreateBulk is the builder for creating many CycleRecharge entities in bulk.
type CycleRechargeCreateBulk struct {
	config
	builders []*CycleRechargeCreate
}

// Save creates the CycleRecharge entities in the database.
func (crcb *CycleRechargeCreateBulk) Save(ctx context.Context) ([]*CycleRecharge, error) {
	specs := make([]*sqlgraph.CreateSpec, len(crcb.builders))
	nodes := make([]*CycleRecharge, len(crcb.builders))
	mutators := make([]Mutator, len(crcb.builders))
	for i := range crcb.builders {
		func(i int, root context.Context) {
			builder := crcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CycleRechargeMutation)
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
					_, err = mutators[i+1].Mutate(root, crcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, crcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, crcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (crcb *CycleRechargeCreateBulk) SaveX(ctx context.Context) []*CycleRecharge {
	v, err := crcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (crcb *CycleRechargeCreateBulk) Exec(ctx context.Context) error {
	_, err := crcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crcb *CycleRechargeCreateBulk) ExecX(ctx context.Context) {
	if err := crcb.Exec(ctx); err != nil {
		panic(err)
	}
}
