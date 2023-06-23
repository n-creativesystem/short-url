// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// Oauth2ClientColumns holds the columns for the "oauth2_client" table.
	Oauth2ClientColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 255},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "secret", Type: field.TypeString, Size: 255},
		{Name: "domain", Type: field.TypeString, Size: 255},
		{Name: "public", Type: field.TypeBool},
		{Name: "user_id", Type: field.TypeString, Size: 255},
		{Name: "app_name", Type: field.TypeString, Size: 255},
	}
	// Oauth2ClientTable holds the schema information for the "oauth2_client" table.
	Oauth2ClientTable = &schema.Table{
		Name:       "oauth2_client",
		Columns:    Oauth2ClientColumns,
		PrimaryKey: []*schema.Column{Oauth2ClientColumns[0]},
	}
	// Oauth2TokenColumns holds the columns for the "oauth2_token" table.
	Oauth2TokenColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "expired_at", Type: field.TypeInt64},
		{Name: "code", Type: field.TypeString, Nullable: true, Size: 512},
		{Name: "access", Type: field.TypeString, Nullable: true, Size: 512},
		{Name: "refresh", Type: field.TypeString, Nullable: true, Size: 512},
		{Name: "data", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "text", "postgres": "text", "sqlite3": "text"}},
	}
	// Oauth2TokenTable holds the schema information for the "oauth2_token" table.
	Oauth2TokenTable = &schema.Table{
		Name:       "oauth2_token",
		Columns:    Oauth2TokenColumns,
		PrimaryKey: []*schema.Column{Oauth2TokenColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "idx_code",
				Unique:  true,
				Columns: []*schema.Column{Oauth2TokenColumns[4]},
			},
			{
				Name:    "idx_access",
				Unique:  true,
				Columns: []*schema.Column{Oauth2TokenColumns[5]},
			},
			{
				Name:    "idx_refresh",
				Unique:  true,
				Columns: []*schema.Column{Oauth2TokenColumns[6]},
			},
			{
				Name:    "idx_expired_at",
				Unique:  true,
				Columns: []*schema.Column{Oauth2TokenColumns[3]},
			},
		},
	}
	// ShortsColumns holds the columns for the "shorts" table.
	ShortsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "key", Type: field.TypeString, Size: 255},
		{Name: "url", Type: field.TypeString},
		{Name: "author", Type: field.TypeString, Size: 255},
	}
	// ShortsTable holds the schema information for the "shorts" table.
	ShortsTable = &schema.Table{
		Name:       "shorts",
		Columns:    ShortsColumns,
		PrimaryKey: []*schema.Column{ShortsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		Oauth2ClientTable,
		Oauth2TokenTable,
		ShortsTable,
	}
)

func init() {
	Oauth2ClientTable.Annotation = &entsql.Annotation{
		Table: "oauth2_client",
	}
	Oauth2TokenTable.Annotation = &entsql.Annotation{
		Table: "oauth2_token",
	}
}
