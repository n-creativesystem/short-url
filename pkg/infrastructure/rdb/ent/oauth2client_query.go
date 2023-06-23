// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent/oauth2client"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent/predicate"
)

// OAuth2ClientQuery is the builder for querying OAuth2Client entities.
type OAuth2ClientQuery struct {
	config
	ctx        *QueryContext
	order      []oauth2client.OrderOption
	inters     []Interceptor
	predicates []predicate.OAuth2Client
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OAuth2ClientQuery builder.
func (oq *OAuth2ClientQuery) Where(ps ...predicate.OAuth2Client) *OAuth2ClientQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit the number of records to be returned by this query.
func (oq *OAuth2ClientQuery) Limit(limit int) *OAuth2ClientQuery {
	oq.ctx.Limit = &limit
	return oq
}

// Offset to start from.
func (oq *OAuth2ClientQuery) Offset(offset int) *OAuth2ClientQuery {
	oq.ctx.Offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *OAuth2ClientQuery) Unique(unique bool) *OAuth2ClientQuery {
	oq.ctx.Unique = &unique
	return oq
}

// Order specifies how the records should be ordered.
func (oq *OAuth2ClientQuery) Order(o ...oauth2client.OrderOption) *OAuth2ClientQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// First returns the first OAuth2Client entity from the query.
// Returns a *NotFoundError when no OAuth2Client was found.
func (oq *OAuth2ClientQuery) First(ctx context.Context) (*OAuth2Client, error) {
	nodes, err := oq.Limit(1).All(setContextOp(ctx, oq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{oauth2client.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OAuth2ClientQuery) FirstX(ctx context.Context) *OAuth2Client {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OAuth2Client ID from the query.
// Returns a *NotFoundError when no OAuth2Client ID was found.
func (oq *OAuth2ClientQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = oq.Limit(1).IDs(setContextOp(ctx, oq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{oauth2client.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *OAuth2ClientQuery) FirstIDX(ctx context.Context) string {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OAuth2Client entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OAuth2Client entity is found.
// Returns a *NotFoundError when no OAuth2Client entities are found.
func (oq *OAuth2ClientQuery) Only(ctx context.Context) (*OAuth2Client, error) {
	nodes, err := oq.Limit(2).All(setContextOp(ctx, oq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{oauth2client.Label}
	default:
		return nil, &NotSingularError{oauth2client.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OAuth2ClientQuery) OnlyX(ctx context.Context) *OAuth2Client {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OAuth2Client ID in the query.
// Returns a *NotSingularError when more than one OAuth2Client ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *OAuth2ClientQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = oq.Limit(2).IDs(setContextOp(ctx, oq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{oauth2client.Label}
	default:
		err = &NotSingularError{oauth2client.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OAuth2ClientQuery) OnlyIDX(ctx context.Context) string {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OAuth2Clients.
func (oq *OAuth2ClientQuery) All(ctx context.Context) ([]*OAuth2Client, error) {
	ctx = setContextOp(ctx, oq.ctx, "All")
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OAuth2Client, *OAuth2ClientQuery]()
	return withInterceptors[[]*OAuth2Client](ctx, oq, qr, oq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oq *OAuth2ClientQuery) AllX(ctx context.Context) []*OAuth2Client {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OAuth2Client IDs.
func (oq *OAuth2ClientQuery) IDs(ctx context.Context) (ids []string, err error) {
	if oq.ctx.Unique == nil && oq.path != nil {
		oq.Unique(true)
	}
	ctx = setContextOp(ctx, oq.ctx, "IDs")
	if err = oq.Select(oauth2client.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OAuth2ClientQuery) IDsX(ctx context.Context) []string {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OAuth2ClientQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oq.ctx, "Count")
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oq, querierCount[*OAuth2ClientQuery](), oq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OAuth2ClientQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OAuth2ClientQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oq.ctx, "Exist")
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OAuth2ClientQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OAuth2ClientQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OAuth2ClientQuery) Clone() *OAuth2ClientQuery {
	if oq == nil {
		return nil
	}
	return &OAuth2ClientQuery{
		config:     oq.config,
		ctx:        oq.ctx.Clone(),
		order:      append([]oauth2client.OrderOption{}, oq.order...),
		inters:     append([]Interceptor{}, oq.inters...),
		predicates: append([]predicate.OAuth2Client{}, oq.predicates...),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OAuth2Client.Query().
//		GroupBy(oauth2client.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oq *OAuth2ClientQuery) GroupBy(field string, fields ...string) *OAuth2ClientGroupBy {
	oq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OAuth2ClientGroupBy{build: oq}
	grbuild.flds = &oq.ctx.Fields
	grbuild.label = oauth2client.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.OAuth2Client.Query().
//		Select(oauth2client.FieldCreateTime).
//		Scan(ctx, &v)
func (oq *OAuth2ClientQuery) Select(fields ...string) *OAuth2ClientSelect {
	oq.ctx.Fields = append(oq.ctx.Fields, fields...)
	sbuild := &OAuth2ClientSelect{OAuth2ClientQuery: oq}
	sbuild.label = oauth2client.Label
	sbuild.flds, sbuild.scan = &oq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OAuth2ClientSelect configured with the given aggregations.
func (oq *OAuth2ClientQuery) Aggregate(fns ...AggregateFunc) *OAuth2ClientSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *OAuth2ClientQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oq); err != nil {
				return err
			}
		}
	}
	for _, f := range oq.ctx.Fields {
		if !oauth2client.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *OAuth2ClientQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OAuth2Client, error) {
	var (
		nodes = []*OAuth2Client{}
		_spec = oq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OAuth2Client).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OAuth2Client{config: oq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (oq *OAuth2ClientQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	_spec.Node.Columns = oq.ctx.Fields
	if len(oq.ctx.Fields) > 0 {
		_spec.Unique = oq.ctx.Unique != nil && *oq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OAuth2ClientQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(oauth2client.Table, oauth2client.Columns, sqlgraph.NewFieldSpec(oauth2client.FieldID, field.TypeString))
	_spec.From = oq.sql
	if unique := oq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oq.path != nil {
		_spec.Unique = true
	}
	if fields := oq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauth2client.FieldID)
		for i := range fields {
			if fields[i] != oauth2client.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OAuth2ClientQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(oauth2client.Table)
	columns := oq.ctx.Fields
	if len(columns) == 0 {
		columns = oauth2client.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.ctx.Unique != nil && *oq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OAuth2ClientGroupBy is the group-by builder for OAuth2Client entities.
type OAuth2ClientGroupBy struct {
	selector
	build *OAuth2ClientQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OAuth2ClientGroupBy) Aggregate(fns ...AggregateFunc) *OAuth2ClientGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the selector query and scans the result into the given value.
func (ogb *OAuth2ClientGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ogb.build.ctx, "GroupBy")
	if err := ogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OAuth2ClientQuery, *OAuth2ClientGroupBy](ctx, ogb.build, ogb, ogb.build.inters, v)
}

func (ogb *OAuth2ClientGroupBy) sqlScan(ctx context.Context, root *OAuth2ClientQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ogb.flds)+len(ogb.fns))
		for _, f := range *ogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OAuth2ClientSelect is the builder for selecting fields of OAuth2Client entities.
type OAuth2ClientSelect struct {
	*OAuth2ClientQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *OAuth2ClientSelect) Aggregate(fns ...AggregateFunc) *OAuth2ClientSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *OAuth2ClientSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, os.ctx, "Select")
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OAuth2ClientQuery, *OAuth2ClientSelect](ctx, os.OAuth2ClientQuery, os, os.inters, v)
}

func (os *OAuth2ClientSelect) sqlScan(ctx context.Context, root *OAuth2ClientQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
