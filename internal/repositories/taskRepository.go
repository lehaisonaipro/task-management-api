package repositories

import (
	"context"

	"github.com/lehaisonaipro/task-management-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) *TaskRepository {
	return &TaskRepository{collection}
}

func (repo *TaskRepository) CreateTask(task *models.Task) error {
	_, err := repo.collection.InsertOne(context.Background(), task)
	return err
}

func (repo *TaskRepository) GetTasks(filter bson.M) ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := repo.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
