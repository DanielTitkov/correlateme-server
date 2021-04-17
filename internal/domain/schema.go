package domain

import "time"

type (
	// User holds user data
	User struct {
		ID           int // id is passed to domain model for simplicity
		Username     string
		Password     string
		PasswordHash string
		Email        string // TODO: add ent validation
		// Service      bool // if user is a service
	}
	// UserSettings holds user app/site preferences
	UserSettings struct {
	}
	// Indicator is a set of user-created data
	Indicator struct {
		ID           int
		Code         string // unique code for external systems
		Title        string
		Description  string
		Active       bool // FIXME move to dataset?
		BuiltIn      bool // if Indicator is created by the service
		External     bool // if Indicator is populated by the user or external system
		Scale        *Scale
		Author       *User
		UserDataset  *Dataset           // dataset for a specific user
		ValueMapping map[float64]string // aliases for nomial and ordinal scales
		CreateTime   time.Time
		UpdateTime   time.Time
	}
	// Scale is a type of a scale for an Indicator
	Scale struct {
		ID          int
		Type        string `json:"type"` // numeric, ordinal or nomial
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	// Dataset is an intance of an Indicator populated by user data.
	// Each User can have one Dataset for each Indicator
	Dataset struct {
		ID           int
		User         *User
		Indicator    *Indicator
		CreateTime   time.Time
		UpdateTime   time.Time
		Observations []*Observation
		Params       *DatasetParams
		Source       string // user input or external system
		Shared       bool   // if dataset can be shared between all users
	}
	// DatasetParams holds various dataset options
	DatasetParams struct {
		Style       DatasetStyle
		Aggregation DatasetAggregation
	}
	// DatasetStyle holds dataset apperance params for site/app
	DatasetStyle struct {
		Color string `json:"color,omitempty"`
	}
	// DatasetAggregation holds params for aggregation
	DatasetAggregation struct {
		IncludeZeroValues bool `json:"includeZeroValues"` // use zero values in aggregations
	}
	// Observation is a data point for an indicator
	Observation struct {
		ID          int
		Value       float64
		Dataset     *Dataset
		Date        *time.Time
		Granularity string
		CreateTime  time.Time
		UpdateTime  time.Time
	}
	// Correlation is a corr value of a pair of Datasets
	Correlation struct {
		ID          int
		Left        *Dataset
		Right       *Dataset
		Coef        float64 // correlation coef
		P           float64 // p-value
		R2          float64 // determination coef
		Type        string  // Pearson or whatever
		Granularity string
		CreateTime  time.Time
		UpdateTime  time.Time
	}
	// CorrelationMatrix holds correlations of user datasets to be shown in the app.
	CorrelationMatrix struct {
		Header []CorrelationMatrixHeaderItem
		Body   [][]CorrelationMatrixBodyItem
	}
	// CorrelationMatrixHeaderItem holds a matrix subject description
	CorrelationMatrixHeaderItem struct {
		IndicatorID    int
		DatasetID      int
		IndicatorTitle string
		DatasetShared  bool
	}
	// CorrelationMatrixBodyItem holds a correlation on the matrix
	CorrelationMatrixBodyItem struct {
		CorrelationID int
		Coef          float64 // correlation coef
		P             float64 // p-value
		R2            float64 // determination coef
		Type          string  // Pearson or whatever
		UpdateTime    time.Time
	}
	// ServiceStats struct holds service stats
	ServiceStats struct {
		Users        int `json:"users,omitempty"`
		Scales       int `json:"scales,omitempty"`
		Indicators   int `json:"indicators,omitempty"`
		Datasets     int `json:"datasets,omitempty"`
		Observations int `json:"observations,omitempty"`
		Correlations int `json:"correlations,omitempty"`
	}
)
