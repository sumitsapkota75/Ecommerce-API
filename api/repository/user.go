package repository

import (
	"travel/infrastructure"
	"travel/models"

	"gorm.io/gorm"
)

// UserRepository database structure
type UserRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewUserRepository -> creates a new user repository
func NewUserRepository(db infrastructure.Database, logger infrastructure.Logger) UserRepository {
	return UserRepository{
		db:     db,
		logger: logger,
	}
}

// WithTrx -> enables repository with transaction
func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		r.logger.Zap.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.db.DB = trxHandle
	return r
}

// CreateUser -> creates new user
func (u UserRepository) CreateUser(user models.User) (models.User, error) {
	return user, u.db.DB.Create(&user).Error
}

//GetUserByID -> gets the user by uid
func (u UserRepository) GetUserByID(userID string) (user models.User, err error) {
	return user, u.db.DB.Where("id = ?", userID).First(&user).Error
}
