// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/indicator"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/indicatorparams"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/predicate"
)

// IndicatorParamsUpdate is the builder for updating IndicatorParams entities.
type IndicatorParamsUpdate struct {
	config
	hooks    []Hook
	mutation *IndicatorParamsMutation
}

// Where adds a new predicate for the IndicatorParamsUpdate builder.
func (ipu *IndicatorParamsUpdate) Where(ps ...predicate.IndicatorParams) *IndicatorParamsUpdate {
	ipu.mutation.predicates = append(ipu.mutation.predicates, ps...)
	return ipu
}

// SetValueMapping sets the "value_mapping" field.
func (ipu *IndicatorParamsUpdate) SetValueMapping(m map[string]string) *IndicatorParamsUpdate {
	ipu.mutation.SetValueMapping(m)
	return ipu
}

// ClearValueMapping clears the value of the "value_mapping" field.
func (ipu *IndicatorParamsUpdate) ClearValueMapping() *IndicatorParamsUpdate {
	ipu.mutation.ClearValueMapping()
	return ipu
}

// SetValueParams sets the "value_params" field.
func (ipu *IndicatorParamsUpdate) SetValueParams(dvp domain.IndicatorValueParams) *IndicatorParamsUpdate {
	ipu.mutation.SetValueParams(dvp)
	return ipu
}

// SetIndicatorID sets the "indicator" edge to the Indicator entity by ID.
func (ipu *IndicatorParamsUpdate) SetIndicatorID(id int) *IndicatorParamsUpdate {
	ipu.mutation.SetIndicatorID(id)
	return ipu
}

// SetIndicator sets the "indicator" edge to the Indicator entity.
func (ipu *IndicatorParamsUpdate) SetIndicator(i *Indicator) *IndicatorParamsUpdate {
	return ipu.SetIndicatorID(i.ID)
}

// Mutation returns the IndicatorParamsMutation object of the builder.
func (ipu *IndicatorParamsUpdate) Mutation() *IndicatorParamsMutation {
	return ipu.mutation
}

// ClearIndicator clears the "indicator" edge to the Indicator entity.
func (ipu *IndicatorParamsUpdate) ClearIndicator() *IndicatorParamsUpdate {
	ipu.mutation.ClearIndicator()
	return ipu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ipu *IndicatorParamsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ipu.hooks) == 0 {
		if err = ipu.check(); err != nil {
			return 0, err
		}
		affected, err = ipu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IndicatorParamsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ipu.check(); err != nil {
				return 0, err
			}
			ipu.mutation = mutation
			affected, err = ipu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ipu.hooks) - 1; i >= 0; i-- {
			mut = ipu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ipu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ipu *IndicatorParamsUpdate) SaveX(ctx context.Context) int {
	affected, err := ipu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ipu *IndicatorParamsUpdate) Exec(ctx context.Context) error {
	_, err := ipu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ipu *IndicatorParamsUpdate) ExecX(ctx context.Context) {
	if err := ipu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ipu *IndicatorParamsUpdate) check() error {
	if _, ok := ipu.mutation.IndicatorID(); ipu.mutation.IndicatorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"indicator\"")
	}
	return nil
}

