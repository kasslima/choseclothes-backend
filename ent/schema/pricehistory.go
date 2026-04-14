package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type PriceHistory struct {
	ent.Schema
}

func (PriceHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Float("price"),
		field.Float("discount_price").Optional(),
		field.Time("recorded_at").Default(time.Now),
	}
}

func (PriceHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("listing", Listing.Type).
			Ref("price_history").
			Unique().
			Required(),
	}
}

func (PriceHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "price_history"},
	}
}

func (PriceHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("recorded_at"),
	}
}
