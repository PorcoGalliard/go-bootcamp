package repositories

import (
	"context"
	"hrapi/internal/domain/query"
	"hrapi/internal/models"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindAll(ctx context.Context) ([]*models.Employee, error)
	FindByName(ctx context.Context, firstName, lastName string) ([]*models.Employee, error)
	FindByID(ctx context.Context, id uint) (*models.Employee, error)
	Create(ctx context.Context, employee *models.Employee) error
	Update(ctx context.Context, employee *models.Employee) error
	Delete(ctx context.Context, id uint) error
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{
		Q: query.Use(db),
	}
}

type employeeRepository struct {
	Q *query.Query
}

func (r *employeeRepository) FindAll(ctx context.Context) ([]*models.Employee, error) {
	var employees []*models.Employee
	employees, err := r.Q.Employee.WithContext(ctx).Find()
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *employeeRepository) FindByName(ctx context.Context, firstName, lastName string) ([]*models.Employee, error) {
	employees, err := r.Q.Employee.WithContext(ctx).Where(
		r.Q.Employee.FirstName.Like("%" + firstName + "%"),
	).Or(
		r.Q.Employee.LastName.Like("%" + lastName + "%"),
	).Find()
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *employeeRepository) FindByID(ctx context.Context, id uint) (*models.Employee, error) {
	employee, err := r.Q.Employee.WithContext(ctx).Where(r.Q.Employee.EmployeeID.Eq(int32(id))).First()
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *employeeRepository) Create(ctx context.Context, employee *models.Employee) error {
	return r.Q.Employee.WithContext(ctx).Create(employee)
}

func (r *employeeRepository) Update(ctx context.Context, employee *models.Employee) error {
	return r.Q.Employee.WithContext(ctx).Save(employee)
}

func (r *employeeRepository) Delete(ctx context.Context, id uint) error {
	result, err := r.Q.Employee.WithContext(ctx).Where(r.Q.Employee.EmployeeID.Eq(int32(id))).Delete(&models.Employee{})
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
