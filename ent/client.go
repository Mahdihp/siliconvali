// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"siliconvali/ent/migrate"

	"siliconvali/ent/deviceiot"
	"siliconvali/ent/mainiot"
	"siliconvali/ent/role"
	"siliconvali/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// DeviceIot is the client for interacting with the DeviceIot builders.
	DeviceIot *DeviceIotClient
	// MainIot is the client for interacting with the MainIot builders.
	MainIot *MainIotClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.DeviceIot = NewDeviceIotClient(c.config)
	c.MainIot = NewMainIotClient(c.config)
	c.Role = NewRoleClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		DeviceIot: NewDeviceIotClient(cfg),
		MainIot:   NewMainIotClient(cfg),
		Role:      NewRoleClient(cfg),
		User:      NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		DeviceIot: NewDeviceIotClient(cfg),
		MainIot:   NewMainIotClient(cfg),
		Role:      NewRoleClient(cfg),
		User:      NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		DeviceIot.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.DeviceIot.Use(hooks...)
	c.MainIot.Use(hooks...)
	c.Role.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.DeviceIot.Intercept(interceptors...)
	c.MainIot.Intercept(interceptors...)
	c.Role.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *DeviceIotMutation:
		return c.DeviceIot.mutate(ctx, m)
	case *MainIotMutation:
		return c.MainIot.mutate(ctx, m)
	case *RoleMutation:
		return c.Role.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// DeviceIotClient is a client for the DeviceIot schema.
type DeviceIotClient struct {
	config
}

