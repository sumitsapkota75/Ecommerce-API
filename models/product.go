package models

//ProductSearchParams -> search Product Params
type ProductSearchParams struct {
	Keyword string `json:"keyword"`
}

//Product -> model
type Product struct {
	UintBase
	Name               string   `json:"name"`
	CategoryID         uint     `json:"category_id"`
	Category           Category `json:"category"`
	BrandID            uint     `json:"brand_id"`
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
}

// TableName  -> returns table name of model
func (p Product) TableName() string {
	return "products"
}
