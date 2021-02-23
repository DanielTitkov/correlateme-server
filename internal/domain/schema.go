package domain

import "time"

type (
	// User holds user data
	User struct {
		ID           int  // id is passed to domain model for simplicity
		Service      bool // if user is a service // TODO: only service users create items
		Username     string
		Password     string
		PasswordHash string
		Email        string // TODO: add ent validation
	}
	// Indicator is a set of user-created data
	Indicator struct {
		ID          int
		Code        string // unique code for external systems
		Slug        string
		Title       string
		Description string
		Type        string
		Active      bool
		BuiltIn     bool // if Indicator is created by the service
		External    bool // if Indicator is populated by the user or external system
		Author      User
		CreateTime  time.Time
		UpdateTime  time.Time
	}
	// Dataset is an intance of an Indicator populated by user data.
	// Each User can have one Dataset for each Indicator
	Dataset struct {
		ID           int
		User         User
		Indicator    Indicator
		CreateTime   time.Time
		UpdateTime   time.Time
		Observations []*Observation
		Source       string // user input or external system
		Shared       bool   // if dataset can be shared between all users
	}
	// Observation is a data point for an indicator
	Observation struct {
		ID         int
		Dataset    Dataset
		CreateTime time.Time
		Value      float64
	}
	// Correlation is a corr value of a pair of Datasets
	Correlation struct {
		ID         int
		Left       Dataset
		Right      Dataset
		Coef       float64 // correlation coef
		P          float64 // p-value
		R2         float64 // determination coef
		Type       string  // Pearson or whatever
		CreateTime time.Time
		UpdateTime time.Time
	}
)
