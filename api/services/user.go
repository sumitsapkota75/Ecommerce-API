package services

import (
	"travel/api/repository"
	"travel/models"

	"gorm.io/gorm"
)

// UserService struct
type UserService struct {
	repository repository.UserRepository
}

//NewUserService -> creates a new user service
func NewUserService(repository repository.UserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

// WithTrx -> enables repository with transaction
func (s UserService) WithTrx(trxHandle *gorm.DB) UserService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// CreateUser -> call to create a user
func (s UserService) CreateUser(user models.User) (models.User, error) {
	return s.repository.CreateUser(user)
}
