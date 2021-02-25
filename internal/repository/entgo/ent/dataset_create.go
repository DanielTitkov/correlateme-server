// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/correlation"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/dataset"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/indicator"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/observation"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/user"
)

// DatasetCreate is the builder for creating a Dataset entity.
type DatasetCreate struct {
	config
	mutation *DatasetMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (dc *DatasetCreate) SetCreateTime(t time.Time) *DatasetCreate {
	dc.mutation.SetCreateTime(t)
	return dc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (dc *DatasetCreate) SetNillableCreateTime(t *time.Time) *DatasetCreate {
	if t != nil {
		dc.SetCreateTime(*t)
	}
	return dc
}

// SetUpdateTime sets the "update_time" field.
func (dc *DatasetCreate) SetUpdateTime(t time.Time) *DatasetCreate {
	dc.mutation.SetUpdateTime(t)
	return dc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (dc *DatasetCreate) SetNillableUpdateTime(t *time.Time) *DatasetCreate {
	if t != nil {
		dc.SetUpdateTime(*t)
	}
	return dc
}

// SetShared sets the "shared" field.
func (dc *DatasetCreate) SetShared(b bool) *DatasetCreate {
	dc.mutation.SetShared(b)
	return dc
}

// SetNillableShared sets the "shared" field if the given value is not nil.
func (dc *DatasetCreate) SetNillableShared(b *bool) *DatasetCreate {
	if b != nil {
		dc.SetShared(*b)
	}
	return dc
}

// SetSource sets the "source" field.
func (dc *DatasetCreate) SetSource(s string) *DatasetCreate {
	dc.mutation.SetSource(s)
	return dc
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (dc *DatasetCreate) SetNillableSource(s *string) *DatasetCreate {
	if s != nil {
		dc.SetSource(*s)
	}
	return dc
}

// AddLeftIDs adds the "left" edge to the Correlation entity by IDs.
func (dc *DatasetCreate) AddLeftIDs(ids ...int) *DatasetCreate {
	dc.mutation.AddLeftIDs(ids...)
	return dc
}

// AddLeft adds the "left" edges to the Correlation entity.
func (dc *DatasetCreate) AddLeft(c ...*Correlation) *DatasetCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return dc.AddLeftIDs(ids...)
}

// AddRightIDs adds the "right" edge to the Correlation entity by IDs.
func (dc *DatasetCreate) AddRightIDs(ids ...int) *DatasetCreate {
	dc.mutation.AddRightIDs(ids...)
	return dc
}

// AddRight adds the "right" edges to the Correlation entity.
func (dc *DatasetCreate) AddRight(c ...*Correlation) *DatasetCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return dc.AddRightIDs(ids...)
}

// AddObservationIDs adds the "observations" edge to the Observation entity by IDs.
func (dc *DatasetCreate) AddObservationIDs(ids ...int) *DatasetCreate {
	dc.mutation.AddObservationIDs(ids...)
	return dc
}

// AddObservations adds the "observations" edges to the Observation entity.
func (dc *DatasetCreate) AddObservations(o ...*Observation) *DatasetCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return dc.AddObservationIDs(ids...)
}

// SetIndicatorID sets the "indicator" edge to the Indicator entity by ID.
func (dc *DatasetCreate) SetIndicatorID(id int) *DatasetCreate {
	dc.mutation.SetIndicatorID(id)
	return dc
}

// SetIndicator sets the "indicator" edge to the Indicator entity.
func (dc *DatasetCreate) SetIndicator(i *Indicator) *DatasetCreate {
	return dc.SetIndicatorID(i.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (dc *DatasetCreate) SetUserID(id int) *DatasetCreate {
	dc.mutation.SetUserID(id)
	return dc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (dc *DatasetCreate) SetNillableUserID(id *int) *DatasetCreate {
	if id != nil {
		dc = dc.SetUserID(*id)
	}
	return dc
}

// SetUser sets the "user" edge to the User entity.
func (dc *DatasetCreate) SetUser(u *User) *DatasetCreate {
	return dc.SetUserID(u.ID)
}

// Mutation returns the DatasetMutation object of the builder.
func (dc *DatasetCreate) Mutation() *DatasetMutation {
	return dc.mutation
}

// Save creates the Dataset in the database.
func (dc *DatasetCreate) Save(ctx context.Context) (*Dataset, error) {
	var (
		err  error
		node *Dataset
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DatasetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DatasetCreate) SaveX(ctx context.Context) *Dataset {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (dc *DatasetCreate) defaults() {
	if _, ok := dc.mutation.CreateTime(); !ok {
		v := dataset.DefaultCreateTime()
		dc.mutation.SetCreateTime(v)
	}
	if _, ok := dc.mutation.UpdateTime(); !ok {
		v := dataset.DefaultUpdateTime()
		dc.mutation.SetUpdateTime(v)
	}
	if _, ok := dc.mutation.Shared(); !ok {
		v := dataset.DefaultShared
		dc.mutation.SetShared(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DatasetCreate) check() error {
	if _, ok := dc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := dc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := dc.mutation.Shared(); !ok {
		return &ValidationError{Name: "shared", err: errors.New("ent: missing required field \"shared\"")}
	}
	if _, ok := dc.mutation.IndicatorID(); !ok {
		return &ValidationError{Name: "indicator", err: errors.New("ent: missing required edge \"indicator\"")}
	}
	return nil
}

func (dc *DatasetCreate) sqlSave(ctx context.Context) (*Dataset, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dc *DatasetCreate) createSpec() (*Dataset, *sqlgraph.CreateSpec) {
	var (
		_node = &Dataset{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dataset.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dataset.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dataset.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := dc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dataset.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := dc.mutation.Shared(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: dataset.FieldShared,
		})
		_node.Shared = value
	}
	if value, ok := dc.mutation.Source(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dataset.FieldSource,
		})
		_node.Source = &value
	}
	if nodes := dc.mutation.LeftIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dataset.LeftTable,
			Columns: []string{dataset.LeftColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: correlation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.RightIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dataset.RightTable,
			Columns: []string{dataset.RightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: correlation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ObservationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dataset.ObservationsTable,
			Columns: []string{dataset.ObservationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: observation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.IndicatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dataset.IndicatorTable,
			Columns: []string{dataset.IndicatorColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dataset.UserTable,
			Columns: []string{dataset.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
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

// DatasetCreateBulk is the builder for creating many Dataset entities in bulk.
type DatasetCreateBulk struct {
	config
	builders []*DatasetCreate
}

// Save creates the Dataset entities in the database.
func (dcb *DatasetCreateBulk) Save(ctx context.Context) ([]*Dataset, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Dataset, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DatasetMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DatasetCreateBulk) SaveX(ctx context.Context) []*Dataset {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
