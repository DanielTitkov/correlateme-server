// Code generated by entc, DO NOT EDIT.

package indicator

import (
	"time"
)

const (
	// Label holds the string label denoting the indicator type in the database.
	Label = "indicator"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldBuiltIn holds the string denoting the built_in field in the database.
	FieldBuiltIn = "built_in"
	// FieldExternal holds the string denoting the external field in the database.
	FieldExternal = "external"

	// EdgeDatasets holds the string denoting the datasets edge name in mutations.
	EdgeDatasets = "datasets"
	// EdgeIndicatorParams holds the string denoting the indicator_params edge name in mutations.
	EdgeIndicatorParams = "indicator_params"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeScale holds the string denoting the scale edge name in mutations.
	EdgeScale = "scale"

	// Table holds the table name of the indicator in the database.
	Table = "indicators"
	// DatasetsTable is the table the holds the datasets relation/edge.
	DatasetsTable = "datasets"
	// DatasetsInverseTable is the table name for the Dataset entity.
	// It exists in this package in order to avoid circular dependency with the "dataset" package.
	DatasetsInverseTable = "datasets"
	// DatasetsColumn is the table column denoting the datasets relation/edge.
	DatasetsColumn = "indicator_datasets"
	// IndicatorParamsTable is the table the holds the indicator_params relation/edge.
	IndicatorParamsTable = "indicator_params"
	// IndicatorParamsInverseTable is the table name for the IndicatorParams entity.
	// It exists in this package in order to avoid circular dependency with the "indicatorparams" package.
	IndicatorParamsInverseTable = "indicator_params"
	// IndicatorParamsColumn is the table column denoting the indicator_params relation/edge.
	IndicatorParamsColumn = "indicator_indicator_params"
	// AuthorTable is the table the holds the author relation/edge.
	AuthorTable = "indicators"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "user_indicators"
	// ScaleTable is the table the holds the scale relation/edge.
	ScaleTable = "indicators"
	// ScaleInverseTable is the table name for the Scale entity.
	// It exists in this package in order to avoid circular dependency with the "scale" package.
	ScaleInverseTable = "scales"
	// ScaleColumn is the table column denoting the scale relation/edge.
	ScaleColumn = "scale_indicators"
)

// Columns holds all SQL columns for indicator fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldCode,
	FieldTitle,
	FieldDescription,
	FieldActive,
	FieldBuiltIn,
	FieldExternal,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Indicator type.
var ForeignKeys = []string{
	"scale_indicators",
	"user_indicators",
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// DefaultBuiltIn holds the default value on creation for the "built_in" field.
	DefaultBuiltIn bool
	// DefaultExternal holds the default value on creation for the "external" field.
	DefaultExternal bool
)
