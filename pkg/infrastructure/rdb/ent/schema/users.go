package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials"
	"github.com/n-creativesystem/short-url/pkg/utils/hash"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("Subject").MaxLen(256),
		field.String("profile"),
		field.String("email").MaxLen(256).GoType(credentials.EncryptString{}),
		field.Other("email_hash", hash.Hash{}).Default(hash.NewHash("")).SchemaType(map[string]string{
			dialect.MySQL:    "text",
			dialect.Postgres: "text",
			dialect.SQLite:   "text",
		}),
		field.Bool("email_verified"),
		field.String("username").MaxLen(256).Optional().GoType(credentials.EncryptString{}),
		field.String("picture").Optional(),
		field.Bytes("claims").Optional().GoType(credentials.EncryptString{}),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}

func (Users) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").Unique(),
	}
}

func (Users) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
