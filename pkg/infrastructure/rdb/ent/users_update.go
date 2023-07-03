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
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent/users"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials"
	"github.com/n-creativesystem/short-url/pkg/utils/hash"
)

// UsersUpdate is the builder for updating Users entities.
type UsersUpdate struct {
	config
	hooks    []Hook
	mutation *UsersMutation
}

// Where appends a list predicates to the UsersUpdate builder.
func (uu *UsersUpdate) Where(ps ...predicate.Users) *UsersUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdateTime sets the "update_time" field.
func (uu *UsersUpdate) SetUpdateTime(t time.Time) *UsersUpdate {
	uu.mutation.SetUpdateTime(t)
	return uu
}

// SetSubject sets the "Subject" field.
func (uu *UsersUpdate) SetSubject(s string) *UsersUpdate {
	uu.mutation.SetSubject(s)
	return uu
}

// SetProfile sets the "profile" field.
func (uu *UsersUpdate) SetProfile(s string) *UsersUpdate {
	uu.mutation.SetProfile(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UsersUpdate) SetEmail(cs credentials.EncryptString) *UsersUpdate {
	uu.mutation.SetEmail(cs)
	return uu
}

// SetEmailHash sets the "email_hash" field.
func (uu *UsersUpdate) SetEmailHash(h hash.Hash) *UsersUpdate {
	uu.mutation.SetEmailHash(h)
	return uu
}

// SetNillableEmailHash sets the "email_hash" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableEmailHash(h *hash.Hash) *UsersUpdate {
	if h != nil {
		uu.SetEmailHash(*h)
	}
	return uu
}

