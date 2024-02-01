package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return nil
}
