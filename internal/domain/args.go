package domain

import "time"

type (
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
	CreateObservationArgs struct {
		UserID      int
		IndicatorID int
		Value       float64
		Date        *time.Time
	}
)
