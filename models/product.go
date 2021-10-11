package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//ProductSearchParams -> search Product Params
type ProductSearchParams struct {
	Keyword string `json:"keyword"`
}

//Product -> model
type Product struct {
	Base
	Vendor
	Name               string   `json:"name"`
	CategoryID         BINARY16 `json:"category_id"`
	Category           Category `json:"category"`
	BrandID            BINARY16 `json:"brand_id"`
	Brand              Brand    `json:"brand"`
	CostPrice          string   `json:"cost_price"`
	Price              string   `json:"price"`
	Slug               string   `json:"slug"`
	Code               string   `json:"code"`
	Quantity           uint     `json:"quantity"`
	Description        string   `json:"description"`
	Specification      string   `json:"specification"`
	TopSelling         int      `json:"top_selling"`
	NewArrival         int      `json:"new_arrival"`
	DailyDeal          int      `json:"daily_deal"`
	OrderLimit         int      `json:"order_limit"`
	StockAlert         int      `json:"stock_alert"`
	Refundable         int      `json:"refundable"`
	IsActive           int      `json:"is_active"`
	FeaturedCollection int      `json:"featured_collection"`
	Thumbnail          string   `json:"thumbnail"`
	OnSale             int      `json:"on_sale"`
	SalePrice          string   `json:"sale_price"`
	SaleTime           string   `json:"sale_time"`
}

// TableName  -> returns table name of model
func (p Product) TableName() string {
	return "products"
}

// BeforeCreate -> Called before inserting record into Column Table
func (u *Product) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	u.ID = BINARY16(id)
	return err
}
