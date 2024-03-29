// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/scale"
)

// Scale is the model entity for the Scale schema.
type Scale struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ScaleQuery when eager-loading is set.
	Edges ScaleEdges `json:"edges"`
}

// ScaleEdges holds the relations/edges for other nodes in the graph.
type ScaleEdges struct {
	// Indicators holds the value of the indicators edge.
	Indicators []*Indicator `json:"indicators,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// IndicatorsOrErr returns the Indicators value or an error if the edge
// was not loaded in eager-loading.
func (e ScaleEdges) IndicatorsOrErr() ([]*Indicator, error) {
	if e.loadedTypes[0] {
		return e.Indicators, nil
	}
	return nil, &NotLoadedError{edge: "indicators"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Scale) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case scale.FieldID:
			values[i] = &sql.NullInt64{}
		case scale.FieldType, scale.FieldTitle, scale.FieldDescription:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Scale", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Scale fields.
func (s *Scale) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case scale.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case scale.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				s.Type = value.String
			}
		case scale.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				s.Title = value.String
			}
		case scale.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		}
	}
	return nil
}

// QueryIndicators queries the "indicators" edge of the Scale entity.
func (s *Scale) QueryIndicators() *IndicatorQuery {
	return (&ScaleClient{config: s.config}).QueryIndicators(s)
}

// Update returns a builder for updating this Scale.
// Note that you need to call Scale.Unwrap() before calling this method if this Scale
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Scale) Update() *ScaleUpdateOne {
	return (&ScaleClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Scale entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Scale) Unwrap() *Scale {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Scale is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Scale) String() string {
	var builder strings.Builder
	builder.WriteString("Scale(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", type=")
	builder.WriteString(s.Type)
	builder.WriteString(", title=")
	builder.WriteString(s.Title)
	builder.WriteString(", description=")
	builder.WriteString(s.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Scales is a parsable slice of Scale.
type Scales []*Scale

func (s Scales) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
