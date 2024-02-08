package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// UserPaymentPlan holds the schema definition for the UserPaymentPlan entity.
type UserPaymentPlan struct {
	ent.Schema
}

// Fields of the UserPaymentPlan.
func (UserPaymentPlan) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("amount").Optional().Comment("مبلغ"),
		field.String("reference_number").Unique().Comment("کد پیگیری"),
		field.String("transaction_number").Optional().Nillable().Comment("کد تراکنش"),
		field.String("source_account_number").Optional().Nillable().Comment("شماره حساب مبدا"),
		field.String("destination_account_number").Optional().Nillable().Comment("شماره حساب مقصد"),
		field.Bool("deleted").Default(false).Comment("حذف منطقی"),

		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp(6)", // Override Postgres.
			}).Default(time.Now).Immutable(),
	}
}

// Edges of the UserPaymentPlan.
func (UserPaymentPlan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_id", User.Type).Ref("userpaymentplans").Unique().Required(),
		edge.From("plan_id", Plan.Type).Ref("userpaymentplans").Unique().Required(),
	}
}
