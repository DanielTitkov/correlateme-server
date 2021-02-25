// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/migrate"

	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/correlation"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/dataset"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/indicator"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/observation"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/scale"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Correlation is the client for interacting with the Correlation builders.
	Correlation *CorrelationClient
	// Dataset is the client for interacting with the Dataset builders.
	Dataset *DatasetClient
	// Indicator is the client for interacting with the Indicator builders.
	Indicator *IndicatorClient
	// Observation is the client for interacting with the Observation builders.
	Observation *ObservationClient
	// Scale is the client for interacting with the Scale builders.
	Scale *ScaleClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Correlation = NewCorrelationClient(c.config)
	c.Dataset = NewDatasetClient(c.config)
	c.Indicator = NewIndicatorClient(c.config)
	c.Observation = NewObservationClient(c.config)
	c.Scale = NewScaleClient(c.config)
	c.User = NewUserClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Correlation: NewCorrelationClient(cfg),
		Dataset:     NewDatasetClient(cfg),
		Indicator:   NewIndicatorClient(cfg),
		Observation: NewObservationClient(cfg),
		Scale:       NewScaleClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:      cfg,
		Correlation: NewCorrelationClient(cfg),
		Dataset:     NewDatasetClient(cfg),
		Indicator:   NewIndicatorClient(cfg),
		Observation: NewObservationClient(cfg),
		Scale:       NewScaleClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Correlation.
//		Query().
//		Count(ctx)
//
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
	c.Correlation.Use(hooks...)
	c.Dataset.Use(hooks...)
	c.Indicator.Use(hooks...)
	c.Observation.Use(hooks...)
	c.Scale.Use(hooks...)
	c.User.Use(hooks...)
}

// CorrelationClient is a client for the Correlation schema.
type CorrelationClient struct {
	config
}

