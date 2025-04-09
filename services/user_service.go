package services

import (
	"microservicio/models"
	"microservicio/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUserService(user models.User) (*mongo.InsertOneResult, error) {
	return repositories.CreateUser(user)
}

func GetUserByIDService(id string) (models.User, error) {
	return repositories.GetUserByID(id)
}

func GetAllUsersService() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func UpdateUserService(id string, user models.User) error {
	return repositories.UpdateUser(id, user)
}

func DeleteUserService(id string) error {
	return repositories.DeleteUser(id)
}
