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

// WorkerNetworkSettingsUpdate is the builder for updating WorkerNetworkSettings entities.
type WorkerNetworkSettingsUpdate struct {
	config
	hooks    []Hook
	mutation *WorkerNetworkSettingsMutation
}

// Where appends a list predicates to the WorkerNetworkSettingsUpdate builder.
func (wnsu *WorkerNetworkSettingsUpdate) Where(ps ...predicate.WorkerNetworkSettings) *WorkerNetworkSettingsUpdate {
	wnsu.mutation.Where(ps...)
	return wnsu
}

// SetWorkerContainedInformationID sets the "worker_contained_information" edge to the WorkerContainedInformation entity by ID.
func (wnsu *WorkerNetworkSettingsUpdate) SetWorkerContainedInformationID(id int) *WorkerNetworkSettingsUpdate {
	wnsu.mutation.SetWorkerContainedInformationID(id)
	return wnsu
}

// SetNillableWorkerContainedInformationID sets the "worker_contained_information" edge to the WorkerContainedInformation entity by ID if the given value is not nil.
func (wnsu *WorkerNetworkSettingsUpdate) SetNillableWorkerContainedInformationID(id *int) *WorkerNetworkSettingsUpdate {
	if id != nil {
		wnsu = wnsu.SetWorkerContainedInformationID(*id)
	}
	return wnsu
}

// SetWorkerContainedInformation sets the "worker_contained_information" edge to the WorkerContainedInformation entity.
func (wnsu *WorkerNetworkSettingsUpdate) SetWorkerContainedInformation(w *WorkerContainedInformation) *WorkerNetworkSettingsUpdate {
	return wnsu.SetWorkerContainedInformationID(w.ID)
}

// Mutation returns the WorkerNetworkSettingsMutation object of the builder.
func (wnsu *WorkerNetworkSettingsUpdate) Mutation() *WorkerNetworkSettingsMutation {
	return wnsu.mutation
}

// ClearWorkerContainedInformation clears the "worker_contained_information" edge to the WorkerContainedInformation entity.
func (wnsu *WorkerNetworkSettingsUpdate) ClearWorkerContainedInformation() *WorkerNetworkSettingsUpdate {
	wnsu.mutation.ClearWorkerContainedInformation()
	return wnsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wnsu *WorkerNetworkSettingsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wnsu.hooks) == 0 {
		affected, err = wnsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkerNetworkSettingsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wnsu.mutation = mutation
			affected, err = wnsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wnsu.hooks) - 1; i >= 0; i-- {
			if wnsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wnsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wnsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wnsu *WorkerNetworkSettingsUpdate) SaveX(ctx context.Context) int {
	affected, err := wnsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wnsu *WorkerNetworkSettingsUpdate) Exec(ctx context.Context) error {
	_, err := wnsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wnsu *WorkerNetworkSettingsUpdate) ExecX(ctx context.Context) {
	if err := wnsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wnsu *WorkerNetworkSettingsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workernetworksettings.Table,
			Columns: workernetworksettings.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workernetworksettings.FieldID,
			},
		},
	}
	if ps := wnsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if wnsu.mutation.WorkerContainedInformationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   workernetworksettings.WorkerContainedInformationTable,
			Columns: []string{workernetworksettings.WorkerContainedInformationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workercontainedinformation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wnsu.mutation.WorkerContainedInformationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   workernetworksettings.WorkerContainedInformationTable,
			Columns: []string{workernetworksettings.WorkerContainedInformationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workercontainedinformation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wnsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workernetworksettings.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// WorkerNetworkSettingsUpdateOne is the builder for updating a single WorkerNetworkSettings entity.
type WorkerNetworkSettingsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkerNetworkSettingsMutation
}

// SetWorkerContainedInformationID sets the "worker_contained_information" edge to the WorkerContainedInformation entity by ID.
func (wnsuo *WorkerNetworkSettingsUpdateOne) SetWorkerContainedInformationID(id int) *WorkerNetworkSettingsUpdateOne {
	wnsuo.mutation.SetWorkerContainedInformationID(id)
	return wnsuo
}

// SetNillableWorkerContainedInformationID sets the "worker_contained_information" edge to the WorkerContainedInformation entity by ID if the given value is not nil.
func (wnsuo *WorkerNetworkSettingsUpdateOne) SetNillableWorkerContainedInformationID(id *int) *WorkerNetworkSettingsUpdateOne {
	if id != nil {
		wnsuo = wnsuo.SetWorkerContainedInformationID(*id)
	}
	return wnsuo
}

// SetWorkerContainedInformation sets the "worker_contained_information" edge to the WorkerContainedInformation entity.
func (wnsuo *WorkerNetworkSettingsUpdateOne) SetWorkerContainedInformation(w *WorkerContainedInformation) *WorkerNetworkSettingsUpdateOne {
	return wnsuo.SetWorkerContainedInformationID(w.ID)
}

// Mutation returns the WorkerNetworkSettingsMutation object of the builder.
func (wnsuo *WorkerNetworkSettingsUpdateOne) Mutation() *WorkerNetworkSettingsMutation {
	return wnsuo.mutation
}

// ClearWorkerContainedInformation clears the "worker_contained_information" edge to the WorkerContainedInformation entity.
func (wnsuo *WorkerNetworkSettingsUpdateOne) ClearWorkerContainedInformation() *WorkerNetworkSettingsUpdateOne {
	wnsuo.mutation.ClearWorkerContainedInformation()
	return wnsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wnsuo *WorkerNetworkSettingsUpdateOne) Select(field string, fields ...string) *WorkerNetworkSettingsUpdateOne {
	wnsuo.fields = append([]string{field}, fields...)
	return wnsuo
}

// Save executes the query and returns the updated WorkerNetworkSettings entity.
func (wnsuo *WorkerNetworkSettingsUpdateOne) Save(ctx context.Context) (*WorkerNetworkSettings, error) {
	var (
		err  error
		node *WorkerNetworkSettings
	)
	if len(wnsuo.hooks) == 0 {
		node, err = wnsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkerNetworkSettingsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wnsuo.mutation = mutation
			node, err = wnsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wnsuo.hooks) - 1; i >= 0; i-- {
			if wnsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wnsuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, wnsuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*WorkerNetworkSettings)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from WorkerNetworkSettingsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wnsuo *WorkerNetworkSettingsUpdateOne) SaveX(ctx context.Context) *WorkerNetworkSettings {
	node, err := wnsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wnsuo *WorkerNetworkSettingsUpdateOne) Exec(ctx context.Context) error {
	_, err := wnsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wnsuo *WorkerNetworkSettingsUpdateOne) ExecX(ctx context.Context) {
	if err := wnsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wnsuo *WorkerNetworkSettingsUpdateOne) sqlSave(ctx context.Context) (_node *WorkerNetworkSettings, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workernetworksettings.Table,
			Columns: workernetworksettings.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workernetworksettings.FieldID,
			},
		},
	}
	id, ok := wnsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WorkerNetworkSettings.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wnsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workernetworksettings.FieldID)
		for _, f := range fields {
			if !workernetworksettings.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workernetworksettings.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wnsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if wnsuo.mutation.WorkerContainedInformationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   workernetworksettings.WorkerContainedInformationTable,
			Columns: []string{workernetworksettings.WorkerContainedInformationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workercontainedinformation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wnsuo.mutation.WorkerContainedInformationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   workernetworksettings.WorkerContainedInformationTable,
			Columns: []string{workernetworksettings.WorkerContainedInformationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workercontainedinformation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WorkerNetworkSettings{config: wnsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wnsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workernetworksettings.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
