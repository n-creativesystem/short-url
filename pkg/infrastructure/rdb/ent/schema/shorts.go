package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials"
)

// Shorts holds the schema definition for the Shorts entity.
type Shorts struct {
	ent.Schema
}

// Fields of the Shorts.
func (Shorts) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("key").MaxLen(255),
		field.String("url").NotEmpty().GoType(credentials.EncryptString{}),
		field.String("author").MaxLen(255).NotEmpty(),
	}
}

// Edges of the Shorts.
func (Shorts) Edges() []ent.Edge {
	return nil
}

func (Shorts) Index() []ent.Index {
	return []ent.Index{
		index.Fields("key").Unique(),
	}
}

func (Shorts) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
