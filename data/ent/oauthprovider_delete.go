// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"hertz-admin/data/ent/oauthprovider"
	"hertz-admin/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OauthProviderDelete is the builder for deleting a OauthProvider entity.
type OauthProviderDelete struct {
	config
	hooks    []Hook
	mutation *OauthProviderMutation
}

// Where appends a list predicates to the OauthProviderDelete builder.
func (opd *OauthProviderDelete) Where(ps ...predicate.OauthProvider) *OauthProviderDelete {
	opd.mutation.Where(ps...)
	return opd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (opd *OauthProviderDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, OauthProviderMutation](ctx, opd.sqlExec, opd.mutation, opd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (opd *OauthProviderDelete) ExecX(ctx context.Context) int {
	n, err := opd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (opd *OauthProviderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: oauthprovider.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: oauthprovider.FieldID,
			},
		},
	}
	if ps := opd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, opd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	opd.mutation.done = true
	return affected, err
}

// OauthProviderDeleteOne is the builder for deleting a single OauthProvider entity.
type OauthProviderDeleteOne struct {
	opd *OauthProviderDelete
}

// Exec executes the deletion query.
func (opdo *OauthProviderDeleteOne) Exec(ctx context.Context) error {
	n, err := opdo.opd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{oauthprovider.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (opdo *OauthProviderDeleteOne) ExecX(ctx context.Context) {
	opdo.opd.ExecX(ctx)
}
