package repository

import (
	"travel/infrastructure"
	"travel/models"

	"gorm.io/gorm"
)

// ShippingAddressRepository -> struct
type ShippingAddressRepository struct {
	logger infrastructure.Logger
	db     infrastructure.Database
}

//NewShippingAddressRepository -> constructor
func NewShippingAddressRepository(logger infrastructure.Logger, db infrastructure.Database) ShippingAddressRepository {
	return ShippingAddressRepository{
		logger: logger,
		db:     db,
	}
}

// WithTrx -> enables repository with transaction
func (b ShippingAddressRepository) WithTrx(trxHandle *gorm.DB) ShippingAddressRepository {
	if trxHandle == nil {
		b.logger.Zap.Error("Transaction Database not found in gin context. ")
		return b
	}
	b.db.DB = trxHandle
	return b
}

// GetShippingAddresses of customer
func (s ShippingAddressRepository) GetShippingAddresses(shippingAddress models.ShippingAddress) (addresses []models.ShippingAddress, err error) {
	return addresses, s.db.DB.Model(&models.ShippingAddress{}).Preload("User").Where("user_id = ?", shippingAddress.UserID).Find(&addresses).Error
}

// GetShippingAddressByID
func (s ShippingAddressRepository) GetShippingAddressByID(shippingAddress models.ShippingAddress) (addresses models.ShippingAddress, err error) {
	return addresses, s.db.DB.Model(&models.ShippingAddress{}).Where("id = ?", shippingAddress.ID).First(&addresses).Error
}

// AddShippingAddress
func (s ShippingAddressRepository) AddShippingAddress(shippingAddress models.ShippingAddress) error {
	return s.db.DB.Create(&shippingAddress).Error
}

// DeleteShippingAddress
func (s ShippingAddressRepository) DeleteShippingAddress(shippingAddress models.ShippingAddress) error {
	return s.db.DB.Where("id = ?", shippingAddress.ID).Delete(&models.ShippingAddress{}).Error
}

// UpdateShippingAddress -> updates the product
func (p ShippingAddressRepository) UpdateShippingAddress(shippingAddress models.ShippingAddress) error {
	return p.db.DB.Model(&models.ShippingAddress{}).Where("id = ?", shippingAddress.ID).
		Updates(map[string]interface{}{
			"user_id":  shippingAddress.UserID,
			"district": shippingAddress.District,
			"address":  shippingAddress.Address,
			"phone":    shippingAddress.Phone,
		}).Error
}
