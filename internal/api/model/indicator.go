package model

type (
	CreateIndicatorRequest struct {
		Title       string `json:"title" validate:"required"`
		ScaleType   string `json:"scaleType" validate:"required"`
		Description string `json:"description"`
	}
)
