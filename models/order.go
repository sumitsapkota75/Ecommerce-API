package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//OrderSearchParams -> order search struct
type OrderSearchParams struct {
	Keyword string `json:"keyword"`
}

// Order struct-> models
type Order struct {
	Base
	User          User           `json:"user"`
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
	OrderStatus   string         `json:"order_status"`
	OrderItem     []OrderItem    `json:"order_item"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName gives table name of model
func (p Order) TableName() string {
	return "orders"
}

// BeforeCreate -> Called before inserting record into Column Table
func (u *Order) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	u.ID = BINARY16(id)
	return err
}

//OrderItem model -> struct
type OrderItem struct {
	Base
	ProductID BINARY16       `json:"product_id"`
	Product   Product        `json:"product"`
	Quantity  int            `json:"quantity"`
	OrderID   BINARY16       `json:"order_id"`
	Price     float64        `json:"price"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// TableName gives table name of model
func (o OrderItem) TableName() string {
	return "order_items"
}

// BeforeCreate -> Called before inserting record into Column Table
func (u *OrderItem) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	u.ID = BINARY16(id)
	return err
}
