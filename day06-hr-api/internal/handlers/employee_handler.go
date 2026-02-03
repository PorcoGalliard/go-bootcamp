package handlers

import (
	"fmt"
	"hrapi/internal/dto"
	"hrapi/internal/response"
	"hrapi/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	service services.EmployeeService
}

func NewEmployeeHandler(svc services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		service: svc,
	}
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var req dto.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid Input")
		fmt.Print(err.Error())
		return
	}

	resp, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, "Failed to create new employee data")
		return
	}

	response.SendResponse(c, http.StatusCreated, "Employee data created successfully", resp)
}

func (h *EmployeeHandler) GetAllEmployees(c *gin.Context) {
	resp, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to fetch employees: "+err.Error())
		return
	}

	response.SendResponse(c, http.StatusOK, "Employees retrieved successfully", resp)
}

func (h *EmployeeHandler) SearchEmployees(c *gin.Context) {
	name := c.Query("q")
	if name == "" {
		response.SendError(c, http.StatusBadRequest, "Search query required")
		return
	}

	resp, err := h.service.GetByName(c.Request.Context(), name)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to search employees: "+err.Error())
		return
	}

	if len(resp) == 0 {
		response.SendError(c, http.StatusNotFound, "No employees found")
		return
	}

	response.SendResponse(c, http.StatusOK, "Employees retrieved successfully", resp)
}

func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	var req dto.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	resp, err := h.service.Update(c.Request.Context(), uint(id), &req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Failed to update employee: "+err.Error())
		return
	}

	response.SendResponse(c, http.StatusOK, "Employee updated successfully", resp)
}

func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		response.SendError(c, http.StatusNotFound, "Employee not found: "+err.Error())
		return
	}

	response.SendResponse(c, http.StatusOK, "Employee deleted successfully", nil)
}
