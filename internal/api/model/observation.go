package model

import "time"

type (
	// Observation is a common model to use in various methods
	Observation struct {
		ID    int        `json:"id"`
		Value float64    `json:"value"`
		Date  *time.Time `json:"date"`
	}
)

type (
	// UpdateAggregationsRequest is just for testing
	UpdateAggregationsRequest struct {
		DatasetID int `json:"datasetID"`
	}
	CreateOrUpdateObservationRequest struct {
		IndicatorID int
		Value       float64
		Date        *time.Time
	}
)
