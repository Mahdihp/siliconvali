package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DeviceDetails holds the schema definition for the DeviceDetails entity.
type DeviceDetails struct {
	ent.Schema
}

// Fields of the DeviceDetails.
func (DeviceDetails) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Bool("light_status").Optional().Nillable().Comment("روشن بودن"),
	}
}

// Edges of the DeviceDetails.
func (DeviceDetails) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("deviceiot_id", DeviceIot.Type).Ref("devicedetails").Unique(),
	}
}
