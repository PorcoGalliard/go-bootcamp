package dto

import "hrapi/types"

type CreateEmployeeRequest struct {
	FirstName    *string    `json:"first_name" validate:"omitempty,min=2,max=20"`
	LastName     string     `json:"last_name" validate:"required,min=2,max=25"`
	Email        string     `json:"email" validate:"required,email,max=100"`
	PhoneNumber  *string    `json:"phone_number" validate:"omitempty,max=20"`
	HireDate     types.Date `json:"hire_date" validate:"required"`
	JobID        int32      `json:"job_id" validate:"required,gt=0"`
	Salary       float64    `json:"salary" validate:"required,gt=0"`
	ManagerID    *int32     `json:"manager_id" validate:"omitempty,gt=0"`
	DepartmentID *int32     `json:"department_id" validate:"omitempty,gt=0"`
}

type UpdateEmployeeRequest struct {
	FirstName    *string     `json:"first_name,omitempty" validate:"omitempty,min=2,max=20"`
	LastName     *string     `json:"last_name,omitempty" validate:"omitempty,min=2,max=25"`
	Email        *string     `json:"email,omitempty" validate:"omitempty,email,max=100"`
	PhoneNumber  *string     `json:"phone_number,omitempty" validate:"omitempty,max=20"`
	HireDate     *types.Date `json:"hire_date,omitempty" validate:"omitempty"`
	JobID        *int32      `json:"job_id,omitempty" validate:"omitempty,gt=0"`
	Salary       *float64    `json:"salary,omitempty" validate:"omitempty,gt=0"`
	ManagerID    *int32      `json:"manager_id,omitempty" validate:"omitempty,gt=0"`
	DepartmentID *int32      `json:"department_id,omitempty" validate:"omitempty,gt=0"`
}

type EmployeeResponse struct {
	EmployeeID   int32      `json:"employee_id"`
	FirstName    *string    `json:"first_name,omitempty"`
	LastName     string     `json:"last_name"`
	Email        string     `json:"email"`
	PhoneNumber  *string    `json:"phone_number,omitempty"`
	HireDate     types.Date `json:"hire_date"`
	JobID        int32      `json:"job_id"`
	Salary       float64    `json:"salary"`
	ManagerID    *int32     `json:"manager_id,omitempty"`
	DepartmentID *int32     `json:"department_id,omitempty"`
}
