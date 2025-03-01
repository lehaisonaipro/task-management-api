package main

import (
	"context"
	"log"

	"github.com/lehaisonaipro/task-management-api/internal/controllers"
	"github.com/lehaisonaipro/task-management-api/internal/repositories"
	"github.com/lehaisonaipro/task-management-api/internal/routes"
	"github.com/lehaisonaipro/task-management-api/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	taskRepo := repositories.NewTaskRepository(client.Database("taskmanager").Collection("tasks"))
	taskService := services.NewTaskService(taskRepo)
	taskController := controllers.NewTaskController(taskService)

	// Set up the routes
	r := routes.SetupRouter(taskController)
	r.Run(":8080")
}
