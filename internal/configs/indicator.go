package configs

type IndicatorConfig struct {
	DefaultObservationLimit    int64 `yaml:"defaultObservationLimit"`
	MaxCorrelationObservations int64 `yaml:"maxCorrelationObservations"`
	MinCorrelationObservations int64 `yaml:"minCorrelationObservations"`
}
