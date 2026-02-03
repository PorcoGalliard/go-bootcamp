package handlers

import (
	"hrapi/internal/models"
	"hrapi/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	jobService services.JobService
}

func NewJobHandler(jobService services.JobService) *JobHandler {
	return &JobHandler{
		jobService: jobService,
	}
}

func (j *JobHandler) GetJobs(ctx *gin.Context) {
	regions, err := j.jobService.GetAllJobs(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to get all jobs",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    regions,
	})
}

func (j *JobHandler) GetJob(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid Job ID",
		})
		return
	}

	job, err := j.jobService.GetJobByID(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": "Job Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    job,
	})
}

func (j *JobHandler) CreateJob(ctx *gin.Context) {
	var job models.Job
	if err := ctx.ShouldBindJSON(&job); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "invalid input",
			"details": err.Error(),
		})
		return
	}

	if err := j.jobService.CreateJob(ctx.Request.Context(), &job); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to create job",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Job created successfully",
		"data":    job,
	})
}

func (j *JobHandler) UpdateJob(ctx *gin.Context) {

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid input ID",
			"details": err.Error(),
		})
		return
	}

	var job models.Job
	if err := ctx.ShouldBindJSON(&job); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "invalid input",
			"details": err.Error(),
		})
		return
	}

	job.JobID = int32(id)
	if err := j.jobService.UpdateJob(ctx.Request.Context(), &job); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to Update Job",
			"details": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully updating job",
		"data":    job,
	})
}

func (j *JobHandler) DeleteJob(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid ID Input",
			"details": err.Error(),
		})
		return
	}

	if err := j.jobService.DeleteJob(ctx.Request.Context(), uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to delete job",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully deleting job",
	})
}
