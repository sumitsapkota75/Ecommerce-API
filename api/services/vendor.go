package services

import (
	"travel/api/repository"
	"travel/models"
	"travel/utils"

	"gorm.io/gorm"
)

// VendorService struct
type VendorService struct {
	repository repository.VendorRepository
}

//NewVendorService -> creates a new Vendor service
func NewVendorService(repository repository.VendorRepository) VendorService {
	return VendorService{
		repository: repository,
	}
}

// WithTrx -> enables repository with transaction
func (s VendorService) WithTrx(trxHandle *gorm.DB) VendorService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// CreateVendor -> call to create a Vendor
func (s VendorService) CreateVendor(Vendor models.Vendor) (models.Vendor, error) {
	return s.repository.CreateVendor(Vendor)
}

// GetVendorByID -> returns a Vendor by UID
func (s VendorService) GetVendorByID(VendorID string) (models.Vendor, error) {
	return s.repository.GetVendorByID(VendorID)
}

// GetAllVendors -> returns a list of Vendor
func (s VendorService) GetAllVendors(pagination utils.Pagination) ([]models.Vendor, int64, error) {
	return s.repository.GetAllVendors(pagination)
}

// UpdateVendor -> updates the Vendor data
func (s VendorService) UpdateVendor(vendor models.Vendor) error {
	return s.repository.UpdateVendor(vendor)
}
