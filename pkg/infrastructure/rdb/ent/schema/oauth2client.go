package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials"
)

// OAuth2Client holds the schema definition for the OAuth2Client entity.
type OAuth2Client struct {
	ent.Schema
}

// Fields of the OAuth2Client.
func (OAuth2Client) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(255).NotEmpty(),
		field.String("secret").MaxLen(255).NotEmpty().GoType(credentials.EncryptString{}),
		field.String("domain").MaxLen(255).GoType(credentials.EncryptString{}),
		field.Bool("public"),
		field.String("user_id").MaxLen(255).NotEmpty(),
		field.String("app_name").MaxLen(255),
	}
}

// Edges of the OAuth2Client.
func (OAuth2Client) Edges() []ent.Edge {
	return nil
}

func (OAuth2Client) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "oauth2_client"},
	}
}

func (OAuth2Client) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
