package configs

type AppConfig struct {
	DefaultObservationLimit    int `yaml:"defaultObservationLimit"`
	MaxCorrelationObservations int `yaml:"maxCorrelationObservations"`
	MinCorrelationObservations int `yaml:"minCorrelationObservations"`
	UpdateCorrelationsBuffer   int `yaml:"updateCorrelationsBuffer"`
}
