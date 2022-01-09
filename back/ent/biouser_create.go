// Code generated by entc, DO NOT EDIT.

package ent

import (
	"back/ent/biouser"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BioUserCreate is the builder for creating a BioUser entity.
type BioUserCreate struct {
	config
	mutation *BioUserMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (buc *BioUserCreate) SetName(s string) *BioUserCreate {
	buc.mutation.SetName(s)
	return buc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (buc *BioUserCreate) SetNillableName(s *string) *BioUserCreate {
	if s != nil {
		buc.SetName(*s)
	}
	return buc
}

// Mutation returns the BioUserMutation object of the builder.
func (buc *BioUserCreate) Mutation() *BioUserMutation {
	return buc.mutation
}

// Save creates the BioUser in the database.
func (buc *BioUserCreate) Save(ctx context.Context) (*BioUser, error) {
	var (
		err  error
		node *BioUser
	)
	buc.defaults()
	if len(buc.hooks) == 0 {
		if err = buc.check(); err != nil {
			return nil, err
		}
		node, err = buc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BioUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = buc.check(); err != nil {
				return nil, err
			}
			buc.mutation = mutation
			if node, err = buc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(buc.hooks) - 1; i >= 0; i-- {
			if buc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (buc *BioUserCreate) SaveX(ctx context.Context) *BioUser {
	v, err := buc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (buc *BioUserCreate) Exec(ctx context.Context) error {
	_, err := buc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buc *BioUserCreate) ExecX(ctx context.Context) {
	if err := buc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buc *BioUserCreate) defaults() {
	if _, ok := buc.mutation.Name(); !ok {
		v := biouser.DefaultName
		buc.mutation.SetName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buc *BioUserCreate) check() error {
	if _, ok := buc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	return nil
}

func (buc *BioUserCreate) sqlSave(ctx context.Context) (*BioUser, error) {
	_node, _spec := buc.createSpec()
	if err := sqlgraph.CreateNode(ctx, buc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (buc *BioUserCreate) createSpec() (*BioUser, *sqlgraph.CreateSpec) {
	var (
		_node = &BioUser{config: buc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: biouser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: biouser.FieldID,
			},
		}
	)
	if value, ok := buc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: biouser.FieldName,
		})
		_node.Name = value
	}
	return _node, _spec
}

// BioUserCreateBulk is the builder for creating many BioUser entities in bulk.
type BioUserCreateBulk struct {
	config
	builders []*BioUserCreate
}

// Save creates the BioUser entities in the database.
func (bucb *BioUserCreateBulk) Save(ctx context.Context) ([]*BioUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bucb.builders))
	nodes := make([]*BioUser, len(bucb.builders))
	mutators := make([]Mutator, len(bucb.builders))
	for i := range bucb.builders {
		func(i int, root context.Context) {
			builder := bucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BioUserMutation)
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
					_, err = mutators[i+1].Mutate(root, bucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bucb *BioUserCreateBulk) SaveX(ctx context.Context) []*BioUser {
	v, err := bucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bucb *BioUserCreateBulk) Exec(ctx context.Context) error {
	_, err := bucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bucb *BioUserCreateBulk) ExecX(ctx context.Context) {
	if err := bucb.Exec(ctx); err != nil {
		panic(err)
	}
}