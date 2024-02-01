// Code generated by ent, DO NOT EDIT.

package deviceiot

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the deviceiot type in the database.
	Label = "device_iot"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldSerialNumber holds the string denoting the serial_number field in the database.
	FieldSerialNumber = "serial_number"
	// FieldTypeDevice holds the string denoting the type_device field in the database.
	FieldTypeDevice = "type_device"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldLat holds the string denoting the lat field in the database.
	FieldLat = "lat"
	// FieldLon holds the string denoting the lon field in the database.
	FieldLon = "lon"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeMainiotID holds the string denoting the mainiot_id edge name in mutations.
	EdgeMainiotID = "mainiot_id"
	// EdgeDevicedetails holds the string denoting the devicedetails edge name in mutations.
	EdgeDevicedetails = "devicedetails"
	// Table holds the table name of the deviceiot in the database.
	Table = "device_iots"
	// MainiotIDTable is the table that holds the mainiot_id relation/edge.
	MainiotIDTable = "device_iots"
	// MainiotIDInverseTable is the table name for the MainIot entity.
	// It exists in this package in order to avoid circular dependency with the "mainiot" package.
	MainiotIDInverseTable = "main_iots"
	// MainiotIDColumn is the table column denoting the mainiot_id relation/edge.
	MainiotIDColumn = "main_iot_deviceiots"
	// DevicedetailsTable is the table that holds the devicedetails relation/edge.
	DevicedetailsTable = "device_details"
	// DevicedetailsInverseTable is the table name for the DeviceDetails entity.
	// It exists in this package in order to avoid circular dependency with the "devicedetails" package.
	DevicedetailsInverseTable = "device_details"
	// DevicedetailsColumn is the table column denoting the devicedetails relation/edge.
	DevicedetailsColumn = "device_iot_devicedetails"
)

// Columns holds all SQL columns for deviceiot fields.
var Columns = []string{
	FieldID,
	FieldDisplayName,
	FieldSerialNumber,
	FieldTypeDevice,
	FieldStatus,
	FieldActive,
	FieldLat,
	FieldLon,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "device_iots"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"main_iot_deviceiots",
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

var (
	// DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	DisplayNameValidator func(string) error
	// SerialNumberValidator is a validator for the "serial_number" field. It is called by the builders before save.
	SerialNumberValidator func(string) error
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt time.Time
)

// OrderOption defines the ordering options for the DeviceIot queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDisplayName orders the results by the display_name field.
func ByDisplayName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisplayName, opts...).ToFunc()
}

// BySerialNumber orders the results by the serial_number field.
func BySerialNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSerialNumber, opts...).ToFunc()
}

// ByTypeDevice orders the results by the type_device field.
func ByTypeDevice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTypeDevice, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByActive orders the results by the active field.
func ByActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActive, opts...).ToFunc()
}

// ByLat orders the results by the lat field.
func ByLat(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLat, opts...).ToFunc()
}

// ByLon orders the results by the lon field.
func ByLon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLon, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByMainiotIDField orders the results by mainiot_id field.
func ByMainiotIDField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMainiotIDStep(), sql.OrderByField(field, opts...))
	}
}

// ByDevicedetailsCount orders the results by devicedetails count.
func ByDevicedetailsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDevicedetailsStep(), opts...)
	}
}

// ByDevicedetails orders the results by devicedetails terms.
func ByDevicedetails(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDevicedetailsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMainiotIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MainiotIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MainiotIDTable, MainiotIDColumn),
	)
}
func newDevicedetailsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DevicedetailsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DevicedetailsTable, DevicedetailsColumn),
	)
}
