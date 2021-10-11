package repository

import (
	"travel/infrastructure"
	"travel/models"
	"travel/utils"

	"gorm.io/gorm"
)

// VendorRepository database structure
type VendorRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewVendorRepository -> creates a new Vendor repository
func NewVendorRepository(db infrastructure.Database, logger infrastructure.Logger) VendorRepository {
	return VendorRepository{
		db:     db,
		logger: logger,
	}
}

// WithTrx -> enables repository with transaction
func (r VendorRepository) WithTrx(trxHandle *gorm.DB) VendorRepository {
	if trxHandle == nil {
		r.logger.Zap.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.db.DB = trxHandle
	return r
}

//GetAllVendors -> returns list of vendor
func (u VendorRepository) GetAllVendors(pagination utils.Pagination) ([]models.Vendor, int64, error) {
	var vendors []models.Vendor
	var count int64
	querybuilder := u.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset)
	if pagination.All {
		querybuilder = u.db.DB
	}
	err := querybuilder.Model(&models.Vendor{}).
		Order("created_at asc").
		Where(&vendors).
		Find(&vendors).
		Offset(-1).
		Limit(-1).
		Count(&count).Error
	return vendors, count, err
}

// CreateVendor -> creates new Vendor
func (u VendorRepository) CreateVendor(vendor models.Vendor) (models.Vendor, error) {
	return vendor, u.db.DB.Create(&vendor).Error
}

//GetVendorByID -> gets the Vendor by uid
func (u VendorRepository) GetVendorByID(VendorID string) (vendor models.Vendor, err error) {
	return vendor, u.db.DB.Where("id = ?", VendorID).First(&vendor).Error
}

//UpdateVendor -> updates the vendor data
func (u VendorRepository) UpdateVendor(vendor models.Vendor) error {
	return u.db.DB.Model(&models.Vendor{}).
		Where("id = ?", vendor.ID).
		Updates(map[string]interface{}{
			"name":          vendor.Name,
			"email":         vendor.Email,
			"address":       vendor.Address,
			"phone":         vendor.Phone,
			"store_name":    vendor.StoreName,
			"document_type": vendor.DocumentType,
			"document_id":   vendor.DocumentID,
			"thumbnail":     vendor.Thumbnail,
		}).Error
}
