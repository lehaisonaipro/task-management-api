package repositories_test

import (
	"context"
	"testing"

	"github.com/lehaisonaipro/task-management-api/internal/models"
	"github.com/lehaisonaipro/task-management-api/internal/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Test CreateTask
func TestCreateTask(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Error("[ERROR] Error connecting to MongoDB: ", err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		t.Error("[ERROR] Error pinging MongoDB: ", err)
	}

	collection := client.Database("task_management").Collection("tasks")
	repo := repositories.NewTaskRepository(collection)
	err = repo.CreateTask(&models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "Pending",
	})
}

// Test GetTasks
func TestGetTasks(t *testing.T) {

}
