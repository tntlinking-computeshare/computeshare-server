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
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/userresourcelimit"
)

// UserResourceLimitQuery is the builder for querying UserResourceLimit entities.
type UserResourceLimitQuery struct {
	config
	ctx        *QueryContext
	order      []userresourcelimit.OrderOption
	inters     []Interceptor
	predicates []predicate.UserResourceLimit
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserResourceLimitQuery builder.
func (urlq *UserResourceLimitQuery) Where(ps ...predicate.UserResourceLimit) *UserResourceLimitQuery {
	urlq.predicates = append(urlq.predicates, ps...)
	return urlq
}

// Limit the number of records to be returned by this query.
func (urlq *UserResourceLimitQuery) Limit(limit int) *UserResourceLimitQuery {
	urlq.ctx.Limit = &limit
	return urlq
}

// Offset to start from.
func (urlq *UserResourceLimitQuery) Offset(offset int) *UserResourceLimitQuery {
	urlq.ctx.Offset = &offset
	return urlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (urlq *UserResourceLimitQuery) Unique(unique bool) *UserResourceLimitQuery {
	urlq.ctx.Unique = &unique
	return urlq
}

// Order specifies how the records should be ordered.
func (urlq *UserResourceLimitQuery) Order(o ...userresourcelimit.OrderOption) *UserResourceLimitQuery {
	urlq.order = append(urlq.order, o...)
	return urlq
}

// First returns the first UserResourceLimit entity from the query.
// Returns a *NotFoundError when no UserResourceLimit was found.
func (urlq *UserResourceLimitQuery) First(ctx context.Context) (*UserResourceLimit, error) {
	nodes, err := urlq.Limit(1).All(setContextOp(ctx, urlq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userresourcelimit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (urlq *UserResourceLimitQuery) FirstX(ctx context.Context) *UserResourceLimit {
	node, err := urlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserResourceLimit ID from the query.
// Returns a *NotFoundError when no UserResourceLimit ID was found.
func (urlq *UserResourceLimitQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = urlq.Limit(1).IDs(setContextOp(ctx, urlq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userresourcelimit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (urlq *UserResourceLimitQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := urlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserResourceLimit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserResourceLimit entity is found.
// Returns a *NotFoundError when no UserResourceLimit entities are found.
func (urlq *UserResourceLimitQuery) Only(ctx context.Context) (*UserResourceLimit, error) {
	nodes, err := urlq.Limit(2).All(setContextOp(ctx, urlq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userresourcelimit.Label}
	default:
		return nil, &NotSingularError{userresourcelimit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (urlq *UserResourceLimitQuery) OnlyX(ctx context.Context) *UserResourceLimit {
	node, err := urlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserResourceLimit ID in the query.
// Returns a *NotSingularError when more than one UserResourceLimit ID is found.
// Returns a *NotFoundError when no entities are found.
func (urlq *UserResourceLimitQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = urlq.Limit(2).IDs(setContextOp(ctx, urlq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userresourcelimit.Label}
	default:
		err = &NotSingularError{userresourcelimit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (urlq *UserResourceLimitQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := urlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserResourceLimits.
func (urlq *UserResourceLimitQuery) All(ctx context.Context) ([]*UserResourceLimit, error) {
	ctx = setContextOp(ctx, urlq.ctx, "All")
	if err := urlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserResourceLimit, *UserResourceLimitQuery]()
	return withInterceptors[[]*UserResourceLimit](ctx, urlq, qr, urlq.inters)
}

// AllX is like All, but panics if an error occurs.
func (urlq *UserResourceLimitQuery) AllX(ctx context.Context) []*UserResourceLimit {
	nodes, err := urlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserResourceLimit IDs.
func (urlq *UserResourceLimitQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if urlq.ctx.Unique == nil && urlq.path != nil {
		urlq.Unique(true)
	}
	ctx = setContextOp(ctx, urlq.ctx, "IDs")
	if err = urlq.Select(userresourcelimit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (urlq *UserResourceLimitQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := urlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (urlq *UserResourceLimitQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, urlq.ctx, "Count")
	if err := urlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, urlq, querierCount[*UserResourceLimitQuery](), urlq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (urlq *UserResourceLimitQuery) CountX(ctx context.Context) int {
	count, err := urlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (urlq *UserResourceLimitQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, urlq.ctx, "Exist")
	switch _, err := urlq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (urlq *UserResourceLimitQuery) ExistX(ctx context.Context) bool {
	exist, err := urlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserResourceLimitQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (urlq *UserResourceLimitQuery) Clone() *UserResourceLimitQuery {
	if urlq == nil {
		return nil
	}
	return &UserResourceLimitQuery{
		config:     urlq.config,
		ctx:        urlq.ctx.Clone(),
		order:      append([]userresourcelimit.OrderOption{}, urlq.order...),
		inters:     append([]Interceptor{}, urlq.inters...),
		predicates: append([]predicate.UserResourceLimit{}, urlq.predicates...),
		// clone intermediate query.
		sql:  urlq.sql.Clone(),
		path: urlq.path,
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
//	client.UserResourceLimit.Query().
//		GroupBy(userresourcelimit.FieldFkUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (urlq *UserResourceLimitQuery) GroupBy(field string, fields ...string) *UserResourceLimitGroupBy {
	urlq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserResourceLimitGroupBy{build: urlq}
	grbuild.flds = &urlq.ctx.Fields
	grbuild.label = userresourcelimit.Label
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
//	client.UserResourceLimit.Query().
//		Select(userresourcelimit.FieldFkUserID).
//		Scan(ctx, &v)
func (urlq *UserResourceLimitQuery) Select(fields ...string) *UserResourceLimitSelect {
	urlq.ctx.Fields = append(urlq.ctx.Fields, fields...)
	sbuild := &UserResourceLimitSelect{UserResourceLimitQuery: urlq}
	sbuild.label = userresourcelimit.Label
	sbuild.flds, sbuild.scan = &urlq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserResourceLimitSelect configured with the given aggregations.
func (urlq *UserResourceLimitQuery) Aggregate(fns ...AggregateFunc) *UserResourceLimitSelect {
	return urlq.Select().Aggregate(fns...)
}

func (urlq *UserResourceLimitQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range urlq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, urlq); err != nil {
				return err
			}
		}
	}
	for _, f := range urlq.ctx.Fields {
		if !userresourcelimit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if urlq.path != nil {
		prev, err := urlq.path(ctx)
		if err != nil {
			return err
		}
		urlq.sql = prev
	}
	return nil
}

func (urlq *UserResourceLimitQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserResourceLimit, error) {
	var (
		nodes = []*UserResourceLimit{}
		_spec = urlq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserResourceLimit).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserResourceLimit{config: urlq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, urlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (urlq *UserResourceLimitQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := urlq.querySpec()
	_spec.Node.Columns = urlq.ctx.Fields
	if len(urlq.ctx.Fields) > 0 {
		_spec.Unique = urlq.ctx.Unique != nil && *urlq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, urlq.driver, _spec)
}

func (urlq *UserResourceLimitQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userresourcelimit.Table, userresourcelimit.Columns, sqlgraph.NewFieldSpec(userresourcelimit.FieldID, field.TypeUUID))
	_spec.From = urlq.sql
	if unique := urlq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if urlq.path != nil {
		_spec.Unique = true
	}
	if fields := urlq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userresourcelimit.FieldID)
		for i := range fields {
			if fields[i] != userresourcelimit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := urlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := urlq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := urlq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := urlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (urlq *UserResourceLimitQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(urlq.driver.Dialect())
	t1 := builder.Table(userresourcelimit.Table)
	columns := urlq.ctx.Fields
	if len(columns) == 0 {
		columns = userresourcelimit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if urlq.sql != nil {
		selector = urlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if urlq.ctx.Unique != nil && *urlq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range urlq.predicates {
		p(selector)
	}
	for _, p := range urlq.order {
		p(selector)
	}
	if offset := urlq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := urlq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserResourceLimitGroupBy is the group-by builder for UserResourceLimit entities.
type UserResourceLimitGroupBy struct {
	selector
	build *UserResourceLimitQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (urlgb *UserResourceLimitGroupBy) Aggregate(fns ...AggregateFunc) *UserResourceLimitGroupBy {
	urlgb.fns = append(urlgb.fns, fns...)
	return urlgb
}

// Scan applies the selector query and scans the result into the given value.
func (urlgb *UserResourceLimitGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, urlgb.build.ctx, "GroupBy")
	if err := urlgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserResourceLimitQuery, *UserResourceLimitGroupBy](ctx, urlgb.build, urlgb, urlgb.build.inters, v)
}

func (urlgb *UserResourceLimitGroupBy) sqlScan(ctx context.Context, root *UserResourceLimitQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(urlgb.fns))
	for _, fn := range urlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*urlgb.flds)+len(urlgb.fns))
		for _, f := range *urlgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*urlgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := urlgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserResourceLimitSelect is the builder for selecting fields of UserResourceLimit entities.
type UserResourceLimitSelect struct {
	*UserResourceLimitQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (urls *UserResourceLimitSelect) Aggregate(fns ...AggregateFunc) *UserResourceLimitSelect {
	urls.fns = append(urls.fns, fns...)
	return urls
}

// Scan applies the selector query and scans the result into the given value.
func (urls *UserResourceLimitSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, urls.ctx, "Select")
	if err := urls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserResourceLimitQuery, *UserResourceLimitSelect](ctx, urls.UserResourceLimitQuery, urls, urls.inters, v)
}

func (urls *UserResourceLimitSelect) sqlScan(ctx context.Context, root *UserResourceLimitQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(urls.fns))
	for _, fn := range urls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*urls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := urls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
