package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid/v5"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(func() uuid.UUID {
				return uuid.Must(uuid.NewV7())
			}),
		field.String("google_id").Optional().Unique(),
		field.String("email").Unique().NotEmpty(),
		field.String("password_hash").Optional(),
		field.String("name").Optional(),
		field.String("profile_picture").Optional(),

		field.Bool("is_subscriber").Default(false),

		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

