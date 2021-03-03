// Code generated by entc, DO NOT EDIT.

package indicatorvaluealias

const (
	// Label holds the string label denoting the indicatorvaluealias type in the database.
	Label = "indicator_value_alias"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValueMapping holds the string denoting the value_mapping field in the database.
	FieldValueMapping = "value_mapping"

	// EdgeIndicator holds the string denoting the indicator edge name in mutations.
	EdgeIndicator = "indicator"

	// Table holds the table name of the indicatorvaluealias in the database.
	Table = "indicator_value_alias"
	// IndicatorTable is the table the holds the indicator relation/edge.
	IndicatorTable = "indicator_value_alias"
	// IndicatorInverseTable is the table name for the Indicator entity.
	// It exists in this package in order to avoid circular dependency with the "indicator" package.
	IndicatorInverseTable = "indicators"
	// IndicatorColumn is the table column denoting the indicator relation/edge.
	IndicatorColumn = "indicator_indicator_value_alias"
)

// Columns holds all SQL columns for indicatorvaluealias fields.
var Columns = []string{
	FieldID,
	FieldValueMapping,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the IndicatorValueAlias type.
var ForeignKeys = []string{
	"indicator_indicator_value_alias",
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