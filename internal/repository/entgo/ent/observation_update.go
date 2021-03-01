// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/dataset"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/observation"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/predicate"
)

// ObservationUpdate is the builder for updating Observation entities.
type ObservationUpdate struct {
	config
	hooks    []Hook
	mutation *ObservationMutation
}

// Where adds a new predicate for the ObservationUpdate builder.
func (ou *ObservationUpdate) Where(ps ...predicate.Observation) *ObservationUpdate {
	ou.mutation.predicates = append(ou.mutation.predicates, ps...)
	return ou
}

// SetValue sets the "value" field.
func (ou *ObservationUpdate) SetValue(f float64) *ObservationUpdate {
	ou.mutation.ResetValue()
	ou.mutation.SetValue(f)
	return ou
}

// AddValue adds f to the "value" field.
func (ou *ObservationUpdate) AddValue(f float64) *ObservationUpdate {
	ou.mutation.AddValue(f)
	return ou
}

// SetDate sets the "date" field.
func (ou *ObservationUpdate) SetDate(t time.Time) *ObservationUpdate {
	ou.mutation.SetDate(t)
	return ou
}

// SetDatasetID sets the "dataset" edge to the Dataset entity by ID.
func (ou *ObservationUpdate) SetDatasetID(id int) *ObservationUpdate {
	ou.mutation.SetDatasetID(id)
	return ou
}

// SetDataset sets the "dataset" edge to the Dataset entity.
func (ou *ObservationUpdate) SetDataset(d *Dataset) *ObservationUpdate {
	return ou.SetDatasetID(d.ID)
}

// Mutation returns the ObservationMutation object of the builder.
func (ou *ObservationUpdate) Mutation() *ObservationMutation {
	return ou.mutation
}

// ClearDataset clears the "dataset" edge to the Dataset entity.
func (ou *ObservationUpdate) ClearDataset() *ObservationUpdate {
	ou.mutation.ClearDataset()
	return ou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *ObservationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ou.defaults()
	if len(ou.hooks) == 0 {
		if err = ou.check(); err != nil {
			return 0, err
		}
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ObservationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ou.check(); err != nil {
				return 0, err
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *ObservationUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *ObservationUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *ObservationUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *ObservationUpdate) defaults() {
	if _, ok := ou.mutation.UpdateTime(); !ok {
		v := observation.UpdateDefaultUpdateTime()
		ou.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *ObservationUpdate) check() error {
	if _, ok := ou.mutation.DatasetID(); ou.mutation.DatasetCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"dataset\"")
	}
	return nil
}

func (ou *ObservationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   observation.Table,
			Columns: observation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: observation.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: observation.FieldUpdateTime,
		})
	}
	if value, ok := ou.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: observation.FieldValue,
		})
	}
	if value, ok := ou.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: observation.FieldValue,
		})
	}
	if value, ok := ou.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: observation.FieldDate,
		})
	}
	if ou.mutation.DatasetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   observation.DatasetTable,
			Columns: []string{observation.DatasetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dataset.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.DatasetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   observation.DatasetTable,
			Columns: []string{observation.DatasetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dataset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{observation.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ObservationUpdateOne is the builder for updating a single Observation entity.
type ObservationUpdateOne struct {
	config
	hooks    []Hook
	mutation *ObservationMutation
}

// SetValue sets the "value" field.
func (ouo *ObservationUpdateOne) SetValue(f float64) *ObservationUpdateOne {
	ouo.mutation.ResetValue()
	ouo.mutation.SetValue(f)
	return ouo
}

// AddValue adds f to the "value" field.
func (ouo *ObservationUpdateOne) AddValue(f float64) *ObservationUpdateOne {
	ouo.mutation.AddValue(f)
	return ouo
}

// SetDate sets the "date" field.
func (ouo *ObservationUpdateOne) SetDate(t time.Time) *ObservationUpdateOne {
	ouo.mutation.SetDate(t)
	return ouo
}

// SetDatasetID sets the "dataset" edge to the Dataset entity by ID.
func (ouo *ObservationUpdateOne) SetDatasetID(id int) *ObservationUpdateOne {
	ouo.mutation.SetDatasetID(id)
	return ouo
}

// SetDataset sets the "dataset" edge to the Dataset entity.
func (ouo *ObservationUpdateOne) SetDataset(d *Dataset) *ObservationUpdateOne {
	return ouo.SetDatasetID(d.ID)
}

// Mutation returns the ObservationMutation object of the builder.
func (ouo *ObservationUpdateOne) Mutation() *ObservationMutation {
	return ouo.mutation
}

// ClearDataset clears the "dataset" edge to the Dataset entity.
func (ouo *ObservationUpdateOne) ClearDataset() *ObservationUpdateOne {
	ouo.mutation.ClearDataset()
	return ouo
}

// Save executes the query and returns the updated Observation entity.
func (ouo *ObservationUpdateOne) Save(ctx context.Context) (*Observation, error) {
	var (
		err  error
		node *Observation
	)
	ouo.defaults()
	if len(ouo.hooks) == 0 {
		if err = ouo.check(); err != nil {
			return nil, err
		}
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ObservationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ouo.check(); err != nil {
				return nil, err
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *ObservationUpdateOne) SaveX(ctx context.Context) *Observation {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *ObservationUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *ObservationUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *ObservationUpdateOne) defaults() {
	if _, ok := ouo.mutation.UpdateTime(); !ok {
		v := observation.UpdateDefaultUpdateTime()
		ouo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *ObservationUpdateOne) check() error {
	if _, ok := ouo.mutation.DatasetID(); ouo.mutation.DatasetCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"dataset\"")
	}
	return nil
}

func (ouo *ObservationUpdateOne) sqlSave(ctx context.Context) (_node *Observation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   observation.Table,
			Columns: observation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: observation.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Observation.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: observation.FieldUpdateTime,
		})
	}
	if value, ok := ouo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: observation.FieldValue,
		})
	}
	if value, ok := ouo.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: observation.FieldValue,
		})
	}
	if value, ok := ouo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: observation.FieldDate,
		})
	}
	if ouo.mutation.DatasetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   observation.DatasetTable,
			Columns: []string{observation.DatasetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dataset.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.DatasetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   observation.DatasetTable,
			Columns: []string{observation.DatasetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dataset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Observation{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{observation.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