// SetEmailVerified sets the "email_verified" field.
func (uu *UsersUpdate) SetEmailVerified(b bool) *UsersUpdate {
	uu.mutation.SetEmailVerified(b)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UsersUpdate) SetUsername(cs credentials.EncryptString) *UsersUpdate {
	uu.mutation.SetUsername(cs)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableUsername(cs *credentials.EncryptString) *UsersUpdate {
	if cs != nil {
		uu.SetUsername(*cs)
	}
	return uu
}

// ClearUsername clears the value of the "username" field.
func (uu *UsersUpdate) ClearUsername() *UsersUpdate {
	uu.mutation.ClearUsername()
	return uu
}

// SetPicture sets the "picture" field.
func (uu *UsersUpdate) SetPicture(s string) *UsersUpdate {
	uu.mutation.SetPicture(s)
	return uu
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (uu *UsersUpdate) SetNillablePicture(s *string) *UsersUpdate {
	if s != nil {
		uu.SetPicture(*s)
	}
	return uu
}

// ClearPicture clears the value of the "picture" field.
func (uu *UsersUpdate) ClearPicture() *UsersUpdate {
	uu.mutation.ClearPicture()
	return uu
}

// SetClaims sets the "claims" field.
func (uu *UsersUpdate) SetClaims(cs credentials.EncryptString) *UsersUpdate {
	uu.mutation.SetClaims(cs)
	return uu
}

// SetNillableClaims sets the "claims" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableClaims(cs *credentials.EncryptString) *UsersUpdate {
	if cs != nil {
		uu.SetClaims(*cs)
	}
	return uu
}

// ClearClaims clears the value of the "claims" field.
func (uu *UsersUpdate) ClearClaims() *UsersUpdate {
	uu.mutation.ClearClaims()
	return uu
}

// Mutation returns the UsersMutation object of the builder.
func (uu *UsersUpdate) Mutation() *UsersMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UsersUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UsersUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UsersUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UsersUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UsersUpdate) defaults() {
	if _, ok := uu.mutation.UpdateTime(); !ok {
		v := users.UpdateDefaultUpdateTime()
		uu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UsersUpdate) check() error {
	if v, ok := uu.mutation.Subject(); ok {
		if err := users.SubjectValidator(v); err != nil {
			return &ValidationError{Name: "Subject", err: fmt.Errorf(`ent: validator failed for field "Users.Subject": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := users.EmailValidator(v.String()); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Users.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Username(); ok {
		if err := users.UsernameValidator(v.String()); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Users.username": %w`, err)}
		}
	}
	return nil
}

func (uu *UsersUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(users.Table, users.Columns, sqlgraph.NewFieldSpec(users.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdateTime(); ok {
		_spec.SetField(users.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Subject(); ok {
		_spec.SetField(users.FieldSubject, field.TypeString, value)
	}
	if value, ok := uu.mutation.Profile(); ok {
		_spec.SetField(users.FieldProfile, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(users.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.EmailHash(); ok {
		_spec.SetField(users.FieldEmailHash, field.TypeOther, value)
	}
	if value, ok := uu.mutation.EmailVerified(); ok {
		_spec.SetField(users.FieldEmailVerified, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(users.FieldUsername, field.TypeString, value)
	}
	if uu.mutation.UsernameCleared() {
		_spec.ClearField(users.FieldUsername, field.TypeString)
	}
	if value, ok := uu.mutation.Picture(); ok {
		_spec.SetField(users.FieldPicture, field.TypeString, value)
	}
	if uu.mutation.PictureCleared() {
		_spec.ClearField(users.FieldPicture, field.TypeString)
	}
	if value, ok := uu.mutation.Claims(); ok {
		_spec.SetField(users.FieldClaims, field.TypeBytes, value)
	}
	if uu.mutation.ClaimsCleared() {
		_spec.ClearField(users.FieldClaims, field.TypeBytes)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{users.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UsersUpdateOne is the builder for updating a single Users entity.
type UsersUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UsersMutation
}

// SetUpdateTime sets the "update_time" field.
func (uuo *UsersUpdateOne) SetUpdateTime(t time.Time) *UsersUpdateOne {
	uuo.mutation.SetUpdateTime(t)
	return uuo
}

// SetSubject sets the "Subject" field.
func (uuo *UsersUpdateOne) SetSubject(s string) *UsersUpdateOne {
	uuo.mutation.SetSubject(s)
	return uuo
}

// SetProfile sets the "profile" field.
func (uuo *UsersUpdateOne) SetProfile(s string) *UsersUpdateOne {
	uuo.mutation.SetProfile(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UsersUpdateOne) SetEmail(cs credentials.EncryptString) *UsersUpdateOne {
	uuo.mutation.SetEmail(cs)
	return uuo
}

// SetEmailHash sets the "email_hash" field.
func (uuo *UsersUpdateOne) SetEmailHash(h hash.Hash) *UsersUpdateOne {
	uuo.mutation.SetEmailHash(h)
	return uuo
}

// SetNillableEmailHash sets the "email_hash" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableEmailHash(h *hash.Hash) *UsersUpdateOne {
	if h != nil {
		uuo.SetEmailHash(*h)
	}
	return uuo
}

// SetEmailVerified sets the "email_verified" field.
func (uuo *UsersUpdateOne) SetEmailVerified(b bool) *UsersUpdateOne {
	uuo.mutation.SetEmailVerified(b)
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UsersUpdateOne) SetUsername(cs credentials.EncryptString) *UsersUpdateOne {
	uuo.mutation.SetUsername(cs)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableUsername(cs *credentials.EncryptString) *UsersUpdateOne {
	if cs != nil {
		uuo.SetUsername(*cs)
	}
	return uuo
}

// ClearUsername clears the value of the "username" field.
func (uuo *UsersUpdateOne) ClearUsername() *UsersUpdateOne {
	uuo.mutation.ClearUsername()
	return uuo
}

// SetPicture sets the "picture" field.
func (uuo *UsersUpdateOne) SetPicture(s string) *UsersUpdateOne {
	uuo.mutation.SetPicture(s)
	return uuo
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillablePicture(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetPicture(*s)
	}
	return uuo
}

// ClearPicture clears the value of the "picture" field.
func (uuo *UsersUpdateOne) ClearPicture() *UsersUpdateOne {
	uuo.mutation.ClearPicture()
	return uuo
}

// SetClaims sets the "claims" field.
func (uuo *UsersUpdateOne) SetClaims(cs credentials.EncryptString) *UsersUpdateOne {
	uuo.mutation.SetClaims(cs)
	return uuo
}

// SetNillableClaims sets the "claims" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableClaims(cs *credentials.EncryptString) *UsersUpdateOne {
	if cs != nil {
		uuo.SetClaims(*cs)
	}
	return uuo
}

// ClearClaims clears the value of the "claims" field.
func (uuo *UsersUpdateOne) ClearClaims() *UsersUpdateOne {
	uuo.mutation.ClearClaims()
	return uuo
}

// Mutation returns the UsersMutation object of the builder.
func (uuo *UsersUpdateOne) Mutation() *UsersMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UsersUpdate builder.
func (uuo *UsersUpdateOne) Where(ps ...predicate.Users) *UsersUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UsersUpdateOne) Select(field string, fields ...string) *UsersUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Users entity.
func (uuo *UsersUpdateOne) Save(ctx context.Context) (*Users, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UsersUpdateOne) SaveX(ctx context.Context) *Users {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UsersUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UsersUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UsersUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdateTime(); !ok {
		v := users.UpdateDefaultUpdateTime()
		uuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UsersUpdateOne) check() error {
	if v, ok := uuo.mutation.Subject(); ok {
		if err := users.SubjectValidator(v); err != nil {
			return &ValidationError{Name: "Subject", err: fmt.Errorf(`ent: validator failed for field "Users.Subject": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := users.EmailValidator(v.String()); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Users.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Username(); ok {
		if err := users.UsernameValidator(v.String()); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Users.username": %w`, err)}
		}
	}
	return nil
}

func (uuo *UsersUpdateOne) sqlSave(ctx context.Context) (_node *Users, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(users.Table, users.Columns, sqlgraph.NewFieldSpec(users.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Users.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, users.FieldID)
		for _, f := range fields {
			if !users.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != users.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdateTime(); ok {
		_spec.SetField(users.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Subject(); ok {
		_spec.SetField(users.FieldSubject, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Profile(); ok {
		_spec.SetField(users.FieldProfile, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(users.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.EmailHash(); ok {
		_spec.SetField(users.FieldEmailHash, field.TypeOther, value)
	}
	if value, ok := uuo.mutation.EmailVerified(); ok {
		_spec.SetField(users.FieldEmailVerified, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(users.FieldUsername, field.TypeString, value)
	}
	if uuo.mutation.UsernameCleared() {
		_spec.ClearField(users.FieldUsername, field.TypeString)
	}
	if value, ok := uuo.mutation.Picture(); ok {
		_spec.SetField(users.FieldPicture, field.TypeString, value)
	}
	if uuo.mutation.PictureCleared() {
		_spec.ClearField(users.FieldPicture, field.TypeString)
	}
	if value, ok := uuo.mutation.Claims(); ok {
		_spec.SetField(users.FieldClaims, field.TypeBytes, value)
	}
	if uuo.mutation.ClaimsCleared() {
		_spec.ClearField(users.FieldClaims, field.TypeBytes)
	}
	_node = &Users{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{users.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}