// Code generated by ent, DO NOT EDIT.

package users

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/n-creativesystem/short-url/pkg/utils/hash"
)

const (
	// Label holds the string label denoting the users type in the database.
	Label = "users"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldSubject holds the string denoting the subject field in the database.
	FieldSubject = "subject"
	// FieldProfile holds the string denoting the profile field in the database.
	FieldProfile = "profile"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldEmailHash holds the string denoting the email_hash field in the database.
	FieldEmailHash = "email_hash"
	// FieldEmailVerified holds the string denoting the email_verified field in the database.
	FieldEmailVerified = "email_verified"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPicture holds the string denoting the picture field in the database.
	FieldPicture = "picture"
	// FieldClaims holds the string denoting the claims field in the database.
	FieldClaims = "claims"
	// Table holds the table name of the users in the database.
	Table = "users"
)

// Columns holds all SQL columns for users fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldSubject,
	FieldProfile,
	FieldEmail,
	FieldEmailHash,
	FieldEmailVerified,
	FieldUsername,
	FieldPicture,
	FieldClaims,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// SubjectValidator is a validator for the "Subject" field. It is called by the builders before save.
	SubjectValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultEmailHash holds the default value on creation for the "email_hash" field.
	DefaultEmailHash hash.Hash
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Users queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// BySubject orders the results by the Subject field.
func BySubject(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubject, opts...).ToFunc()
}

// ByProfile orders the results by the profile field.
func ByProfile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProfile, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByEmailHash orders the results by the email_hash field.
func ByEmailHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailHash, opts...).ToFunc()
}

// ByEmailVerified orders the results by the email_verified field.
func ByEmailVerified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailVerified, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByPicture orders the results by the picture field.
func ByPicture(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPicture, opts...).ToFunc()
}
