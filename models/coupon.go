package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Coupon struct {
	Base
	Code             string    `json:"code"`
	Device           string    `json:"device"`
	Discount         int64     `json:"discount"`
	DiscountType     string    `json:"discount_type"`
	ValidFrom        time.Time `json:"valid_from"`
	ValidTo          time.Time `json:"valid_to"`
	MinCheckoutValue float64   `json:"min_checkout_value"`
	IsActive         int       `json:"is_active"`
}

// TableName  -> returns table name of model
func (c Coupon) TableName() string {
	return "coupons"
}

// BeforeCreate -> Called before inserting record into Column Table
func (u *Coupon) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	u.ID = BINARY16(id)
	return err
}
