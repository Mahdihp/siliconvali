// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DeviceIotsColumns holds the columns for the "device_iots" table.
	DeviceIotsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "display_name", Type: field.TypeString, Size: 50},
		{Name: "serial_number", Type: field.TypeString, Nullable: true, Size: 20},
		{Name: "type_device", Type: field.TypeInt, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp(6)"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp(6)"}},
		{Name: "main_iot_deviceiots", Type: field.TypeInt64, Nullable: true},
	}
	// DeviceIotsTable holds the schema information for the "device_iots" table.
	DeviceIotsTable = &schema.Table{
		Name:       "device_iots",
		Columns:    DeviceIotsColumns,
		PrimaryKey: []*schema.Column{DeviceIotsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "device_iots_main_iots_deviceiots",
				Columns:    []*schema.Column{DeviceIotsColumns[6]},
				RefColumns: []*schema.Column{MainIotsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MainIotsColumns holds the columns for the "main_iots" table.
	MainIotsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "display_name", Type: field.TypeString, Size: 50},
		{Name: "lat", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "point"}},
		{Name: "lon", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "point"}},
		{Name: "address", Type: field.TypeString, Nullable: true, Size: 500, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "serial_number", Type: field.TypeString, Nullable: true, Size: 15},
		{Name: "mac_address", Type: field.TypeString, Nullable: true, Size: 15},
		{Name: "ip_remote", Type: field.TypeString, Nullable: true, Size: 40},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp(6)"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp(6)"}},
		{Name: "user_mainiots", Type: field.TypeInt64, Nullable: true},
	}
	// MainIotsTable holds the schema information for the "main_iots" table.
	MainIotsTable = &schema.Table{
		Name:       "main_iots",
		Columns:    MainIotsColumns,
		PrimaryKey: []*schema.Column{MainIotsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "main_iots_users_mainiots",
				Columns:    []*schema.Column{MainIotsColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt16, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 25},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 100},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "role_id",
				Unique:  true,
				Columns: []*schema.Column{RolesColumns[0]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 50},
		{Name: "firstname", Type: field.TypeString, Nullable: true, Size: 100},
		{Name: "lastname", Type: field.TypeString, Nullable: true, Size: 100},
		{Name: "mobile", Type: field.TypeString, Nullable: true, Size: 11},
		{Name: "national_code", Type: field.TypeString, Nullable: true, Size: 10},
		{Name: "address", Type: field.TypeString, Nullable: true, Size: 500, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp(6)"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp(6)"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_id",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[0]},
			},
		},
	}
	// RoleUsersColumns holds the columns for the "role_users" table.
	RoleUsersColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeInt16},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// RoleUsersTable holds the schema information for the "role_users" table.
	RoleUsersTable = &schema.Table{
		Name:       "role_users",
		Columns:    RoleUsersColumns,
		PrimaryKey: []*schema.Column{RoleUsersColumns[0], RoleUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_users_role_id",
				Columns:    []*schema.Column{RoleUsersColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_users_user_id",
				Columns:    []*schema.Column{RoleUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DeviceIotsTable,
		MainIotsTable,
		RolesTable,
		UsersTable,
		RoleUsersTable,
	}
)

func init() {
	DeviceIotsTable.ForeignKeys[0].RefTable = MainIotsTable
	MainIotsTable.ForeignKeys[0].RefTable = UsersTable
	RoleUsersTable.ForeignKeys[0].RefTable = RolesTable
	RoleUsersTable.ForeignKeys[1].RefTable = UsersTable
}
