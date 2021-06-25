package services

import (
	"gin-user-management/models"
	"gin-user-management/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

func (userService UserService) GetById(id interface{}) (user models.User, error error) {
	return user, userService.repository.Find(&user, "id = ?", id).Error
}