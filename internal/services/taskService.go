package services

import (
	"github.com/lehaisonaipro/task-management-api/internal/models"
	"github.com/lehaisonaipro/task-management-api/internal/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

type TaskService struct {
	taskRepo *repositories.TaskRepository
}

func NewTaskService(taskRepo *repositories.TaskRepository) *TaskService {
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
