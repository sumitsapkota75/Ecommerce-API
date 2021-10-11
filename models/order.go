package models

import (
	"gorm.io/gorm"
)

//OrderSearchParams -> order search struct
type OrderSearchParams struct {
	Keyword string `json:"keyword"`
}

// Order struct-> models
type Order struct {
	UintBase
	User
	UserID        string         `json:"user_id"`
	FirstName     string         `json:"first_name"`
	LastName      string         `json:"last_name"`
	CompanyName   string         `json:"company_name"`
	StreetAddress string         `json:"street_address"`
	City          string         `json:"city"`
	State         string         `json:"state"`
	Zip           string         `json:"zip"`
	Phone         string         `json:"phone"`
	Email         string         `json:"email"`
	Notes         string         `json:"notes"`
	TotalAmount   float64        `json:"total_amount"`
	PaidAmount    float64        `json:"paid_amount"`
	OrderItem     []OrderItem    `json:"order_item"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName gives table name of model
func (p Order) TableName() string {
	return "orders"
}

//OrderItem model -> struct
type OrderItem struct {
	UintBase
	ProductID int            `json:"product_id"`
	Product   Product        `json:"product"`
	Quantity  int            `json:"quantity"`
	OrderID   int            `json:"order_id"`
	Price     float64        `json:"price"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// TableName gives table name of model
func (o OrderItem) TableName() string {
	return "order_items"
}
