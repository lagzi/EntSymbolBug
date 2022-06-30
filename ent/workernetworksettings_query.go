// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lagzi/EntSymbolBug/ent/predicate"
	"github.com/lagzi/EntSymbolBug/ent/workercontainedinformation"
	"github.com/lagzi/EntSymbolBug/ent/workernetworksettings"
)

// WorkerNetworkSettingsQuery is the builder for querying WorkerNetworkSettings entities.
type WorkerNetworkSettingsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.WorkerNetworkSettings
	// eager-loading edges.
	withWorkerContainedInformation *WorkerContainedInformationQuery
	withFKs                        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkerNetworkSettingsQuery builder.
func (wnsq *WorkerNetworkSettingsQuery) Where(ps ...predicate.WorkerNetworkSettings) *WorkerNetworkSettingsQuery {
	wnsq.predicates = append(wnsq.predicates, ps...)
	return wnsq
}

// Limit adds a limit step to the query.
func (wnsq *WorkerNetworkSettingsQuery) Limit(limit int) *WorkerNetworkSettingsQuery {
	wnsq.limit = &limit
	return wnsq
}

// Offset adds an offset step to the query.
func (wnsq *WorkerNetworkSettingsQuery) Offset(offset int) *WorkerNetworkSettingsQuery {
	wnsq.offset = &offset
	return wnsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wnsq *WorkerNetworkSettingsQuery) Unique(unique bool) *WorkerNetworkSettingsQuery {
	wnsq.unique = &unique
	return wnsq
}

// Order adds an order step to the query.
func (wnsq *WorkerNetworkSettingsQuery) Order(o ...OrderFunc) *WorkerNetworkSettingsQuery {
	wnsq.order = append(wnsq.order, o...)
	return wnsq
}

