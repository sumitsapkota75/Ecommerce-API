package repository

import (
	"travel/infrastructure"
	"travel/models"
	"travel/utils"

	"gorm.io/gorm"
)

// Product Repository -> struct
type ProductRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewProductRepository -> creates a new user repository
func NewProductRepository(db infrastructure.Database, logger infrastructure.Logger) ProductRepository {
	return ProductRepository{
		db:     db,
		logger: logger,
	}
}

// WithTrx -> enables repository with transaction
func (p ProductRepository) WithTrx(trxHandle *gorm.DB) ProductRepository {
	if trxHandle == nil {
		p.logger.Zap.Error("Transaction Database not found in gin context. ")
		return p
	}
	p.db.DB = trxHandle
	return p
}

// GetAllProducts -> returns list of all products
func (p ProductRepository) GetAllProducts(searchParams models.ProductSearchParams, pagination utils.Pagination) ([]models.Product, int64, error) {
	var products []models.Product
	var count int64
	queryBuilder := p.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset)
	if pagination.All {
		queryBuilder = p.db.DB
	}
	if searchParams.Keyword != "" {
		query := "%" + searchParams.Keyword + "%"
		queryBuilder = queryBuilder.Where(
			p.db.DB.Where("name LIKE ? ", query))

	}
	err := queryBuilder.Model(&models.Product{}).
		Preload("Category").
		Preload("Brand").
		Order("updated_at desc").
		Where(&products).
		Find(&products).
		Offset(-1).
		Limit(-1).
		Count(&count).Error
	return products, count, err
}

// AddProduct -> creates a new product
func (p ProductRepository) AddProduct(product models.Product) error {
	return p.db.DB.Create(&product).Error
}

// GetProductByID -> gets a product by ID
func (p ProductRepository) GetProductByID(productID int) (product models.Product, err error) {
	return product, p.db.DB.Model(&models.Product{}).Preload("Category").Preload("Brand").Where("id = ?", productID).First(&product).Error
}

// UpdateProduct -> updates the product
func (p ProductRepository) UpdateProduct(product models.Product) error {
	return p.db.DB.Model(&models.Product{}).Where("id=?", product.ID).
		Updates(map[string]interface{}{
			"name":                product.Name,
			"category_id":         product.CategoryID,
			"brand_id":            product.BrandID,
			"cost_price":          product.CostPrice,
			"price":               product.Price,
			"slug":                product.Slug,
			"code":                product.Code,
			"quantity":            product.Quantity,
			"description":         product.Description,
			"specification":       product.Specification,
			"top_selling":         product.TopSelling,
			"new_arrival":         product.NewArrival,
			"daily_deal":          product.DailyDeal,
			"order_limit":         product.OrderLimit,
			"stock_alert":         product.StockAlert,
			"refundable":          product.Refundable,
			"featured_collection": product.FeaturedCollection,
			"thumbnail":           product.Thumbnail,
			"is_active":           product.IsActive,
		}).Error
}

// DeleteProduct deletes the given user
func (p ProductRepository) DeleteProduct(ProductID int) error {
	return p.db.DB.Where("id = ?", ProductID).Delete(&models.Product{}).Error
}
