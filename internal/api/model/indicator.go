package model

import "time"

type (
	CreateIndicatorRequest struct {
		Title       string `json:"title" validate:"required"`
		ScaleType   string `json:"scaleType" validate:"required"`
		Description string `json:"description"`
	}
	GetIndicatorsRequest struct {
		Filter GetIndicatorsRequestFilter `json:"filter"`
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
	}
)
