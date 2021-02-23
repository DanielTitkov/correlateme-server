package domain

type (
	FilterAnomaliesArgs struct {
		JobID     int
		Processed *bool
		// TODO: add other fields
	}
	FilterDetectionJobsArgs struct {
		ID        int
		SiteID    string
		Scheduled bool
		// TODO: add other fields
	}
)
