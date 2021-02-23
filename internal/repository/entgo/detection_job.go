package entgo

import (
	"context"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent/detectionjob"
)

func (r *EntgoRepository) CreateDetectionJob(j *domain.DetectionJob) (*domain.DetectionJob, error) {
	var schedule *string
	if j.Schedule != "" {
		schedule = &j.Schedule
	}

	job, err := r.client.DetectionJob.
		Create().
		SetNillableSchedule(schedule).
		SetSiteID(j.SiteID).
		SetTimeAgo(j.TimeAgo).
		SetTimeStep(j.TimeStep).
		SetMethod(j.Method).
		SetMetric(j.Metric).
		SetAttribute(j.Attribute).
		SetDescription(j.Description).
		Save(context.TODO())

	if err != nil {
		return nil, err
	}

	j.ID = job.ID
	return j, nil
}

func (r *EntgoRepository) DeleteDetectionJobByID(id int) error {
	return r.client.DetectionJob.DeleteOneID(id).Exec(context.TODO())
}

func (r *EntgoRepository) GetDetectionJobByID(id int) (*domain.DetectionJob, error) {
	job, err := r.client.DetectionJob.Query().
		Where(detectionjob.IDEQ(id)).
		Only(context.TODO())

	if err != nil {
		return nil, err
	}

	return entToDomainDetectionJob(job), nil
}

func (r *EntgoRepository) FilterDetectionJobs(args *domain.FilterDetectionJobsArgs) ([]*domain.DetectionJob, error) {
	query := r.client.DetectionJob.Query()

	if args.ID != 0 {
		query = query.Where(detectionjob.IDEQ(args.ID))
	}

	if args.SiteID != "" {
		query = query.Where(detectionjob.SiteIDEQ(args.SiteID))
	}

	if args.Scheduled {
		query = query.Where(detectionjob.ScheduleNotNil())
	}

	jobs, err := query.All(context.TODO())
	if err != nil {
		return []*domain.DetectionJob{}, err
	}

	var res []*domain.DetectionJob
	for _, job := range jobs {
		res = append(res, entToDomainDetectionJob(job))
	}

	return res, nil
}

func entToDomainDetectionJob(job *ent.DetectionJob) *domain.DetectionJob {
	var schedule string
	if job.Schedule != nil {
		schedule = *job.Schedule
	}

	return &domain.DetectionJob{
		ID:          job.ID,
		Schedule:    schedule,
		Method:      job.Method,
		SiteID:      job.SiteID,
		Metric:      job.Metric,
		Attribute:   job.Attribute,
		TimeAgo:     job.TimeAgo,
		TimeStep:    job.TimeStep,
		Description: job.Description,
	}
}
