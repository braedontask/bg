// Code generated by entc, DO NOT EDIT.

package ent

import (
	"back/ent/biouser"
	"back/ent/predicate"
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BioUserQuery is the builder for querying BioUser entities.
type BioUserQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BioUser
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BioUserQuery builder.
func (buq *BioUserQuery) Where(ps ...predicate.BioUser) *BioUserQuery {
	buq.predicates = append(buq.predicates, ps...)
	return buq
}

// Limit adds a limit step to the query.
func (buq *BioUserQuery) Limit(limit int) *BioUserQuery {
	buq.limit = &limit
	return buq
}

// Offset adds an offset step to the query.
func (buq *BioUserQuery) Offset(offset int) *BioUserQuery {
	buq.offset = &offset
	return buq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (buq *BioUserQuery) Unique(unique bool) *BioUserQuery {
	buq.unique = &unique
	return buq
}

// Order adds an order step to the query.
func (buq *BioUserQuery) Order(o ...OrderFunc) *BioUserQuery {
	buq.order = append(buq.order, o...)
	return buq
}

// First returns the first BioUser entity from the query.
// Returns a *NotFoundError when no BioUser was found.
func (buq *BioUserQuery) First(ctx context.Context) (*BioUser, error) {
	nodes, err := buq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{biouser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (buq *BioUserQuery) FirstX(ctx context.Context) *BioUser {
	node, err := buq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BioUser ID from the query.
// Returns a *NotFoundError when no BioUser ID was found.
func (buq *BioUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = buq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{biouser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (buq *BioUserQuery) FirstIDX(ctx context.Context) int {
	id, err := buq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BioUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one BioUser entity is not found.
// Returns a *NotFoundError when no BioUser entities are found.
func (buq *BioUserQuery) Only(ctx context.Context) (*BioUser, error) {
	nodes, err := buq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{biouser.Label}
	default:
		return nil, &NotSingularError{biouser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (buq *BioUserQuery) OnlyX(ctx context.Context) *BioUser {
	node, err := buq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BioUser ID in the query.
// Returns a *NotSingularError when exactly one BioUser ID is not found.
// Returns a *NotFoundError when no entities are found.
func (buq *BioUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = buq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{biouser.Label}
	default:
		err = &NotSingularError{biouser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (buq *BioUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := buq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BioUsers.
func (buq *BioUserQuery) All(ctx context.Context) ([]*BioUser, error) {
	if err := buq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return buq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (buq *BioUserQuery) AllX(ctx context.Context) []*BioUser {
	nodes, err := buq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BioUser IDs.
func (buq *BioUserQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := buq.Select(biouser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (buq *BioUserQuery) IDsX(ctx context.Context) []int {
	ids, err := buq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (buq *BioUserQuery) Count(ctx context.Context) (int, error) {
	if err := buq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return buq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (buq *BioUserQuery) CountX(ctx context.Context) int {
	count, err := buq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (buq *BioUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := buq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return buq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (buq *BioUserQuery) ExistX(ctx context.Context) bool {
	exist, err := buq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BioUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (buq *BioUserQuery) Clone() *BioUserQuery {
	if buq == nil {
		return nil
	}
	return &BioUserQuery{
		config:     buq.config,
		limit:      buq.limit,
		offset:     buq.offset,
		order:      append([]OrderFunc{}, buq.order...),
		predicates: append([]predicate.BioUser{}, buq.predicates...),
		// clone intermediate query.
		sql:  buq.sql.Clone(),
		path: buq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BioUser.Query().
//		GroupBy(biouser.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (buq *BioUserQuery) GroupBy(field string, fields ...string) *BioUserGroupBy {
	group := &BioUserGroupBy{config: buq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := buq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return buq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.BioUser.Query().
//		Select(biouser.FieldName).
//		Scan(ctx, &v)
//
func (buq *BioUserQuery) Select(fields ...string) *BioUserSelect {
	buq.fields = append(buq.fields, fields...)
	return &BioUserSelect{BioUserQuery: buq}
}

func (buq *BioUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range buq.fields {
		if !biouser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if buq.path != nil {
		prev, err := buq.path(ctx)
		if err != nil {
			return err
		}
		buq.sql = prev
	}
	return nil
}

func (buq *BioUserQuery) sqlAll(ctx context.Context) ([]*BioUser, error) {
	var (
		nodes = []*BioUser{}
		_spec = buq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &BioUser{config: buq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, buq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (buq *BioUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := buq.querySpec()
	return sqlgraph.CountNodes(ctx, buq.driver, _spec)
}

func (buq *BioUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := buq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (buq *BioUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   biouser.Table,
			Columns: biouser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: biouser.FieldID,
			},
		},
		From:   buq.sql,
		Unique: true,
	}
	if unique := buq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := buq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, biouser.FieldID)
		for i := range fields {
			if fields[i] != biouser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := buq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := buq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := buq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := buq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (buq *BioUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(buq.driver.Dialect())
	t1 := builder.Table(biouser.Table)
	columns := buq.fields
	if len(columns) == 0 {
		columns = biouser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if buq.sql != nil {
		selector = buq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range buq.predicates {
		p(selector)
	}
	for _, p := range buq.order {
		p(selector)
	}
	if offset := buq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := buq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BioUserGroupBy is the group-by builder for BioUser entities.
type BioUserGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bugb *BioUserGroupBy) Aggregate(fns ...AggregateFunc) *BioUserGroupBy {
	bugb.fns = append(bugb.fns, fns...)
	return bugb
}

// Scan applies the group-by query and scans the result into the given value.
func (bugb *BioUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bugb.path(ctx)
	if err != nil {
		return err
	}
	bugb.sql = query
	return bugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bugb *BioUserGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (bugb *BioUserGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bugb.fields) > 1 {
		return nil, errors.New("ent: BioUserGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bugb *BioUserGroupBy) StringsX(ctx context.Context) []string {
	v, err := bugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bugb *BioUserGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bugb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{biouser.Label}
	default:
		err = fmt.Errorf("ent: BioUserGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bugb *BioUserGroupBy) StringX(ctx context.Context) string {
	v, err := bugb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (bugb *BioUserGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bugb.fields) > 1 {
		return nil, errors.New("ent: BioUserGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bugb *BioUserGroupBy) IntsX(ctx context.Context) []int {
	v, err := bugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bugb *BioUserGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bugb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{biouser.Label}
	default:
		err = fmt.Errorf("ent: BioUserGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bugb *BioUserGroupBy) IntX(ctx context.Context) int {
	v, err := bugb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (bugb *BioUserGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bugb.fields) > 1 {
		return nil, errors.New("ent: BioUserGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bugb *BioUserGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bugb *BioUserGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bugb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{biouser.Label}
	default:
		err = fmt.Errorf("ent: BioUserGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bugb *BioUserGroupBy) Float64X(ctx context.Context) float64 {
	v, err := bugb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (bugb *BioUserGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bugb.fields) > 1 {
		return nil, errors.New("ent: BioUserGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bugb *BioUserGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bugb *BioUserGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bugb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{biouser.Label}
	default:
		err = fmt.Errorf("ent: BioUserGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bugb *BioUserGroupBy) BoolX(ctx context.Context) bool {
	v, err := bugb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bugb *BioUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bugb.fields {
		if !biouser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bugb *BioUserGroupBy) sqlQuery() *sql.Selector {
	selector := bugb.sql.Select()
	aggregation := make([]string, 0, len(bugb.fns))
	for _, fn := range bugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bugb.fields)+len(bugb.fns))
		for _, f := range bugb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bugb.fields...)...)
}

// BioUserSelect is the builder for selecting fields of BioUser entities.
type BioUserSelect struct {
	*BioUserQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bus *BioUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bus.prepareQuery(ctx); err != nil {
		return err
	}
	bus.sql = bus.BioUserQuery.sqlQuery(ctx)
	return bus.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bus *BioUserSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bus.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (bus *BioUserSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bus.fields) > 1 {
		return nil, errors.New("ent: BioUserSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bus *BioUserSelect) StringsX(ctx context.Context) []string {
	v, err := bus.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (bus *BioUserSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bus.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{biouser.Label}
	default:
		err = fmt.Errorf("ent: BioUserSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bus *BioUserSelect) StringX(ctx context.Context) string {
	v, err := bus.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (bus *BioUserSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bus.fields) > 1 {
		return nil, errors.New("ent: BioUserSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bus *BioUserSelect) IntsX(ctx context.Context) []int {
	v, err := bus.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (bus *BioUserSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bus.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{biouser.Label}
	default:
		err = fmt.Errorf("ent: BioUserSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bus *BioUserSelect) IntX(ctx context.Context) int {
	v, err := bus.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (bus *BioUserSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bus.fields) > 1 {
		return nil, errors.New("ent: BioUserSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bus *BioUserSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bus.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (bus *BioUserSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bus.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{biouser.Label}
	default:
		err = fmt.Errorf("ent: BioUserSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bus *BioUserSelect) Float64X(ctx context.Context) float64 {
	v, err := bus.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (bus *BioUserSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bus.fields) > 1 {
		return nil, errors.New("ent: BioUserSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bus *BioUserSelect) BoolsX(ctx context.Context) []bool {
	v, err := bus.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (bus *BioUserSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bus.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{biouser.Label}
	default:
		err = fmt.Errorf("ent: BioUserSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bus *BioUserSelect) BoolX(ctx context.Context) bool {
	v, err := bus.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bus *BioUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bus.sql.Query()
	if err := bus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}