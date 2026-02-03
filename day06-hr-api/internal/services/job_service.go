package services

import (
	"context"
	"hrapi/internal/models"
	"hrapi/internal/repositories"
)

type JobService interface {
	GetAllJobs(ctx context.Context) ([]*models.Job, error)
	GetJobByID(ctx context.Context, id uint) (*models.Job, error)
	CreateJob(ctx context.Context, job *models.Job) error
	UpdateJob(ctx context.Context, job *models.Job) error
	DeleteJob(ctx context.Context, id uint) error
}

type jobService struct {
	jobRepo repositories.JobRepository
}

func NewJobService(jobRepository repositories.JobRepository) JobService {
	return &jobService{
		jobRepo: jobRepository,
	}
}

func (j *jobService) GetAllJobs(ctx context.Context) ([]*models.Job, error) {
	return j.jobRepo.FindAll(ctx)
}

func (j *jobService) GetJobByID(ctx context.Context, id uint) (*models.Job, error) {
	return j.jobRepo.FindByID(ctx, id)
}

func (j *jobService) CreateJob(ctx context.Context, job *models.Job) error {
	return j.jobRepo.Create(ctx, job)
}

func (j *jobService) UpdateJob(ctx context.Context, job *models.Job) error {
	return j.jobRepo.Update(ctx, job)
}

func (j *jobService) DeleteJob(ctx context.Context, id uint) error {
	return j.jobRepo.Delete(ctx, id)
}
