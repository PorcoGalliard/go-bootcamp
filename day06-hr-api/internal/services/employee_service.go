package services

import (
	"context"
	"hrapi/internal/dto"
	"hrapi/internal/errors"
	"hrapi/internal/models"
	"hrapi/internal/repositories"
	"hrapi/types"
	"strings"

	"github.com/go-playground/validator/v10"
)

type EmployeeService interface {
	GetAll(ctx context.Context) ([]dto.EmployeeResponse, error)
	GetByName(ctx context.Context, name string) ([]dto.EmployeeResponse, error)
	Create(ctx context.Context, req *dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error)
	Update(ctx context.Context, id uint, req *dto.UpdateEmployeeRequest) (*dto.EmployeeResponse, error)
	Delete(ctx context.Context, id uint) error
}

type employeeService struct {
	repo      repositories.EmployeeRepository
	validator *validator.Validate
}

func NewEmployeeService(repo repositories.EmployeeRepository) EmployeeService {
	return &employeeService{
		repo:      repo,
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (s *employeeService) GetAll(ctx context.Context) ([]dto.EmployeeResponse, error) {
	employees, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return s.mapResponseList(employees), nil
}

func (s *employeeService) GetByName(ctx context.Context, name string) ([]dto.EmployeeResponse, error) {
	if name == "" {
		return s.GetAll(ctx)
	}

	name = strings.TrimSpace(name)
	nameParts := strings.Fields(name)

	var firstName, lastName string

	if len(nameParts) == 1 {
		firstName = nameParts[0]
		lastName = nameParts[0]
	} else if len(nameParts) >= 2 {
		firstName = nameParts[0]
		lastName = strings.Join(nameParts[1:], " ")
	}

	employees, err := s.repo.FindByName(ctx, firstName, lastName)
	if err != nil {
		return nil, err
	}

	return s.mapResponseList(employees), nil
}

func (s *employeeService) Create(ctx context.Context, req *dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, errors.ErrInvalidInput
	}

	employee := &models.Employee{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
		HireDate:     req.HireDate.ToTime(),
		JobID:        req.JobID,
		Salary:       req.Salary,
		ManagerID:    req.ManagerID,
		DepartmentID: req.DepartmentID,
	}

	if err := s.repo.Create(ctx, employee); err != nil {
		return nil, err
	}

	return s.mapResponse(employee), nil
}

func (s *employeeService) Update(ctx context.Context, id uint, req *dto.UpdateEmployeeRequest) (*dto.EmployeeResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, errors.ErrInvalidInput
	}

	employee, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.FirstName != nil {
		employee.FirstName = req.FirstName
	}
	if req.LastName != nil {
		employee.LastName = *req.LastName
	}
	if req.Email != nil {
		employee.Email = *req.Email
	}
	if req.PhoneNumber != nil {
		employee.PhoneNumber = req.PhoneNumber
	}
	if req.HireDate != nil {
		employee.HireDate = req.HireDate.ToTime()
	}
	if req.JobID != nil {
		employee.JobID = *req.JobID
	}
	if req.Salary != nil {
		employee.Salary = *req.Salary
	}
	if req.ManagerID != nil {
		employee.ManagerID = req.ManagerID
	}
	if req.DepartmentID != nil {
		employee.DepartmentID = req.DepartmentID
	}

	if err := s.repo.Update(ctx, employee); err != nil {
		return nil, err
	}

	return s.mapResponse(employee), nil
}

func (s *employeeService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *employeeService) mapResponse(employee *models.Employee) *dto.EmployeeResponse {
	return &dto.EmployeeResponse{
		EmployeeID:   employee.EmployeeID,
		FirstName:    employee.FirstName,
		LastName:     employee.LastName,
		Email:        employee.Email,
		PhoneNumber:  employee.PhoneNumber,
		HireDate:     types.Date(employee.HireDate),
		JobID:        employee.JobID,
		Salary:       employee.Salary,
		ManagerID:    employee.ManagerID,
		DepartmentID: employee.DepartmentID,
	}
}

func (s *employeeService) mapResponseList(employeeList []*models.Employee) []dto.EmployeeResponse {
	var responses []dto.EmployeeResponse
	for _, emp := range employeeList {
		responses = append(responses, *s.mapResponse(emp))
	}

	return responses
}
