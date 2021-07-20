package services

import (
	"travel/api/repository"
	"travel/models"
	"travel/utils"
)

//ProductService -> struct
type ProductService struct {
	repository repository.ProductRepository
}

//NewProductService -> constructor
func NewProductService(repository repository.ProductRepository) ProductService {
	return ProductService{
		repository: repository,
	}
}

// GetAllCategory -> returns all ProductGetAllProducts
func (c ProductService) GetAllProducts(searchParams models.ProductSearchParams, pagination utils.Pagination) ([]models.Product, int64, error) {
	return c.repository.GetAllProducts(searchParams, pagination)
}
