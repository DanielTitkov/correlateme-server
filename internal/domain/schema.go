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
)

type (
	Anomaly struct {
		ID                     int
		DetectionJobID         int     // related job
		DetectionJobInstanceID int     // related instance
		Type                   string  // warning or alarm
		Value                  float64 // outlier item value
		Processed              bool    // if anomaly is accepted/approved
		PeriodStart            time.Time
		PeriodEnd              time.Time
	}
	DetectionJob struct {
		ID          int
		Schedule    string // cron string
		Method      string // e.g, 3-sigmas
		SiteID      string
		Metric      string
		Attribute   string
		TimeAgo     string // e.g. 30d
		TimeStep    string // e.g. 1d
		Description string
	}
	DetectionJobInstance struct {
		ID             int
		DetectionJobID int // related job
		StartedAt      time.Time
		FinishedAt     time.Time
	}
	Dataset struct {
		SiteID    string
		Metric    string
		Attribute string
		StartDate time.Time
		EndDate   time.Time
		Data      []DataItem
	}
	DataItem struct {
		Timestamp time.Time
		Value     float64
	}
	Notification struct {
		ChannelID string
		Warnings  Anomaly
		Alarms    Anomaly
	}
)
