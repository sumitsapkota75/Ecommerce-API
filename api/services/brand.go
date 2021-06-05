package services

import (
	"travel/api/repository"
	"travel/models"
	"travel/utils"
)

//BrandService -> struct
type BrandService struct {
	repository repository.BrandRepository
}

//NewBrandService -> constructor
func NewBrandService(repository repository.BrandRepository) BrandService {
	return BrandService{
		repository: repository,
	}
}

// GetAllCategory -> returns all BrandGetAllBrands
func (c BrandService) GetAllBrands(searchParams models.BrandSearchParams, pagination utils.Pagination) ([]models.Brand, int64, error) {
	return c.repository.GetAllBrands(searchParams, pagination)
}

//CreateBrand -> creates new Brand
func (c BrandService) CreateBrand(brand models.Brand) error {
	return c.repository.CreateBrand(brand)
}

//GetBrandByID -> gets a Brand by ir
func (c BrandService) GetBrandByID(ID int) (models.Brand, error) {
	return c.repository.GetBrandByID(ID)
}

//UpdateBrand -> updates the Brand data
func (c BrandService) UpdateBrand(brand models.Brand) error {
	return c.repository.UpdateBrand(brand)
}

//DeleteBrand -> deletes the Brand
func (c BrandService) DeleteBrand(brand models.Brand) error {
	return c.repository.DeleteBrand(brand)
}
