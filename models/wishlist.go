package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Wishlist Model -> struct
type Wishlist struct {
	Base
	UserID    string `json:"user_id"`
	ProductID string `json:"product_id"`
}

// TableName  -> returns table name of model
func (u Wishlist) TableName() string {
	return "wishlists"
}

// BeforeCreate -> Called before inserting record into Column Table
func (u *Wishlist) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	u.ID = BINARY16(id)
	return err
}
