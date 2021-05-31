package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User -> Model

type User struct {
	Base
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserType string `json:"user_type,omitempty"`
	Phone    uint64 `json:"phone"`
	Address  string `json:"address"`
}

// TableName  -> returns table name of model
func (u User) TableName() string {
	return "users"
}

//BeforeCreate -> Called before inserting record into Column Table
func (u *User) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	u.ID = BINARY16(id)
	return err
}
