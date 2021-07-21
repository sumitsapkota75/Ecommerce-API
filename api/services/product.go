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

// AddProduct -> add product service
func (p ProductService) AddProduct(product models.Product) error {
	return p.repository.AddProduct(product)
}

// GetProductById -> returns a single id
func (p ProductService) GetProductByID(productId int) (models.Product, error) {
	return p.repository.GetProductByID(productId)
}

// UpdateProduct -> updates the product
func (p ProductService) UpdateProduct(product models.Product) error {
	return p.repository.UpdateProduct(product)
}

// DeleteProduct -> deletes the product
func (p ProductService) DeleteProduct(ProductID int) error {
	return p.repository.DeleteProduct(ProductID)
}
