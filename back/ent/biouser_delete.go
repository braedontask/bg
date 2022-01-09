// Code generated by entc, DO NOT EDIT.

package ent

import (
	"back/ent/biouser"
	"back/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BioUserDelete is the builder for deleting a BioUser entity.
type BioUserDelete struct {
	config
	hooks    []Hook
	mutation *BioUserMutation
}

// Where appends a list predicates to the BioUserDelete builder.
func (bud *BioUserDelete) Where(ps ...predicate.BioUser) *BioUserDelete {
	bud.mutation.Where(ps...)
	return bud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bud *BioUserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bud.hooks) == 0 {
		affected, err = bud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BioUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bud.mutation = mutation
			affected, err = bud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bud.hooks) - 1; i >= 0; i-- {
			if bud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bud *BioUserDelete) ExecX(ctx context.Context) int {
	n, err := bud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bud *BioUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: biouser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: biouser.FieldID,
			},
		},
	}
	if ps := bud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bud.driver, _spec)
}

// BioUserDeleteOne is the builder for deleting a single BioUser entity.
type BioUserDeleteOne struct {
	bud *BioUserDelete
}

// Exec executes the deletion query.
func (budo *BioUserDeleteOne) Exec(ctx context.Context) error {
	n, err := budo.bud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{biouser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (budo *BioUserDeleteOne) ExecX(ctx context.Context) {
	budo.bud.ExecX(ctx)
}
