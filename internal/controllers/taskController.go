package controllers

import (
	"net/http"

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
	filters := make(map[string]interface{})
	// Add filtering logic based on query params
	tasks, err := ctrl.service.GetAllTasks(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