// QueryWorkerContainedInformation chains the current query on the "worker_contained_information" edge.
func (wnsq *WorkerNetworkSettingsQuery) QueryWorkerContainedInformation() *WorkerContainedInformationQuery {
	query := &WorkerContainedInformationQuery{config: wnsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wnsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wnsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workernetworksettings.Table, workernetworksettings.FieldID, selector),
			sqlgraph.To(workercontainedinformation.Table, workercontainedinformation.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, workernetworksettings.WorkerContainedInformationTable, workernetworksettings.WorkerContainedInformationColumn),
		)
		fromU = sqlgraph.SetNeighbors(wnsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WorkerNetworkSettings entity from the query.
// Returns a *NotFoundError when no WorkerNetworkSettings was found.
func (wnsq *WorkerNetworkSettingsQuery) First(ctx context.Context) (*WorkerNetworkSettings, error) {
	nodes, err := wnsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workernetworksettings.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wnsq *WorkerNetworkSettingsQuery) FirstX(ctx context.Context) *WorkerNetworkSettings {
	node, err := wnsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkerNetworkSettings ID from the query.
// Returns a *NotFoundError when no WorkerNetworkSettings ID was found.
func (wnsq *WorkerNetworkSettingsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wnsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workernetworksettings.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wnsq *WorkerNetworkSettingsQuery) FirstIDX(ctx context.Context) int {
	id, err := wnsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WorkerNetworkSettings entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WorkerNetworkSettings entity is found.
// Returns a *NotFoundError when no WorkerNetworkSettings entities are found.
func (wnsq *WorkerNetworkSettingsQuery) Only(ctx context.Context) (*WorkerNetworkSettings, error) {
	nodes, err := wnsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workernetworksettings.Label}
	default:
		return nil, &NotSingularError{workernetworksettings.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wnsq *WorkerNetworkSettingsQuery) OnlyX(ctx context.Context) *WorkerNetworkSettings {
	node, err := wnsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WorkerNetworkSettings ID in the query.
// Returns a *NotSingularError when more than one WorkerNetworkSettings ID is found.
// Returns a *NotFoundError when no entities are found.
func (wnsq *WorkerNetworkSettingsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wnsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workernetworksettings.Label}
	default:
		err = &NotSingularError{workernetworksettings.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wnsq *WorkerNetworkSettingsQuery) OnlyIDX(ctx context.Context) int {
	id, err := wnsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkerNetworkSettingsSlice.
func (wnsq *WorkerNetworkSettingsQuery) All(ctx context.Context) ([]*WorkerNetworkSettings, error) {
	if err := wnsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wnsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wnsq *WorkerNetworkSettingsQuery) AllX(ctx context.Context) []*WorkerNetworkSettings {
	nodes, err := wnsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkerNetworkSettings IDs.
func (wnsq *WorkerNetworkSettingsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wnsq.Select(workernetworksettings.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wnsq *WorkerNetworkSettingsQuery) IDsX(ctx context.Context) []int {
	ids, err := wnsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wnsq *WorkerNetworkSettingsQuery) Count(ctx context.Context) (int, error) {
	if err := wnsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wnsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wnsq *WorkerNetworkSettingsQuery) CountX(ctx context.Context) int {
	count, err := wnsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wnsq *WorkerNetworkSettingsQuery) Exist(ctx context.Context) (bool, error) {
	if err := wnsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wnsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wnsq *WorkerNetworkSettingsQuery) ExistX(ctx context.Context) bool {
	exist, err := wnsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkerNetworkSettingsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wnsq *WorkerNetworkSettingsQuery) Clone() *WorkerNetworkSettingsQuery {
	if wnsq == nil {
		return nil
	}
	return &WorkerNetworkSettingsQuery{
		config:                         wnsq.config,
		limit:                          wnsq.limit,
		offset:                         wnsq.offset,
		order:                          append([]OrderFunc{}, wnsq.order...),
		predicates:                     append([]predicate.WorkerNetworkSettings{}, wnsq.predicates...),
		withWorkerContainedInformation: wnsq.withWorkerContainedInformation.Clone(),
		// clone intermediate query.
		sql:    wnsq.sql.Clone(),
		path:   wnsq.path,
		unique: wnsq.unique,
	}
}

// WithWorkerContainedInformation tells the query-builder to eager-load the nodes that are connected to
// the "worker_contained_information" edge. The optional arguments are used to configure the query builder of the edge.
func (wnsq *WorkerNetworkSettingsQuery) WithWorkerContainedInformation(opts ...func(*WorkerContainedInformationQuery)) *WorkerNetworkSettingsQuery {
	query := &WorkerContainedInformationQuery{config: wnsq.config}
	for _, opt := range opts {
		opt(query)
	}
	wnsq.withWorkerContainedInformation = query
	return wnsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (wnsq *WorkerNetworkSettingsQuery) GroupBy(field string, fields ...string) *WorkerNetworkSettingsGroupBy {
	grbuild := &WorkerNetworkSettingsGroupBy{config: wnsq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wnsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wnsq.sqlQuery(ctx), nil
	}
	grbuild.label = workernetworksettings.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (wnsq *WorkerNetworkSettingsQuery) Select(fields ...string) *WorkerNetworkSettingsSelect {
	wnsq.fields = append(wnsq.fields, fields...)
	selbuild := &WorkerNetworkSettingsSelect{WorkerNetworkSettingsQuery: wnsq}
	selbuild.label = workernetworksettings.Label
	selbuild.flds, selbuild.scan = &wnsq.fields, selbuild.Scan
	return selbuild
}

func (wnsq *WorkerNetworkSettingsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wnsq.fields {
		if !workernetworksettings.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wnsq.path != nil {
		prev, err := wnsq.path(ctx)
		if err != nil {
			return err
		}
		wnsq.sql = prev
	}
	return nil
}

func (wnsq *WorkerNetworkSettingsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WorkerNetworkSettings, error) {
	var (
		nodes       = []*WorkerNetworkSettings{}
		withFKs     = wnsq.withFKs
		_spec       = wnsq.querySpec()
		loadedTypes = [1]bool{
			wnsq.withWorkerContainedInformation != nil,
		}
	)
	if wnsq.withWorkerContainedInformation != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, workernetworksettings.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*WorkerNetworkSettings).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &WorkerNetworkSettings{config: wnsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wnsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wnsq.withWorkerContainedInformation; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*WorkerNetworkSettings)
		for i := range nodes {
			if nodes[i].worker_contained_information_network_settings == nil {
				continue
			}
			fk := *nodes[i].worker_contained_information_network_settings
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(workercontainedinformation.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "worker_contained_information_network_settings" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.WorkerContainedInformation = n
			}
		}
	}

	return nodes, nil
}

func (wnsq *WorkerNetworkSettingsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wnsq.querySpec()
	_spec.Node.Columns = wnsq.fields
	if len(wnsq.fields) > 0 {
		_spec.Unique = wnsq.unique != nil && *wnsq.unique
	}
	return sqlgraph.CountNodes(ctx, wnsq.driver, _spec)
}

func (wnsq *WorkerNetworkSettingsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wnsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wnsq *WorkerNetworkSettingsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workernetworksettings.Table,
			Columns: workernetworksettings.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workernetworksettings.FieldID,
			},
		},
		From:   wnsq.sql,
		Unique: true,
	}
	if unique := wnsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wnsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workernetworksettings.FieldID)
		for i := range fields {
			if fields[i] != workernetworksettings.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wnsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wnsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wnsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wnsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wnsq *WorkerNetworkSettingsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wnsq.driver.Dialect())
	t1 := builder.Table(workernetworksettings.Table)
	columns := wnsq.fields
	if len(columns) == 0 {
		columns = workernetworksettings.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wnsq.sql != nil {
		selector = wnsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wnsq.unique != nil && *wnsq.unique {
		selector.Distinct()
	}
	for _, p := range wnsq.predicates {
		p(selector)
	}
	for _, p := range wnsq.order {
		p(selector)
	}
	if offset := wnsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wnsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WorkerNetworkSettingsGroupBy is the group-by builder for WorkerNetworkSettings entities.
type WorkerNetworkSettingsGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wnsgb *WorkerNetworkSettingsGroupBy) Aggregate(fns ...AggregateFunc) *WorkerNetworkSettingsGroupBy {
	wnsgb.fns = append(wnsgb.fns, fns...)
	return wnsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (wnsgb *WorkerNetworkSettingsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wnsgb.path(ctx)
	if err != nil {
		return err
	}
	wnsgb.sql = query
	return wnsgb.sqlScan(ctx, v)
}

func (wnsgb *WorkerNetworkSettingsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wnsgb.fields {
		if !workernetworksettings.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wnsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wnsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wnsgb *WorkerNetworkSettingsGroupBy) sqlQuery() *sql.Selector {
	selector := wnsgb.sql.Select()
	aggregation := make([]string, 0, len(wnsgb.fns))
	for _, fn := range wnsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(wnsgb.fields)+len(wnsgb.fns))
		for _, f := range wnsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(wnsgb.fields...)...)
}

// WorkerNetworkSettingsSelect is the builder for selecting fields of WorkerNetworkSettings entities.
type WorkerNetworkSettingsSelect struct {
	*WorkerNetworkSettingsQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (wnss *WorkerNetworkSettingsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := wnss.prepareQuery(ctx); err != nil {
		return err
	}
	wnss.sql = wnss.WorkerNetworkSettingsQuery.sqlQuery(ctx)
	return wnss.sqlScan(ctx, v)
}

func (wnss *WorkerNetworkSettingsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wnss.sql.Query()
	if err := wnss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
