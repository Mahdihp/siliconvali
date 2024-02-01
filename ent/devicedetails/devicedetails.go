// Code generated by ent, DO NOT EDIT.

package devicedetails

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the devicedetails type in the database.
	Label = "device_details"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLightStatus holds the string denoting the light_status field in the database.
	FieldLightStatus = "light_status"
	// EdgeDeviceiotID holds the string denoting the deviceiot_id edge name in mutations.
	EdgeDeviceiotID = "deviceiot_id"
	// Table holds the table name of the devicedetails in the database.
	Table = "device_details"
	// DeviceiotIDTable is the table that holds the deviceiot_id relation/edge.
	DeviceiotIDTable = "device_details"
	// DeviceiotIDInverseTable is the table name for the DeviceIot entity.
	// It exists in this package in order to avoid circular dependency with the "deviceiot" package.
	DeviceiotIDInverseTable = "device_iots"
	// DeviceiotIDColumn is the table column denoting the deviceiot_id relation/edge.
	DeviceiotIDColumn = "device_iot_devicedetails"
)

// Columns holds all SQL columns for devicedetails fields.
var Columns = []string{
	FieldID,
	FieldLightStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "device_details"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"device_iot_devicedetails",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the DeviceDetails queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByLightStatus orders the results by the light_status field.
func ByLightStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLightStatus, opts...).ToFunc()
}

// ByDeviceiotIDField orders the results by deviceiot_id field.
func ByDeviceiotIDField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeviceiotIDStep(), sql.OrderByField(field, opts...))
	}
}
func newDeviceiotIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeviceiotIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DeviceiotIDTable, DeviceiotIDColumn),
	)
}
