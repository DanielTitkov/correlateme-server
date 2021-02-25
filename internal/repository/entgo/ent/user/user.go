// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldService holds the string denoting the service field in the database.
	FieldService = "service"

	// EdgeIndicators holds the string denoting the indicators edge name in mutations.
	EdgeIndicators = "indicators"
	// EdgeDatasets holds the string denoting the datasets edge name in mutations.
	EdgeDatasets = "datasets"

	// Table holds the table name of the user in the database.
	Table = "users"
	// IndicatorsTable is the table the holds the indicators relation/edge.
	IndicatorsTable = "indicators"
	// IndicatorsInverseTable is the table name for the Indicator entity.
	// It exists in this package in order to avoid circular dependency with the "indicator" package.
	IndicatorsInverseTable = "indicators"
	// IndicatorsColumn is the table column denoting the indicators relation/edge.
	IndicatorsColumn = "user_indicators"
	// DatasetsTable is the table the holds the datasets relation/edge.
	DatasetsTable = "datasets"
	// DatasetsInverseTable is the table name for the Dataset entity.
	// It exists in this package in order to avoid circular dependency with the "dataset" package.
	DatasetsInverseTable = "datasets"
	// DatasetsColumn is the table column denoting the datasets relation/edge.
	DatasetsColumn = "user_datasets"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldUsername,
	FieldEmail,
	FieldPasswordHash,
	FieldService,
}

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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultService holds the default value on creation for the "service" field.
	DefaultService bool
)
