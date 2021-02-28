package domain

import "time"

type (
	// User holds user data
	User struct {
		ID           int // id is passed to domain model for simplicity
		Username     string
		Password     string
		PasswordHash string
		Email        string // TODO: add ent validation
		// Service      bool // if user is a service
	}
	// Indicator is a set of user-created data
	Indicator struct {
		ID          int
		Code        string // unique code for external systems
		Title       string
		Description string
		Active      bool
		BuiltIn     bool // if Indicator is created by the service
		External    bool // if Indicator is populated by the user or external system
		Scale       *Scale
		Author      *User
		CreateTime  time.Time
		UpdateTime  time.Time
	}
	// Scale is a type of a scale for an Indicator
	Scale struct {
		ID          int
		Type        string // numeric, ordinal or nomial
		Title       string
		Description string
	}
	// Dataset is an intance of an Indicator populated by user data.
	// Each User can have one Dataset for each Indicator
	Dataset struct {
		ID           int
		User         *User
		Indicator    *Indicator
		CreateTime   time.Time
		UpdateTime   time.Time
		Observations []*Observation
		Source       string // user input or external system
		Shared       bool   // if dataset can be shared between all users
	}
	// Observation is a data point for an indicator
	Observation struct {
		ID         int
		Value      float64
		Dataset    *Dataset
		CreateTime time.Time
		UpdateTime time.Time
	}
	// Correlation is a corr value of a pair of Datasets
	Correlation struct {
		ID         int
		Left       *Dataset
		Right      *Dataset
		Coef       float64 // correlation coef
		P          float64 // p-value
		R2         float64 // determination coef
		Type       string  // Pearson or whatever
		CreateTime time.Time
		UpdateTime time.Time
	}
)
