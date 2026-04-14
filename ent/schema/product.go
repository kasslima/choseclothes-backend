package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid/v5"
)

type Product struct {
	ent.Schema
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(func() uuid.UUID {
				return uuid.Must(uuid.NewV7())
			}),
		field.String("title").NotEmpty(),
		field.String("description").Optional(),
		field.String("category").Optional(),
		field.String("image_url").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}


func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("listings", Listing.Type),
	}
}
