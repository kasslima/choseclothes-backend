package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/gofrs/uuid/v5"
)

type Listing struct {
	ent.Schema
}

func (Listing) Fields() []ent.Field {
	return []ent.Field{
		field.String("external_id").NotEmpty(),
		field.String("title").NotEmpty(),
		field.String("seller_name").Optional(),
		field.Float("seller_rating").Optional(),

		field.String("product_url").NotEmpty(),
		field.String("currency").Default("USD"),

		field.Bool("is_active").Default(true),
		field.Time("created_at").Default(time.Now),
	}
}

func (Listing) ID() ent.Field {
	return field.UUID("id", uuid.UUID{}).
		Default(func() uuid.UUID {
			return uuid.Must(uuid.NewV7())
		})
}

func (Listing) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).
			Ref("listings").
			Unique().
			Required(),

		edge.From("source", Source.Type).
			Ref("listings").
			Unique().
			Required(),

		edge.To("price_alerts", PriceAlert.Type),

		edge.To("price_history", PriceHistory.Type),
	}
}

func (Listing) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("external_id").
			Edges("source").
			Unique(),
	}
}
