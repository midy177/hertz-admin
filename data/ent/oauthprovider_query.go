// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hertz-admin/data/ent/oauthprovider"
	"hertz-admin/data/ent/predicate"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OauthProviderQuery is the builder for querying OauthProvider entities.
type OauthProviderQuery struct {
	config
	ctx        *QueryContext
	order      []oauthprovider.OrderOption
	inters     []Interceptor
	predicates []predicate.OauthProvider
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OauthProviderQuery builder.
func (opq *OauthProviderQuery) Where(ps ...predicate.OauthProvider) *OauthProviderQuery {
	opq.predicates = append(opq.predicates, ps...)
	return opq
}

// Limit the number of records to be returned by this query.
func (opq *OauthProviderQuery) Limit(limit int) *OauthProviderQuery {
	opq.ctx.Limit = &limit
	return opq
}

// Offset to start from.
func (opq *OauthProviderQuery) Offset(offset int) *OauthProviderQuery {
	opq.ctx.Offset = &offset
	return opq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (opq *OauthProviderQuery) Unique(unique bool) *OauthProviderQuery {
	opq.ctx.Unique = &unique
	return opq
}

// Order specifies how the records should be ordered.
func (opq *OauthProviderQuery) Order(o ...oauthprovider.OrderOption) *OauthProviderQuery {
	opq.order = append(opq.order, o...)
	return opq
}

// First returns the first OauthProvider entity from the query.
// Returns a *NotFoundError when no OauthProvider was found.
func (opq *OauthProviderQuery) First(ctx context.Context) (*OauthProvider, error) {
	nodes, err := opq.Limit(1).All(setContextOp(ctx, opq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{oauthprovider.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (opq *OauthProviderQuery) FirstX(ctx context.Context) *OauthProvider {
	node, err := opq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OauthProvider ID from the query.
// Returns a *NotFoundError when no OauthProvider ID was found.
func (opq *OauthProviderQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = opq.Limit(1).IDs(setContextOp(ctx, opq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{oauthprovider.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (opq *OauthProviderQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := opq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OauthProvider entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OauthProvider entity is found.
// Returns a *NotFoundError when no OauthProvider entities are found.
func (opq *OauthProviderQuery) Only(ctx context.Context) (*OauthProvider, error) {
	nodes, err := opq.Limit(2).All(setContextOp(ctx, opq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{oauthprovider.Label}
	default:
		return nil, &NotSingularError{oauthprovider.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (opq *OauthProviderQuery) OnlyX(ctx context.Context) *OauthProvider {
	node, err := opq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OauthProvider ID in the query.
// Returns a *NotSingularError when more than one OauthProvider ID is found.
// Returns a *NotFoundError when no entities are found.
func (opq *OauthProviderQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = opq.Limit(2).IDs(setContextOp(ctx, opq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{oauthprovider.Label}
	default:
		err = &NotSingularError{oauthprovider.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (opq *OauthProviderQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := opq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OauthProviders.
func (opq *OauthProviderQuery) All(ctx context.Context) ([]*OauthProvider, error) {
	ctx = setContextOp(ctx, opq.ctx, ent.OpQueryAll)
	if err := opq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OauthProvider, *OauthProviderQuery]()
	return withInterceptors[[]*OauthProvider](ctx, opq, qr, opq.inters)
}

// AllX is like All, but panics if an error occurs.
func (opq *OauthProviderQuery) AllX(ctx context.Context) []*OauthProvider {
	nodes, err := opq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OauthProvider IDs.
func (opq *OauthProviderQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if opq.ctx.Unique == nil && opq.path != nil {
		opq.Unique(true)
	}
	ctx = setContextOp(ctx, opq.ctx, ent.OpQueryIDs)
	if err = opq.Select(oauthprovider.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (opq *OauthProviderQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := opq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (opq *OauthProviderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, opq.ctx, ent.OpQueryCount)
	if err := opq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, opq, querierCount[*OauthProviderQuery](), opq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (opq *OauthProviderQuery) CountX(ctx context.Context) int {
	count, err := opq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (opq *OauthProviderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, opq.ctx, ent.OpQueryExist)
	switch _, err := opq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (opq *OauthProviderQuery) ExistX(ctx context.Context) bool {
	exist, err := opq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OauthProviderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (opq *OauthProviderQuery) Clone() *OauthProviderQuery {
	if opq == nil {
		return nil
	}
	return &OauthProviderQuery{
		config:     opq.config,
		ctx:        opq.ctx.Clone(),
		order:      append([]oauthprovider.OrderOption{}, opq.order...),
		inters:     append([]Interceptor{}, opq.inters...),
		predicates: append([]predicate.OauthProvider{}, opq.predicates...),
		// clone intermediate query.
		sql:  opq.sql.Clone(),
		path: opq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OauthProvider.Query().
//		GroupBy(oauthprovider.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (opq *OauthProviderQuery) GroupBy(field string, fields ...string) *OauthProviderGroupBy {
	opq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OauthProviderGroupBy{build: opq}
	grbuild.flds = &opq.ctx.Fields
	grbuild.label = oauthprovider.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.OauthProvider.Query().
//		Select(oauthprovider.FieldCreatedAt).
//		Scan(ctx, &v)
func (opq *OauthProviderQuery) Select(fields ...string) *OauthProviderSelect {
	opq.ctx.Fields = append(opq.ctx.Fields, fields...)
	sbuild := &OauthProviderSelect{OauthProviderQuery: opq}
	sbuild.label = oauthprovider.Label
	sbuild.flds, sbuild.scan = &opq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OauthProviderSelect configured with the given aggregations.
func (opq *OauthProviderQuery) Aggregate(fns ...AggregateFunc) *OauthProviderSelect {
	return opq.Select().Aggregate(fns...)
}

func (opq *OauthProviderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range opq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, opq); err != nil {
				return err
			}
		}
	}
	for _, f := range opq.ctx.Fields {
		if !oauthprovider.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if opq.path != nil {
		prev, err := opq.path(ctx)
		if err != nil {
			return err
		}
		opq.sql = prev
	}
	return nil
}

func (opq *OauthProviderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OauthProvider, error) {
	var (
		nodes = []*OauthProvider{}
		_spec = opq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OauthProvider).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OauthProvider{config: opq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, opq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (opq *OauthProviderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := opq.querySpec()
	_spec.Node.Columns = opq.ctx.Fields
	if len(opq.ctx.Fields) > 0 {
		_spec.Unique = opq.ctx.Unique != nil && *opq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, opq.driver, _spec)
}

func (opq *OauthProviderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(oauthprovider.Table, oauthprovider.Columns, sqlgraph.NewFieldSpec(oauthprovider.FieldID, field.TypeUint64))
	_spec.From = opq.sql
	if unique := opq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if opq.path != nil {
		_spec.Unique = true
	}
	if fields := opq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauthprovider.FieldID)
		for i := range fields {
			if fields[i] != oauthprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := opq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := opq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := opq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := opq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (opq *OauthProviderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(opq.driver.Dialect())
	t1 := builder.Table(oauthprovider.Table)
	columns := opq.ctx.Fields
	if len(columns) == 0 {
		columns = oauthprovider.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if opq.sql != nil {
		selector = opq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if opq.ctx.Unique != nil && *opq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range opq.predicates {
		p(selector)
	}
	for _, p := range opq.order {
		p(selector)
	}
	if offset := opq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := opq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OauthProviderGroupBy is the group-by builder for OauthProvider entities.
type OauthProviderGroupBy struct {
	selector
	build *OauthProviderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (opgb *OauthProviderGroupBy) Aggregate(fns ...AggregateFunc) *OauthProviderGroupBy {
	opgb.fns = append(opgb.fns, fns...)
	return opgb
}

// Scan applies the selector query and scans the result into the given value.
func (opgb *OauthProviderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, opgb.build.ctx, ent.OpQueryGroupBy)
	if err := opgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OauthProviderQuery, *OauthProviderGroupBy](ctx, opgb.build, opgb, opgb.build.inters, v)
}

func (opgb *OauthProviderGroupBy) sqlScan(ctx context.Context, root *OauthProviderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(opgb.fns))
	for _, fn := range opgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*opgb.flds)+len(opgb.fns))
		for _, f := range *opgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*opgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := opgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OauthProviderSelect is the builder for selecting fields of OauthProvider entities.
type OauthProviderSelect struct {
	*OauthProviderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ops *OauthProviderSelect) Aggregate(fns ...AggregateFunc) *OauthProviderSelect {
	ops.fns = append(ops.fns, fns...)
	return ops
}

// Scan applies the selector query and scans the result into the given value.
func (ops *OauthProviderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ops.ctx, ent.OpQuerySelect)
	if err := ops.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OauthProviderQuery, *OauthProviderSelect](ctx, ops.OauthProviderQuery, ops, ops.inters, v)
}

func (ops *OauthProviderSelect) sqlScan(ctx context.Context, root *OauthProviderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ops.fns))
	for _, fn := range ops.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ops.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ops.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
