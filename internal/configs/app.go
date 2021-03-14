package configs

type AppConfig struct {
	DefaultObservationLimit         int `yaml:"defaultObservationLimit"`
	MaxCorrelationObservations      int `yaml:"maxCorrelationObservations"`
	MinCorrelationObservations      int `yaml:"minCorrelationObservations"`
	UpdateCorrelationsBuffer        int `yaml:"updateCorrelationsBuffer"`
	UpdateAggregationsBuffer        int `yaml:"updateAggregationsBuffer"`
	MaxWeekAggregationObservations  int `yaml:"maxWeekAggregationObservations"`
	MaxMonthAggregationObservations int `yaml:"maxMonthAggregationObservations"`
}
