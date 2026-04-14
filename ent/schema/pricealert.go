package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PriceAlert struct {
	ent.Schema
}

func (PriceAlert) Fields() []ent.Field {
	return []ent.Field{
		field.Float("target_price").Optional(),

		field.Bool("is_active").Default(true),
		field.Time("last_notified_at").Optional().Nillable(),
		field.Time("created_at").Default(time.Now),
	}
}

func (PriceAlert) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("channel", Channel.Type).
			Ref("price_alerts").
			Unique().
			Required(),

		edge.From("listing", Listing.Type).
			Ref("price_alerts").
			Unique().
			Required(),
	}
}
