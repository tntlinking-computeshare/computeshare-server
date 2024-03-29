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
	"github.com/mohaijiang/computeshare-server/internal/data/ent/s3bucket"
)

// S3BucketQuery is the builder for querying S3Bucket entities.
type S3BucketQuery struct {
	config
	ctx        *QueryContext
	order      []s3bucket.OrderOption
	inters     []Interceptor
	predicates []predicate.S3Bucket
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the S3BucketQuery builder.
func (sq *S3BucketQuery) Where(ps ...predicate.S3Bucket) *S3BucketQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *S3BucketQuery) Limit(limit int) *S3BucketQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *S3BucketQuery) Offset(offset int) *S3BucketQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *S3BucketQuery) Unique(unique bool) *S3BucketQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *S3BucketQuery) Order(o ...s3bucket.OrderOption) *S3BucketQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// First returns the first S3Bucket entity from the query.
// Returns a *NotFoundError when no S3Bucket was found.
func (sq *S3BucketQuery) First(ctx context.Context) (*S3Bucket, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{s3bucket.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *S3BucketQuery) FirstX(ctx context.Context) *S3Bucket {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first S3Bucket ID from the query.
// Returns a *NotFoundError when no S3Bucket ID was found.
func (sq *S3BucketQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{s3bucket.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *S3BucketQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single S3Bucket entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one S3Bucket entity is found.
// Returns a *NotFoundError when no S3Bucket entities are found.
func (sq *S3BucketQuery) Only(ctx context.Context) (*S3Bucket, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{s3bucket.Label}
	default:
		return nil, &NotSingularError{s3bucket.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *S3BucketQuery) OnlyX(ctx context.Context) *S3Bucket {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only S3Bucket ID in the query.
// Returns a *NotSingularError when more than one S3Bucket ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *S3BucketQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{s3bucket.Label}
	default:
		err = &NotSingularError{s3bucket.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *S3BucketQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of S3Buckets.
func (sq *S3BucketQuery) All(ctx context.Context) ([]*S3Bucket, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*S3Bucket, *S3BucketQuery]()
	return withInterceptors[[]*S3Bucket](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *S3BucketQuery) AllX(ctx context.Context) []*S3Bucket {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of S3Bucket IDs.
func (sq *S3BucketQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(s3bucket.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *S3BucketQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *S3BucketQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*S3BucketQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *S3BucketQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *S3BucketQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *S3BucketQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the S3BucketQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *S3BucketQuery) Clone() *S3BucketQuery {
	if sq == nil {
		return nil
	}
	return &S3BucketQuery{
		config:     sq.config,
		ctx:        sq.ctx.Clone(),
		order:      append([]s3bucket.OrderOption{}, sq.order...),
		inters:     append([]Interceptor{}, sq.inters...),
		predicates: append([]predicate.S3Bucket{}, sq.predicates...),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
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
//	client.S3Bucket.Query().
//		GroupBy(s3bucket.FieldFkUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *S3BucketQuery) GroupBy(field string, fields ...string) *S3BucketGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &S3BucketGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = s3bucket.Label
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
//	client.S3Bucket.Query().
//		Select(s3bucket.FieldFkUserID).
//		Scan(ctx, &v)
func (sq *S3BucketQuery) Select(fields ...string) *S3BucketSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &S3BucketSelect{S3BucketQuery: sq}
	sbuild.label = s3bucket.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a S3BucketSelect configured with the given aggregations.
func (sq *S3BucketQuery) Aggregate(fns ...AggregateFunc) *S3BucketSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *S3BucketQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !s3bucket.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *S3BucketQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*S3Bucket, error) {
	var (
		nodes = []*S3Bucket{}
		_spec = sq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*S3Bucket).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &S3Bucket{config: sq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sq *S3BucketQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *S3BucketQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(s3bucket.Table, s3bucket.Columns, sqlgraph.NewFieldSpec(s3bucket.FieldID, field.TypeUUID))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, s3bucket.FieldID)
		for i := range fields {
			if fields[i] != s3bucket.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *S3BucketQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(s3bucket.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = s3bucket.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// S3BucketGroupBy is the group-by builder for S3Bucket entities.
type S3BucketGroupBy struct {
	selector
	build *S3BucketQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *S3BucketGroupBy) Aggregate(fns ...AggregateFunc) *S3BucketGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *S3BucketGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*S3BucketQuery, *S3BucketGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *S3BucketGroupBy) sqlScan(ctx context.Context, root *S3BucketQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// S3BucketSelect is the builder for selecting fields of S3Bucket entities.
type S3BucketSelect struct {
	*S3BucketQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *S3BucketSelect) Aggregate(fns ...AggregateFunc) *S3BucketSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *S3BucketSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*S3BucketQuery, *S3BucketSelect](ctx, ss.S3BucketQuery, ss, ss.inters, v)
}

func (ss *S3BucketSelect) sqlScan(ctx context.Context, root *S3BucketQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
