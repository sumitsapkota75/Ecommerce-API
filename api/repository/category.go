package repository

import (
	"travel/infrastructure"
	"travel/models"
	"travel/utils"

	"gorm.io/gorm"
)

//CategoryRepository -> struct
type CategoryRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewUserRepository -> creates a new user repository
func NewCategoryRepository(db infrastructure.Database, logger infrastructure.Logger) CategoryRepository {
	return CategoryRepository{
		db:     db,
		logger: logger,
	}
}

// WithTrx -> enables repository with transaction
func (c CategoryRepository) WithTrx(trxHandle *gorm.DB) CategoryRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}

// GetAllCategories -> lists all category
func (c CategoryRepository) GetAllCategories(searchParams models.CategorySearchParams, pagination utils.Pagination) ([]models.Category, int64, error) {
	var categories []models.Category
	var count int64

	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset)
	if pagination.All {
		queryBuilder = c.db.DB
	}
	if searchParams.Keyword != "" {
		query := "%" + searchParams.Keyword + "%"
		queryBuilder = queryBuilder.Where(
			c.db.DB.Where("name LIKE ? ", query))
	}
	err := queryBuilder.Model(&models.Category{}).
		Order("updated_at desc").
		Where(&categories).
		Find(&categories).
		Offset(-1).
		Limit(-1).
		Count(&count).Error
	return categories, count, err
}

// CreateCategory -> adds new category
func (c CategoryRepository) CreateCategory(category models.Category) error {
	return c.db.DB.Create(&category).Error
}

//GetCategoryByID -> gets a category by ID
func (c CategoryRepository) GetCategoryByID(ID int) (category models.Category, err error) {
	return category, c.db.DB.Model(&models.Category{}).Where("id = ?", ID).First(&category).Error
}

// UpdateCategory -> updates category detail
func (c CategoryRepository) UpdateCategory(category models.Category) error {
	return c.db.DB.Model(&models.Category{}).
		Where("id = ?", category.ID).
		Updates(map[string]interface{}{
			"name":        category.Name,
			"description": category.Description,
			"thumbnail":   category.Thumbnail,
		}).Error
}
