package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// OAuth2Token holds the schema definition for the OAuth2Token entity.
type OAuth2Token struct {
	ent.Schema
}

// Fields of the OAuth2Token.
func (OAuth2Token) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("expired_at"),
		field.String("code").MaxLen(512).Optional(),
		field.String("access").MaxLen(512).Optional(),
		field.String("refresh").MaxLen(512).Optional(),
		field.String("data").SchemaType(
			map[string]string{
				dialect.MySQL:    "text",
				dialect.Postgres: "text",
				dialect.SQLite:   "text",
			},
		).Optional(),
	}
}

// Edges of the OAuth2Token.
func (OAuth2Token) Edges() []ent.Edge {
	return nil
}

func (OAuth2Token) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").Unique().StorageKey("idx_code"),
		index.Fields("access").Unique().StorageKey("idx_access"),
		index.Fields("refresh").Unique().StorageKey("idx_refresh"),
		index.Fields("expired_at").Unique().StorageKey("idx_expired_at"),
	}
}

func (OAuth2Token) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "oauth2_token"},
	}
}

func (OAuth2Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
