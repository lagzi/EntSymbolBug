// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lagzi/EntSymbolBug/ent/predicate"
	"github.com/lagzi/EntSymbolBug/ent/workercontainedinformation"
	"github.com/lagzi/EntSymbolBug/ent/workernetworksettings"
)

// WorkerContainedInformationQuery is the builder for querying WorkerContainedInformation entities.
type WorkerContainedInformationQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.WorkerContainedInformation
	// eager-loading edges.
	withNetworkSettings *WorkerNetworkSettingsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkerContainedInformationQuery builder.
func (wciq *WorkerContainedInformationQuery) Where(ps ...predicate.WorkerContainedInformation) *WorkerContainedInformationQuery {
	wciq.predicates = append(wciq.predicates, ps...)
	return wciq
}

// Limit adds a limit step to the query.
func (wciq *WorkerContainedInformationQuery) Limit(limit int) *WorkerContainedInformationQuery {
	wciq.limit = &limit
	return wciq
}

// Offset adds an offset step to the query.
func (wciq *WorkerContainedInformationQuery) Offset(offset int) *WorkerContainedInformationQuery {
	wciq.offset = &offset
	return wciq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wciq *WorkerContainedInformationQuery) Unique(unique bool) *WorkerContainedInformationQuery {
	wciq.unique = &unique
	return wciq
}

// Order adds an order step to the query.
func (wciq *WorkerContainedInformationQuery) Order(o ...OrderFunc) *WorkerContainedInformationQuery {
	wciq.order = append(wciq.order, o...)
	return wciq
}

