package model

import "time"

type (
	CreateOrUpdateObservationRequest struct {
		IndicatorID int
		Value       float64
		Date        *time.Time
	}
)
