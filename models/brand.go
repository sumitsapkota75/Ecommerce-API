package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//BrandSearchParams -> search Category Params
type BrandSearchParams struct {
	Keyword string `json:"keyword"`
}

//Category -> category model
type Brand struct {
	Base
	Name        string `json:"name"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}

// TableName  -> returns table name of model
func (b Brand) TableName() string {
	return "brands"
}

// BeforeCreate -> Called before inserting record into Column Table
func (u *Brand) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	u.ID = BINARY16(id)
	return err
}
