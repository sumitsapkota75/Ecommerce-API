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
		Order("updated_at desc").
		Where(&products).
		Find(&products).
		Offset(-1).
		Limit(-1).
		Count(&count).Error
	return products, count, err
}
