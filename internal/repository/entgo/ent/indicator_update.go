// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/dataset"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/indicator"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/scale"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/user"
)

// IndicatorUpdate is the builder for updating Indicator entities.
type IndicatorUpdate struct {
	config
	hooks    []Hook
	mutation *IndicatorMutation
}

// Where adds a new predicate for the IndicatorUpdate builder.
func (iu *IndicatorUpdate) Where(ps ...predicate.Indicator) *IndicatorUpdate {
	iu.mutation.predicates = append(iu.mutation.predicates, ps...)
	return iu
}

// SetCode sets the "code" field.
func (iu *IndicatorUpdate) SetCode(s string) *IndicatorUpdate {
	iu.mutation.SetCode(s)
	return iu
}

// SetTitle sets the "title" field.
func (iu *IndicatorUpdate) SetTitle(s string) *IndicatorUpdate {
	iu.mutation.SetTitle(s)
	return iu
}

// SetDescription sets the "description" field.
func (iu *IndicatorUpdate) SetDescription(s string) *IndicatorUpdate {
	iu.mutation.SetDescription(s)
	return iu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iu *IndicatorUpdate) SetNillableDescription(s *string) *IndicatorUpdate {
	if s != nil {
		iu.SetDescription(*s)
	}
	return iu
}

// ClearDescription clears the value of the "description" field.
func (iu *IndicatorUpdate) ClearDescription() *IndicatorUpdate {
	iu.mutation.ClearDescription()
	return iu
}

