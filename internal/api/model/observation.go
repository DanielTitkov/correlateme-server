package model

import "time"

type (
	CreateObservationRequest struct {
		IndicatorID int
		Value       float64
		Timestamp   time.Time
	}
)
