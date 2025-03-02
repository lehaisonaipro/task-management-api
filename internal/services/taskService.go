package services

import (
	"errors"

	"github.com/lehaisonaipro/task-management-api/internal/models"
	"github.com/lehaisonaipro/task-management-api/internal/repositories"
)

type ITaskService interface {
	CreateTask(task *models.Task) error
	UpdateTaskStatus(employeeID string, taskID string, status string) error
	AssignTask(taskID string, employeeID string) error
	GetAssignedTasks(employeeID string) ([]*models.Task, error)
	GetAllTasks(filters map[string]string, sorts map[string]int) ([]*models.Task, error)
	GetTaskSummary(employeeID string) ([]*models.TaskSummary, error)
	GetAllTaskSummary() (*models.TaskSummary, error)
}

type TaskService struct {
	taskRepo repositories.ITaskRepository
}

// AssignTask implements ITaskService.
func (t *TaskService) AssignTask(taskID string, employeeID string) error {
	_, err := t.taskRepo.GetTask(taskID)
	if err != nil {
		return errors.New("NOT FOUND TASK")
	}
	return t.taskRepo.AssignTask(taskID, employeeID)
}

// CreateTask implements ITaskService.
func (t *TaskService) CreateTask(task *models.Task) error {
	return t.taskRepo.CreateTask(task)
}

// GetAllTaskSummary implements ITaskService.
func (t *TaskService) GetAllTaskSummary() (*models.TaskSummary, error) {
	panic("unimplemented")
}

// GetAllTasks implements ITaskService.
func (t *TaskService) GetAllTasks(filters map[string]string, sorts map[string]int) ([]*models.Task, error) {
	employeeID := filters["assigned_to"]
	status := filters["status"]
	sortDateType := ""
	sortDateAsc := false
	sortStatus := false
	sortStatusAsc := false
	for k, v := range sorts {
		if k == "created_at" || k == "due_date" {
			sortDateType = k
			if v == 1 {
				sortDateAsc = true
			}
		}
		if k == "status" {
			sortStatus = true
			if v == 1 {
				sortStatusAsc = true
			}
		}
	}
	return t.taskRepo.ViewTasks(employeeID, status, sortDateType, sortDateAsc, sortStatus, sortStatusAsc)
}

// GetAssignedTasks implements ITaskService.
func (t *TaskService) GetAssignedTasks(employeeID string) ([]*models.Task, error) {
	panic("unimplemented")
}

// GetTaskSummary implements ITaskService.
func (t *TaskService) GetTaskSummary(employeeID string) ([]*models.TaskSummary, error) {
	panic("unimplemented")
}

// UpdateTaskStatus implements ITaskService.
func (t *TaskService) UpdateTaskStatus(employeeID string, taskID string, status string) error {
	panic("unimplemented")
}

func NewTaskService(taskRepo repositories.ITaskRepository) ITaskService {
	return &TaskService{taskRepo}
}
