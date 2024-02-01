package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// DeviceIot holds the schema definition for the DeviceIot entity.
type DeviceIot struct {
	ent.Schema
}

// Fields of the DeviceIot.
func (DeviceIot) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("display_name").MaxLen(50),
		field.String("serial_number").MaxLen(20).Optional().Nillable(),
		field.Int("type_device").Optional().Nillable(),
		field.String("status").Optional().Nillable().Comment("وضعیت"),
		field.Bool("active").Default(true).Comment("فعال بودن"),

		field.Float("lat").
			SchemaType(map[string]string{
				dialect.Postgres: "point",
			}).Optional().Nillable(),

		field.Float("lon").
			SchemaType(map[string]string{
				dialect.Postgres: "point",
			}).Optional().Nillable(),

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

// Edges of the DeviceIot.
func (DeviceIot) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("mainiot_id", MainIot.Type).Ref("deviceiots").Unique(),

		edge.To("devicedetails", DeviceDetails.Type),
	}
}
