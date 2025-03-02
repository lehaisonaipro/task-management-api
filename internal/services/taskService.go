package services

import (
	"github.com/lehaisonaipro/task-management-api/internal/models"
	"github.com/lehaisonaipro/task-management-api/internal/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

type ITaskService interface {
	CreateTask(task *models.Task) error
	UpdateTaskStatus(employeeID string, taskID string, status string) error
	ViewTaskSummary(employeeID string) (map[string]int, error)
}

type TaskService struct {
	taskRepo repositories.ITaskRepository
}

// ViewTaskSummary implements ITaskService.
func (service *TaskService) ViewTaskSummary(employeeID string) (map[string]int, error) {

}

// ViewTasks implements ITaskService.
func (service *TaskService) ViewTasks(employeeID string, status string, sortDateType string) ([]models.Task, error) {
	panic("unimplemented")
}

func NewTaskService(taskRepo repositories.ITaskRepository) ITaskService {
	return &TaskService{taskRepo}
}

func (service *TaskService) CreateTask(task *models.Task) error {
	return service.taskRepo.CreateTask(task)
}

func (service *TaskService) GetAssignedTasks(userId string) ([]models.Task, error) {
	filter := bson.M{"assigned_to": userId}
	return service.taskRepo.GetTasks(filter)
}

func (service *TaskService) GetAllTasks(filters bson.M) ([]models.Task, error) {
	return service.taskRepo.GetTasks(filters)
}
