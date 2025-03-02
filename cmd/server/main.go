package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/lehaisonaipro/task-management-api/internal/config"
	"github.com/lehaisonaipro/task-management-api/internal/controllers"
	"github.com/lehaisonaipro/task-management-api/internal/repositories"
	"github.com/lehaisonaipro/task-management-api/internal/routes"
	"github.com/lehaisonaipro/task-management-api/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	configFile := flag.String("config", "config.yaml", "Path to the configuration file")
	flag.Parse()

	conf, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatal("Error loading the configuration file: ", err)
	}

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(conf.Database.MongoDB.URI))
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	taskRepo := repositories.NewTaskRepository(client.Database(conf.Database.MongoDB.Database).Collection(conf.Database.MongoDB.Collection))
	taskService := services.NewTaskService(taskRepo)
	taskController := controllers.NewTaskController(taskService)

	// Set up the routes
	r := routes.SetupRouter(taskController)
	log.Println("Server is running on port", conf.Server.Port)
	r.Run(":" + fmt.Sprint(conf.Server.Port))

}
