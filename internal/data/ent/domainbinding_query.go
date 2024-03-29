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
	"github.com/mohaijiang/computeshare-server/internal/data/ent/domainbinding"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// DomainBindingQuery is the builder for querying DomainBinding entities.
type DomainBindingQuery struct {
	config
	ctx        *QueryContext
	order      []domainbinding.OrderOption
	inters     []Interceptor
	predicates []predicate.DomainBinding
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DomainBindingQuery builder.
func (dbq *DomainBindingQuery) Where(ps ...predicate.DomainBinding) *DomainBindingQuery {
	dbq.predicates = append(dbq.predicates, ps...)
	return dbq
}

// Limit the number of records to be returned by this query.
func (dbq *DomainBindingQuery) Limit(limit int) *DomainBindingQuery {
	dbq.ctx.Limit = &limit
	return dbq
}

// Offset to start from.
func (dbq *DomainBindingQuery) Offset(offset int) *DomainBindingQuery {
	dbq.ctx.Offset = &offset
	return dbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dbq *DomainBindingQuery) Unique(unique bool) *DomainBindingQuery {
	dbq.ctx.Unique = &unique
	return dbq
}

// Order specifies how the records should be ordered.
func (dbq *DomainBindingQuery) Order(o ...domainbinding.OrderOption) *DomainBindingQuery {
	dbq.order = append(dbq.order, o...)
	return dbq
}

// First returns the first DomainBinding entity from the query.
// Returns a *NotFoundError when no DomainBinding was found.
func (dbq *DomainBindingQuery) First(ctx context.Context) (*DomainBinding, error) {
	nodes, err := dbq.Limit(1).All(setContextOp(ctx, dbq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{domainbinding.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dbq *DomainBindingQuery) FirstX(ctx context.Context) *DomainBinding {
	node, err := dbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DomainBinding ID from the query.
// Returns a *NotFoundError when no DomainBinding ID was found.
func (dbq *DomainBindingQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dbq.Limit(1).IDs(setContextOp(ctx, dbq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{domainbinding.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dbq *DomainBindingQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DomainBinding entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DomainBinding entity is found.
// Returns a *NotFoundError when no DomainBinding entities are found.
func (dbq *DomainBindingQuery) Only(ctx context.Context) (*DomainBinding, error) {
	nodes, err := dbq.Limit(2).All(setContextOp(ctx, dbq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{domainbinding.Label}
	default:
		return nil, &NotSingularError{domainbinding.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dbq *DomainBindingQuery) OnlyX(ctx context.Context) *DomainBinding {
	node, err := dbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DomainBinding ID in the query.
// Returns a *NotSingularError when more than one DomainBinding ID is found.
// Returns a *NotFoundError when no entities are found.
func (dbq *DomainBindingQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dbq.Limit(2).IDs(setContextOp(ctx, dbq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{domainbinding.Label}
	default:
		err = &NotSingularError{domainbinding.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dbq *DomainBindingQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DomainBindings.
func (dbq *DomainBindingQuery) All(ctx context.Context) ([]*DomainBinding, error) {
	ctx = setContextOp(ctx, dbq.ctx, "All")
	if err := dbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DomainBinding, *DomainBindingQuery]()
	return withInterceptors[[]*DomainBinding](ctx, dbq, qr, dbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dbq *DomainBindingQuery) AllX(ctx context.Context) []*DomainBinding {
	nodes, err := dbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DomainBinding IDs.
func (dbq *DomainBindingQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if dbq.ctx.Unique == nil && dbq.path != nil {
		dbq.Unique(true)
	}
	ctx = setContextOp(ctx, dbq.ctx, "IDs")
	if err = dbq.Select(domainbinding.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dbq *DomainBindingQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dbq *DomainBindingQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dbq.ctx, "Count")
	if err := dbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dbq, querierCount[*DomainBindingQuery](), dbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dbq *DomainBindingQuery) CountX(ctx context.Context) int {
	count, err := dbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dbq *DomainBindingQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dbq.ctx, "Exist")
	switch _, err := dbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dbq *DomainBindingQuery) ExistX(ctx context.Context) bool {
	exist, err := dbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DomainBindingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dbq *DomainBindingQuery) Clone() *DomainBindingQuery {
	if dbq == nil {
		return nil
	}
	return &DomainBindingQuery{
		config:     dbq.config,
		ctx:        dbq.ctx.Clone(),
		order:      append([]domainbinding.OrderOption{}, dbq.order...),
		inters:     append([]Interceptor{}, dbq.inters...),
		predicates: append([]predicate.DomainBinding{}, dbq.predicates...),
		// clone intermediate query.
		sql:  dbq.sql.Clone(),
		path: dbq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DomainBinding.Query().
//		GroupBy(domainbinding.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dbq *DomainBindingQuery) GroupBy(field string, fields ...string) *DomainBindingGroupBy {
	dbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DomainBindingGroupBy{build: dbq}
	grbuild.flds = &dbq.ctx.Fields
	grbuild.label = domainbinding.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//	}
//
//	client.DomainBinding.Query().
//		Select(domainbinding.FieldUserID).
//		Scan(ctx, &v)
func (dbq *DomainBindingQuery) Select(fields ...string) *DomainBindingSelect {
	dbq.ctx.Fields = append(dbq.ctx.Fields, fields...)
	sbuild := &DomainBindingSelect{DomainBindingQuery: dbq}
	sbuild.label = domainbinding.Label
	sbuild.flds, sbuild.scan = &dbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DomainBindingSelect configured with the given aggregations.
func (dbq *DomainBindingQuery) Aggregate(fns ...AggregateFunc) *DomainBindingSelect {
	return dbq.Select().Aggregate(fns...)
}

func (dbq *DomainBindingQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dbq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dbq); err != nil {
				return err
			}
		}
	}
	for _, f := range dbq.ctx.Fields {
		if !domainbinding.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dbq.path != nil {
		prev, err := dbq.path(ctx)
		if err != nil {
			return err
		}
		dbq.sql = prev
	}
	return nil
}

func (dbq *DomainBindingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DomainBinding, error) {
	var (
		nodes = []*DomainBinding{}
		_spec = dbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DomainBinding).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DomainBinding{config: dbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dbq *DomainBindingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dbq.querySpec()
	_spec.Node.Columns = dbq.ctx.Fields
	if len(dbq.ctx.Fields) > 0 {
		_spec.Unique = dbq.ctx.Unique != nil && *dbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dbq.driver, _spec)
}

func (dbq *DomainBindingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(domainbinding.Table, domainbinding.Columns, sqlgraph.NewFieldSpec(domainbinding.FieldID, field.TypeUUID))
	_spec.From = dbq.sql
	if unique := dbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dbq.path != nil {
		_spec.Unique = true
	}
	if fields := dbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, domainbinding.FieldID)
		for i := range fields {
			if fields[i] != domainbinding.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dbq *DomainBindingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dbq.driver.Dialect())
	t1 := builder.Table(domainbinding.Table)
	columns := dbq.ctx.Fields
	if len(columns) == 0 {
		columns = domainbinding.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dbq.sql != nil {
		selector = dbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dbq.ctx.Unique != nil && *dbq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dbq.predicates {
		p(selector)
	}
	for _, p := range dbq.order {
		p(selector)
	}
	if offset := dbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DomainBindingGroupBy is the group-by builder for DomainBinding entities.
type DomainBindingGroupBy struct {
	selector
	build *DomainBindingQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dbgb *DomainBindingGroupBy) Aggregate(fns ...AggregateFunc) *DomainBindingGroupBy {
	dbgb.fns = append(dbgb.fns, fns...)
	return dbgb
}

// Scan applies the selector query and scans the result into the given value.
func (dbgb *DomainBindingGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dbgb.build.ctx, "GroupBy")
	if err := dbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DomainBindingQuery, *DomainBindingGroupBy](ctx, dbgb.build, dbgb, dbgb.build.inters, v)
}

func (dbgb *DomainBindingGroupBy) sqlScan(ctx context.Context, root *DomainBindingQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dbgb.fns))
	for _, fn := range dbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dbgb.flds)+len(dbgb.fns))
		for _, f := range *dbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DomainBindingSelect is the builder for selecting fields of DomainBinding entities.
type DomainBindingSelect struct {
	*DomainBindingQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dbs *DomainBindingSelect) Aggregate(fns ...AggregateFunc) *DomainBindingSelect {
	dbs.fns = append(dbs.fns, fns...)
	return dbs
}

// Scan applies the selector query and scans the result into the given value.
func (dbs *DomainBindingSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dbs.ctx, "Select")
	if err := dbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DomainBindingQuery, *DomainBindingSelect](ctx, dbs.DomainBindingQuery, dbs, dbs.inters, v)
}

func (dbs *DomainBindingSelect) sqlScan(ctx context.Context, root *DomainBindingQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dbs.fns))
	for _, fn := range dbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
