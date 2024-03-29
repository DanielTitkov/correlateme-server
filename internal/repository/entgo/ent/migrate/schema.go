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
		{Name: "granularity", Type: field.TypeEnum, Enums: []string{"day", "week", "month"}, Default: "day"},
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
				Columns: []*schema.Column{CorrelationsColumns[8]},

				RefColumns: []*schema.Column{DatasetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "correlations_datasets_right",
				Columns: []*schema.Column{CorrelationsColumns[9]},

				RefColumns: []*schema.Column{DatasetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "correlation_granularity_dataset_left_dataset_right",
				Unique:  true,
				Columns: []*schema.Column{CorrelationsColumns[7], CorrelationsColumns[8], CorrelationsColumns[9]},
			},
			{
				Name:    "correlation_granularity_dataset_right_dataset_left",
				Unique:  true,
				Columns: []*schema.Column{CorrelationsColumns[7], CorrelationsColumns[9], CorrelationsColumns[8]},
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
	// DatasetParamsColumns holds the columns for the "dataset_params" table.
	DatasetParamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "style", Type: field.TypeJSON},
		{Name: "aggregation", Type: field.TypeJSON},
		{Name: "dataset_dataset_params", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// DatasetParamsTable holds the schema information for the "dataset_params" table.
	DatasetParamsTable = &schema.Table{
		Name:       "dataset_params",
		Columns:    DatasetParamsColumns,
		PrimaryKey: []*schema.Column{DatasetParamsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "dataset_params_datasets_dataset_params",
				Columns: []*schema.Column{DatasetParamsColumns[3]},

				RefColumns: []*schema.Column{DatasetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DictionariesColumns holds the columns for the "dictionaries" table.
	DictionariesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// DictionariesTable holds the schema information for the "dictionaries" table.
	DictionariesTable = &schema.Table{
		Name:        "dictionaries",
		Columns:     DictionariesColumns,
		PrimaryKey:  []*schema.Column{DictionariesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DictionaryEntriesColumns holds the columns for the "dictionary_entries" table.
	DictionaryEntriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "code", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "dictionary_entries", Type: field.TypeInt, Nullable: true},
	}
	// DictionaryEntriesTable holds the schema information for the "dictionary_entries" table.
	DictionaryEntriesTable = &schema.Table{
		Name:       "dictionary_entries",
		Columns:    DictionaryEntriesColumns,
		PrimaryKey: []*schema.Column{DictionaryEntriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "dictionary_entries_dictionaries_entries",
				Columns: []*schema.Column{DictionaryEntriesColumns[3]},

				RefColumns: []*schema.Column{DictionariesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "dictionaryentry_code_dictionary_entries",
				Unique:  true,
				Columns: []*schema.Column{DictionaryEntriesColumns[1], DictionaryEntriesColumns[3]},
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
	// IndicatorParamsColumns holds the columns for the "indicator_params" table.
	IndicatorParamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value_mapping", Type: field.TypeJSON, Nullable: true},
		{Name: "value_params", Type: field.TypeJSON},
		{Name: "indicator_indicator_params", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// IndicatorParamsTable holds the schema information for the "indicator_params" table.
	IndicatorParamsTable = &schema.Table{
		Name:       "indicator_params",
		Columns:    IndicatorParamsColumns,
		PrimaryKey: []*schema.Column{IndicatorParamsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "indicator_params_indicators_indicator_params",
				Columns: []*schema.Column{IndicatorParamsColumns[3]},

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
		{Name: "granularity", Type: field.TypeEnum, Enums: []string{"day", "week", "month"}, Default: "day"},
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
				Columns: []*schema.Column{ObservationsColumns[6]},

				RefColumns: []*schema.Column{DatasetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "observation_date_granularity_dataset_observations",
				Unique:  true,
				Columns: []*schema.Column{ObservationsColumns[4], ObservationsColumns[5], ObservationsColumns[6]},
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
	// UserSettingsColumns holds the columns for the "user_settings" table.
	UserSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_settings", Type: field.TypeInt, Nullable: true},
	}
	// UserSettingsTable holds the schema information for the "user_settings" table.
	UserSettingsTable = &schema.Table{
		Name:       "user_settings",
		Columns:    UserSettingsColumns,
		PrimaryKey: []*schema.Column{UserSettingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_settings_users_settings",
				Columns: []*schema.Column{UserSettingsColumns[1]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CorrelationsTable,
		DatasetsTable,
		DatasetParamsTable,
		DictionariesTable,
		DictionaryEntriesTable,
		IndicatorsTable,
		IndicatorParamsTable,
		ObservationsTable,
		ScalesTable,
		UsersTable,
		UserSettingsTable,
	}
)

func init() {
	CorrelationsTable.ForeignKeys[0].RefTable = DatasetsTable
	CorrelationsTable.ForeignKeys[1].RefTable = DatasetsTable
	DatasetsTable.ForeignKeys[0].RefTable = IndicatorsTable
	DatasetsTable.ForeignKeys[1].RefTable = UsersTable
	DatasetParamsTable.ForeignKeys[0].RefTable = DatasetsTable
	DictionaryEntriesTable.ForeignKeys[0].RefTable = DictionariesTable
	IndicatorsTable.ForeignKeys[0].RefTable = ScalesTable
	IndicatorsTable.ForeignKeys[1].RefTable = UsersTable
	IndicatorParamsTable.ForeignKeys[0].RefTable = IndicatorsTable
	ObservationsTable.ForeignKeys[0].RefTable = DatasetsTable
	UserSettingsTable.ForeignKeys[0].RefTable = UsersTable
}