func (ipu *IndicatorParamsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   indicatorparams.Table,
			Columns: indicatorparams.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: indicatorparams.FieldID,
			},
		},
	}
	if ps := ipu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ipu.mutation.ValueMapping(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: indicatorparams.FieldValueMapping,
		})
	}
	if ipu.mutation.ValueMappingCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: indicatorparams.FieldValueMapping,
		})
	}
	if value, ok := ipu.mutation.ValueParams(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: indicatorparams.FieldValueParams,
		})
	}
	if ipu.mutation.IndicatorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   indicatorparams.IndicatorTable,
			Columns: []string{indicatorparams.IndicatorColumn},
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
	if nodes := ipu.mutation.IndicatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   indicatorparams.IndicatorTable,
			Columns: []string{indicatorparams.IndicatorColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ipu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{indicatorparams.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// IndicatorParamsUpdateOne is the builder for updating a single IndicatorParams entity.
type IndicatorParamsUpdateOne struct {
	config
	hooks    []Hook
	mutation *IndicatorParamsMutation
}

// SetValueMapping sets the "value_mapping" field.
func (ipuo *IndicatorParamsUpdateOne) SetValueMapping(m map[string]string) *IndicatorParamsUpdateOne {
	ipuo.mutation.SetValueMapping(m)
	return ipuo
}

// ClearValueMapping clears the value of the "value_mapping" field.
func (ipuo *IndicatorParamsUpdateOne) ClearValueMapping() *IndicatorParamsUpdateOne {
	ipuo.mutation.ClearValueMapping()
	return ipuo
}

// SetValueParams sets the "value_params" field.
func (ipuo *IndicatorParamsUpdateOne) SetValueParams(dvp domain.IndicatorValueParams) *IndicatorParamsUpdateOne {
	ipuo.mutation.SetValueParams(dvp)
	return ipuo
}

// SetIndicatorID sets the "indicator" edge to the Indicator entity by ID.
func (ipuo *IndicatorParamsUpdateOne) SetIndicatorID(id int) *IndicatorParamsUpdateOne {
	ipuo.mutation.SetIndicatorID(id)
	return ipuo
}

// SetIndicator sets the "indicator" edge to the Indicator entity.
func (ipuo *IndicatorParamsUpdateOne) SetIndicator(i *Indicator) *IndicatorParamsUpdateOne {
	return ipuo.SetIndicatorID(i.ID)
}

// Mutation returns the IndicatorParamsMutation object of the builder.
func (ipuo *IndicatorParamsUpdateOne) Mutation() *IndicatorParamsMutation {
	return ipuo.mutation
}

// ClearIndicator clears the "indicator" edge to the Indicator entity.
func (ipuo *IndicatorParamsUpdateOne) ClearIndicator() *IndicatorParamsUpdateOne {
	ipuo.mutation.ClearIndicator()
	return ipuo
}

// Save executes the query and returns the updated IndicatorParams entity.
func (ipuo *IndicatorParamsUpdateOne) Save(ctx context.Context) (*IndicatorParams, error) {
	var (
		err  error
		node *IndicatorParams
	)
	if len(ipuo.hooks) == 0 {
		if err = ipuo.check(); err != nil {
			return nil, err
		}
		node, err = ipuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IndicatorParamsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ipuo.check(); err != nil {
				return nil, err
			}
			ipuo.mutation = mutation
			node, err = ipuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ipuo.hooks) - 1; i >= 0; i-- {
			mut = ipuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ipuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ipuo *IndicatorParamsUpdateOne) SaveX(ctx context.Context) *IndicatorParams {
	node, err := ipuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ipuo *IndicatorParamsUpdateOne) Exec(ctx context.Context) error {
	_, err := ipuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ipuo *IndicatorParamsUpdateOne) ExecX(ctx context.Context) {
	if err := ipuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ipuo *IndicatorParamsUpdateOne) check() error {
	if _, ok := ipuo.mutation.IndicatorID(); ipuo.mutation.IndicatorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"indicator\"")
	}
	return nil
}

func (ipuo *IndicatorParamsUpdateOne) sqlSave(ctx context.Context) (_node *IndicatorParams, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   indicatorparams.Table,
			Columns: indicatorparams.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: indicatorparams.FieldID,
			},
		},
	}
	id, ok := ipuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing IndicatorParams.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := ipuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ipuo.mutation.ValueMapping(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: indicatorparams.FieldValueMapping,
		})
	}
	if ipuo.mutation.ValueMappingCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: indicatorparams.FieldValueMapping,
		})
	}
	if value, ok := ipuo.mutation.ValueParams(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: indicatorparams.FieldValueParams,
		})
	}
	if ipuo.mutation.IndicatorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   indicatorparams.IndicatorTable,
			Columns: []string{indicatorparams.IndicatorColumn},
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
	if nodes := ipuo.mutation.IndicatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   indicatorparams.IndicatorTable,
			Columns: []string{indicatorparams.IndicatorColumn},
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
	_node = &IndicatorParams{config: ipuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ipuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{indicatorparams.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