// SetActive sets the "active" field.
func (iu *IndicatorUpdate) SetActive(b bool) *IndicatorUpdate {
	iu.mutation.SetActive(b)
	return iu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (iu *IndicatorUpdate) SetNillableActive(b *bool) *IndicatorUpdate {
	if b != nil {
		iu.SetActive(*b)
	}
	return iu
}

// SetBuiltIn sets the "built_in" field.
func (iu *IndicatorUpdate) SetBuiltIn(b bool) *IndicatorUpdate {
	iu.mutation.SetBuiltIn(b)
	return iu
}

// SetNillableBuiltIn sets the "built_in" field if the given value is not nil.
func (iu *IndicatorUpdate) SetNillableBuiltIn(b *bool) *IndicatorUpdate {
	if b != nil {
		iu.SetBuiltIn(*b)
	}
	return iu
}

// SetExternal sets the "external" field.
func (iu *IndicatorUpdate) SetExternal(b bool) *IndicatorUpdate {
	iu.mutation.SetExternal(b)
	return iu
}

// SetNillableExternal sets the "external" field if the given value is not nil.
func (iu *IndicatorUpdate) SetNillableExternal(b *bool) *IndicatorUpdate {
	if b != nil {
		iu.SetExternal(*b)
	}
	return iu
}

// AddDatasetIDs adds the "datasets" edge to the Dataset entity by IDs.
func (iu *IndicatorUpdate) AddDatasetIDs(ids ...int) *IndicatorUpdate {
	iu.mutation.AddDatasetIDs(ids...)
	return iu
}

// AddDatasets adds the "datasets" edges to the Dataset entity.
func (iu *IndicatorUpdate) AddDatasets(d ...*Dataset) *IndicatorUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iu.AddDatasetIDs(ids...)
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (iu *IndicatorUpdate) SetAuthorID(id int) *IndicatorUpdate {
	iu.mutation.SetAuthorID(id)
	return iu
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (iu *IndicatorUpdate) SetNillableAuthorID(id *int) *IndicatorUpdate {
	if id != nil {
		iu = iu.SetAuthorID(*id)
	}
	return iu
}

// SetAuthor sets the "author" edge to the User entity.
func (iu *IndicatorUpdate) SetAuthor(u *User) *IndicatorUpdate {
	return iu.SetAuthorID(u.ID)
}

// SetScaleID sets the "scale" edge to the Scale entity by ID.
func (iu *IndicatorUpdate) SetScaleID(id int) *IndicatorUpdate {
	iu.mutation.SetScaleID(id)
	return iu
}

// SetScale sets the "scale" edge to the Scale entity.
func (iu *IndicatorUpdate) SetScale(s *Scale) *IndicatorUpdate {
	return iu.SetScaleID(s.ID)
}

// Mutation returns the IndicatorMutation object of the builder.
func (iu *IndicatorUpdate) Mutation() *IndicatorMutation {
	return iu.mutation
}

// ClearDatasets clears all "datasets" edges to the Dataset entity.
func (iu *IndicatorUpdate) ClearDatasets() *IndicatorUpdate {
	iu.mutation.ClearDatasets()
	return iu
}

// RemoveDatasetIDs removes the "datasets" edge to Dataset entities by IDs.
func (iu *IndicatorUpdate) RemoveDatasetIDs(ids ...int) *IndicatorUpdate {
	iu.mutation.RemoveDatasetIDs(ids...)
	return iu
}

// RemoveDatasets removes "datasets" edges to Dataset entities.
func (iu *IndicatorUpdate) RemoveDatasets(d ...*Dataset) *IndicatorUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iu.RemoveDatasetIDs(ids...)
}

// ClearAuthor clears the "author" edge to the User entity.
func (iu *IndicatorUpdate) ClearAuthor() *IndicatorUpdate {
	iu.mutation.ClearAuthor()
	return iu
}

// ClearScale clears the "scale" edge to the Scale entity.
func (iu *IndicatorUpdate) ClearScale() *IndicatorUpdate {
	iu.mutation.ClearScale()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IndicatorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	iu.defaults()
	if len(iu.hooks) == 0 {
		if err = iu.check(); err != nil {
			return 0, err
		}
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IndicatorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iu.check(); err != nil {
				return 0, err
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IndicatorUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IndicatorUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IndicatorUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *IndicatorUpdate) defaults() {
	if _, ok := iu.mutation.UpdateTime(); !ok {
		v := indicator.UpdateDefaultUpdateTime()
		iu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *IndicatorUpdate) check() error {
	if v, ok := iu.mutation.Code(); ok {
		if err := indicator.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf("ent: validator failed for field \"code\": %w", err)}
		}
	}
	if v, ok := iu.mutation.Title(); ok {
		if err := indicator.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if _, ok := iu.mutation.ScaleID(); iu.mutation.ScaleCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"scale\"")
	}
	return nil
}

func (iu *IndicatorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   indicator.Table,
			Columns: indicator.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: indicator.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: indicator.FieldUpdateTime,
		})
	}
	if value, ok := iu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: indicator.FieldCode,
		})
	}
	if value, ok := iu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: indicator.FieldTitle,
		})
	}
	if value, ok := iu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: indicator.FieldDescription,
		})
	}
	if iu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: indicator.FieldDescription,
		})
	}
	if value, ok := iu.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: indicator.FieldActive,
		})
	}
	if value, ok := iu.mutation.BuiltIn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: indicator.FieldBuiltIn,
		})
	}
	if value, ok := iu.mutation.External(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: indicator.FieldExternal,
		})
	}
	if iu.mutation.DatasetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   indicator.DatasetsTable,
			Columns: []string{indicator.DatasetsColumn},
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
	if nodes := iu.mutation.RemovedDatasetsIDs(); len(nodes) > 0 && !iu.mutation.DatasetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   indicator.DatasetsTable,
			Columns: []string{indicator.DatasetsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.DatasetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   indicator.DatasetsTable,
			Columns: []string{indicator.DatasetsColumn},
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
	if iu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   indicator.AuthorTable,
			Columns: []string{indicator.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   indicator.AuthorTable,
			Columns: []string{indicator.AuthorColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.ScaleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   indicator.ScaleTable,
			Columns: []string{indicator.ScaleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scale.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.ScaleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   indicator.ScaleTable,
			Columns: []string{indicator.ScaleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scale.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{indicator.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// IndicatorUpdateOne is the builder for updating a single Indicator entity.
type IndicatorUpdateOne struct {
	config
	hooks    []Hook
	mutation *IndicatorMutation
}

// SetCode sets the "code" field.
func (iuo *IndicatorUpdateOne) SetCode(s string) *IndicatorUpdateOne {
	iuo.mutation.SetCode(s)
	return iuo
}

// SetTitle sets the "title" field.
func (iuo *IndicatorUpdateOne) SetTitle(s string) *IndicatorUpdateOne {
	iuo.mutation.SetTitle(s)
	return iuo
}

// SetDescription sets the "description" field.
func (iuo *IndicatorUpdateOne) SetDescription(s string) *IndicatorUpdateOne {
	iuo.mutation.SetDescription(s)
	return iuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iuo *IndicatorUpdateOne) SetNillableDescription(s *string) *IndicatorUpdateOne {
	if s != nil {
		iuo.SetDescription(*s)
	}
	return iuo
}

// ClearDescription clears the value of the "description" field.
func (iuo *IndicatorUpdateOne) ClearDescription() *IndicatorUpdateOne {
	iuo.mutation.ClearDescription()
	return iuo
}

// SetActive sets the "active" field.
func (iuo *IndicatorUpdateOne) SetActive(b bool) *IndicatorUpdateOne {
	iuo.mutation.SetActive(b)
	return iuo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (iuo *IndicatorUpdateOne) SetNillableActive(b *bool) *IndicatorUpdateOne {
	if b != nil {
		iuo.SetActive(*b)
	}
	return iuo
}

// SetBuiltIn sets the "built_in" field.
func (iuo *IndicatorUpdateOne) SetBuiltIn(b bool) *IndicatorUpdateOne {
	iuo.mutation.SetBuiltIn(b)
	return iuo
}

// SetNillableBuiltIn sets the "built_in" field if the given value is not nil.
func (iuo *IndicatorUpdateOne) SetNillableBuiltIn(b *bool) *IndicatorUpdateOne {
	if b != nil {
		iuo.SetBuiltIn(*b)
	}
	return iuo
}

// SetExternal sets the "external" field.
func (iuo *IndicatorUpdateOne) SetExternal(b bool) *IndicatorUpdateOne {
	iuo.mutation.SetExternal(b)
	return iuo
}

// SetNillableExternal sets the "external" field if the given value is not nil.
func (iuo *IndicatorUpdateOne) SetNillableExternal(b *bool) *IndicatorUpdateOne {
	if b != nil {
		iuo.SetExternal(*b)
	}
	return iuo
}

// AddDatasetIDs adds the "datasets" edge to the Dataset entity by IDs.
func (iuo *IndicatorUpdateOne) AddDatasetIDs(ids ...int) *IndicatorUpdateOne {
	iuo.mutation.AddDatasetIDs(ids...)
	return iuo
}

// AddDatasets adds the "datasets" edges to the Dataset entity.
func (iuo *IndicatorUpdateOne) AddDatasets(d ...*Dataset) *IndicatorUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iuo.AddDatasetIDs(ids...)
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (iuo *IndicatorUpdateOne) SetAuthorID(id int) *IndicatorUpdateOne {
	iuo.mutation.SetAuthorID(id)
	return iuo
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (iuo *IndicatorUpdateOne) SetNillableAuthorID(id *int) *IndicatorUpdateOne {
	if id != nil {
		iuo = iuo.SetAuthorID(*id)
	}
	return iuo
}

// SetAuthor sets the "author" edge to the User entity.
func (iuo *IndicatorUpdateOne) SetAuthor(u *User) *IndicatorUpdateOne {
	return iuo.SetAuthorID(u.ID)
}

// SetScaleID sets the "scale" edge to the Scale entity by ID.
func (iuo *IndicatorUpdateOne) SetScaleID(id int) *IndicatorUpdateOne {
	iuo.mutation.SetScaleID(id)
	return iuo
}

// SetScale sets the "scale" edge to the Scale entity.
func (iuo *IndicatorUpdateOne) SetScale(s *Scale) *IndicatorUpdateOne {
	return iuo.SetScaleID(s.ID)
}

// Mutation returns the IndicatorMutation object of the builder.
func (iuo *IndicatorUpdateOne) Mutation() *IndicatorMutation {
	return iuo.mutation
}

// ClearDatasets clears all "datasets" edges to the Dataset entity.
func (iuo *IndicatorUpdateOne) ClearDatasets() *IndicatorUpdateOne {
	iuo.mutation.ClearDatasets()
	return iuo
}

// RemoveDatasetIDs removes the "datasets" edge to Dataset entities by IDs.
func (iuo *IndicatorUpdateOne) RemoveDatasetIDs(ids ...int) *IndicatorUpdateOne {
	iuo.mutation.RemoveDatasetIDs(ids...)
	return iuo
}

// RemoveDatasets removes "datasets" edges to Dataset entities.
func (iuo *IndicatorUpdateOne) RemoveDatasets(d ...*Dataset) *IndicatorUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iuo.RemoveDatasetIDs(ids...)
}

// ClearAuthor clears the "author" edge to the User entity.
func (iuo *IndicatorUpdateOne) ClearAuthor() *IndicatorUpdateOne {
	iuo.mutation.ClearAuthor()
	return iuo
}

// ClearScale clears the "scale" edge to the Scale entity.
func (iuo *IndicatorUpdateOne) ClearScale() *IndicatorUpdateOne {
	iuo.mutation.ClearScale()
	return iuo
}

// Save executes the query and returns the updated Indicator entity.
func (iuo *IndicatorUpdateOne) Save(ctx context.Context) (*Indicator, error) {
	var (
		err  error
		node *Indicator
	)
	iuo.defaults()
	if len(iuo.hooks) == 0 {
		if err = iuo.check(); err != nil {
			return nil, err
		}
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IndicatorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iuo.check(); err != nil {
				return nil, err
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IndicatorUpdateOne) SaveX(ctx context.Context) *Indicator {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IndicatorUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IndicatorUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *IndicatorUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdateTime(); !ok {
		v := indicator.UpdateDefaultUpdateTime()
		iuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *IndicatorUpdateOne) check() error {
	if v, ok := iuo.mutation.Code(); ok {
		if err := indicator.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf("ent: validator failed for field \"code\": %w", err)}
		}
	}
	if v, ok := iuo.mutation.Title(); ok {
		if err := indicator.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if _, ok := iuo.mutation.ScaleID(); iuo.mutation.ScaleCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"scale\"")
	}
	return nil
}

func (iuo *IndicatorUpdateOne) sqlSave(ctx context.Context) (_node *Indicator, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   indicator.Table,
			Columns: indicator.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: indicator.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Indicator.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: indicator.FieldUpdateTime,
		})
	}
	if value, ok := iuo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: indicator.FieldCode,
		})
	}
	if value, ok := iuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: indicator.FieldTitle,
		})
	}
	if value, ok := iuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: indicator.FieldDescription,
		})
	}
	if iuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: indicator.FieldDescription,
		})
	}
	if value, ok := iuo.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: indicator.FieldActive,
		})
	}
	if value, ok := iuo.mutation.BuiltIn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: indicator.FieldBuiltIn,
		})
	}
	if value, ok := iuo.mutation.External(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: indicator.FieldExternal,
		})
	}
	if iuo.mutation.DatasetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   indicator.DatasetsTable,
			Columns: []string{indicator.DatasetsColumn},
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
	if nodes := iuo.mutation.RemovedDatasetsIDs(); len(nodes) > 0 && !iuo.mutation.DatasetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   indicator.DatasetsTable,
			Columns: []string{indicator.DatasetsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.DatasetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   indicator.DatasetsTable,
			Columns: []string{indicator.DatasetsColumn},
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
	if iuo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   indicator.AuthorTable,
			Columns: []string{indicator.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   indicator.AuthorTable,
			Columns: []string{indicator.AuthorColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.ScaleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   indicator.ScaleTable,
			Columns: []string{indicator.ScaleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scale.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.ScaleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   indicator.ScaleTable,
			Columns: []string{indicator.ScaleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scale.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Indicator{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{indicator.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
