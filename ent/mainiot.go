// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"siliconvali/ent/mainiot"
	"siliconvali/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// MainIot is the model entity for the MainIot schema.
type MainIot struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// Lat holds the value of the "lat" field.
	Lat *float64 `json:"lat,omitempty"`
	// Lon holds the value of the "lon" field.
	Lon *float64 `json:"lon,omitempty"`
	// Address holds the value of the "address" field.
	Address *string `json:"address,omitempty"`
	// SerialNumber holds the value of the "serial_number" field.
	SerialNumber *string `json:"serial_number,omitempty"`
	// MACAddress holds the value of the "mac_address" field.
	MACAddress *string `json:"mac_address,omitempty"`
	// IPRemote holds the value of the "ip_remote" field.
	IPRemote *string `json:"ip_remote,omitempty"`
	// وضعیت
	Status *string `json:"status,omitempty"`
	// فعال بودن
	Active bool `json:"active,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MainIotQuery when eager-loading is set.
	Edges         MainIotEdges `json:"edges"`
	user_mainiots *int64
	selectValues  sql.SelectValues
}

// MainIotEdges holds the relations/edges for other nodes in the graph.
type MainIotEdges struct {
	// Deviceiots holds the value of the deviceiots edge.
	Deviceiots []*DeviceIot `json:"deviceiots,omitempty"`
	// UserID holds the value of the user_id edge.
	UserID *User `json:"user_id,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DeviceiotsOrErr returns the Deviceiots value or an error if the edge
// was not loaded in eager-loading.
func (e MainIotEdges) DeviceiotsOrErr() ([]*DeviceIot, error) {
	if e.loadedTypes[0] {
		return e.Deviceiots, nil
	}
	return nil, &NotLoadedError{edge: "deviceiots"}
}

// UserIDOrErr returns the UserID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MainIotEdges) UserIDOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.UserID == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.UserID, nil
	}
	return nil, &NotLoadedError{edge: "user_id"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MainIot) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case mainiot.FieldActive:
			values[i] = new(sql.NullBool)
		case mainiot.FieldLat, mainiot.FieldLon:
			values[i] = new(sql.NullFloat64)
		case mainiot.FieldID:
			values[i] = new(sql.NullInt64)
		case mainiot.FieldDisplayName, mainiot.FieldAddress, mainiot.FieldSerialNumber, mainiot.FieldMACAddress, mainiot.FieldIPRemote, mainiot.FieldStatus:
			values[i] = new(sql.NullString)
		case mainiot.FieldCreatedAt, mainiot.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case mainiot.ForeignKeys[0]: // user_mainiots
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MainIot fields.
func (mi *MainIot) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mainiot.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mi.ID = int64(value.Int64)
		case mainiot.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				mi.DisplayName = value.String
			}
		case mainiot.FieldLat:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field lat", values[i])
			} else if value.Valid {
				mi.Lat = new(float64)
				*mi.Lat = value.Float64
			}
		case mainiot.FieldLon:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field lon", values[i])
			} else if value.Valid {
				mi.Lon = new(float64)
				*mi.Lon = value.Float64
			}
		case mainiot.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				mi.Address = new(string)
				*mi.Address = value.String
			}
		case mainiot.FieldSerialNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial_number", values[i])
			} else if value.Valid {
				mi.SerialNumber = new(string)
				*mi.SerialNumber = value.String
			}
		case mainiot.FieldMACAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mac_address", values[i])
			} else if value.Valid {
				mi.MACAddress = new(string)
				*mi.MACAddress = value.String
			}
		case mainiot.FieldIPRemote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip_remote", values[i])
			} else if value.Valid {
				mi.IPRemote = new(string)
				*mi.IPRemote = value.String
			}
		case mainiot.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				mi.Status = new(string)
				*mi.Status = value.String
			}
		case mainiot.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				mi.Active = value.Bool
			}
		case mainiot.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mi.CreatedAt = value.Time
			}
		case mainiot.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mi.UpdatedAt = value.Time
			}
		case mainiot.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_mainiots", value)
			} else if value.Valid {
				mi.user_mainiots = new(int64)
				*mi.user_mainiots = int64(value.Int64)
			}
		default:
			mi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MainIot.
// This includes values selected through modifiers, order, etc.
func (mi *MainIot) Value(name string) (ent.Value, error) {
	return mi.selectValues.Get(name)
}

// QueryDeviceiots queries the "deviceiots" edge of the MainIot entity.
func (mi *MainIot) QueryDeviceiots() *DeviceIotQuery {
	return NewMainIotClient(mi.config).QueryDeviceiots(mi)
}

// QueryUserID queries the "user_id" edge of the MainIot entity.
func (mi *MainIot) QueryUserID() *UserQuery {
	return NewMainIotClient(mi.config).QueryUserID(mi)
}

// Update returns a builder for updating this MainIot.
// Note that you need to call MainIot.Unwrap() before calling this method if this MainIot
// was returned from a transaction, and the transaction was committed or rolled back.
func (mi *MainIot) Update() *MainIotUpdateOne {
	return NewMainIotClient(mi.config).UpdateOne(mi)
}

// Unwrap unwraps the MainIot entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mi *MainIot) Unwrap() *MainIot {
	_tx, ok := mi.config.driver.(*txDriver)
	if !ok {
		panic("ent: MainIot is not a transactional entity")
	}
	mi.config.driver = _tx.drv
	return mi
}

// String implements the fmt.Stringer.
func (mi *MainIot) String() string {
	var builder strings.Builder
	builder.WriteString("MainIot(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mi.ID))
	builder.WriteString("display_name=")
	builder.WriteString(mi.DisplayName)
	builder.WriteString(", ")
	if v := mi.Lat; v != nil {
		builder.WriteString("lat=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := mi.Lon; v != nil {
		builder.WriteString("lon=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := mi.Address; v != nil {
		builder.WriteString("address=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := mi.SerialNumber; v != nil {
		builder.WriteString("serial_number=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := mi.MACAddress; v != nil {
		builder.WriteString("mac_address=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := mi.IPRemote; v != nil {
		builder.WriteString("ip_remote=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := mi.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", mi.Active))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(mi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mi.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MainIots is a parsable slice of MainIot.
type MainIots []*MainIot
