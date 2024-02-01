// Code generated by ent, DO NOT EDIT.

package mainiot

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the mainiot type in the database.
	Label = "main_iot"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldLat holds the string denoting the lat field in the database.
	FieldLat = "lat"
	// FieldLon holds the string denoting the lon field in the database.
	FieldLon = "lon"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldSerialNumber holds the string denoting the serial_number field in the database.
	FieldSerialNumber = "serial_number"
	// FieldMACAddress holds the string denoting the mac_address field in the database.
	FieldMACAddress = "mac_address"
	// FieldIPRemote holds the string denoting the ip_remote field in the database.
	FieldIPRemote = "ip_remote"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeDeviceiots holds the string denoting the deviceiots edge name in mutations.
	EdgeDeviceiots = "deviceiots"
	// EdgeUserID holds the string denoting the user_id edge name in mutations.
	EdgeUserID = "user_id"
	// Table holds the table name of the mainiot in the database.
	Table = "main_iots"
	// DeviceiotsTable is the table that holds the deviceiots relation/edge.
	DeviceiotsTable = "device_iots"
	// DeviceiotsInverseTable is the table name for the DeviceIot entity.
	// It exists in this package in order to avoid circular dependency with the "deviceiot" package.
	DeviceiotsInverseTable = "device_iots"
	// DeviceiotsColumn is the table column denoting the deviceiots relation/edge.
	DeviceiotsColumn = "main_iot_deviceiots"
	// UserIDTable is the table that holds the user_id relation/edge.
	UserIDTable = "main_iots"
	// UserIDInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserIDInverseTable = "users"
	// UserIDColumn is the table column denoting the user_id relation/edge.
	UserIDColumn = "user_mainiots"
)

// Columns holds all SQL columns for mainiot fields.
var Columns = []string{
	FieldID,
	FieldDisplayName,
	FieldLat,
	FieldLon,
	FieldAddress,
	FieldSerialNumber,
	FieldMACAddress,
	FieldIPRemote,
	FieldStatus,
	FieldActive,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "main_iots"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_mainiots",
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
	// AddressValidator is a validator for the "address" field. It is called by the builders before save.
	AddressValidator func(string) error
	// SerialNumberValidator is a validator for the "serial_number" field. It is called by the builders before save.
	SerialNumberValidator func(string) error
	// MACAddressValidator is a validator for the "mac_address" field. It is called by the builders before save.
	MACAddressValidator func(string) error
	// IPRemoteValidator is a validator for the "ip_remote" field. It is called by the builders before save.
	IPRemoteValidator func(string) error
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt time.Time
)

// OrderOption defines the ordering options for the MainIot queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDisplayName orders the results by the display_name field.
func ByDisplayName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisplayName, opts...).ToFunc()
}

// ByLat orders the results by the lat field.
func ByLat(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLat, opts...).ToFunc()
}

// ByLon orders the results by the lon field.
func ByLon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLon, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// BySerialNumber orders the results by the serial_number field.
func BySerialNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSerialNumber, opts...).ToFunc()
}

// ByMACAddress orders the results by the mac_address field.
func ByMACAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMACAddress, opts...).ToFunc()
}

// ByIPRemote orders the results by the ip_remote field.
func ByIPRemote(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIPRemote, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByActive orders the results by the active field.
func ByActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActive, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeviceiotsCount orders the results by deviceiots count.
func ByDeviceiotsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDeviceiotsStep(), opts...)
	}
}

// ByDeviceiots orders the results by deviceiots terms.
func ByDeviceiots(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeviceiotsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserIDField orders the results by user_id field.
func ByUserIDField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserIDStep(), sql.OrderByField(field, opts...))
	}
}
func newDeviceiotsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeviceiotsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DeviceiotsTable, DeviceiotsColumn),
	)
}
func newUserIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserIDTable, UserIDColumn),
	)
}
