// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/indicator"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/scale"
)

// ScaleUpdate is the builder for updating Scale entities.
type ScaleUpdate struct {
	config
	hooks    []Hook
	mutation *ScaleMutation
}

// Where adds a new predicate for the ScaleUpdate builder.
func (su *ScaleUpdate) Where(ps ...predicate.Scale) *ScaleUpdate {
	su.mutation.predicates = append(su.mutation.predicates, ps...)
	return su
}

// SetDescription sets the "description" field.
func (su *ScaleUpdate) SetDescription(s string) *ScaleUpdate {
	su.mutation.SetDescription(s)
	return su
}

// AddIndicatorIDs adds the "indicators" edge to the Indicator entity by IDs.
func (su *ScaleUpdate) AddIndicatorIDs(ids ...int) *ScaleUpdate {
	su.mutation.AddIndicatorIDs(ids...)
	return su
}

// AddIndicators adds the "indicators" edges to the Indicator entity.
func (su *ScaleUpdate) AddIndicators(i ...*Indicator) *ScaleUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return su.AddIndicatorIDs(ids...)
}

// Mutation returns the ScaleMutation object of the builder.
func (su *ScaleUpdate) Mutation() *ScaleMutation {
	return su.mutation
}

// ClearIndicators clears all "indicators" edges to the Indicator entity.
func (su *ScaleUpdate) ClearIndicators() *ScaleUpdate {
	su.mutation.ClearIndicators()
	return su
}

// RemoveIndicatorIDs removes the "indicators" edge to Indicator entities by IDs.
func (su *ScaleUpdate) RemoveIndicatorIDs(ids ...int) *ScaleUpdate {
	su.mutation.RemoveIndicatorIDs(ids...)
	return su
}

// RemoveIndicators removes "indicators" edges to Indicator entities.
func (su *ScaleUpdate) RemoveIndicators(i ...*Indicator) *ScaleUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return su.RemoveIndicatorIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ScaleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScaleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *ScaleUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ScaleUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ScaleUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *ScaleUpdate) check() error {
	if v, ok := su.mutation.Description(); ok {
		if err := scale.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("ent: validator failed for field \"description\": %w", err)}
		}
	}
	return nil
}

func (su *ScaleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   scale.Table,
			Columns: scale.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: scale.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: scale.FieldDescription,
		})
	}
	if su.mutation.IndicatorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scale.IndicatorsTable,
			Columns: []string{scale.IndicatorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: indicator.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedIndicatorsIDs(); len(nodes) > 0 && !su.mutation.IndicatorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scale.IndicatorsTable,
			Columns: []string{scale.IndicatorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: indicator.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.IndicatorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scale.IndicatorsTable,
			Columns: []string{scale.IndicatorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: indicator.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{scale.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ScaleUpdateOne is the builder for updating a single Scale entity.
type ScaleUpdateOne struct {
	config
	hooks    []Hook
	mutation *ScaleMutation
}

// SetDescription sets the "description" field.
func (suo *ScaleUpdateOne) SetDescription(s string) *ScaleUpdateOne {
	suo.mutation.SetDescription(s)
	return suo
}

// AddIndicatorIDs adds the "indicators" edge to the Indicator entity by IDs.
func (suo *ScaleUpdateOne) AddIndicatorIDs(ids ...int) *ScaleUpdateOne {
	suo.mutation.AddIndicatorIDs(ids...)
	return suo
}

// AddIndicators adds the "indicators" edges to the Indicator entity.
func (suo *ScaleUpdateOne) AddIndicators(i ...*Indicator) *ScaleUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return suo.AddIndicatorIDs(ids...)
}

// Mutation returns the ScaleMutation object of the builder.
func (suo *ScaleUpdateOne) Mutation() *ScaleMutation {
	return suo.mutation
}

// ClearIndicators clears all "indicators" edges to the Indicator entity.
func (suo *ScaleUpdateOne) ClearIndicators() *ScaleUpdateOne {
	suo.mutation.ClearIndicators()
	return suo
}

// RemoveIndicatorIDs removes the "indicators" edge to Indicator entities by IDs.
func (suo *ScaleUpdateOne) RemoveIndicatorIDs(ids ...int) *ScaleUpdateOne {
	suo.mutation.RemoveIndicatorIDs(ids...)
	return suo
}

// RemoveIndicators removes "indicators" edges to Indicator entities.
func (suo *ScaleUpdateOne) RemoveIndicators(i ...*Indicator) *ScaleUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return suo.RemoveIndicatorIDs(ids...)
}

// Save executes the query and returns the updated Scale entity.
func (suo *ScaleUpdateOne) Save(ctx context.Context) (*Scale, error) {
	var (
		err  error
		node *Scale
	)
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScaleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ScaleUpdateOne) SaveX(ctx context.Context) *Scale {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ScaleUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ScaleUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *ScaleUpdateOne) check() error {
	if v, ok := suo.mutation.Description(); ok {
		if err := scale.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("ent: validator failed for field \"description\": %w", err)}
		}
	}
	return nil
}

func (suo *ScaleUpdateOne) sqlSave(ctx context.Context) (_node *Scale, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   scale.Table,
			Columns: scale.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: scale.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Scale.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: scale.FieldDescription,
		})
	}
	if suo.mutation.IndicatorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scale.IndicatorsTable,
			Columns: []string{scale.IndicatorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: indicator.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedIndicatorsIDs(); len(nodes) > 0 && !suo.mutation.IndicatorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scale.IndicatorsTable,
			Columns: []string{scale.IndicatorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: indicator.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.IndicatorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scale.IndicatorsTable,
			Columns: []string{scale.IndicatorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: indicator.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Scale{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{scale.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