// NewCorrelationClient returns a client for the Correlation from the given config.
func NewCorrelationClient(c config) *CorrelationClient {
	return &CorrelationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `correlation.Hooks(f(g(h())))`.
func (c *CorrelationClient) Use(hooks ...Hook) {
	c.hooks.Correlation = append(c.hooks.Correlation, hooks...)
}

// Create returns a create builder for Correlation.
func (c *CorrelationClient) Create() *CorrelationCreate {
	mutation := newCorrelationMutation(c.config, OpCreate)
	return &CorrelationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Correlation entities.
func (c *CorrelationClient) CreateBulk(builders ...*CorrelationCreate) *CorrelationCreateBulk {
	return &CorrelationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Correlation.
func (c *CorrelationClient) Update() *CorrelationUpdate {
	mutation := newCorrelationMutation(c.config, OpUpdate)
	return &CorrelationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CorrelationClient) UpdateOne(co *Correlation) *CorrelationUpdateOne {
	mutation := newCorrelationMutation(c.config, OpUpdateOne, withCorrelation(co))
	return &CorrelationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CorrelationClient) UpdateOneID(id int) *CorrelationUpdateOne {
	mutation := newCorrelationMutation(c.config, OpUpdateOne, withCorrelationID(id))
	return &CorrelationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Correlation.
func (c *CorrelationClient) Delete() *CorrelationDelete {
	mutation := newCorrelationMutation(c.config, OpDelete)
	return &CorrelationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CorrelationClient) DeleteOne(co *Correlation) *CorrelationDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CorrelationClient) DeleteOneID(id int) *CorrelationDeleteOne {
	builder := c.Delete().Where(correlation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CorrelationDeleteOne{builder}
}

// Query returns a query builder for Correlation.
func (c *CorrelationClient) Query() *CorrelationQuery {
	return &CorrelationQuery{config: c.config}
}

// Get returns a Correlation entity by its id.
func (c *CorrelationClient) Get(ctx context.Context, id int) (*Correlation, error) {
	return c.Query().Where(correlation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CorrelationClient) GetX(ctx context.Context, id int) *Correlation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryLeft queries the left edge of a Correlation.
func (c *CorrelationClient) QueryLeft(co *Correlation) *DatasetQuery {
	query := &DatasetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(correlation.Table, correlation.FieldID, id),
			sqlgraph.To(dataset.Table, dataset.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, correlation.LeftTable, correlation.LeftColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRight queries the right edge of a Correlation.
func (c *CorrelationClient) QueryRight(co *Correlation) *DatasetQuery {
	query := &DatasetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(correlation.Table, correlation.FieldID, id),
			sqlgraph.To(dataset.Table, dataset.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, correlation.RightTable, correlation.RightColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CorrelationClient) Hooks() []Hook {
	return c.hooks.Correlation
}

// DatasetClient is a client for the Dataset schema.
type DatasetClient struct {
	config
}

// NewDatasetClient returns a client for the Dataset from the given config.
func NewDatasetClient(c config) *DatasetClient {
	return &DatasetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dataset.Hooks(f(g(h())))`.
func (c *DatasetClient) Use(hooks ...Hook) {
	c.hooks.Dataset = append(c.hooks.Dataset, hooks...)
}

// Create returns a create builder for Dataset.
func (c *DatasetClient) Create() *DatasetCreate {
	mutation := newDatasetMutation(c.config, OpCreate)
	return &DatasetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Dataset entities.
func (c *DatasetClient) CreateBulk(builders ...*DatasetCreate) *DatasetCreateBulk {
	return &DatasetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Dataset.
func (c *DatasetClient) Update() *DatasetUpdate {
	mutation := newDatasetMutation(c.config, OpUpdate)
	return &DatasetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DatasetClient) UpdateOne(d *Dataset) *DatasetUpdateOne {
	mutation := newDatasetMutation(c.config, OpUpdateOne, withDataset(d))
	return &DatasetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DatasetClient) UpdateOneID(id int) *DatasetUpdateOne {
	mutation := newDatasetMutation(c.config, OpUpdateOne, withDatasetID(id))
	return &DatasetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Dataset.
func (c *DatasetClient) Delete() *DatasetDelete {
	mutation := newDatasetMutation(c.config, OpDelete)
	return &DatasetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DatasetClient) DeleteOne(d *Dataset) *DatasetDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DatasetClient) DeleteOneID(id int) *DatasetDeleteOne {
	builder := c.Delete().Where(dataset.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DatasetDeleteOne{builder}
}

// Query returns a query builder for Dataset.
func (c *DatasetClient) Query() *DatasetQuery {
	return &DatasetQuery{config: c.config}
}

// Get returns a Dataset entity by its id.
func (c *DatasetClient) Get(ctx context.Context, id int) (*Dataset, error) {
	return c.Query().Where(dataset.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DatasetClient) GetX(ctx context.Context, id int) *Dataset {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryLeft queries the left edge of a Dataset.
func (c *DatasetClient) QueryLeft(d *Dataset) *CorrelationQuery {
	query := &CorrelationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dataset.Table, dataset.FieldID, id),
			sqlgraph.To(correlation.Table, correlation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dataset.LeftTable, dataset.LeftColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRight queries the right edge of a Dataset.
func (c *DatasetClient) QueryRight(d *Dataset) *CorrelationQuery {
	query := &CorrelationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dataset.Table, dataset.FieldID, id),
			sqlgraph.To(correlation.Table, correlation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dataset.RightTable, dataset.RightColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryObservations queries the observations edge of a Dataset.
func (c *DatasetClient) QueryObservations(d *Dataset) *ObservationQuery {
	query := &ObservationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dataset.Table, dataset.FieldID, id),
			sqlgraph.To(observation.Table, observation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dataset.ObservationsTable, dataset.ObservationsColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryIndicator queries the indicator edge of a Dataset.
func (c *DatasetClient) QueryIndicator(d *Dataset) *IndicatorQuery {
	query := &IndicatorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dataset.Table, dataset.FieldID, id),
			sqlgraph.To(indicator.Table, indicator.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dataset.IndicatorTable, dataset.IndicatorColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Dataset.
func (c *DatasetClient) QueryUser(d *Dataset) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dataset.Table, dataset.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dataset.UserTable, dataset.UserColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DatasetClient) Hooks() []Hook {
	return c.hooks.Dataset
}

// IndicatorClient is a client for the Indicator schema.
type IndicatorClient struct {
	config
}

// NewIndicatorClient returns a client for the Indicator from the given config.
func NewIndicatorClient(c config) *IndicatorClient {
	return &IndicatorClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `indicator.Hooks(f(g(h())))`.
func (c *IndicatorClient) Use(hooks ...Hook) {
	c.hooks.Indicator = append(c.hooks.Indicator, hooks...)
}

// Create returns a create builder for Indicator.
func (c *IndicatorClient) Create() *IndicatorCreate {
	mutation := newIndicatorMutation(c.config, OpCreate)
	return &IndicatorCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Indicator entities.
func (c *IndicatorClient) CreateBulk(builders ...*IndicatorCreate) *IndicatorCreateBulk {
	return &IndicatorCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Indicator.
func (c *IndicatorClient) Update() *IndicatorUpdate {
	mutation := newIndicatorMutation(c.config, OpUpdate)
	return &IndicatorUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *IndicatorClient) UpdateOne(i *Indicator) *IndicatorUpdateOne {
	mutation := newIndicatorMutation(c.config, OpUpdateOne, withIndicator(i))
	return &IndicatorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *IndicatorClient) UpdateOneID(id int) *IndicatorUpdateOne {
	mutation := newIndicatorMutation(c.config, OpUpdateOne, withIndicatorID(id))
	return &IndicatorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Indicator.
func (c *IndicatorClient) Delete() *IndicatorDelete {
	mutation := newIndicatorMutation(c.config, OpDelete)
	return &IndicatorDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *IndicatorClient) DeleteOne(i *Indicator) *IndicatorDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *IndicatorClient) DeleteOneID(id int) *IndicatorDeleteOne {
	builder := c.Delete().Where(indicator.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &IndicatorDeleteOne{builder}
}

// Query returns a query builder for Indicator.
func (c *IndicatorClient) Query() *IndicatorQuery {
	return &IndicatorQuery{config: c.config}
}

// Get returns a Indicator entity by its id.
func (c *IndicatorClient) Get(ctx context.Context, id int) (*Indicator, error) {
	return c.Query().Where(indicator.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *IndicatorClient) GetX(ctx context.Context, id int) *Indicator {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDatasets queries the datasets edge of a Indicator.
func (c *IndicatorClient) QueryDatasets(i *Indicator) *DatasetQuery {
	query := &DatasetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(indicator.Table, indicator.FieldID, id),
			sqlgraph.To(dataset.Table, dataset.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, indicator.DatasetsTable, indicator.DatasetsColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuthor queries the author edge of a Indicator.
func (c *IndicatorClient) QueryAuthor(i *Indicator) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(indicator.Table, indicator.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, indicator.AuthorTable, indicator.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryScale queries the scale edge of a Indicator.
func (c *IndicatorClient) QueryScale(i *Indicator) *ScaleQuery {
	query := &ScaleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(indicator.Table, indicator.FieldID, id),
			sqlgraph.To(scale.Table, scale.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, indicator.ScaleTable, indicator.ScaleColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *IndicatorClient) Hooks() []Hook {
	return c.hooks.Indicator
}

// ObservationClient is a client for the Observation schema.
type ObservationClient struct {
	config
}

// NewObservationClient returns a client for the Observation from the given config.
func NewObservationClient(c config) *ObservationClient {
	return &ObservationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `observation.Hooks(f(g(h())))`.
func (c *ObservationClient) Use(hooks ...Hook) {
	c.hooks.Observation = append(c.hooks.Observation, hooks...)
}

// Create returns a create builder for Observation.
func (c *ObservationClient) Create() *ObservationCreate {
	mutation := newObservationMutation(c.config, OpCreate)
	return &ObservationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Observation entities.
func (c *ObservationClient) CreateBulk(builders ...*ObservationCreate) *ObservationCreateBulk {
	return &ObservationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Observation.
func (c *ObservationClient) Update() *ObservationUpdate {
	mutation := newObservationMutation(c.config, OpUpdate)
	return &ObservationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ObservationClient) UpdateOne(o *Observation) *ObservationUpdateOne {
	mutation := newObservationMutation(c.config, OpUpdateOne, withObservation(o))
	return &ObservationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ObservationClient) UpdateOneID(id int) *ObservationUpdateOne {
	mutation := newObservationMutation(c.config, OpUpdateOne, withObservationID(id))
	return &ObservationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Observation.
func (c *ObservationClient) Delete() *ObservationDelete {
	mutation := newObservationMutation(c.config, OpDelete)
	return &ObservationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ObservationClient) DeleteOne(o *Observation) *ObservationDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ObservationClient) DeleteOneID(id int) *ObservationDeleteOne {
	builder := c.Delete().Where(observation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ObservationDeleteOne{builder}
}

// Query returns a query builder for Observation.
func (c *ObservationClient) Query() *ObservationQuery {
	return &ObservationQuery{config: c.config}
}

// Get returns a Observation entity by its id.
func (c *ObservationClient) Get(ctx context.Context, id int) (*Observation, error) {
	return c.Query().Where(observation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ObservationClient) GetX(ctx context.Context, id int) *Observation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDataset queries the dataset edge of a Observation.
func (c *ObservationClient) QueryDataset(o *Observation) *DatasetQuery {
	query := &DatasetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(observation.Table, observation.FieldID, id),
			sqlgraph.To(dataset.Table, dataset.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, observation.DatasetTable, observation.DatasetColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ObservationClient) Hooks() []Hook {
	return c.hooks.Observation
}

// ScaleClient is a client for the Scale schema.
type ScaleClient struct {
	config
}

// NewScaleClient returns a client for the Scale from the given config.
func NewScaleClient(c config) *ScaleClient {
	return &ScaleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `scale.Hooks(f(g(h())))`.
func (c *ScaleClient) Use(hooks ...Hook) {
	c.hooks.Scale = append(c.hooks.Scale, hooks...)
}

// Create returns a create builder for Scale.
func (c *ScaleClient) Create() *ScaleCreate {
	mutation := newScaleMutation(c.config, OpCreate)
	return &ScaleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Scale entities.
func (c *ScaleClient) CreateBulk(builders ...*ScaleCreate) *ScaleCreateBulk {
	return &ScaleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Scale.
func (c *ScaleClient) Update() *ScaleUpdate {
	mutation := newScaleMutation(c.config, OpUpdate)
	return &ScaleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ScaleClient) UpdateOne(s *Scale) *ScaleUpdateOne {
	mutation := newScaleMutation(c.config, OpUpdateOne, withScale(s))
	return &ScaleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ScaleClient) UpdateOneID(id int) *ScaleUpdateOne {
	mutation := newScaleMutation(c.config, OpUpdateOne, withScaleID(id))
	return &ScaleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Scale.
func (c *ScaleClient) Delete() *ScaleDelete {
	mutation := newScaleMutation(c.config, OpDelete)
	return &ScaleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ScaleClient) DeleteOne(s *Scale) *ScaleDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ScaleClient) DeleteOneID(id int) *ScaleDeleteOne {
	builder := c.Delete().Where(scale.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ScaleDeleteOne{builder}
}

// Query returns a query builder for Scale.
func (c *ScaleClient) Query() *ScaleQuery {
	return &ScaleQuery{config: c.config}
}

// Get returns a Scale entity by its id.
func (c *ScaleClient) Get(ctx context.Context, id int) (*Scale, error) {
	return c.Query().Where(scale.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ScaleClient) GetX(ctx context.Context, id int) *Scale {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryIndicators queries the indicators edge of a Scale.
func (c *ScaleClient) QueryIndicators(s *Scale) *IndicatorQuery {
	query := &IndicatorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(scale.Table, scale.FieldID, id),
			sqlgraph.To(indicator.Table, indicator.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, scale.IndicatorsTable, scale.IndicatorsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ScaleClient) Hooks() []Hook {
	return c.hooks.Scale
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

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
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
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryIndicators queries the indicators edge of a User.
func (c *UserClient) QueryIndicators(u *User) *IndicatorQuery {
	query := &IndicatorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(indicator.Table, indicator.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.IndicatorsTable, user.IndicatorsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDatasets queries the datasets edge of a User.
func (c *UserClient) QueryDatasets(u *User) *DatasetQuery {
	query := &DatasetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(dataset.Table, dataset.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.DatasetsTable, user.DatasetsColumn),
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
