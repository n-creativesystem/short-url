// Code generated by ent, DO NOT EDIT.

package oauth2token

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the oauth2token type in the database.
	Label = "oauth2token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldExpiredAt holds the string denoting the expired_at field in the database.
	FieldExpiredAt = "expired_at"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldAccess holds the string denoting the access field in the database.
	FieldAccess = "access"
	// FieldRefresh holds the string denoting the refresh field in the database.
	FieldRefresh = "refresh"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// Table holds the table name of the oauth2token in the database.
	Table = "oauth2_token"
)

// Columns holds all SQL columns for oauth2token fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldExpiredAt,
	FieldCode,
	FieldAccess,
	FieldRefresh,
	FieldData,
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
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// AccessValidator is a validator for the "access" field. It is called by the builders before save.
	AccessValidator func(string) error
	// RefreshValidator is a validator for the "refresh" field. It is called by the builders before save.
	RefreshValidator func(string) error
)

// OrderOption defines the ordering options for the OAuth2Token queries.
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

// ByExpiredAt orders the results by the expired_at field.
func ByExpiredAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiredAt, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByAccess orders the results by the access field.
func ByAccess(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccess, opts...).ToFunc()
}

// ByRefresh orders the results by the refresh field.
func ByRefresh(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefresh, opts...).ToFunc()
}

// ByData orders the results by the data field.
func ByData(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldData, opts...).ToFunc()
}
