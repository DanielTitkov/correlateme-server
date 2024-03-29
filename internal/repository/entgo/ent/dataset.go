// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/dataset"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/datasetparams"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/indicator"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/user"
)

// Dataset is the model entity for the Dataset schema.
type Dataset struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Shared holds the value of the "shared" field.
	Shared bool `json:"shared,omitempty"`
	// Source holds the value of the "source" field.
	Source *string `json:"source,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DatasetQuery when eager-loading is set.
	Edges              DatasetEdges `json:"edges"`
	indicator_datasets *int
	user_datasets      *int
}

// DatasetEdges holds the relations/edges for other nodes in the graph.
type DatasetEdges struct {
	// Left holds the value of the left edge.
	Left []*Correlation `json:"left,omitempty"`
	// Right holds the value of the right edge.
	Right []*Correlation `json:"right,omitempty"`
	// Observations holds the value of the observations edge.
	Observations []*Observation `json:"observations,omitempty"`
	// DatasetParams holds the value of the dataset_params edge.
	DatasetParams *DatasetParams `json:"dataset_params,omitempty"`
	// Indicator holds the value of the indicator edge.
	Indicator *Indicator `json:"indicator,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// LeftOrErr returns the Left value or an error if the edge
// was not loaded in eager-loading.
func (e DatasetEdges) LeftOrErr() ([]*Correlation, error) {
	if e.loadedTypes[0] {
		return e.Left, nil
	}
	return nil, &NotLoadedError{edge: "left"}
}

// RightOrErr returns the Right value or an error if the edge
// was not loaded in eager-loading.
func (e DatasetEdges) RightOrErr() ([]*Correlation, error) {
	if e.loadedTypes[1] {
		return e.Right, nil
	}
	return nil, &NotLoadedError{edge: "right"}
}

// ObservationsOrErr returns the Observations value or an error if the edge
// was not loaded in eager-loading.
func (e DatasetEdges) ObservationsOrErr() ([]*Observation, error) {
	if e.loadedTypes[2] {
		return e.Observations, nil
	}
	return nil, &NotLoadedError{edge: "observations"}
}

// DatasetParamsOrErr returns the DatasetParams value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DatasetEdges) DatasetParamsOrErr() (*DatasetParams, error) {
	if e.loadedTypes[3] {
		if e.DatasetParams == nil {
			// The edge dataset_params was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: datasetparams.Label}
		}
		return e.DatasetParams, nil
	}
	return nil, &NotLoadedError{edge: "dataset_params"}
}

// IndicatorOrErr returns the Indicator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DatasetEdges) IndicatorOrErr() (*Indicator, error) {
	if e.loadedTypes[4] {
		if e.Indicator == nil {
			// The edge indicator was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: indicator.Label}
		}
		return e.Indicator, nil
	}
	return nil, &NotLoadedError{edge: "indicator"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DatasetEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[5] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Dataset) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case dataset.FieldShared:
			values[i] = &sql.NullBool{}
		case dataset.FieldID:
			values[i] = &sql.NullInt64{}
		case dataset.FieldSource:
			values[i] = &sql.NullString{}
		case dataset.FieldCreateTime, dataset.FieldUpdateTime:
			values[i] = &sql.NullTime{}
		case dataset.ForeignKeys[0]: // indicator_datasets
			values[i] = &sql.NullInt64{}
		case dataset.ForeignKeys[1]: // user_datasets
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Dataset", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Dataset fields.
func (d *Dataset) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dataset.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case dataset.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				d.CreateTime = value.Time
			}
		case dataset.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				d.UpdateTime = value.Time
			}
		case dataset.FieldShared:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field shared", values[i])
			} else if value.Valid {
				d.Shared = value.Bool
			}
		case dataset.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				d.Source = new(string)
				*d.Source = value.String
			}
		case dataset.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field indicator_datasets", value)
			} else if value.Valid {
				d.indicator_datasets = new(int)
				*d.indicator_datasets = int(value.Int64)
			}
		case dataset.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_datasets", value)
			} else if value.Valid {
				d.user_datasets = new(int)
				*d.user_datasets = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryLeft queries the "left" edge of the Dataset entity.
func (d *Dataset) QueryLeft() *CorrelationQuery {
	return (&DatasetClient{config: d.config}).QueryLeft(d)
}

// QueryRight queries the "right" edge of the Dataset entity.
func (d *Dataset) QueryRight() *CorrelationQuery {
	return (&DatasetClient{config: d.config}).QueryRight(d)
}

// QueryObservations queries the "observations" edge of the Dataset entity.
func (d *Dataset) QueryObservations() *ObservationQuery {
	return (&DatasetClient{config: d.config}).QueryObservations(d)
}

// QueryDatasetParams queries the "dataset_params" edge of the Dataset entity.
func (d *Dataset) QueryDatasetParams() *DatasetParamsQuery {
	return (&DatasetClient{config: d.config}).QueryDatasetParams(d)
}

// QueryIndicator queries the "indicator" edge of the Dataset entity.
func (d *Dataset) QueryIndicator() *IndicatorQuery {
	return (&DatasetClient{config: d.config}).QueryIndicator(d)
}

// QueryUser queries the "user" edge of the Dataset entity.
func (d *Dataset) QueryUser() *UserQuery {
	return (&DatasetClient{config: d.config}).QueryUser(d)
}

// Update returns a builder for updating this Dataset.
// Note that you need to call Dataset.Unwrap() before calling this method if this Dataset
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dataset) Update() *DatasetUpdateOne {
	return (&DatasetClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Dataset entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Dataset) Unwrap() *Dataset {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dataset is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dataset) String() string {
	var builder strings.Builder
	builder.WriteString("Dataset(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(d.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(d.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", shared=")
	builder.WriteString(fmt.Sprintf("%v", d.Shared))
	if v := d.Source; v != nil {
		builder.WriteString(", source=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Datasets is a parsable slice of Dataset.
type Datasets []*Dataset

func (d Datasets) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
