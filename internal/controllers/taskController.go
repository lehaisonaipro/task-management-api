package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lehaisonaipro/task-management-api/internal/models"
	"github.com/lehaisonaipro/task-management-api/internal/services"
)

type TaskController struct {
	service services.ITaskService
}

func NewTaskController(service services.ITaskService) *TaskController {
	return &TaskController{service}
}

func (ctrl *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.service.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully"})
}

func (ctrl *TaskController) GetAssignedTasks(c *gin.Context) {
	userId := c.Param("userId")
	tasks, err := ctrl.service.GetAssignedTasks(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (ctrl *TaskController) GetAllTasks(c *gin.Context) {

	filters := make(map[string]string)
	sorts := make(map[string]int)
	// Extract filters from query parameters
	if assignedTo := c.Query("assigned_to"); assignedTo != "" {
		filters["assigned_to"] = strings.Trim(assignedTo, `"`) // Remove quotes
	}
	if status := c.Query("status"); status != "" {
		filters["status"] = strings.Trim(status, `"`) // Remove quotes
	}

	// Extract sorting options
	if sortDateType := c.Query("sort_date_type"); sortDateType != "" {
		sortOrder := 1 // Default to ascending
		if c.Query("sort_date_asc") == "false" {
			sortOrder = -1 // Descending order
		}
		sorts[strings.Trim(sortDateType, `"`)] = sortOrder
	}

	if c.Query("sort_by_status") == "true" {
		sortOrder := 1
		if c.Query("sort_by_status_asc") == "false" {
			sortOrder = -1
		}
		sorts["status"] = sortOrder
	}

	tasks, err := ctrl.service.GetAllTasks(filters, sorts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (ctrl *TaskController) GetTaskSummary(c *gin.Context) {

}
