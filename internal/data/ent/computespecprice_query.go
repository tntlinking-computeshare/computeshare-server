// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/computespecprice"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// ComputeSpecPriceQuery is the builder for querying ComputeSpecPrice entities.
type ComputeSpecPriceQuery struct {
	config
	ctx        *QueryContext
	order      []computespecprice.OrderOption
	inters     []Interceptor
	predicates []predicate.ComputeSpecPrice
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ComputeSpecPriceQuery builder.
func (cspq *ComputeSpecPriceQuery) Where(ps ...predicate.ComputeSpecPrice) *ComputeSpecPriceQuery {
	cspq.predicates = append(cspq.predicates, ps...)
	return cspq
}

// Limit the number of records to be returned by this query.
func (cspq *ComputeSpecPriceQuery) Limit(limit int) *ComputeSpecPriceQuery {
	cspq.ctx.Limit = &limit
	return cspq
}

// Offset to start from.
func (cspq *ComputeSpecPriceQuery) Offset(offset int) *ComputeSpecPriceQuery {
	cspq.ctx.Offset = &offset
	return cspq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cspq *ComputeSpecPriceQuery) Unique(unique bool) *ComputeSpecPriceQuery {
	cspq.ctx.Unique = &unique
	return cspq
}

// Order specifies how the records should be ordered.
func (cspq *ComputeSpecPriceQuery) Order(o ...computespecprice.OrderOption) *ComputeSpecPriceQuery {
	cspq.order = append(cspq.order, o...)
	return cspq
}

// First returns the first ComputeSpecPrice entity from the query.
// Returns a *NotFoundError when no ComputeSpecPrice was found.
func (cspq *ComputeSpecPriceQuery) First(ctx context.Context) (*ComputeSpecPrice, error) {
	nodes, err := cspq.Limit(1).All(setContextOp(ctx, cspq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{computespecprice.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cspq *ComputeSpecPriceQuery) FirstX(ctx context.Context) *ComputeSpecPrice {
	node, err := cspq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ComputeSpecPrice ID from the query.
// Returns a *NotFoundError when no ComputeSpecPrice ID was found.
func (cspq *ComputeSpecPriceQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = cspq.Limit(1).IDs(setContextOp(ctx, cspq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{computespecprice.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cspq *ComputeSpecPriceQuery) FirstIDX(ctx context.Context) int32 {
	id, err := cspq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ComputeSpecPrice entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ComputeSpecPrice entity is found.
// Returns a *NotFoundError when no ComputeSpecPrice entities are found.
func (cspq *ComputeSpecPriceQuery) Only(ctx context.Context) (*ComputeSpecPrice, error) {
	nodes, err := cspq.Limit(2).All(setContextOp(ctx, cspq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{computespecprice.Label}
	default:
		return nil, &NotSingularError{computespecprice.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cspq *ComputeSpecPriceQuery) OnlyX(ctx context.Context) *ComputeSpecPrice {
	node, err := cspq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ComputeSpecPrice ID in the query.
// Returns a *NotSingularError when more than one ComputeSpecPrice ID is found.
// Returns a *NotFoundError when no entities are found.
func (cspq *ComputeSpecPriceQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = cspq.Limit(2).IDs(setContextOp(ctx, cspq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{computespecprice.Label}
	default:
		err = &NotSingularError{computespecprice.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cspq *ComputeSpecPriceQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := cspq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ComputeSpecPrices.
func (cspq *ComputeSpecPriceQuery) All(ctx context.Context) ([]*ComputeSpecPrice, error) {
	ctx = setContextOp(ctx, cspq.ctx, "All")
	if err := cspq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ComputeSpecPrice, *ComputeSpecPriceQuery]()
	return withInterceptors[[]*ComputeSpecPrice](ctx, cspq, qr, cspq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cspq *ComputeSpecPriceQuery) AllX(ctx context.Context) []*ComputeSpecPrice {
	nodes, err := cspq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ComputeSpecPrice IDs.
func (cspq *ComputeSpecPriceQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if cspq.ctx.Unique == nil && cspq.path != nil {
		cspq.Unique(true)
	}
	ctx = setContextOp(ctx, cspq.ctx, "IDs")
	if err = cspq.Select(computespecprice.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cspq *ComputeSpecPriceQuery) IDsX(ctx context.Context) []int32 {
	ids, err := cspq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cspq *ComputeSpecPriceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cspq.ctx, "Count")
	if err := cspq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cspq, querierCount[*ComputeSpecPriceQuery](), cspq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cspq *ComputeSpecPriceQuery) CountX(ctx context.Context) int {
	count, err := cspq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cspq *ComputeSpecPriceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cspq.ctx, "Exist")
	switch _, err := cspq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cspq *ComputeSpecPriceQuery) ExistX(ctx context.Context) bool {
	exist, err := cspq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ComputeSpecPriceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cspq *ComputeSpecPriceQuery) Clone() *ComputeSpecPriceQuery {
	if cspq == nil {
		return nil
	}
	return &ComputeSpecPriceQuery{
		config:     cspq.config,
		ctx:        cspq.ctx.Clone(),
		order:      append([]computespecprice.OrderOption{}, cspq.order...),
		inters:     append([]Interceptor{}, cspq.inters...),
		predicates: append([]predicate.ComputeSpecPrice{}, cspq.predicates...),
		// clone intermediate query.
		sql:  cspq.sql.Clone(),
		path: cspq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FkComputeSpecID int32 `json:"fk_compute_spec_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ComputeSpecPrice.Query().
//		GroupBy(computespecprice.FieldFkComputeSpecID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cspq *ComputeSpecPriceQuery) GroupBy(field string, fields ...string) *ComputeSpecPriceGroupBy {
	cspq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ComputeSpecPriceGroupBy{build: cspq}
	grbuild.flds = &cspq.ctx.Fields
	grbuild.label = computespecprice.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FkComputeSpecID int32 `json:"fk_compute_spec_id,omitempty"`
//	}
//
//	client.ComputeSpecPrice.Query().
//		Select(computespecprice.FieldFkComputeSpecID).
//		Scan(ctx, &v)
func (cspq *ComputeSpecPriceQuery) Select(fields ...string) *ComputeSpecPriceSelect {
	cspq.ctx.Fields = append(cspq.ctx.Fields, fields...)
	sbuild := &ComputeSpecPriceSelect{ComputeSpecPriceQuery: cspq}
	sbuild.label = computespecprice.Label
	sbuild.flds, sbuild.scan = &cspq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ComputeSpecPriceSelect configured with the given aggregations.
func (cspq *ComputeSpecPriceQuery) Aggregate(fns ...AggregateFunc) *ComputeSpecPriceSelect {
	return cspq.Select().Aggregate(fns...)
}

func (cspq *ComputeSpecPriceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cspq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cspq); err != nil {
				return err
			}
		}
	}
	for _, f := range cspq.ctx.Fields {
		if !computespecprice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cspq.path != nil {
		prev, err := cspq.path(ctx)
		if err != nil {
			return err
		}
		cspq.sql = prev
	}
	return nil
}

func (cspq *ComputeSpecPriceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ComputeSpecPrice, error) {
	var (
		nodes = []*ComputeSpecPrice{}
		_spec = cspq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ComputeSpecPrice).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ComputeSpecPrice{config: cspq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cspq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cspq *ComputeSpecPriceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cspq.querySpec()
	_spec.Node.Columns = cspq.ctx.Fields
	if len(cspq.ctx.Fields) > 0 {
		_spec.Unique = cspq.ctx.Unique != nil && *cspq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cspq.driver, _spec)
}

func (cspq *ComputeSpecPriceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(computespecprice.Table, computespecprice.Columns, sqlgraph.NewFieldSpec(computespecprice.FieldID, field.TypeInt32))
	_spec.From = cspq.sql
	if unique := cspq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cspq.path != nil {
		_spec.Unique = true
	}
	if fields := cspq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, computespecprice.FieldID)
		for i := range fields {
			if fields[i] != computespecprice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cspq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cspq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cspq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cspq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cspq *ComputeSpecPriceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cspq.driver.Dialect())
	t1 := builder.Table(computespecprice.Table)
	columns := cspq.ctx.Fields
	if len(columns) == 0 {
		columns = computespecprice.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cspq.sql != nil {
		selector = cspq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cspq.ctx.Unique != nil && *cspq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cspq.predicates {
		p(selector)
	}
	for _, p := range cspq.order {
		p(selector)
	}
	if offset := cspq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cspq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ComputeSpecPriceGroupBy is the group-by builder for ComputeSpecPrice entities.
type ComputeSpecPriceGroupBy struct {
	selector
	build *ComputeSpecPriceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cspgb *ComputeSpecPriceGroupBy) Aggregate(fns ...AggregateFunc) *ComputeSpecPriceGroupBy {
	cspgb.fns = append(cspgb.fns, fns...)
	return cspgb
}

// Scan applies the selector query and scans the result into the given value.
func (cspgb *ComputeSpecPriceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cspgb.build.ctx, "GroupBy")
	if err := cspgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ComputeSpecPriceQuery, *ComputeSpecPriceGroupBy](ctx, cspgb.build, cspgb, cspgb.build.inters, v)
}

func (cspgb *ComputeSpecPriceGroupBy) sqlScan(ctx context.Context, root *ComputeSpecPriceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cspgb.fns))
	for _, fn := range cspgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cspgb.flds)+len(cspgb.fns))
		for _, f := range *cspgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cspgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cspgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ComputeSpecPriceSelect is the builder for selecting fields of ComputeSpecPrice entities.
type ComputeSpecPriceSelect struct {
	*ComputeSpecPriceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (csps *ComputeSpecPriceSelect) Aggregate(fns ...AggregateFunc) *ComputeSpecPriceSelect {
	csps.fns = append(csps.fns, fns...)
	return csps
}

// Scan applies the selector query and scans the result into the given value.
func (csps *ComputeSpecPriceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, csps.ctx, "Select")
	if err := csps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ComputeSpecPriceQuery, *ComputeSpecPriceSelect](ctx, csps.ComputeSpecPriceQuery, csps, csps.inters, v)
}

func (csps *ComputeSpecPriceSelect) sqlScan(ctx context.Context, root *ComputeSpecPriceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(csps.fns))
	for _, fn := range csps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*csps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := csps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
