package repositories

import (
	"github.com/lehaisonaipro/task-management-api/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	GetEmployees() ([]models.User, error)
}

type UserRepository struct {
	collection *mongo.Collection
}
