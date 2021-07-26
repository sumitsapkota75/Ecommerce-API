package models

import (
	"gorm.io/gorm"
)

// Order struct-> models
type Order struct {
	Base
	FirstName     string         `json:"first_name"`
	LastName      string         `json:"last_name"`
	CompanyName   string         `json:"company_name"`
	Country       string         `json:"country"`
	StreetAddress string         `json:"street_address"`
	City          string         `json:"city"`
	State         string         `json:"state"`
	Zip           string         `json:"zip"`
	Phone         string         `json:"phone"`
	Email         string         `json:"email"`
	Notes         string         `json:"notes"`
	TotalAmount   string         `json:"total_amount"`
	PaidAmount    string         `json:"paid_amount"`
	OrderItem     []OrderItem    `json:"order_item"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName gives table name of model
func (p Order) TableName() string {
	return "orders"
}

//OrderItem model -> struct
type OrderItem struct {
	Base
	ProductID string  `json:"product_id"`
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity"`
	OrderID   string  `json:"order_id"`
}

// TableName gives table name of model
func (o OrderItem) TableName() string {
	return "order_items"
}
