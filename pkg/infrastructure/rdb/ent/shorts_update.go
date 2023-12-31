// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent/predicate"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent/shorts"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials"
)

// ShortsUpdate is the builder for updating Shorts entities.
type ShortsUpdate struct {
	config
	hooks    []Hook
	mutation *ShortsMutation
}

// Where appends a list predicates to the ShortsUpdate builder.
func (su *ShortsUpdate) Where(ps ...predicate.Shorts) *ShortsUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *ShortsUpdate) SetUpdateTime(t time.Time) *ShortsUpdate {
	su.mutation.SetUpdateTime(t)
	return su
}

// SetKey sets the "key" field.
func (su *ShortsUpdate) SetKey(s string) *ShortsUpdate {
	su.mutation.SetKey(s)
	return su
}

// SetURL sets the "url" field.
func (su *ShortsUpdate) SetURL(cs credentials.EncryptString) *ShortsUpdate {
	su.mutation.SetURL(cs)
	return su
}

// SetAuthor sets the "author" field.
func (su *ShortsUpdate) SetAuthor(s string) *ShortsUpdate {
	su.mutation.SetAuthor(s)
	return su
}

// Mutation returns the ShortsMutation object of the builder.
func (su *ShortsUpdate) Mutation() *ShortsMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ShortsUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *ShortsUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ShortsUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ShortsUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *ShortsUpdate) defaults() {
	if _, ok := su.mutation.UpdateTime(); !ok {
		v := shorts.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *ShortsUpdate) check() error {
	if v, ok := su.mutation.Key(); ok {
		if err := shorts.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf(`ent: validator failed for field "Shorts.key": %w`, err)}
		}
	}
	if v, ok := su.mutation.URL(); ok {
		if err := shorts.URLValidator(v.String()); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Shorts.url": %w`, err)}
		}
	}
	if v, ok := su.mutation.Author(); ok {
		if err := shorts.AuthorValidator(v); err != nil {
			return &ValidationError{Name: "author", err: fmt.Errorf(`ent: validator failed for field "Shorts.author": %w`, err)}
		}
	}
	return nil
}

func (su *ShortsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(shorts.Table, shorts.Columns, sqlgraph.NewFieldSpec(shorts.FieldID, field.TypeInt64))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.SetField(shorts.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.Key(); ok {
		_spec.SetField(shorts.FieldKey, field.TypeString, value)
	}
	if value, ok := su.mutation.URL(); ok {
		_spec.SetField(shorts.FieldURL, field.TypeString, value)
	}
	if value, ok := su.mutation.Author(); ok {
		_spec.SetField(shorts.FieldAuthor, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shorts.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// ShortsUpdateOne is the builder for updating a single Shorts entity.
type ShortsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ShortsMutation
}

// SetUpdateTime sets the "update_time" field.
func (suo *ShortsUpdateOne) SetUpdateTime(t time.Time) *ShortsUpdateOne {
	suo.mutation.SetUpdateTime(t)
	return suo
}

// SetKey sets the "key" field.
func (suo *ShortsUpdateOne) SetKey(s string) *ShortsUpdateOne {
	suo.mutation.SetKey(s)
	return suo
}

// SetURL sets the "url" field.
func (suo *ShortsUpdateOne) SetURL(cs credentials.EncryptString) *ShortsUpdateOne {
	suo.mutation.SetURL(cs)
	return suo
}

// SetAuthor sets the "author" field.
func (suo *ShortsUpdateOne) SetAuthor(s string) *ShortsUpdateOne {
	suo.mutation.SetAuthor(s)
	return suo
}

// Mutation returns the ShortsMutation object of the builder.
func (suo *ShortsUpdateOne) Mutation() *ShortsMutation {
	return suo.mutation
}

// Where appends a list predicates to the ShortsUpdate builder.
func (suo *ShortsUpdateOne) Where(ps ...predicate.Shorts) *ShortsUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ShortsUpdateOne) Select(field string, fields ...string) *ShortsUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Shorts entity.
func (suo *ShortsUpdateOne) Save(ctx context.Context) (*Shorts, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ShortsUpdateOne) SaveX(ctx context.Context) *Shorts {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ShortsUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ShortsUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *ShortsUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		v := shorts.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *ShortsUpdateOne) check() error {
	if v, ok := suo.mutation.Key(); ok {
		if err := shorts.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf(`ent: validator failed for field "Shorts.key": %w`, err)}
		}
	}
	if v, ok := suo.mutation.URL(); ok {
		if err := shorts.URLValidator(v.String()); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Shorts.url": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Author(); ok {
		if err := shorts.AuthorValidator(v); err != nil {
			return &ValidationError{Name: "author", err: fmt.Errorf(`ent: validator failed for field "Shorts.author": %w`, err)}
		}
	}
	return nil
}

func (suo *ShortsUpdateOne) sqlSave(ctx context.Context) (_node *Shorts, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(shorts.Table, shorts.Columns, sqlgraph.NewFieldSpec(shorts.FieldID, field.TypeInt64))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Shorts.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shorts.FieldID)
		for _, f := range fields {
			if !shorts.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != shorts.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.SetField(shorts.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Key(); ok {
		_spec.SetField(shorts.FieldKey, field.TypeString, value)
	}
	if value, ok := suo.mutation.URL(); ok {
		_spec.SetField(shorts.FieldURL, field.TypeString, value)
	}
	if value, ok := suo.mutation.Author(); ok {
		_spec.SetField(shorts.FieldAuthor, field.TypeString, value)
	}
	_node = &Shorts{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shorts.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
