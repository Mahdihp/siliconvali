// Code generated by ent, DO NOT EDIT.

package ent

import (
	"siliconvali/ent/deviceiot"
	"siliconvali/ent/mainiot"
	"siliconvali/ent/plan"
	"siliconvali/ent/role"
	"siliconvali/ent/schema"
	"siliconvali/ent/user"
	"siliconvali/ent/userpaymentplan"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	deviceiotFields := schema.DeviceIot{}.Fields()
	_ = deviceiotFields
	// deviceiotDescDisplayName is the schema descriptor for display_name field.
	deviceiotDescDisplayName := deviceiotFields[1].Descriptor()
	// deviceiot.DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	deviceiot.DisplayNameValidator = deviceiotDescDisplayName.Validators[0].(func(string) error)
	// deviceiotDescSerialNumber is the schema descriptor for serial_number field.
	deviceiotDescSerialNumber := deviceiotFields[2].Descriptor()
	// deviceiot.SerialNumberValidator is a validator for the "serial_number" field. It is called by the builders before save.
	deviceiot.SerialNumberValidator = deviceiotDescSerialNumber.Validators[0].(func(string) error)
	// deviceiotDescActive is the schema descriptor for active field.
	deviceiotDescActive := deviceiotFields[5].Descriptor()
	// deviceiot.DefaultActive holds the default value on creation for the active field.
	deviceiot.DefaultActive = deviceiotDescActive.Default.(bool)
	// deviceiotDescCreatedAt is the schema descriptor for created_at field.
	deviceiotDescCreatedAt := deviceiotFields[8].Descriptor()
	// deviceiot.DefaultCreatedAt holds the default value on creation for the created_at field.
	deviceiot.DefaultCreatedAt = deviceiotDescCreatedAt.Default.(func() time.Time)
	// deviceiotDescUpdatedAt is the schema descriptor for updated_at field.
	deviceiotDescUpdatedAt := deviceiotFields[9].Descriptor()
	// deviceiot.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	deviceiot.DefaultUpdatedAt = deviceiotDescUpdatedAt.Default.(time.Time)
	mainiotFields := schema.MainIot{}.Fields()
	_ = mainiotFields
	// mainiotDescDisplayName is the schema descriptor for display_name field.
	mainiotDescDisplayName := mainiotFields[1].Descriptor()
	// mainiot.DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	mainiot.DisplayNameValidator = mainiotDescDisplayName.Validators[0].(func(string) error)
	// mainiotDescAddress is the schema descriptor for address field.
	mainiotDescAddress := mainiotFields[4].Descriptor()
	// mainiot.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	mainiot.AddressValidator = mainiotDescAddress.Validators[0].(func(string) error)
	// mainiotDescSerialNumber is the schema descriptor for serial_number field.
	mainiotDescSerialNumber := mainiotFields[5].Descriptor()
	// mainiot.SerialNumberValidator is a validator for the "serial_number" field. It is called by the builders before save.
	mainiot.SerialNumberValidator = mainiotDescSerialNumber.Validators[0].(func(string) error)
	// mainiotDescMACAddress is the schema descriptor for mac_address field.
	mainiotDescMACAddress := mainiotFields[6].Descriptor()
	// mainiot.MACAddressValidator is a validator for the "mac_address" field. It is called by the builders before save.
	mainiot.MACAddressValidator = mainiotDescMACAddress.Validators[0].(func(string) error)
	// mainiotDescIPRemote is the schema descriptor for ip_remote field.
	mainiotDescIPRemote := mainiotFields[7].Descriptor()
	// mainiot.IPRemoteValidator is a validator for the "ip_remote" field. It is called by the builders before save.
	mainiot.IPRemoteValidator = mainiotDescIPRemote.Validators[0].(func(string) error)
	// mainiotDescActive is the schema descriptor for active field.
	mainiotDescActive := mainiotFields[9].Descriptor()
	// mainiot.DefaultActive holds the default value on creation for the active field.
	mainiot.DefaultActive = mainiotDescActive.Default.(bool)
	// mainiotDescCreatedAt is the schema descriptor for created_at field.
	mainiotDescCreatedAt := mainiotFields[10].Descriptor()
	// mainiot.DefaultCreatedAt holds the default value on creation for the created_at field.
	mainiot.DefaultCreatedAt = mainiotDescCreatedAt.Default.(func() time.Time)
	// mainiotDescUpdatedAt is the schema descriptor for updated_at field.
	mainiotDescUpdatedAt := mainiotFields[11].Descriptor()
	// mainiot.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	mainiot.DefaultUpdatedAt = mainiotDescUpdatedAt.Default.(time.Time)
	planFields := schema.Plan{}.Fields()
	_ = planFields
	// planDescName is the schema descriptor for name field.
	planDescName := planFields[1].Descriptor()
	// plan.NameValidator is a validator for the "name" field. It is called by the builders before save.
	plan.NameValidator = planDescName.Validators[0].(func(string) error)
	// planDescActive is the schema descriptor for active field.
	planDescActive := planFields[4].Descriptor()
	// plan.DefaultActive holds the default value on creation for the active field.
	plan.DefaultActive = planDescActive.Default.(bool)
	// planDescDescription is the schema descriptor for description field.
	planDescDescription := planFields[5].Descriptor()
	// plan.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	plan.DescriptionValidator = planDescDescription.Validators[0].(func(string) error)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescName is the schema descriptor for name field.
	roleDescName := roleFields[1].Descriptor()
	// role.NameValidator is a validator for the "name" field. It is called by the builders before save.
	role.NameValidator = roleDescName.Validators[0].(func(string) error)
	// roleDescDescription is the schema descriptor for description field.
	roleDescDescription := roleFields[2].Descriptor()
	// role.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	role.DescriptionValidator = roleDescDescription.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescFirstname is the schema descriptor for firstname field.
	userDescFirstname := userFields[2].Descriptor()
	// user.FirstnameValidator is a validator for the "firstname" field. It is called by the builders before save.
	user.FirstnameValidator = userDescFirstname.Validators[0].(func(string) error)
	// userDescLastname is the schema descriptor for lastname field.
	userDescLastname := userFields[3].Descriptor()
	// user.LastnameValidator is a validator for the "lastname" field. It is called by the builders before save.
	user.LastnameValidator = userDescLastname.Validators[0].(func(string) error)
	// userDescMobile is the schema descriptor for mobile field.
	userDescMobile := userFields[4].Descriptor()
	// user.MobileValidator is a validator for the "mobile" field. It is called by the builders before save.
	user.MobileValidator = userDescMobile.Validators[0].(func(string) error)
	// userDescNationalCode is the schema descriptor for national_code field.
	userDescNationalCode := userFields[5].Descriptor()
	// user.NationalCodeValidator is a validator for the "national_code" field. It is called by the builders before save.
	user.NationalCodeValidator = userDescNationalCode.Validators[0].(func(string) error)
	// userDescActive is the schema descriptor for active field.
	userDescActive := userFields[6].Descriptor()
	// user.DefaultActive holds the default value on creation for the active field.
	user.DefaultActive = userDescActive.Default.(bool)
	// userDescDeleted is the schema descriptor for deleted field.
	userDescDeleted := userFields[7].Descriptor()
	// user.DefaultDeleted holds the default value on creation for the deleted field.
	user.DefaultDeleted = userDescDeleted.Default.(bool)
	// userDescAddress is the schema descriptor for address field.
	userDescAddress := userFields[8].Descriptor()
	// user.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	user.AddressValidator = userDescAddress.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[9].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[10].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(time.Time)
	userpaymentplanFields := schema.UserPaymentPlan{}.Fields()
	_ = userpaymentplanFields
	// userpaymentplanDescDeleted is the schema descriptor for deleted field.
	userpaymentplanDescDeleted := userpaymentplanFields[6].Descriptor()
	// userpaymentplan.DefaultDeleted holds the default value on creation for the deleted field.
	userpaymentplan.DefaultDeleted = userpaymentplanDescDeleted.Default.(bool)
	// userpaymentplanDescCreatedAt is the schema descriptor for created_at field.
	userpaymentplanDescCreatedAt := userpaymentplanFields[7].Descriptor()
	// userpaymentplan.DefaultCreatedAt holds the default value on creation for the created_at field.
	userpaymentplan.DefaultCreatedAt = userpaymentplanDescCreatedAt.Default.(func() time.Time)
}
