// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lagzi/EntSymbolBug/ent/predicate"
	"github.com/lagzi/EntSymbolBug/ent/workercontainedinformation"
	"github.com/lagzi/EntSymbolBug/ent/workernetworksettings"
)

// WorkerContainedInformationUpdate is the builder for updating WorkerContainedInformation entities.
type WorkerContainedInformationUpdate struct {
	config
	hooks    []Hook
	mutation *WorkerContainedInformationMutation
}

// Where appends a list predicates to the WorkerContainedInformationUpdate builder.
func (wciu *WorkerContainedInformationUpdate) Where(ps ...predicate.WorkerContainedInformation) *WorkerContainedInformationUpdate {
	wciu.mutation.Where(ps...)
	return wciu
}

// SetNetworkSettingsID sets the "network_settings" edge to the WorkerNetworkSettings entity by ID.
func (wciu *WorkerContainedInformationUpdate) SetNetworkSettingsID(id int) *WorkerContainedInformationUpdate {
	wciu.mutation.SetNetworkSettingsID(id)
	return wciu
}

// SetNillableNetworkSettingsID sets the "network_settings" edge to the WorkerNetworkSettings entity by ID if the given value is not nil.
func (wciu *WorkerContainedInformationUpdate) SetNillableNetworkSettingsID(id *int) *WorkerContainedInformationUpdate {
	if id != nil {
		wciu = wciu.SetNetworkSettingsID(*id)
	}
	return wciu
}

// SetNetworkSettings sets the "network_settings" edge to the WorkerNetworkSettings entity.
func (wciu *WorkerContainedInformationUpdate) SetNetworkSettings(w *WorkerNetworkSettings) *WorkerContainedInformationUpdate {
	return wciu.SetNetworkSettingsID(w.ID)
}

// Mutation returns the WorkerContainedInformationMutation object of the builder.
func (wciu *WorkerContainedInformationUpdate) Mutation() *WorkerContainedInformationMutation {
	return wciu.mutation
}

// ClearNetworkSettings clears the "network_settings" edge to the WorkerNetworkSettings entity.
func (wciu *WorkerContainedInformationUpdate) ClearNetworkSettings() *WorkerContainedInformationUpdate {
	wciu.mutation.ClearNetworkSettings()
	return wciu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wciu *WorkerContainedInformationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wciu.hooks) == 0 {
		affected, err = wciu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkerContainedInformationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wciu.mutation = mutation
			affected, err = wciu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wciu.hooks) - 1; i >= 0; i-- {
			if wciu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wciu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wciu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wciu *WorkerContainedInformationUpdate) SaveX(ctx context.Context) int {
	affected, err := wciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wciu *WorkerContainedInformationUpdate) Exec(ctx context.Context) error {
	_, err := wciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wciu *WorkerContainedInformationUpdate) ExecX(ctx context.Context) {
	if err := wciu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wciu *WorkerContainedInformationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workercontainedinformation.Table,
			Columns: workercontainedinformation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workercontainedinformation.FieldID,
			},
		},
	}
	if ps := wciu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if wciu.mutation.NetworkSettingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   workercontainedinformation.NetworkSettingsTable,
			Columns: []string{workercontainedinformation.NetworkSettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workernetworksettings.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wciu.mutation.NetworkSettingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   workercontainedinformation.NetworkSettingsTable,
			Columns: []string{workercontainedinformation.NetworkSettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workernetworksettings.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workercontainedinformation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// WorkerContainedInformationUpdateOne is the builder for updating a single WorkerContainedInformation entity.
type WorkerContainedInformationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkerContainedInformationMutation
}

// SetNetworkSettingsID sets the "network_settings" edge to the WorkerNetworkSettings entity by ID.
func (wciuo *WorkerContainedInformationUpdateOne) SetNetworkSettingsID(id int) *WorkerContainedInformationUpdateOne {
	wciuo.mutation.SetNetworkSettingsID(id)
	return wciuo
}

// SetNillableNetworkSettingsID sets the "network_settings" edge to the WorkerNetworkSettings entity by ID if the given value is not nil.
func (wciuo *WorkerContainedInformationUpdateOne) SetNillableNetworkSettingsID(id *int) *WorkerContainedInformationUpdateOne {
	if id != nil {
		wciuo = wciuo.SetNetworkSettingsID(*id)
	}
	return wciuo
}

// SetNetworkSettings sets the "network_settings" edge to the WorkerNetworkSettings entity.
func (wciuo *WorkerContainedInformationUpdateOne) SetNetworkSettings(w *WorkerNetworkSettings) *WorkerContainedInformationUpdateOne {
	return wciuo.SetNetworkSettingsID(w.ID)
}

// Mutation returns the WorkerContainedInformationMutation object of the builder.
func (wciuo *WorkerContainedInformationUpdateOne) Mutation() *WorkerContainedInformationMutation {
	return wciuo.mutation
}

// ClearNetworkSettings clears the "network_settings" edge to the WorkerNetworkSettings entity.
func (wciuo *WorkerContainedInformationUpdateOne) ClearNetworkSettings() *WorkerContainedInformationUpdateOne {
	wciuo.mutation.ClearNetworkSettings()
	return wciuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wciuo *WorkerContainedInformationUpdateOne) Select(field string, fields ...string) *WorkerContainedInformationUpdateOne {
	wciuo.fields = append([]string{field}, fields...)
	return wciuo
}

// Save executes the query and returns the updated WorkerContainedInformation entity.
func (wciuo *WorkerContainedInformationUpdateOne) Save(ctx context.Context) (*WorkerContainedInformation, error) {
	var (
		err  error
		node *WorkerContainedInformation
	)
	if len(wciuo.hooks) == 0 {
		node, err = wciuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkerContainedInformationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wciuo.mutation = mutation
			node, err = wciuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wciuo.hooks) - 1; i >= 0; i-- {
			if wciuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wciuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, wciuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*WorkerContainedInformation)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from WorkerContainedInformationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wciuo *WorkerContainedInformationUpdateOne) SaveX(ctx context.Context) *WorkerContainedInformation {
	node, err := wciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wciuo *WorkerContainedInformationUpdateOne) Exec(ctx context.Context) error {
	_, err := wciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wciuo *WorkerContainedInformationUpdateOne) ExecX(ctx context.Context) {
	if err := wciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wciuo *WorkerContainedInformationUpdateOne) sqlSave(ctx context.Context) (_node *WorkerContainedInformation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workercontainedinformation.Table,
			Columns: workercontainedinformation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workercontainedinformation.FieldID,
			},
		},
	}
	id, ok := wciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WorkerContainedInformation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workercontainedinformation.FieldID)
		for _, f := range fields {
			if !workercontainedinformation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workercontainedinformation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wciuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if wciuo.mutation.NetworkSettingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   workercontainedinformation.NetworkSettingsTable,
			Columns: []string{workercontainedinformation.NetworkSettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workernetworksettings.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wciuo.mutation.NetworkSettingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   workercontainedinformation.NetworkSettingsTable,
			Columns: []string{workercontainedinformation.NetworkSettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workernetworksettings.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WorkerContainedInformation{config: wciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workercontainedinformation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}