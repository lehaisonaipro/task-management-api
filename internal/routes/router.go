package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lehaisonaipro/task-management-api/internal/controllers"
	"github.com/lehaisonaipro/task-management-api/internal/middlewares"
)

func SetupRouter(taskController *controllers.TaskController) *gin.Engine {
	r := gin.Default()
	r.POST("/tasks", middlewares.RoleAuthorization("Employer"), taskController.CreateTask)
	r.GET("/tasks/assigned/:userId", middlewares.RoleAuthorization("Employee", "Employer"), taskController.GetAssignedTasks)
	r.GET("/tasks", middlewares.RoleAuthorization("Employer"), taskController.GetAllTasks)
	return r
}
