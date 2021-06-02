package services

import (
	"travel/api/repository"
	"travel/models"
	"travel/utils"
)

//CategoryService -> struct
type CategoryService struct {
	repository repository.CategoryRepository
}

//NewCategoryService -> constructor
func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return CategoryService{
		repository: repository,
	}
}

// GetAllCategory -> returns all categories
func (c CategoryService) GetAllCategories(searchParams models.CategorySearchParams, pagination utils.Pagination) ([]models.Category, int64, error) {
	return c.repository.GetAllCategories(searchParams, pagination)
}

//CreateCategory -> creates new Category
func (c CategoryService) CreateCategory(category models.Category) error {
	return c.repository.CreateCategory(category)
}

//GetCategoryByID -> gets a category by ir
func (c CategoryService) GetCategoryByID(ID int) (models.Category, error) {
	return c.repository.GetCategoryByID(ID)
}

//UpdateCategory -> updates the category data
func (c CategoryService) UpdateCategory(category models.Category) error {
	return c.repository.UpdateCategory(category)
}

//DeleteCategory -> deletes the category
func (c CategoryService) DeleteCategory(category models.Category) error {
	return c.repository.DeleteCategory(category)
}
