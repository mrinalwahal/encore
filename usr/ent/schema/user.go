package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

type Metadata struct{}

// Fields of the User.
func (User) Fields() []ent.Field {

	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("name"),
		field.String("username").Unique().Optional(),
		field.Time("created_at").Default(time.Now()).Immutable(),
		field.String("email").Optional().Unique(),
		field.String("phone").Optional().Unique(),
		field.Bool("disabled").Optional().Default(false),
		field.String("avatar_url").Optional(),
		field.String("locale").Default("en/IN"),
		field.String("password_hash").Optional(),
		field.String("default_role").Default("user"),
		field.Bool("is_anonymous").Default(false),
		field.String("totp_secret").Optional(),
		field.String("active_mfa_type").Optional(),
		field.JSON("metadata", &Metadata{}).Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
