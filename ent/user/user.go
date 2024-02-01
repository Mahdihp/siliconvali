// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldFirstname holds the string denoting the firstname field in the database.
	FieldFirstname = "firstname"
	// FieldLastname holds the string denoting the lastname field in the database.
	FieldLastname = "lastname"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldNationalCode holds the string denoting the national_code field in the database.
	FieldNationalCode = "national_code"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldDeleted holds the string denoting the deleted field in the database.
	FieldDeleted = "deleted"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// EdgeMainiots holds the string denoting the mainiots edge name in mutations.
	EdgeMainiots = "mainiots"
	// EdgeUserpaymentplans holds the string denoting the userpaymentplans edge name in mutations.
	EdgeUserpaymentplans = "userpaymentplans"
	// Table holds the table name of the user in the database.
	Table = "users"
	// RolesTable is the table that holds the roles relation/edge. The primary key declared below.
	RolesTable = "role_users"
	// RolesInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RolesInverseTable = "roles"
	// MainiotsTable is the table that holds the mainiots relation/edge.
	MainiotsTable = "main_iots"
	// MainiotsInverseTable is the table name for the MainIot entity.
	// It exists in this package in order to avoid circular dependency with the "mainiot" package.
	MainiotsInverseTable = "main_iots"
	// MainiotsColumn is the table column denoting the mainiots relation/edge.
	MainiotsColumn = "user_mainiots"
	// UserpaymentplansTable is the table that holds the userpaymentplans relation/edge.
	UserpaymentplansTable = "user_payment_plans"
	// UserpaymentplansInverseTable is the table name for the UserPaymentPlan entity.
	// It exists in this package in order to avoid circular dependency with the "userpaymentplan" package.
	UserpaymentplansInverseTable = "user_payment_plans"
	// UserpaymentplansColumn is the table column denoting the userpaymentplans relation/edge.
	UserpaymentplansColumn = "user_userpaymentplans"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldFirstname,
	FieldLastname,
	FieldMobile,
	FieldNationalCode,
	FieldActive,
	FieldDeleted,
	FieldAddress,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// RolesPrimaryKey and RolesColumn2 are the table columns denoting the
	// primary key for the roles relation (M2M).
	RolesPrimaryKey = []string{"role_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// FirstnameValidator is a validator for the "firstname" field. It is called by the builders before save.
	FirstnameValidator func(string) error
	// LastnameValidator is a validator for the "lastname" field. It is called by the builders before save.
	LastnameValidator func(string) error
	// MobileValidator is a validator for the "mobile" field. It is called by the builders before save.
	MobileValidator func(string) error
	// NationalCodeValidator is a validator for the "national_code" field. It is called by the builders before save.
	NationalCodeValidator func(string) error
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// DefaultDeleted holds the default value on creation for the "deleted" field.
	DefaultDeleted bool
	// AddressValidator is a validator for the "address" field. It is called by the builders before save.
	AddressValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByFirstname orders the results by the firstname field.
func ByFirstname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstname, opts...).ToFunc()
}

// ByLastname orders the results by the lastname field.
func ByLastname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastname, opts...).ToFunc()
}

// ByMobile orders the results by the mobile field.
func ByMobile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobile, opts...).ToFunc()
}

// ByNationalCode orders the results by the national_code field.
func ByNationalCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNationalCode, opts...).ToFunc()
}

// ByActive orders the results by the active field.
func ByActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActive, opts...).ToFunc()
}

// ByDeleted orders the results by the deleted field.
func ByDeleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleted, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByRolesCount orders the results by roles count.
func ByRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRolesStep(), opts...)
	}
}

// ByRoles orders the results by roles terms.
func ByRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMainiotsCount orders the results by mainiots count.
func ByMainiotsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMainiotsStep(), opts...)
	}
}

// ByMainiots orders the results by mainiots terms.
func ByMainiots(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMainiotsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserpaymentplansCount orders the results by userpaymentplans count.
func ByUserpaymentplansCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserpaymentplansStep(), opts...)
	}
}

// ByUserpaymentplans orders the results by userpaymentplans terms.
func ByUserpaymentplans(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserpaymentplansStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
	)
}
func newMainiotsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MainiotsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MainiotsTable, MainiotsColumn),
	)
}
func newUserpaymentplansStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserpaymentplansInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserpaymentplansTable, UserpaymentplansColumn),
	)
}
