package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique(),
		field.String("Subject").MaxLen(256),
		field.String("profile"),
		field.String("email").MaxLen(256),
		field.Bool("email_verified"),
		field.String("username").MaxLen(256),
		field.String("picture"),
		field.Bytes("claims"),
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
