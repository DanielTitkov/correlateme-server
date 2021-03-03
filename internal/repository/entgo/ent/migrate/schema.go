// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CorrelationsColumns holds the columns for the "correlations" table.
	CorrelationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "coef", Type: field.TypeFloat64},
		{Name: "p", Type: field.TypeFloat64},
		{Name: "r2", Type: field.TypeFloat64},
		{Name: "type", Type: field.TypeString},
		{Name: "dataset_left", Type: field.TypeInt, Nullable: true},
		{Name: "dataset_right", Type: field.TypeInt, Nullable: true},
	}
	// CorrelationsTable holds the schema information for the "correlations" table.
	CorrelationsTable = &schema.Table{
		Name:       "correlations",
		Columns:    CorrelationsColumns,
		PrimaryKey: []*schema.Column{CorrelationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "correlations_datasets_left",
				Columns: []*schema.Column{CorrelationsColumns[7]},

				RefColumns: []*schema.Column{DatasetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "correlations_datasets_right",
				Columns: []*schema.Column{CorrelationsColumns[8]},

				RefColumns: []*schema.Column{DatasetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DatasetsColumns holds the columns for the "datasets" table.
	DatasetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "shared", Type: field.TypeBool},
		{Name: "source", Type: field.TypeString, Nullable: true},
		{Name: "indicator_datasets", Type: field.TypeInt, Nullable: true},
		{Name: "user_datasets", Type: field.TypeInt, Nullable: true},
	}
	// DatasetsTable holds the schema information for the "datasets" table.
	DatasetsTable = &schema.Table{
		Name:       "datasets",
		Columns:    DatasetsColumns,
		PrimaryKey: []*schema.Column{DatasetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "datasets_indicators_datasets",
				Columns: []*schema.Column{DatasetsColumns[5]},

				RefColumns: []*schema.Column{IndicatorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "datasets_users_datasets",
				Columns: []*schema.Column{DatasetsColumns[6]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "dataset_user_datasets_indicator_datasets",
				Unique:  true,
				Columns: []*schema.Column{DatasetsColumns[6], DatasetsColumns[5]},
			},
		},
	}
	// IndicatorsColumns holds the columns for the "indicators" table.
	IndicatorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "built_in", Type: field.TypeBool},
		{Name: "external", Type: field.TypeBool},
		{Name: "scale_indicators", Type: field.TypeInt, Nullable: true},
		{Name: "user_indicators", Type: field.TypeInt, Nullable: true},
	}
	// IndicatorsTable holds the schema information for the "indicators" table.
	IndicatorsTable = &schema.Table{
		Name:       "indicators",
		Columns:    IndicatorsColumns,
		PrimaryKey: []*schema.Column{IndicatorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "indicators_scales_indicators",
				Columns: []*schema.Column{IndicatorsColumns[9]},

				RefColumns: []*schema.Column{ScalesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "indicators_users_indicators",
				Columns: []*schema.Column{IndicatorsColumns[10]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "indicator_title_user_indicators",
				Unique:  true,
				Columns: []*schema.Column{IndicatorsColumns[4], IndicatorsColumns[10]},
			},
		},
	}
	// IndicatorValueAliasColumns holds the columns for the "indicator_value_alias" table.
	IndicatorValueAliasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value_mapping", Type: field.TypeJSON, Nullable: true},
		{Name: "indicator_indicator_value_alias", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// IndicatorValueAliasTable holds the schema information for the "indicator_value_alias" table.
	IndicatorValueAliasTable = &schema.Table{
		Name:       "indicator_value_alias",
		Columns:    IndicatorValueAliasColumns,
		PrimaryKey: []*schema.Column{IndicatorValueAliasColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "indicator_value_alias_indicators_indicator_value_alias",
				Columns: []*schema.Column{IndicatorValueAliasColumns[2]},

				RefColumns: []*schema.Column{IndicatorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ObservationsColumns holds the columns for the "observations" table.
	ObservationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "value", Type: field.TypeFloat64},
		{Name: "date", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATE", "postgres": "date"}},
		{Name: "dataset_observations", Type: field.TypeInt, Nullable: true},
	}
	// ObservationsTable holds the schema information for the "observations" table.
	ObservationsTable = &schema.Table{
		Name:       "observations",
		Columns:    ObservationsColumns,
		PrimaryKey: []*schema.Column{ObservationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "observations_datasets_observations",
				Columns: []*schema.Column{ObservationsColumns[5]},

				RefColumns: []*schema.Column{DatasetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "observation_date_dataset_observations",
				Unique:  true,
				Columns: []*schema.Column{ObservationsColumns[4], ObservationsColumns[5]},
			},
		},
	}
	// ScalesColumns holds the columns for the "scales" table.
	ScalesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Unique: true},
	}
	// ScalesTable holds the schema information for the "scales" table.
	ScalesTable = &schema.Table{
		Name:        "scales",
		Columns:     ScalesColumns,
		PrimaryKey:  []*schema.Column{ScalesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "service", Type: field.TypeBool},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CorrelationsTable,
		DatasetsTable,
		IndicatorsTable,
		IndicatorValueAliasTable,
		ObservationsTable,
		ScalesTable,
		UsersTable,
	}
)

func init() {
	CorrelationsTable.ForeignKeys[0].RefTable = DatasetsTable
	CorrelationsTable.ForeignKeys[1].RefTable = DatasetsTable
	DatasetsTable.ForeignKeys[0].RefTable = IndicatorsTable
	DatasetsTable.ForeignKeys[1].RefTable = UsersTable
	IndicatorsTable.ForeignKeys[0].RefTable = ScalesTable
	IndicatorsTable.ForeignKeys[1].RefTable = UsersTable
	IndicatorValueAliasTable.ForeignKeys[0].RefTable = IndicatorsTable
	ObservationsTable.ForeignKeys[0].RefTable = DatasetsTable
}
