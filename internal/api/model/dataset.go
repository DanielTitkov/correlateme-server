package model

import "github.com/DanielTitkov/correlateme-server/internal/domain"

type (
	// Dataset is a common model to use in various methods
	Dataset struct {
		ID           int                       `json:"id,omitempty"`
		Source       string                    `json:"source,omitempty"`
		Shared       bool                      `json:"shared,omitempty"`
		Observations []Observation             `json:"observations,omitempty"`
		Style        domain.DatasetStyle       `json:"style,omitempty"`
		Aggregation  domain.DatasetAggregation `json:"aggregation,omitempty"`
	}
)
