package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// MainIot holds the schema definition for the MainIot entity.
type MainIot struct {
	ent.Schema
}

// Fields of the MainIot.
func (MainIot) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("display_name").MaxLen(50),

		field.Float("lat").
			SchemaType(map[string]string{
				dialect.Postgres: "point", // Override Postgres.
			}).Optional().Nillable(),

		field.Float("lon").
			SchemaType(map[string]string{
				dialect.Postgres: "point", // Override Postgres.
			}).Optional().Nillable(),

		field.String("address").
			SchemaType(map[string]string{
				dialect.Postgres: "text", // Override Postgres.
			}).MaxLen(500).Optional().Nillable(),

		field.String("serial_number").MaxLen(15).Optional().Nillable(),
		field.String("mac_address").MaxLen(15).Optional().Nillable(),
		field.String("ip_remote").MaxLen(40).Optional().Nillable(),

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

// Edges of the MainIot.
func (MainIot) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("deviceiots", DeviceIot.Type),
		edge.From("owner", User.Type).Ref("mainiots").Unique(),
	}
}
