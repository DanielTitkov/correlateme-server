package model

type (
	// Dataset is a common model to use in various methods
	Dataset struct {
		ID           int           `json:"id,omitempty"`
		Source       string        `json:"source,omitempty"`
		Shared       bool          `json:"shared,omitempty"`
		Observations []Observation `json:"observations,omitempty"`
	}
)
