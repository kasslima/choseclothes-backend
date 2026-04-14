package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid/v5"
)

type Channel struct {
	ent.Schema
}

func (Channel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(func() uuid.UUID {
				return uuid.Must(uuid.NewV7())
			}),
		field.String("type").NotEmpty(),
		field.String("external_id").NotEmpty(),
		field.String("name").Optional(),

		field.Bool("is_active").Default(true),
		field.Time("created_at").Default(time.Now),
	}
}


func (Channel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("price_alerts", PriceAlert.Type),
	}
}