// NewDeviceIotClient returns a client for the DeviceIot from the given config.
func NewDeviceIotClient(c config) *DeviceIotClient {
	return &DeviceIotClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `deviceiot.Hooks(f(g(h())))`.
func (c *DeviceIotClient) Use(hooks ...Hook) {
	c.hooks.DeviceIot = append(c.hooks.DeviceIot, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `deviceiot.Intercept(f(g(h())))`.
func (c *DeviceIotClient) Intercept(interceptors ...Interceptor) {
	c.inters.DeviceIot = append(c.inters.DeviceIot, interceptors...)
}

// Create returns a builder for creating a DeviceIot entity.
func (c *DeviceIotClient) Create() *DeviceIotCreate {
	mutation := newDeviceIotMutation(c.config, OpCreate)
	return &DeviceIotCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of DeviceIot entities.
func (c *DeviceIotClient) CreateBulk(builders ...*DeviceIotCreate) *DeviceIotCreateBulk {
	return &DeviceIotCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *DeviceIotClient) MapCreateBulk(slice any, setFunc func(*DeviceIotCreate, int)) *DeviceIotCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &DeviceIotCreateBulk{err: fmt.Errorf("calling to DeviceIotClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*DeviceIotCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &DeviceIotCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DeviceIot.
func (c *DeviceIotClient) Update() *DeviceIotUpdate {
	mutation := newDeviceIotMutation(c.config, OpUpdate)
	return &DeviceIotUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DeviceIotClient) UpdateOne(di *DeviceIot) *DeviceIotUpdateOne {
	mutation := newDeviceIotMutation(c.config, OpUpdateOne, withDeviceIot(di))
	return &DeviceIotUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DeviceIotClient) UpdateOneID(id int64) *DeviceIotUpdateOne {
	mutation := newDeviceIotMutation(c.config, OpUpdateOne, withDeviceIotID(id))
	return &DeviceIotUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DeviceIot.
func (c *DeviceIotClient) Delete() *DeviceIotDelete {
	mutation := newDeviceIotMutation(c.config, OpDelete)
	return &DeviceIotDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DeviceIotClient) DeleteOne(di *DeviceIot) *DeviceIotDeleteOne {
	return c.DeleteOneID(di.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DeviceIotClient) DeleteOneID(id int64) *DeviceIotDeleteOne {
	builder := c.Delete().Where(deviceiot.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DeviceIotDeleteOne{builder}
}

// Query returns a query builder for DeviceIot.
func (c *DeviceIotClient) Query() *DeviceIotQuery {
	return &DeviceIotQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeDeviceIot},
		inters: c.Interceptors(),
	}
}

// Get returns a DeviceIot entity by its id.
func (c *DeviceIotClient) Get(ctx context.Context, id int64) (*DeviceIot, error) {
	return c.Query().Where(deviceiot.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DeviceIotClient) GetX(ctx context.Context, id int64) *DeviceIot {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a DeviceIot.
func (c *DeviceIotClient) QueryOwner(di *DeviceIot) *MainIotQuery {
	query := (&MainIotClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := di.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deviceiot.Table, deviceiot.FieldID, id),
			sqlgraph.To(mainiot.Table, mainiot.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deviceiot.OwnerTable, deviceiot.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(di.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DeviceIotClient) Hooks() []Hook {
	return c.hooks.DeviceIot
}

// Interceptors returns the client interceptors.
func (c *DeviceIotClient) Interceptors() []Interceptor {
	return c.inters.DeviceIot
}

func (c *DeviceIotClient) mutate(ctx context.Context, m *DeviceIotMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&DeviceIotCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&DeviceIotUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&DeviceIotUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&DeviceIotDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown DeviceIot mutation op: %q", m.Op())
	}
}

// MainIotClient is a client for the MainIot schema.
type MainIotClient struct {
	config
}

// NewMainIotClient returns a client for the MainIot from the given config.
func NewMainIotClient(c config) *MainIotClient {
	return &MainIotClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `mainiot.Hooks(f(g(h())))`.
func (c *MainIotClient) Use(hooks ...Hook) {
	c.hooks.MainIot = append(c.hooks.MainIot, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `mainiot.Intercept(f(g(h())))`.
func (c *MainIotClient) Intercept(interceptors ...Interceptor) {
	c.inters.MainIot = append(c.inters.MainIot, interceptors...)
}

// Create returns a builder for creating a MainIot entity.
func (c *MainIotClient) Create() *MainIotCreate {
	mutation := newMainIotMutation(c.config, OpCreate)
	return &MainIotCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MainIot entities.
func (c *MainIotClient) CreateBulk(builders ...*MainIotCreate) *MainIotCreateBulk {
	return &MainIotCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MainIotClient) MapCreateBulk(slice any, setFunc func(*MainIotCreate, int)) *MainIotCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MainIotCreateBulk{err: fmt.Errorf("calling to MainIotClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MainIotCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MainIotCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MainIot.
func (c *MainIotClient) Update() *MainIotUpdate {
	mutation := newMainIotMutation(c.config, OpUpdate)
	return &MainIotUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MainIotClient) UpdateOne(mi *MainIot) *MainIotUpdateOne {
	mutation := newMainIotMutation(c.config, OpUpdateOne, withMainIot(mi))
	return &MainIotUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MainIotClient) UpdateOneID(id int64) *MainIotUpdateOne {
	mutation := newMainIotMutation(c.config, OpUpdateOne, withMainIotID(id))
	return &MainIotUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MainIot.
func (c *MainIotClient) Delete() *MainIotDelete {
	mutation := newMainIotMutation(c.config, OpDelete)
	return &MainIotDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MainIotClient) DeleteOne(mi *MainIot) *MainIotDeleteOne {
	return c.DeleteOneID(mi.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MainIotClient) DeleteOneID(id int64) *MainIotDeleteOne {
	builder := c.Delete().Where(mainiot.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MainIotDeleteOne{builder}
}

// Query returns a query builder for MainIot.
func (c *MainIotClient) Query() *MainIotQuery {
	return &MainIotQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMainIot},
		inters: c.Interceptors(),
	}
}

// Get returns a MainIot entity by its id.
func (c *MainIotClient) Get(ctx context.Context, id int64) (*MainIot, error) {
	return c.Query().Where(mainiot.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MainIotClient) GetX(ctx context.Context, id int64) *MainIot {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDeviceiots queries the deviceiots edge of a MainIot.
func (c *MainIotClient) QueryDeviceiots(mi *MainIot) *DeviceIotQuery {
	query := (&DeviceIotClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := mi.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(mainiot.Table, mainiot.FieldID, id),
			sqlgraph.To(deviceiot.Table, deviceiot.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, mainiot.DeviceiotsTable, mainiot.DeviceiotsColumn),
		)
		fromV = sqlgraph.Neighbors(mi.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOwner queries the owner edge of a MainIot.
func (c *MainIotClient) QueryOwner(mi *MainIot) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := mi.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(mainiot.Table, mainiot.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, mainiot.OwnerTable, mainiot.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(mi.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MainIotClient) Hooks() []Hook {
	return c.hooks.MainIot
}

// Interceptors returns the client interceptors.
func (c *MainIotClient) Interceptors() []Interceptor {
	return c.inters.MainIot
}

func (c *MainIotClient) mutate(ctx context.Context, m *MainIotMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MainIotCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MainIotUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MainIotUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MainIotDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown MainIot mutation op: %q", m.Op())
	}
}

// RoleClient is a client for the Role schema.
type RoleClient struct {
	config
}

// NewRoleClient returns a client for the Role from the given config.
func NewRoleClient(c config) *RoleClient {
	return &RoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `role.Hooks(f(g(h())))`.
func (c *RoleClient) Use(hooks ...Hook) {
	c.hooks.Role = append(c.hooks.Role, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `role.Intercept(f(g(h())))`.
func (c *RoleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Role = append(c.inters.Role, interceptors...)
}

// Create returns a builder for creating a Role entity.
func (c *RoleClient) Create() *RoleCreate {
	mutation := newRoleMutation(c.config, OpCreate)
	return &RoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Role entities.
func (c *RoleClient) CreateBulk(builders ...*RoleCreate) *RoleCreateBulk {
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RoleClient) MapCreateBulk(slice any, setFunc func(*RoleCreate, int)) *RoleCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RoleCreateBulk{err: fmt.Errorf("calling to RoleClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RoleCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Role.
func (c *RoleClient) Update() *RoleUpdate {
	mutation := newRoleMutation(c.config, OpUpdate)
	return &RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleClient) UpdateOne(r *Role) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRole(r))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleClient) UpdateOneID(id int16) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRoleID(id))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Role.
func (c *RoleClient) Delete() *RoleDelete {
	mutation := newRoleMutation(c.config, OpDelete)
	return &RoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RoleClient) DeleteOne(r *Role) *RoleDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RoleClient) DeleteOneID(id int16) *RoleDeleteOne {
	builder := c.Delete().Where(role.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleDeleteOne{builder}
}

// Query returns a query builder for Role.
func (c *RoleClient) Query() *RoleQuery {
	return &RoleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRole},
		inters: c.Interceptors(),
	}
}

// Get returns a Role entity by its id.
func (c *RoleClient) Get(ctx context.Context, id int16) (*Role, error) {
	return c.Query().Where(role.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleClient) GetX(ctx context.Context, id int16) *Role {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Role.
func (c *RoleClient) QueryUsers(r *Role) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, role.UsersTable, role.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoleClient) Hooks() []Hook {
	return c.hooks.Role
}

// Interceptors returns the client interceptors.
func (c *RoleClient) Interceptors() []Interceptor {
	return c.inters.Role
}

func (c *RoleClient) mutate(ctx context.Context, m *RoleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RoleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RoleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Role mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int64) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int64) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int64) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int64) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRoles queries the roles edge of a User.
func (c *UserClient) QueryRoles(u *User) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.RolesTable, user.RolesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMainiots queries the mainiots edge of a User.
func (c *UserClient) QueryMainiots(u *User) *MainIotQuery {
	query := (&MainIotClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(mainiot.Table, mainiot.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.MainiotsTable, user.MainiotsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		DeviceIot, MainIot, Role, User []ent.Hook
	}
	inters struct {
		DeviceIot, MainIot, Role, User []ent.Interceptor
	}
)
