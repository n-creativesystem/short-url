// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent/users"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials"
	"github.com/n-creativesystem/short-url/pkg/utils/hash"
)

// UsersCreate is the builder for creating a Users entity.
type UsersCreate struct {
	config
	mutation *UsersMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (uc *UsersCreate) SetCreateTime(t time.Time) *UsersCreate {
	uc.mutation.SetCreateTime(t)
	return uc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uc *UsersCreate) SetNillableCreateTime(t *time.Time) *UsersCreate {
	if t != nil {
		uc.SetCreateTime(*t)
	}
	return uc
}

// SetUpdateTime sets the "update_time" field.
func (uc *UsersCreate) SetUpdateTime(t time.Time) *UsersCreate {
	uc.mutation.SetUpdateTime(t)
	return uc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (uc *UsersCreate) SetNillableUpdateTime(t *time.Time) *UsersCreate {
	if t != nil {
		uc.SetUpdateTime(*t)
	}
	return uc
}

// SetSubject sets the "Subject" field.
func (uc *UsersCreate) SetSubject(s string) *UsersCreate {
	uc.mutation.SetSubject(s)
	return uc
}

// SetProfile sets the "profile" field.
func (uc *UsersCreate) SetProfile(s string) *UsersCreate {
	uc.mutation.SetProfile(s)
	return uc
}

// SetEmail sets the "email" field.
func (uc *UsersCreate) SetEmail(cs credentials.EncryptString) *UsersCreate {
	uc.mutation.SetEmail(cs)
	return uc
}

// SetEmailHash sets the "email_hash" field.
func (uc *UsersCreate) SetEmailHash(h hash.Hash) *UsersCreate {
	uc.mutation.SetEmailHash(h)
	return uc
}

// SetNillableEmailHash sets the "email_hash" field if the given value is not nil.
func (uc *UsersCreate) SetNillableEmailHash(h *hash.Hash) *UsersCreate {
	if h != nil {
		uc.SetEmailHash(*h)
	}
	return uc
}

// SetEmailVerified sets the "email_verified" field.
func (uc *UsersCreate) SetEmailVerified(b bool) *UsersCreate {
	uc.mutation.SetEmailVerified(b)
	return uc
}

// SetUsername sets the "username" field.
func (uc *UsersCreate) SetUsername(cs credentials.EncryptString) *UsersCreate {
	uc.mutation.SetUsername(cs)
	return uc
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uc *UsersCreate) SetNillableUsername(cs *credentials.EncryptString) *UsersCreate {
	if cs != nil {
		uc.SetUsername(*cs)
	}
	return uc
}

// SetPicture sets the "picture" field.
func (uc *UsersCreate) SetPicture(s string) *UsersCreate {
	uc.mutation.SetPicture(s)
	return uc
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (uc *UsersCreate) SetNillablePicture(s *string) *UsersCreate {
	if s != nil {
		uc.SetPicture(*s)
	}
	return uc
}

// SetClaims sets the "claims" field.
func (uc *UsersCreate) SetClaims(cs credentials.EncryptString) *UsersCreate {
	uc.mutation.SetClaims(cs)
	return uc
}

// SetNillableClaims sets the "claims" field if the given value is not nil.
func (uc *UsersCreate) SetNillableClaims(cs *credentials.EncryptString) *UsersCreate {
	if cs != nil {
		uc.SetClaims(*cs)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UsersCreate) SetID(u uuid.UUID) *UsersCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UsersCreate) SetNillableID(u *uuid.UUID) *UsersCreate {
	if u != nil {
		uc.SetID(*u)
	}
	return uc
}

// Mutation returns the UsersMutation object of the builder.
func (uc *UsersCreate) Mutation() *UsersMutation {
	return uc.mutation
}

// Save creates the Users in the database.
func (uc *UsersCreate) Save(ctx context.Context) (*Users, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UsersCreate) SaveX(ctx context.Context) *Users {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UsersCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UsersCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UsersCreate) defaults() {
	if _, ok := uc.mutation.CreateTime(); !ok {
		v := users.DefaultCreateTime()
		uc.mutation.SetCreateTime(v)
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		v := users.DefaultUpdateTime()
		uc.mutation.SetUpdateTime(v)
	}
	if _, ok := uc.mutation.EmailHash(); !ok {
		v := users.DefaultEmailHash
		uc.mutation.SetEmailHash(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := users.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UsersCreate) check() error {
	if _, ok := uc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Users.create_time"`)}
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Users.update_time"`)}
	}
	if _, ok := uc.mutation.Subject(); !ok {
		return &ValidationError{Name: "Subject", err: errors.New(`ent: missing required field "Users.Subject"`)}
	}
	if v, ok := uc.mutation.Subject(); ok {
		if err := users.SubjectValidator(v); err != nil {
			return &ValidationError{Name: "Subject", err: fmt.Errorf(`ent: validator failed for field "Users.Subject": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Profile(); !ok {
		return &ValidationError{Name: "profile", err: errors.New(`ent: missing required field "Users.profile"`)}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Users.email"`)}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := users.EmailValidator(v.String()); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Users.email": %w`, err)}
		}
	}
	if _, ok := uc.mutation.EmailHash(); !ok {
		return &ValidationError{Name: "email_hash", err: errors.New(`ent: missing required field "Users.email_hash"`)}
	}
	if _, ok := uc.mutation.EmailVerified(); !ok {
		return &ValidationError{Name: "email_verified", err: errors.New(`ent: missing required field "Users.email_verified"`)}
	}
	if v, ok := uc.mutation.Username(); ok {
		if err := users.UsernameValidator(v.String()); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Users.username": %w`, err)}
		}
	}
	return nil
}

func (uc *UsersCreate) sqlSave(ctx context.Context) (*Users, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UsersCreate) createSpec() (*Users, *sqlgraph.CreateSpec) {
	var (
		_node = &Users{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(users.Table, sqlgraph.NewFieldSpec(users.FieldID, field.TypeUUID))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uc.mutation.CreateTime(); ok {
		_spec.SetField(users.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := uc.mutation.UpdateTime(); ok {
		_spec.SetField(users.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := uc.mutation.Subject(); ok {
		_spec.SetField(users.FieldSubject, field.TypeString, value)
		_node.Subject = value
	}
	if value, ok := uc.mutation.Profile(); ok {
		_spec.SetField(users.FieldProfile, field.TypeString, value)
		_node.Profile = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(users.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.EmailHash(); ok {
		_spec.SetField(users.FieldEmailHash, field.TypeOther, value)
		_node.EmailHash = value
	}
	if value, ok := uc.mutation.EmailVerified(); ok {
		_spec.SetField(users.FieldEmailVerified, field.TypeBool, value)
		_node.EmailVerified = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(users.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Picture(); ok {
		_spec.SetField(users.FieldPicture, field.TypeString, value)
		_node.Picture = value
	}
	if value, ok := uc.mutation.Claims(); ok {
		_spec.SetField(users.FieldClaims, field.TypeBytes, value)
		_node.Claims = value
	}
	return _node, _spec
}

// UsersCreateBulk is the builder for creating many Users entities in bulk.
type UsersCreateBulk struct {
	config
	builders []*UsersCreate
}

// Save creates the Users entities in the database.
func (ucb *UsersCreateBulk) Save(ctx context.Context) ([]*Users, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*Users, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UsersMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UsersCreateBulk) SaveX(ctx context.Context) []*Users {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UsersCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UsersCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
