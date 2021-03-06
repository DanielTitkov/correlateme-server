package domain

import "time"

type (
	// Indicator

	CreateIndicatorArgs struct {
		Username    string // TODO: change to ID
		Title       string
		Description string
		ScaleType   string
	}
	GetIndicatorsArgs struct {
		UserID           int
		WithDataset      bool
		ObservationLimit int
		Granularity      string
		Filter           GetIndicatorsArgsFilter
	}
	GetIndicatorsArgsFilter struct {
		ID        []int
		Code      []string
		Title     []string
		Active    *bool
		BuiltIn   *bool
		AuthorID  *int
		ScaleType *string
		External  *bool // not accesible via API, only for intenal use
	}

	// Observation

	CreateOrUpdateObservationArgs struct {
		UserID      int
		IndicatorID int
		Value       float64
		Date        *time.Time
	}
	UpdateAggregationsArgs struct {
		UserID    int
		DatasetID int
		// Method    string // mean, median, sum
	}

	// Dataset

	GetDatasetsArgs struct {
		// IndicatorID      *int
		UserID           int
		WithIndicator    bool
		WithUser         bool
		ObservationLimit int
		Granularity      string
		Filter           GetDatasetsArgsFilter
	}
	GetDatasetsArgsFilter struct {
		ID         []int
		WithShared bool
	}

	// Correlation

	UpdateCorrelationsArgs struct {
		UserID     int
		WithShared bool
		// Granularity string
		Method string // pearson, spearman or auto
	}

	GetCorrelationMatrixArgs struct {
		UserID      int
		WithShared  bool
		Granularity string
	}

	GetCorrelationArgs struct {
		ID               int
		UserID           int
		WithDatasets     bool
		WithObservations bool
	}
)
