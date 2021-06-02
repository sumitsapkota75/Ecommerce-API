package repository

import (
	"travel/infrastructure"
	"travel/models"
	"travel/utils"

	"gorm.io/gorm"
)

// BrandRepository -> struct
type BrandRepository struct {
	logger infrastructure.Logger
	db     infrastructure.Database
}

//NewBrandRepository -> constructor
func NewBrandRepository(logger infrastructure.Logger, db infrastructure.Database) BrandRepository {
	return BrandRepository{
		logger: logger,
		db:     db,
	}
}

// WithTrx -> enables repository with transaction
func (b BrandRepository) WithTrx(trxHandle *gorm.DB) BrandRepository {
	if trxHandle == nil {
		b.logger.Zap.Error("Transaction Database not found in gin context. ")
		return b
	}
	b.db.DB = trxHandle
	return b
}

// / GetAllBrands -> lists all brands
func (c BrandRepository) GetAllBrands(searchParams models.BrandSearchParams, pagination utils.Pagination) ([]models.Brand, int64, error) {
	var brands []models.Brand
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
	err := queryBuilder.Model(&models.Brand{}).
		Order("updated_at desc").
		Where(&brands).
		Find(&brands).
		Offset(-1).
		Limit(-1).
		Count(&count).Error
	return brands, count, err
}

// / CreateBrand -> adds new Brand
func (c BrandRepository) CreateBrand(brand models.Brand) error {
	return c.db.DB.Create(&brand).Error
}

//GetBrandByID -> gets a Brand by ID
func (c BrandRepository) GetBrandByID(ID int) (brand models.Brand, err error) {
	return brand, c.db.DB.Model(&models.Brand{}).Where("id = ?", ID).First(&brand).Error
}

// UpdateBrand -> updates Brand detail
func (c BrandRepository) UpdateBrand(brand models.Brand) error {
	return c.db.DB.Model(&models.Brand{}).
		Where("id = ?", brand.ID).
		Updates(map[string]interface{}{
			"name":        brand.Name,
			"description": brand.Description,
			"thumbnail":   brand.Thumbnail,
		}).Error
}

// DeleteBrand -> deletes Brand
func (c BrandRepository) DeleteBrand(brand models.Brand) error {
	return c.db.DB.Delete(&models.Brand{}, brand.ID).Error
}
