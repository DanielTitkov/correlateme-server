// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/dataset"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/indicator"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/indicatorparams"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/scale"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/user"
)

// IndicatorQuery is the builder for querying Indicator entities.
type IndicatorQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Indicator
	// eager-loading edges.
	withDatasets        *DatasetQuery
	withIndicatorParams *IndicatorParamsQuery
	withAuthor          *UserQuery
	withScale           *ScaleQuery
	withFKs             bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IndicatorQuery builder.
func (iq *IndicatorQuery) Where(ps ...predicate.Indicator) *IndicatorQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit adds a limit step to the query.
func (iq *IndicatorQuery) Limit(limit int) *IndicatorQuery {
	iq.limit = &limit
	return iq
}

// Offset adds an offset step to the query.
func (iq *IndicatorQuery) Offset(offset int) *IndicatorQuery {
	iq.offset = &offset
	return iq
}

// Order adds an order step to the query.
func (iq *IndicatorQuery) Order(o ...OrderFunc) *IndicatorQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryDatasets chains the current query on the "datasets" edge.
func (iq *IndicatorQuery) QueryDatasets() *DatasetQuery {
	query := &DatasetQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(indicator.Table, indicator.FieldID, selector),
			sqlgraph.To(dataset.Table, dataset.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, indicator.DatasetsTable, indicator.DatasetsColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIndicatorParams chains the current query on the "indicator_params" edge.
func (iq *IndicatorQuery) QueryIndicatorParams() *IndicatorParamsQuery {
	query := &IndicatorParamsQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(indicator.Table, indicator.FieldID, selector),
			sqlgraph.To(indicatorparams.Table, indicatorparams.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, indicator.IndicatorParamsTable, indicator.IndicatorParamsColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAuthor chains the current query on the "author" edge.
func (iq *IndicatorQuery) QueryAuthor() *UserQuery {
	query := &UserQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(indicator.Table, indicator.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, indicator.AuthorTable, indicator.AuthorColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryScale chains the current query on the "scale" edge.
func (iq *IndicatorQuery) QueryScale() *ScaleQuery {
	query := &ScaleQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(indicator.Table, indicator.FieldID, selector),
			sqlgraph.To(scale.Table, scale.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, indicator.ScaleTable, indicator.ScaleColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Indicator entity from the query.
// Returns a *NotFoundError when no Indicator was found.
func (iq *IndicatorQuery) First(ctx context.Context) (*Indicator, error) {
	nodes, err := iq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{indicator.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *IndicatorQuery) FirstX(ctx context.Context) *Indicator {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Indicator ID from the query.
// Returns a *NotFoundError when no Indicator ID was found.
func (iq *IndicatorQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{indicator.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *IndicatorQuery) FirstIDX(ctx context.Context) int {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Indicator entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Indicator entity is not found.
// Returns a *NotFoundError when no Indicator entities are found.
func (iq *IndicatorQuery) Only(ctx context.Context) (*Indicator, error) {
	nodes, err := iq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{indicator.Label}
	default:
		return nil, &NotSingularError{indicator.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *IndicatorQuery) OnlyX(ctx context.Context) *Indicator {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Indicator ID in the query.
// Returns a *NotSingularError when exactly one Indicator ID is not found.
// Returns a *NotFoundError when no entities are found.
func (iq *IndicatorQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{indicator.Label}
	default:
		err = &NotSingularError{indicator.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *IndicatorQuery) OnlyIDX(ctx context.Context) int {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Indicators.
func (iq *IndicatorQuery) All(ctx context.Context) ([]*Indicator, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return iq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iq *IndicatorQuery) AllX(ctx context.Context) []*Indicator {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Indicator IDs.
func (iq *IndicatorQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := iq.Select(indicator.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *IndicatorQuery) IDsX(ctx context.Context) []int {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *IndicatorQuery) Count(ctx context.Context) (int, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return iq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iq *IndicatorQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *IndicatorQuery) Exist(ctx context.Context) (bool, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return iq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *IndicatorQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IndicatorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *IndicatorQuery) Clone() *IndicatorQuery {
	if iq == nil {
		return nil
	}
	return &IndicatorQuery{
		config:              iq.config,
		limit:               iq.limit,
		offset:              iq.offset,
		order:               append([]OrderFunc{}, iq.order...),
		predicates:          append([]predicate.Indicator{}, iq.predicates...),
		withDatasets:        iq.withDatasets.Clone(),
		withIndicatorParams: iq.withIndicatorParams.Clone(),
		withAuthor:          iq.withAuthor.Clone(),
		withScale:           iq.withScale.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithDatasets tells the query-builder to eager-load the nodes that are connected to
// the "datasets" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IndicatorQuery) WithDatasets(opts ...func(*DatasetQuery)) *IndicatorQuery {
	query := &DatasetQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withDatasets = query
	return iq
}

// WithIndicatorParams tells the query-builder to eager-load the nodes that are connected to
// the "indicator_params" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IndicatorQuery) WithIndicatorParams(opts ...func(*IndicatorParamsQuery)) *IndicatorQuery {
	query := &IndicatorParamsQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withIndicatorParams = query
	return iq
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IndicatorQuery) WithAuthor(opts ...func(*UserQuery)) *IndicatorQuery {
	query := &UserQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withAuthor = query
	return iq
}

// WithScale tells the query-builder to eager-load the nodes that are connected to
// the "scale" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IndicatorQuery) WithScale(opts ...func(*ScaleQuery)) *IndicatorQuery {
	query := &ScaleQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withScale = query
	return iq
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
//	client.Indicator.Query().
//		GroupBy(indicator.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (iq *IndicatorQuery) GroupBy(field string, fields ...string) *IndicatorGroupBy {
	group := &IndicatorGroupBy{config: iq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iq.sqlQuery(ctx), nil
	}
	return group
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
//	client.Indicator.Query().
//		Select(indicator.FieldCreateTime).
//		Scan(ctx, &v)
//
func (iq *IndicatorQuery) Select(field string, fields ...string) *IndicatorSelect {
	iq.fields = append([]string{field}, fields...)
	return &IndicatorSelect{IndicatorQuery: iq}
}

func (iq *IndicatorQuery) prepareQuery(ctx context.Context) error {
	for _, f := range iq.fields {
		if !indicator.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *IndicatorQuery) sqlAll(ctx context.Context) ([]*Indicator, error) {
	var (
		nodes       = []*Indicator{}
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [4]bool{
			iq.withDatasets != nil,
			iq.withIndicatorParams != nil,
			iq.withAuthor != nil,
			iq.withScale != nil,
		}
	)
	if iq.withAuthor != nil || iq.withScale != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, indicator.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Indicator{config: iq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := iq.withDatasets; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Indicator)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Datasets = []*Dataset{}
		}
		query.withFKs = true
		query.Where(predicate.Dataset(func(s *sql.Selector) {
			s.Where(sql.InValues(indicator.DatasetsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.indicator_datasets
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "indicator_datasets" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "indicator_datasets" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Datasets = append(node.Edges.Datasets, n)
		}
	}

	if query := iq.withIndicatorParams; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Indicator)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.IndicatorParams(func(s *sql.Selector) {
			s.Where(sql.InValues(indicator.IndicatorParamsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.indicator_indicator_params
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "indicator_indicator_params" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "indicator_indicator_params" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.IndicatorParams = n
		}
	}

	if query := iq.withAuthor; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Indicator)
		for i := range nodes {
			if fk := nodes[i].user_indicators; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_indicators" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Author = n
			}
		}
	}

	if query := iq.withScale; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Indicator)
		for i := range nodes {
			if fk := nodes[i].scale_indicators; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(scale.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "scale_indicators" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Scale = n
			}
		}
	}

	return nodes, nil
}

func (iq *IndicatorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *IndicatorQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := iq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (iq *IndicatorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   indicator.Table,
			Columns: indicator.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: indicator.FieldID,
			},
		},
		From:   iq.sql,
		Unique: true,
	}
	if fields := iq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, indicator.FieldID)
		for i := range fields {
			if fields[i] != indicator.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, indicator.ValidColumn)
			}
		}
	}
	return _spec
}

func (iq *IndicatorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(indicator.Table)
	selector := builder.Select(t1.Columns(indicator.Columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(indicator.Columns...)...)
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector, indicator.ValidColumn)
	}
	if offset := iq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IndicatorGroupBy is the group-by builder for Indicator entities.
type IndicatorGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *IndicatorGroupBy) Aggregate(fns ...AggregateFunc) *IndicatorGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the group-by query and scans the result into the given value.
func (igb *IndicatorGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := igb.path(ctx)
	if err != nil {
		return err
	}
	igb.sql = query
	return igb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (igb *IndicatorGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := igb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *IndicatorGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: IndicatorGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (igb *IndicatorGroupBy) StringsX(ctx context.Context) []string {
	v, err := igb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *IndicatorGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = igb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{indicator.Label}
	default:
		err = fmt.Errorf("ent: IndicatorGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (igb *IndicatorGroupBy) StringX(ctx context.Context) string {
	v, err := igb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *IndicatorGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: IndicatorGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (igb *IndicatorGroupBy) IntsX(ctx context.Context) []int {
	v, err := igb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *IndicatorGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = igb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{indicator.Label}
	default:
		err = fmt.Errorf("ent: IndicatorGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (igb *IndicatorGroupBy) IntX(ctx context.Context) int {
	v, err := igb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *IndicatorGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: IndicatorGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (igb *IndicatorGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := igb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *IndicatorGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = igb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{indicator.Label}
	default:
		err = fmt.Errorf("ent: IndicatorGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (igb *IndicatorGroupBy) Float64X(ctx context.Context) float64 {
	v, err := igb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *IndicatorGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: IndicatorGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (igb *IndicatorGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := igb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *IndicatorGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = igb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{indicator.Label}
	default:
		err = fmt.Errorf("ent: IndicatorGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (igb *IndicatorGroupBy) BoolX(ctx context.Context) bool {
	v, err := igb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (igb *IndicatorGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range igb.fields {
		if !indicator.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := igb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (igb *IndicatorGroupBy) sqlQuery() *sql.Selector {
	selector := igb.sql
	columns := make([]string, 0, len(igb.fields)+len(igb.fns))
	columns = append(columns, igb.fields...)
	for _, fn := range igb.fns {
		columns = append(columns, fn(selector, indicator.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(igb.fields...)
}

// IndicatorSelect is the builder for selecting fields of Indicator entities.
type IndicatorSelect struct {
	*IndicatorQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (is *IndicatorSelect) Scan(ctx context.Context, v interface{}) error {
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	is.sql = is.IndicatorQuery.sqlQuery(ctx)
	return is.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (is *IndicatorSelect) ScanX(ctx context.Context, v interface{}) {
	if err := is.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (is *IndicatorSelect) Strings(ctx context.Context) ([]string, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: IndicatorSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (is *IndicatorSelect) StringsX(ctx context.Context) []string {
	v, err := is.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (is *IndicatorSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = is.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{indicator.Label}
	default:
		err = fmt.Errorf("ent: IndicatorSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (is *IndicatorSelect) StringX(ctx context.Context) string {
	v, err := is.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (is *IndicatorSelect) Ints(ctx context.Context) ([]int, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: IndicatorSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (is *IndicatorSelect) IntsX(ctx context.Context) []int {
	v, err := is.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (is *IndicatorSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = is.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{indicator.Label}
	default:
		err = fmt.Errorf("ent: IndicatorSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (is *IndicatorSelect) IntX(ctx context.Context) int {
	v, err := is.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (is *IndicatorSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: IndicatorSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (is *IndicatorSelect) Float64sX(ctx context.Context) []float64 {
	v, err := is.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (is *IndicatorSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = is.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{indicator.Label}
	default:
		err = fmt.Errorf("ent: IndicatorSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (is *IndicatorSelect) Float64X(ctx context.Context) float64 {
	v, err := is.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (is *IndicatorSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: IndicatorSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (is *IndicatorSelect) BoolsX(ctx context.Context) []bool {
	v, err := is.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (is *IndicatorSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = is.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{indicator.Label}
	default:
		err = fmt.Errorf("ent: IndicatorSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (is *IndicatorSelect) BoolX(ctx context.Context) bool {
	v, err := is.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (is *IndicatorSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := is.sqlQuery().Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (is *IndicatorSelect) sqlQuery() sql.Querier {
	selector := is.sql
	selector.Select(selector.Columns(is.fields...)...)
	return selector
}
