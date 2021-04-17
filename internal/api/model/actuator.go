package model

type (
	ServiceStatsResponse struct {
		Users        int `json:"users,omitempty"`
		Scales       int `json:"scales,omitempty"`
		Indicators   int `json:"indicators,omitempty"`
		Datasets     int `json:"datasets,omitempty"`
		Observations int `json:"observations,omitempty"`
		Correlations int `json:"correlations,omitempty"`
	}
)
