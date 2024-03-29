// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cycleorder"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// CycleOrderQuery is the builder for querying CycleOrder entities.
type CycleOrderQuery struct {
	config
	ctx        *QueryContext
	order      []cycleorder.OrderOption
	inters     []Interceptor
	predicates []predicate.CycleOrder
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CycleOrderQuery builder.
func (coq *CycleOrderQuery) Where(ps ...predicate.CycleOrder) *CycleOrderQuery {
	coq.predicates = append(coq.predicates, ps...)
	return coq
}

// Limit the number of records to be returned by this query.
func (coq *CycleOrderQuery) Limit(limit int) *CycleOrderQuery {
	coq.ctx.Limit = &limit
	return coq
}

// Offset to start from.
func (coq *CycleOrderQuery) Offset(offset int) *CycleOrderQuery {
	coq.ctx.Offset = &offset
	return coq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (coq *CycleOrderQuery) Unique(unique bool) *CycleOrderQuery {
	coq.ctx.Unique = &unique
	return coq
}

// Order specifies how the records should be ordered.
func (coq *CycleOrderQuery) Order(o ...cycleorder.OrderOption) *CycleOrderQuery {
	coq.order = append(coq.order, o...)
	return coq
}

// First returns the first CycleOrder entity from the query.
// Returns a *NotFoundError when no CycleOrder was found.
func (coq *CycleOrderQuery) First(ctx context.Context) (*CycleOrder, error) {
	nodes, err := coq.Limit(1).All(setContextOp(ctx, coq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cycleorder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (coq *CycleOrderQuery) FirstX(ctx context.Context) *CycleOrder {
	node, err := coq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CycleOrder ID from the query.
// Returns a *NotFoundError when no CycleOrder ID was found.
func (coq *CycleOrderQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = coq.Limit(1).IDs(setContextOp(ctx, coq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cycleorder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (coq *CycleOrderQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := coq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CycleOrder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CycleOrder entity is found.
// Returns a *NotFoundError when no CycleOrder entities are found.
func (coq *CycleOrderQuery) Only(ctx context.Context) (*CycleOrder, error) {
	nodes, err := coq.Limit(2).All(setContextOp(ctx, coq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cycleorder.Label}
	default:
		return nil, &NotSingularError{cycleorder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (coq *CycleOrderQuery) OnlyX(ctx context.Context) *CycleOrder {
	node, err := coq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CycleOrder ID in the query.
// Returns a *NotSingularError when more than one CycleOrder ID is found.
// Returns a *NotFoundError when no entities are found.
func (coq *CycleOrderQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = coq.Limit(2).IDs(setContextOp(ctx, coq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cycleorder.Label}
	default:
		err = &NotSingularError{cycleorder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (coq *CycleOrderQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := coq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CycleOrders.
func (coq *CycleOrderQuery) All(ctx context.Context) ([]*CycleOrder, error) {
	ctx = setContextOp(ctx, coq.ctx, "All")
	if err := coq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CycleOrder, *CycleOrderQuery]()
	return withInterceptors[[]*CycleOrder](ctx, coq, qr, coq.inters)
}

// AllX is like All, but panics if an error occurs.
func (coq *CycleOrderQuery) AllX(ctx context.Context) []*CycleOrder {
	nodes, err := coq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CycleOrder IDs.
func (coq *CycleOrderQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if coq.ctx.Unique == nil && coq.path != nil {
		coq.Unique(true)
	}
	ctx = setContextOp(ctx, coq.ctx, "IDs")
	if err = coq.Select(cycleorder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (coq *CycleOrderQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := coq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (coq *CycleOrderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, coq.ctx, "Count")
	if err := coq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, coq, querierCount[*CycleOrderQuery](), coq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (coq *CycleOrderQuery) CountX(ctx context.Context) int {
	count, err := coq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (coq *CycleOrderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, coq.ctx, "Exist")
	switch _, err := coq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (coq *CycleOrderQuery) ExistX(ctx context.Context) bool {
	exist, err := coq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CycleOrderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (coq *CycleOrderQuery) Clone() *CycleOrderQuery {
	if coq == nil {
		return nil
	}
	return &CycleOrderQuery{
		config:     coq.config,
		ctx:        coq.ctx.Clone(),
		order:      append([]cycleorder.OrderOption{}, coq.order...),
		inters:     append([]Interceptor{}, coq.inters...),
		predicates: append([]predicate.CycleOrder{}, coq.predicates...),
		// clone intermediate query.
		sql:  coq.sql.Clone(),
		path: coq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FkUserID uuid.UUID `json:"fk_user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CycleOrder.Query().
//		GroupBy(cycleorder.FieldFkUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (coq *CycleOrderQuery) GroupBy(field string, fields ...string) *CycleOrderGroupBy {
	coq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CycleOrderGroupBy{build: coq}
	grbuild.flds = &coq.ctx.Fields
	grbuild.label = cycleorder.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FkUserID uuid.UUID `json:"fk_user_id,omitempty"`
//	}
//
//	client.CycleOrder.Query().
//		Select(cycleorder.FieldFkUserID).
//		Scan(ctx, &v)
func (coq *CycleOrderQuery) Select(fields ...string) *CycleOrderSelect {
	coq.ctx.Fields = append(coq.ctx.Fields, fields...)
	sbuild := &CycleOrderSelect{CycleOrderQuery: coq}
	sbuild.label = cycleorder.Label
	sbuild.flds, sbuild.scan = &coq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CycleOrderSelect configured with the given aggregations.
func (coq *CycleOrderQuery) Aggregate(fns ...AggregateFunc) *CycleOrderSelect {
	return coq.Select().Aggregate(fns...)
}

func (coq *CycleOrderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range coq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, coq); err != nil {
				return err
			}
		}
	}
	for _, f := range coq.ctx.Fields {
		if !cycleorder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if coq.path != nil {
		prev, err := coq.path(ctx)
		if err != nil {
			return err
		}
		coq.sql = prev
	}
	return nil
}

func (coq *CycleOrderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CycleOrder, error) {
	var (
		nodes = []*CycleOrder{}
		_spec = coq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CycleOrder).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CycleOrder{config: coq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, coq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (coq *CycleOrderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := coq.querySpec()
	_spec.Node.Columns = coq.ctx.Fields
	if len(coq.ctx.Fields) > 0 {
		_spec.Unique = coq.ctx.Unique != nil && *coq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, coq.driver, _spec)
}

func (coq *CycleOrderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(cycleorder.Table, cycleorder.Columns, sqlgraph.NewFieldSpec(cycleorder.FieldID, field.TypeUUID))
	_spec.From = coq.sql
	if unique := coq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if coq.path != nil {
		_spec.Unique = true
	}
	if fields := coq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cycleorder.FieldID)
		for i := range fields {
			if fields[i] != cycleorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := coq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := coq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := coq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := coq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (coq *CycleOrderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(coq.driver.Dialect())
	t1 := builder.Table(cycleorder.Table)
	columns := coq.ctx.Fields
	if len(columns) == 0 {
		columns = cycleorder.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if coq.sql != nil {
		selector = coq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if coq.ctx.Unique != nil && *coq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range coq.predicates {
		p(selector)
	}
	for _, p := range coq.order {
		p(selector)
	}
	if offset := coq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := coq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CycleOrderGroupBy is the group-by builder for CycleOrder entities.
type CycleOrderGroupBy struct {
	selector
	build *CycleOrderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cogb *CycleOrderGroupBy) Aggregate(fns ...AggregateFunc) *CycleOrderGroupBy {
	cogb.fns = append(cogb.fns, fns...)
	return cogb
}

// Scan applies the selector query and scans the result into the given value.
func (cogb *CycleOrderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cogb.build.ctx, "GroupBy")
	if err := cogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CycleOrderQuery, *CycleOrderGroupBy](ctx, cogb.build, cogb, cogb.build.inters, v)
}

func (cogb *CycleOrderGroupBy) sqlScan(ctx context.Context, root *CycleOrderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cogb.fns))
	for _, fn := range cogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cogb.flds)+len(cogb.fns))
		for _, f := range *cogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CycleOrderSelect is the builder for selecting fields of CycleOrder entities.
type CycleOrderSelect struct {
	*CycleOrderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cos *CycleOrderSelect) Aggregate(fns ...AggregateFunc) *CycleOrderSelect {
	cos.fns = append(cos.fns, fns...)
	return cos
}

// Scan applies the selector query and scans the result into the given value.
func (cos *CycleOrderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cos.ctx, "Select")
	if err := cos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CycleOrderQuery, *CycleOrderSelect](ctx, cos.CycleOrderQuery, cos, cos.inters, v)
}

func (cos *CycleOrderSelect) sqlScan(ctx context.Context, root *CycleOrderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cos.fns))
	for _, fn := range cos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
