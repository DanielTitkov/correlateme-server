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
		WithObservations bool
		ObservationLimit int64
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

	// Dataset

	GetDatasetsArgs struct {
		// IndicatorID      *int
		UserID           int
		WithIndicator    bool
		WithUser         bool
		WithObservations bool
		ObservationLimit int64
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
		Method     string // pearson, spearman or auto
	}

	GetCorrelationMatrixArgs struct {
		UserID     int
		WithShared bool
	}

	GetCorrelationArgs struct {
		ID               int
		UserID           int
		WithDatasets     bool
		WithObservations bool
	}
)
