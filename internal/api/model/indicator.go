package model

import "time"

type (
	// Indicator is a common model to use in various methods
	Indicator struct {
		ID          int                           `json:"id"`
		Code        string                        `json:"code"`
		Title       string                        `json:"title"`
		Description string                        `json:"description"`
		Active      bool                          `json:"active"`
		BuiltIn     bool                          `json:"builtIn"`
		External    bool                          `json:"external"`
		ScaleID     int                           `json:"scaleID"`
		AuthorID    int                           `json:"authorID,omitempty"`
		CreateTime  time.Time                     `json:"createTime"`
		UpdateTime  time.Time                     `json:"updateTime"`
		Dataset     *GetIndicatorsResponseDataset `json:"dataset,omitempty"`
	}
)

type (
	CreateIndicatorRequest struct {
		Title       string `json:"title" validate:"required"`
		ScaleType   string `json:"scaleType" validate:"required"`
		Description string `json:"description"`
	}
	GetIndicatorsRequest struct {
		WithDataset      bool                       `json:"withDataset"`
		WithObservations bool                       `json:"withObservations"`
		Filter           GetIndicatorsRequestFilter `json:"filter"`
	}
	GetIndicatorsRequestFilter struct {
		ID        []int    `json:"id"`
		Code      []string `json:"code"`
		Title     []string `json:"title"`
		Active    *bool    `json:"active"`
		BuiltIn   *bool    `json:"builtIn"`
		ScaleType *string  `json:"scaleType"`
	}
	GetIndicatorsResponse     []GetIndicatorsResponseItem
	GetIndicatorsResponseItem struct {
		Indicator
	}
	GetIndicatorsResponseDataset struct {
		ID           int                                `json:"id,omitempty"`
		Source       string                             `json:"source,omitempty"`
		Shared       bool                               `json:"shared,omitempty"`
		Observations []GetIndicatorsResponseObservation `json:"observations,omitempty"`
	}
	GetIndicatorsResponseObservation struct {
		ID    int        `json:"id"`
		Value float64    `json:"value"`
		Date  *time.Time `json:"date"`
	}
)
