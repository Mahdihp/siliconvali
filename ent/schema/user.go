package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("mobile").Unique().MaxLen(11),
		field.String("password").MaxLen(50),
		field.String("firstname").MaxLen(100).Optional().Nillable(),
		field.String("lastname").MaxLen(100).Optional().Nillable(),
		field.String("national_code").MaxLen(10).Optional(),
		field.Bool("active").Default(true).Comment("فعال بودن"),
		field.Bool("deleted").Default(false).Comment("حذف منطقی"),

		field.String("address").Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "text", // Override Postgres.
			}).MaxLen(500).Optional().Nillable(),

		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp(6)", // Override Postgres.
			}).Default(time.Now).Immutable(),

		field.Time("updated_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp(6)", // Override Postgres.
			}).Default(time.Now()),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).Ref("users"),

		edge.To("mainiots", MainIot.Type),

		edge.To("userpaymentplans", UserPaymentPlan.Type),
	}
}
