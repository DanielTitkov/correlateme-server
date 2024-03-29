// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/dataset"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/observation"
)

// ObservationCreate is the builder for creating a Observation entity.
type ObservationCreate struct {
	config
	mutation *ObservationMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (oc *ObservationCreate) SetCreateTime(t time.Time) *ObservationCreate {
	oc.mutation.SetCreateTime(t)
	return oc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (oc *ObservationCreate) SetNillableCreateTime(t *time.Time) *ObservationCreate {
	if t != nil {
		oc.SetCreateTime(*t)
	}
	return oc
}

// SetUpdateTime sets the "update_time" field.
func (oc *ObservationCreate) SetUpdateTime(t time.Time) *ObservationCreate {
	oc.mutation.SetUpdateTime(t)
	return oc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (oc *ObservationCreate) SetNillableUpdateTime(t *time.Time) *ObservationCreate {
	if t != nil {
		oc.SetUpdateTime(*t)
	}
	return oc
}

// SetValue sets the "value" field.
func (oc *ObservationCreate) SetValue(f float64) *ObservationCreate {
	oc.mutation.SetValue(f)
	return oc
}

// SetDate sets the "date" field.
func (oc *ObservationCreate) SetDate(t time.Time) *ObservationCreate {
	oc.mutation.SetDate(t)
	return oc
}

// SetGranularity sets the "granularity" field.
func (oc *ObservationCreate) SetGranularity(o observation.Granularity) *ObservationCreate {
	oc.mutation.SetGranularity(o)
	return oc
}

// SetNillableGranularity sets the "granularity" field if the given value is not nil.
func (oc *ObservationCreate) SetNillableGranularity(o *observation.Granularity) *ObservationCreate {
	if o != nil {
		oc.SetGranularity(*o)
	}
	return oc
}

// SetDatasetID sets the "dataset" edge to the Dataset entity by ID.
func (oc *ObservationCreate) SetDatasetID(id int) *ObservationCreate {
	oc.mutation.SetDatasetID(id)
	return oc
}

// SetDataset sets the "dataset" edge to the Dataset entity.
func (oc *ObservationCreate) SetDataset(d *Dataset) *ObservationCreate {
	return oc.SetDatasetID(d.ID)
}

// Mutation returns the ObservationMutation object of the builder.
func (oc *ObservationCreate) Mutation() *ObservationMutation {
	return oc.mutation
}

// Save creates the Observation in the database.
func (oc *ObservationCreate) Save(ctx context.Context) (*Observation, error) {
	var (
		err  error
		node *Observation
	)
	oc.defaults()
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ObservationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			node, err = oc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *ObservationCreate) SaveX(ctx context.Context) *Observation {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (oc *ObservationCreate) defaults() {
	if _, ok := oc.mutation.CreateTime(); !ok {
		v := observation.DefaultCreateTime()
		oc.mutation.SetCreateTime(v)
	}
	if _, ok := oc.mutation.UpdateTime(); !ok {
		v := observation.DefaultUpdateTime()
		oc.mutation.SetUpdateTime(v)
	}
	if _, ok := oc.mutation.Granularity(); !ok {
		v := observation.DefaultGranularity
		oc.mutation.SetGranularity(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *ObservationCreate) check() error {
	if _, ok := oc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := oc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := oc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New("ent: missing required field \"value\"")}
	}
	if _, ok := oc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New("ent: missing required field \"date\"")}
	}
	if _, ok := oc.mutation.Granularity(); !ok {
		return &ValidationError{Name: "granularity", err: errors.New("ent: missing required field \"granularity\"")}
	}
	if v, ok := oc.mutation.Granularity(); ok {
		if err := observation.GranularityValidator(v); err != nil {
			return &ValidationError{Name: "granularity", err: fmt.Errorf("ent: validator failed for field \"granularity\": %w", err)}
		}
	}
	if _, ok := oc.mutation.DatasetID(); !ok {
		return &ValidationError{Name: "dataset", err: errors.New("ent: missing required edge \"dataset\"")}
	}
	return nil
}

func (oc *ObservationCreate) sqlSave(ctx context.Context) (*Observation, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oc *ObservationCreate) createSpec() (*Observation, *sqlgraph.CreateSpec) {
	var (
		_node = &Observation{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: observation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: observation.FieldID,
			},
		}
	)
	if value, ok := oc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: observation.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := oc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: observation.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := oc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: observation.FieldValue,
		})
		_node.Value = value
	}
	if value, ok := oc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: observation.FieldDate,
		})
		_node.Date = value
	}
	if value, ok := oc.mutation.Granularity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: observation.FieldGranularity,
		})
		_node.Granularity = value
	}
	if nodes := oc.mutation.DatasetIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ObservationCreateBulk is the builder for creating many Observation entities in bulk.
type ObservationCreateBulk struct {
	config
	builders []*ObservationCreate
}

// Save creates the Observation entities in the database.
func (ocb *ObservationCreateBulk) Save(ctx context.Context) ([]*Observation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Observation, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ObservationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *ObservationCreateBulk) SaveX(ctx context.Context) []*Observation {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
