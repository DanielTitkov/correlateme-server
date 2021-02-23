package model

import "github.com/DanielTitkov/correlateme-server/internal/domain"

// Domain models are used for simplicity.
// In production-grade system there would be separate models.

type (
	ListJobsRequest struct {
		Filter domain.FilterDetectionJobsArgs
	}
	ListJobsResponse struct {
		Jobs []*domain.DetectionJob
	}
	AddJobRequest struct {
		Sync bool
		Job  domain.DetectionJob
	}
	AddJobResponse struct {
		Job      *domain.DetectionJob
		Instance *domain.DetectionJobInstance
		Result   []*domain.Anomaly
	}
	DeleteJobRequest struct {
	}
	DeleteJobResponse struct {
	}
)
