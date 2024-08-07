// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"siliconvali/ent/plan"
	"siliconvali/ent/predicate"
	"siliconvali/ent/user"
	"siliconvali/ent/userpaymentplan"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserPaymentPlanQuery is the builder for querying UserPaymentPlan entities.
type UserPaymentPlanQuery struct {
	config
	ctx        *QueryContext
	order      []userpaymentplan.OrderOption
	inters     []Interceptor
	predicates []predicate.UserPaymentPlan
	withUserID *UserQuery
	withPlanID *PlanQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserPaymentPlanQuery builder.
func (uppq *UserPaymentPlanQuery) Where(ps ...predicate.UserPaymentPlan) *UserPaymentPlanQuery {
	uppq.predicates = append(uppq.predicates, ps...)
	return uppq
}

// Limit the number of records to be returned by this query.
func (uppq *UserPaymentPlanQuery) Limit(limit int) *UserPaymentPlanQuery {
	uppq.ctx.Limit = &limit
	return uppq
}

// Offset to start from.
func (uppq *UserPaymentPlanQuery) Offset(offset int) *UserPaymentPlanQuery {
	uppq.ctx.Offset = &offset
	return uppq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uppq *UserPaymentPlanQuery) Unique(unique bool) *UserPaymentPlanQuery {
	uppq.ctx.Unique = &unique
	return uppq
}

// Order specifies how the records should be ordered.
func (uppq *UserPaymentPlanQuery) Order(o ...userpaymentplan.OrderOption) *UserPaymentPlanQuery {
	uppq.order = append(uppq.order, o...)
	return uppq
}

// QueryUserID chains the current query on the "user_id" edge.
func (uppq *UserPaymentPlanQuery) QueryUserID() *UserQuery {
	query := (&UserClient{config: uppq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uppq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userpaymentplan.Table, userpaymentplan.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userpaymentplan.UserIDTable, userpaymentplan.UserIDColumn),
		)
		fromU = sqlgraph.SetNeighbors(uppq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPlanID chains the current query on the "plan_id" edge.
func (uppq *UserPaymentPlanQuery) QueryPlanID() *PlanQuery {
	query := (&PlanClient{config: uppq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uppq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userpaymentplan.Table, userpaymentplan.FieldID, selector),
			sqlgraph.To(plan.Table, plan.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userpaymentplan.PlanIDTable, userpaymentplan.PlanIDColumn),
		)
		fromU = sqlgraph.SetNeighbors(uppq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserPaymentPlan entity from the query.
// Returns a *NotFoundError when no UserPaymentPlan was found.
func (uppq *UserPaymentPlanQuery) First(ctx context.Context) (*UserPaymentPlan, error) {
	nodes, err := uppq.Limit(1).All(setContextOp(ctx, uppq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userpaymentplan.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uppq *UserPaymentPlanQuery) FirstX(ctx context.Context) *UserPaymentPlan {
	node, err := uppq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserPaymentPlan ID from the query.
// Returns a *NotFoundError when no UserPaymentPlan ID was found.
func (uppq *UserPaymentPlanQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = uppq.Limit(1).IDs(setContextOp(ctx, uppq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userpaymentplan.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uppq *UserPaymentPlanQuery) FirstIDX(ctx context.Context) int64 {
	id, err := uppq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserPaymentPlan entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserPaymentPlan entity is found.
// Returns a *NotFoundError when no UserPaymentPlan entities are found.
func (uppq *UserPaymentPlanQuery) Only(ctx context.Context) (*UserPaymentPlan, error) {
	nodes, err := uppq.Limit(2).All(setContextOp(ctx, uppq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userpaymentplan.Label}
	default:
		return nil, &NotSingularError{userpaymentplan.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uppq *UserPaymentPlanQuery) OnlyX(ctx context.Context) *UserPaymentPlan {
	node, err := uppq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserPaymentPlan ID in the query.
// Returns a *NotSingularError when more than one UserPaymentPlan ID is found.
// Returns a *NotFoundError when no entities are found.
func (uppq *UserPaymentPlanQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = uppq.Limit(2).IDs(setContextOp(ctx, uppq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userpaymentplan.Label}
	default:
		err = &NotSingularError{userpaymentplan.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uppq *UserPaymentPlanQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := uppq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserPaymentPlans.
func (uppq *UserPaymentPlanQuery) All(ctx context.Context) ([]*UserPaymentPlan, error) {
	ctx = setContextOp(ctx, uppq.ctx, "All")
	if err := uppq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserPaymentPlan, *UserPaymentPlanQuery]()
	return withInterceptors[[]*UserPaymentPlan](ctx, uppq, qr, uppq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uppq *UserPaymentPlanQuery) AllX(ctx context.Context) []*UserPaymentPlan {
	nodes, err := uppq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserPaymentPlan IDs.
func (uppq *UserPaymentPlanQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if uppq.ctx.Unique == nil && uppq.path != nil {
		uppq.Unique(true)
	}
	ctx = setContextOp(ctx, uppq.ctx, "IDs")
	if err = uppq.Select(userpaymentplan.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uppq *UserPaymentPlanQuery) IDsX(ctx context.Context) []int64 {
	ids, err := uppq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uppq *UserPaymentPlanQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uppq.ctx, "Count")
	if err := uppq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uppq, querierCount[*UserPaymentPlanQuery](), uppq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uppq *UserPaymentPlanQuery) CountX(ctx context.Context) int {
	count, err := uppq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uppq *UserPaymentPlanQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uppq.ctx, "Exist")
	switch _, err := uppq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uppq *UserPaymentPlanQuery) ExistX(ctx context.Context) bool {
	exist, err := uppq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserPaymentPlanQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uppq *UserPaymentPlanQuery) Clone() *UserPaymentPlanQuery {
	if uppq == nil {
		return nil
	}
	return &UserPaymentPlanQuery{
		config:     uppq.config,
		ctx:        uppq.ctx.Clone(),
		order:      append([]userpaymentplan.OrderOption{}, uppq.order...),
		inters:     append([]Interceptor{}, uppq.inters...),
		predicates: append([]predicate.UserPaymentPlan{}, uppq.predicates...),
		withUserID: uppq.withUserID.Clone(),
		withPlanID: uppq.withPlanID.Clone(),
		// clone intermediate query.
		sql:  uppq.sql.Clone(),
		path: uppq.path,
	}
}

// WithUserID tells the query-builder to eager-load the nodes that are connected to
// the "user_id" edge. The optional arguments are used to configure the query builder of the edge.
func (uppq *UserPaymentPlanQuery) WithUserID(opts ...func(*UserQuery)) *UserPaymentPlanQuery {
	query := (&UserClient{config: uppq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uppq.withUserID = query
	return uppq
}

// WithPlanID tells the query-builder to eager-load the nodes that are connected to
// the "plan_id" edge. The optional arguments are used to configure the query builder of the edge.
func (uppq *UserPaymentPlanQuery) WithPlanID(opts ...func(*PlanQuery)) *UserPaymentPlanQuery {
	query := (&PlanClient{config: uppq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uppq.withPlanID = query
	return uppq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Amount int64 `json:"amount,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserPaymentPlan.Query().
//		GroupBy(userpaymentplan.FieldAmount).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uppq *UserPaymentPlanQuery) GroupBy(field string, fields ...string) *UserPaymentPlanGroupBy {
	uppq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserPaymentPlanGroupBy{build: uppq}
	grbuild.flds = &uppq.ctx.Fields
	grbuild.label = userpaymentplan.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Amount int64 `json:"amount,omitempty"`
//	}
//
//	client.UserPaymentPlan.Query().
//		Select(userpaymentplan.FieldAmount).
//		Scan(ctx, &v)
func (uppq *UserPaymentPlanQuery) Select(fields ...string) *UserPaymentPlanSelect {
	uppq.ctx.Fields = append(uppq.ctx.Fields, fields...)
	sbuild := &UserPaymentPlanSelect{UserPaymentPlanQuery: uppq}
	sbuild.label = userpaymentplan.Label
	sbuild.flds, sbuild.scan = &uppq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserPaymentPlanSelect configured with the given aggregations.
func (uppq *UserPaymentPlanQuery) Aggregate(fns ...AggregateFunc) *UserPaymentPlanSelect {
	return uppq.Select().Aggregate(fns...)
}

func (uppq *UserPaymentPlanQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uppq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uppq); err != nil {
				return err
			}
		}
	}
	for _, f := range uppq.ctx.Fields {
		if !userpaymentplan.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uppq.path != nil {
		prev, err := uppq.path(ctx)
		if err != nil {
			return err
		}
		uppq.sql = prev
	}
	return nil
}

func (uppq *UserPaymentPlanQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserPaymentPlan, error) {
	var (
		nodes       = []*UserPaymentPlan{}
		withFKs     = uppq.withFKs
		_spec       = uppq.querySpec()
		loadedTypes = [2]bool{
			uppq.withUserID != nil,
			uppq.withPlanID != nil,
		}
	)
	if uppq.withUserID != nil || uppq.withPlanID != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, userpaymentplan.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserPaymentPlan).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserPaymentPlan{config: uppq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uppq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uppq.withUserID; query != nil {
		if err := uppq.loadUserID(ctx, query, nodes, nil,
			func(n *UserPaymentPlan, e *User) { n.Edges.UserID = e }); err != nil {
			return nil, err
		}
	}
	if query := uppq.withPlanID; query != nil {
		if err := uppq.loadPlanID(ctx, query, nodes, nil,
			func(n *UserPaymentPlan, e *Plan) { n.Edges.PlanID = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uppq *UserPaymentPlanQuery) loadUserID(ctx context.Context, query *UserQuery, nodes []*UserPaymentPlan, init func(*UserPaymentPlan), assign func(*UserPaymentPlan, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*UserPaymentPlan)
	for i := range nodes {
		if nodes[i].user_userpaymentplans == nil {
			continue
		}
		fk := *nodes[i].user_userpaymentplans
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_userpaymentplans" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (uppq *UserPaymentPlanQuery) loadPlanID(ctx context.Context, query *PlanQuery, nodes []*UserPaymentPlan, init func(*UserPaymentPlan), assign func(*UserPaymentPlan, *Plan)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserPaymentPlan)
	for i := range nodes {
		if nodes[i].plan_userpaymentplans == nil {
			continue
		}
		fk := *nodes[i].plan_userpaymentplans
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(plan.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "plan_userpaymentplans" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (uppq *UserPaymentPlanQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uppq.querySpec()
	_spec.Node.Columns = uppq.ctx.Fields
	if len(uppq.ctx.Fields) > 0 {
		_spec.Unique = uppq.ctx.Unique != nil && *uppq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uppq.driver, _spec)
}

func (uppq *UserPaymentPlanQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userpaymentplan.Table, userpaymentplan.Columns, sqlgraph.NewFieldSpec(userpaymentplan.FieldID, field.TypeInt64))
	_spec.From = uppq.sql
	if unique := uppq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uppq.path != nil {
		_spec.Unique = true
	}
	if fields := uppq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userpaymentplan.FieldID)
		for i := range fields {
			if fields[i] != userpaymentplan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uppq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uppq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uppq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uppq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uppq *UserPaymentPlanQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uppq.driver.Dialect())
	t1 := builder.Table(userpaymentplan.Table)
	columns := uppq.ctx.Fields
	if len(columns) == 0 {
		columns = userpaymentplan.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uppq.sql != nil {
		selector = uppq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uppq.ctx.Unique != nil && *uppq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uppq.predicates {
		p(selector)
	}
	for _, p := range uppq.order {
		p(selector)
	}
	if offset := uppq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uppq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserPaymentPlanGroupBy is the group-by builder for UserPaymentPlan entities.
type UserPaymentPlanGroupBy struct {
	selector
	build *UserPaymentPlanQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uppgb *UserPaymentPlanGroupBy) Aggregate(fns ...AggregateFunc) *UserPaymentPlanGroupBy {
	uppgb.fns = append(uppgb.fns, fns...)
	return uppgb
}

// Scan applies the selector query and scans the result into the given value.
func (uppgb *UserPaymentPlanGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uppgb.build.ctx, "GroupBy")
	if err := uppgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserPaymentPlanQuery, *UserPaymentPlanGroupBy](ctx, uppgb.build, uppgb, uppgb.build.inters, v)
}

func (uppgb *UserPaymentPlanGroupBy) sqlScan(ctx context.Context, root *UserPaymentPlanQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(uppgb.fns))
	for _, fn := range uppgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*uppgb.flds)+len(uppgb.fns))
		for _, f := range *uppgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*uppgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uppgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserPaymentPlanSelect is the builder for selecting fields of UserPaymentPlan entities.
type UserPaymentPlanSelect struct {
	*UserPaymentPlanQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (upps *UserPaymentPlanSelect) Aggregate(fns ...AggregateFunc) *UserPaymentPlanSelect {
	upps.fns = append(upps.fns, fns...)
	return upps
}

// Scan applies the selector query and scans the result into the given value.
func (upps *UserPaymentPlanSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, upps.ctx, "Select")
	if err := upps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserPaymentPlanQuery, *UserPaymentPlanSelect](ctx, upps.UserPaymentPlanQuery, upps, upps.inters, v)
}

func (upps *UserPaymentPlanSelect) sqlScan(ctx context.Context, root *UserPaymentPlanQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(upps.fns))
	for _, fn := range upps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*upps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := upps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
