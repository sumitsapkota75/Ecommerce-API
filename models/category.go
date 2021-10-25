package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CategorySearchParams -> search Category Params
type CategorySearchParams struct {
	Keyword string `json:"keyword"`
}

// Category -> category model
type Category struct {
	Base
	Name        string `json:"name"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}

// TableName  -> returns table name of model
func (c Category) TableName() string {
	return "categories"
}

// BeforeCreate -> Called before inserting record into Column Table
func (u *Category) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	u.ID = BINARY16(id)
	return err
}
