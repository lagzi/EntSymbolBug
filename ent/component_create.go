// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lagzi/EntSymbolBug/ent/component"
	"github.com/lagzi/EntSymbolBug/ent/picture"
)

// ComponentCreate is the builder for creating a Component entity.
type ComponentCreate struct {
	config
	mutation *ComponentMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (cc *ComponentCreate) SetCreateTime(t time.Time) *ComponentCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *ComponentCreate) SetNillableCreateTime(t *time.Time) *ComponentCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *ComponentCreate) SetUpdateTime(t time.Time) *ComponentCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *ComponentCreate) SetNillableUpdateTime(t *time.Time) *ComponentCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetSomeID sets the "some_id" field.
func (cc *ComponentCreate) SetSomeID(u uint64) *ComponentCreate {
	cc.mutation.SetSomeID(u)
	return cc
}

// AddPictureIDs adds the "pictures" edge to the Picture entity by IDs.
func (cc *ComponentCreate) AddPictureIDs(ids ...int) *ComponentCreate {
	cc.mutation.AddPictureIDs(ids...)
	return cc
}

// AddPictures adds the "pictures" edges to the Picture entity.
func (cc *ComponentCreate) AddPictures(p ...*Picture) *ComponentCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddPictureIDs(ids...)
}

// Mutation returns the ComponentMutation object of the builder.
func (cc *ComponentCreate) Mutation() *ComponentMutation {
	return cc.mutation
}

// Save creates the Component in the database.
func (cc *ComponentCreate) Save(ctx context.Context) (*Component, error) {
	var (
		err  error
		node *Component
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ComponentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Component)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ComponentMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ComponentCreate) SaveX(ctx context.Context) *Component {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ComponentCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ComponentCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ComponentCreate) defaults() {
	if _, ok := cc.mutation.CreateTime(); !ok {
		v := component.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		v := component.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ComponentCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Component.create_time"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Component.update_time"`)}
	}
	if _, ok := cc.mutation.SomeID(); !ok {
		return &ValidationError{Name: "some_id", err: errors.New(`ent: missing required field "Component.some_id"`)}
	}
	return nil
}

func (cc *ComponentCreate) sqlSave(ctx context.Context) (*Component, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *ComponentCreate) createSpec() (*Component, *sqlgraph.CreateSpec) {
	var (
		_node = &Component{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: component.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: component.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: component.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: component.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.SomeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: component.FieldSomeID,
		})
		_node.SomeID = value
	}
	if nodes := cc.mutation.PicturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   component.PicturesTable,
			Columns: []string{component.PicturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: picture.FieldID,
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

// ComponentCreateBulk is the builder for creating many Component entities in bulk.
type ComponentCreateBulk struct {
	config
	builders []*ComponentCreate
}

// Save creates the Component entities in the database.
func (ccb *ComponentCreateBulk) Save(ctx context.Context) ([]*Component, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Component, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ComponentMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ComponentCreateBulk) SaveX(ctx context.Context) []*Component {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ComponentCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ComponentCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}