package model

import "time"

type (
	// Indicator is a common model to use in various methods
	Indicator struct {
		ID          int       `json:"id"`
		Code        string    `json:"code"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Active      bool      `json:"active"`
		BuiltIn     bool      `json:"builtIn"`
		External    bool      `json:"external"`
		ScaleID     int       `json:"scaleID"`
		AuthorID    int       `json:"authorID,omitempty"`
		CreateTime  time.Time `json:"createTime"`
		UpdateTime  time.Time `json:"updateTime"`
		Dataset     *Dataset  `json:"dataset,omitempty"`
	}
)

type (
	CreateIndicatorRequest struct {
		Title       string `json:"title" validate:"required"`
		ScaleType   string `json:"scaleType" validate:"required"`
		Description string `json:"description"`
	}
	GetIndicatorsRequest struct {
		WithDataset       bool                       `json:"withDataset"`
		ObservationsLimit int                        `json:"observationsLimit"`
		Granularity       string                     `json:"granularity"`
		Filter            GetIndicatorsRequestFilter `json:"filter"`
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
)
