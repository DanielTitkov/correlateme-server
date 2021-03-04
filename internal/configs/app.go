package configs

type AppConfig struct {
	DefaultObservationLimit    int64 `yaml:"defaultObservationLimit"`
	MaxCorrelationObservations int64 `yaml:"maxCorrelationObservations"`
	MinCorrelationObservations int64 `yaml:"minCorrelationObservations"`
	UpdateCorrelationsBuffer   int   `yaml:"updateCorrelationsBuffer"`
}
