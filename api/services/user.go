package services

import (
	"travel/api/repository"
	"travel/models"
	"travel/utils"

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

// GetUserByID -> returns a user by UID
func (s UserService) GetUserByID(userID string) (models.User, error) {
	return s.repository.GetUserByID(userID)
}

// GetAllUsers -> returns a list of user
func (s UserService) GetAllUsers(pagination utils.Pagination) ([]models.User, int64, error) {
	return s.repository.GetAllUsers(pagination)
}

// UpdateUser -> updates the user data
func (s UserService) UpdateUser(user models.User) error {
	return s.repository.UpdateUser(user)
}
