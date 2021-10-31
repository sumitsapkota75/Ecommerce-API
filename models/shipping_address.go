package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShippingAddressSearchParams struct {
	Keyword string `json:"keyword"`
}

type ShippingAddress struct {
	Base
	UserID   string `json:"user_id","-"`
	User     User   `json:"user"`
	District string `json:"district"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

func (s ShippingAddress) TableName() string {
	return "shipping_addresses"
}

// BeforeCreate -> Called before inserting record into Column Table
func (u *ShippingAddress) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	u.ID = BINARY16(id)
	return err
}
