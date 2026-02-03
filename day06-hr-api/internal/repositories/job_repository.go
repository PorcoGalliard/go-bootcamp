package repositories

import (
	"context"
	"hrapi/internal/models"

	"gorm.io/gorm"
)

type JobRepository interface {
	FindAll(ctx context.Context) ([]*models.Job, error)
	FindByID(ctx context.Context, id uint) (*models.Job, error)
	Create(ctx context.Context, job *models.Job) error
	Update(ctx context.Context, job *models.Job) error
	Delete(ctx context.Context, id uint) error
}

type jobRepository struct {
	DB *gorm.DB
}

func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobRepository{
		DB: db,
	}
}

func (j *jobRepository) FindAll(ctx context.Context) ([]*models.Job, error) {
	var jobs []*models.Job
	err := j.DB.WithContext(ctx).Find(jobs).Error
	return jobs, err
}

func (j *jobRepository) FindByID(ctx context.Context, id uint) (*models.Job, error) {
	var job *models.Job
	err := j.DB.WithContext(ctx).First(job, id).Error
	return job, err
}

func (j *jobRepository) Create(ctx context.Context, job *models.Job) error {
	return j.DB.WithContext(ctx).Create(job).Error
}

func (j *jobRepository) Update(ctx context.Context, job *models.Job) error {
	return j.DB.WithContext(ctx).Save(job).Error
}

func (j *jobRepository) Delete(ctx context.Context, id uint) error {
	return j.DB.WithContext(ctx).Delete(&models.Job{}, id).Error
}