// QueryNetworkSettings chains the current query on the "network_settings" edge.
func (wciq *WorkerContainedInformationQuery) QueryNetworkSettings() *WorkerNetworkSettingsQuery {
	query := &WorkerNetworkSettingsQuery{config: wciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workercontainedinformation.Table, workercontainedinformation.FieldID, selector),
			sqlgraph.To(workernetworksettings.Table, workernetworksettings.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, workercontainedinformation.NetworkSettingsTable, workercontainedinformation.NetworkSettingsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WorkerContainedInformation entity from the query.
// Returns a *NotFoundError when no WorkerContainedInformation was found.
func (wciq *WorkerContainedInformationQuery) First(ctx context.Context) (*WorkerContainedInformation, error) {
	nodes, err := wciq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workercontainedinformation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wciq *WorkerContainedInformationQuery) FirstX(ctx context.Context) *WorkerContainedInformation {
	node, err := wciq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkerContainedInformation ID from the query.
// Returns a *NotFoundError when no WorkerContainedInformation ID was found.
func (wciq *WorkerContainedInformationQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wciq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workercontainedinformation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wciq *WorkerContainedInformationQuery) FirstIDX(ctx context.Context) int {
	id, err := wciq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WorkerContainedInformation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WorkerContainedInformation entity is found.
// Returns a *NotFoundError when no WorkerContainedInformation entities are found.
func (wciq *WorkerContainedInformationQuery) Only(ctx context.Context) (*WorkerContainedInformation, error) {
	nodes, err := wciq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workercontainedinformation.Label}
	default:
		return nil, &NotSingularError{workercontainedinformation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wciq *WorkerContainedInformationQuery) OnlyX(ctx context.Context) *WorkerContainedInformation {
	node, err := wciq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WorkerContainedInformation ID in the query.
// Returns a *NotSingularError when more than one WorkerContainedInformation ID is found.
// Returns a *NotFoundError when no entities are found.
func (wciq *WorkerContainedInformationQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wciq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workercontainedinformation.Label}
	default:
		err = &NotSingularError{workercontainedinformation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wciq *WorkerContainedInformationQuery) OnlyIDX(ctx context.Context) int {
	id, err := wciq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkerContainedInformations.
func (wciq *WorkerContainedInformationQuery) All(ctx context.Context) ([]*WorkerContainedInformation, error) {
	if err := wciq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wciq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wciq *WorkerContainedInformationQuery) AllX(ctx context.Context) []*WorkerContainedInformation {
	nodes, err := wciq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkerContainedInformation IDs.
func (wciq *WorkerContainedInformationQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wciq.Select(workercontainedinformation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wciq *WorkerContainedInformationQuery) IDsX(ctx context.Context) []int {
	ids, err := wciq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wciq *WorkerContainedInformationQuery) Count(ctx context.Context) (int, error) {
	if err := wciq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wciq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wciq *WorkerContainedInformationQuery) CountX(ctx context.Context) int {
	count, err := wciq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wciq *WorkerContainedInformationQuery) Exist(ctx context.Context) (bool, error) {
	if err := wciq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wciq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wciq *WorkerContainedInformationQuery) ExistX(ctx context.Context) bool {
	exist, err := wciq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkerContainedInformationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wciq *WorkerContainedInformationQuery) Clone() *WorkerContainedInformationQuery {
	if wciq == nil {
		return nil
	}
	return &WorkerContainedInformationQuery{
		config:              wciq.config,
		limit:               wciq.limit,
		offset:              wciq.offset,
		order:               append([]OrderFunc{}, wciq.order...),
		predicates:          append([]predicate.WorkerContainedInformation{}, wciq.predicates...),
		withNetworkSettings: wciq.withNetworkSettings.Clone(),
		// clone intermediate query.
		sql:    wciq.sql.Clone(),
		path:   wciq.path,
		unique: wciq.unique,
	}
}

// WithNetworkSettings tells the query-builder to eager-load the nodes that are connected to
// the "network_settings" edge. The optional arguments are used to configure the query builder of the edge.
func (wciq *WorkerContainedInformationQuery) WithNetworkSettings(opts ...func(*WorkerNetworkSettingsQuery)) *WorkerContainedInformationQuery {
	query := &WorkerNetworkSettingsQuery{config: wciq.config}
	for _, opt := range opts {
		opt(query)
	}
	wciq.withNetworkSettings = query
	return wciq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (wciq *WorkerContainedInformationQuery) GroupBy(field string, fields ...string) *WorkerContainedInformationGroupBy {
	grbuild := &WorkerContainedInformationGroupBy{config: wciq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wciq.sqlQuery(ctx), nil
	}
	grbuild.label = workercontainedinformation.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (wciq *WorkerContainedInformationQuery) Select(fields ...string) *WorkerContainedInformationSelect {
	wciq.fields = append(wciq.fields, fields...)
	selbuild := &WorkerContainedInformationSelect{WorkerContainedInformationQuery: wciq}
	selbuild.label = workercontainedinformation.Label
	selbuild.flds, selbuild.scan = &wciq.fields, selbuild.Scan
	return selbuild
}

func (wciq *WorkerContainedInformationQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wciq.fields {
		if !workercontainedinformation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wciq.path != nil {
		prev, err := wciq.path(ctx)
		if err != nil {
			return err
		}
		wciq.sql = prev
	}
	return nil
}

func (wciq *WorkerContainedInformationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WorkerContainedInformation, error) {
	var (
		nodes       = []*WorkerContainedInformation{}
		_spec       = wciq.querySpec()
		loadedTypes = [1]bool{
			wciq.withNetworkSettings != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*WorkerContainedInformation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &WorkerContainedInformation{config: wciq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wciq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wciq.withNetworkSettings; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*WorkerContainedInformation)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.WorkerNetworkSettings(func(s *sql.Selector) {
			s.Where(sql.InValues(workercontainedinformation.NetworkSettingsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.worker_contained_information_network_settings
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "worker_contained_information_network_settings" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "worker_contained_information_network_settings" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.NetworkSettings = n
		}
	}

	return nodes, nil
}

func (wciq *WorkerContainedInformationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wciq.querySpec()
	_spec.Node.Columns = wciq.fields
	if len(wciq.fields) > 0 {
		_spec.Unique = wciq.unique != nil && *wciq.unique
	}
	return sqlgraph.CountNodes(ctx, wciq.driver, _spec)
}

func (wciq *WorkerContainedInformationQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wciq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wciq *WorkerContainedInformationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workercontainedinformation.Table,
			Columns: workercontainedinformation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workercontainedinformation.FieldID,
			},
		},
		From:   wciq.sql,
		Unique: true,
	}
	if unique := wciq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wciq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workercontainedinformation.FieldID)
		for i := range fields {
			if fields[i] != workercontainedinformation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wciq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wciq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wciq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wciq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wciq *WorkerContainedInformationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wciq.driver.Dialect())
	t1 := builder.Table(workercontainedinformation.Table)
	columns := wciq.fields
	if len(columns) == 0 {
		columns = workercontainedinformation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wciq.sql != nil {
		selector = wciq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wciq.unique != nil && *wciq.unique {
		selector.Distinct()
	}
	for _, p := range wciq.predicates {
		p(selector)
	}
	for _, p := range wciq.order {
		p(selector)
	}
	if offset := wciq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wciq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WorkerContainedInformationGroupBy is the group-by builder for WorkerContainedInformation entities.
type WorkerContainedInformationGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wcigb *WorkerContainedInformationGroupBy) Aggregate(fns ...AggregateFunc) *WorkerContainedInformationGroupBy {
	wcigb.fns = append(wcigb.fns, fns...)
	return wcigb
}

// Scan applies the group-by query and scans the result into the given value.
func (wcigb *WorkerContainedInformationGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wcigb.path(ctx)
	if err != nil {
		return err
	}
	wcigb.sql = query
	return wcigb.sqlScan(ctx, v)
}

func (wcigb *WorkerContainedInformationGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wcigb.fields {
		if !workercontainedinformation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wcigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wcigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wcigb *WorkerContainedInformationGroupBy) sqlQuery() *sql.Selector {
	selector := wcigb.sql.Select()
	aggregation := make([]string, 0, len(wcigb.fns))
	for _, fn := range wcigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(wcigb.fields)+len(wcigb.fns))
		for _, f := range wcigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(wcigb.fields...)...)
}

// WorkerContainedInformationSelect is the builder for selecting fields of WorkerContainedInformation entities.
type WorkerContainedInformationSelect struct {
	*WorkerContainedInformationQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (wcis *WorkerContainedInformationSelect) Scan(ctx context.Context, v interface{}) error {
	if err := wcis.prepareQuery(ctx); err != nil {
		return err
	}
	wcis.sql = wcis.WorkerContainedInformationQuery.sqlQuery(ctx)
	return wcis.sqlScan(ctx, v)
}

func (wcis *WorkerContainedInformationSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wcis.sql.Query()
	if err := wcis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
