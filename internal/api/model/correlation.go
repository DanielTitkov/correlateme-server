package model

import "time"

type (
	FindUserCorrelationsRequest struct {
		UserID     int  `json:"userID"`
		WithShared bool `json:"withShared"`
	}
	GetCorrelationMatrixRequest struct {
		WithShared bool `json:withShared`
	}
	GetCorrelationMatrixResponse struct {
		Header []GetCorrelationMatrixResponseHeaderItem `json:"header"`
		Body   [][]GetCorrelationMatrixResponseBodyItem `json:"body"`
	}
	GetCorrelationMatrixResponseHeaderItem struct {
		IndicatorID    int    `json:"indicatorID"`
		DatasetID      int    `json:"datasetID"`
		IndicatorTitle string `json:"indicatorTitle"`
		DatasetShared  bool   `json:"datasetShared"`
	}
	GetCorrelationMatrixResponseBodyItem struct {
		CorrelationID int       `json:"correlationID"`
		Coef          float64   `json:"coef"`
		P             float64   `json:"p"`
		R2            float64   `json:"r2"`
		Type          string    `json:"type"`
		UpdateTime    time.Time `json:"updateTime"`
	}
)
