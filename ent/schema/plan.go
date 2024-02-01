package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Plan holds the schema definition for the Plan entity.
type Plan struct {
	ent.Schema
}

// Fields of the Plan.
func (Plan) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name").MaxLen(50).Unique(),
		field.Int64("price").Optional().Nillable().Comment("قیمت"),
		field.Int("period").Optional().Nillable().Comment("مدت استفاده به ماه"),
		field.Bool("active").Default(true).Comment("فعال بودن"),
		field.String("description").Optional().Comment("توضیحات").
			SchemaType(map[string]string{
				dialect.Postgres: "text",
			}).MaxLen(500).Optional().Nillable(),
	}
}

// Edges of the Plan.
func (Plan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("userpaymentplans", UserPaymentPlan.Type),
	}
}
