package controllers

import (
	"net/http"
	"strconv"
	"travel/api/responses"
	"travel/api/services"
	"travel/infrastructure"
	"travel/models"
	"travel/utils"

	"github.com/gin-gonic/gin"
)

// CategoryController -> data type
type CategoryController struct {
	logger          infrastructure.Logger
	categoryService services.CategoryService
}

// NewCategoryController -> creates new user controller
func NewCategoryController(logger infrastructure.Logger, categoryService services.CategoryService, firebaseService services.FirebaseService) CategoryController {
	return CategoryController{
		logger:          logger,
		categoryService: categoryService,
	}
}

// GetAllCategories
func (cc CategoryController) GetAllCategories(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	searchParams := models.CategorySearchParams{
		Keyword: c.Query("keyword"),
	}
	categories, count, err := cc.categoryService.GetAllCategories(searchParams, pagination)
	if err != nil {
		cc.logger.Zap.Error("Failed to get categories::", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to get categories")
		return
	}
	responses.JSONCount(c, http.StatusOK, categories, int(count))
}

// CreateCategory -> creates a new category
func (cc CategoryController) CreateCategory(c *gin.Context) {
	// uid := c.MustGet("uid")
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		cc.logger.Zap.Error("Category params parse error in Controller:", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to parse category params")
		return
	}
	if err := cc.categoryService.CreateCategory(category); err != nil {
		cc.logger.Zap.Error("Failed to save category:", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to save category")
		return
	}
	responses.JSON(c, http.StatusCreated, "Category successfully created")
}

// GetCategoryByID -> gets a new category by ID
func (cc CategoryController) GetCategoryByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		cc.logger.Zap.Error("Error retriving id param:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to retrieve id param")
		return
	}
	category, err := cc.categoryService.GetCategoryByID(id)
	if err != nil {
		cc.logger.Zap.Error("Can not find category:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Can not find category")
		return
	}
	responses.JSON(c, http.StatusOK, category)
}

// UpdateCategory -> updates the new category
func (cc CategoryController) UpdateCategory(c *gin.Context) {
	// uid := c.MustGet("uid")
	var newCategory models.Category
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		cc.logger.Zap.Error("Error retriving id param:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to retrieve id param")
		return
	}
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		cc.logger.Zap.Error("Category params parse error in Controller:", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to parse category params")
		return
	}
	category, err := cc.categoryService.GetCategoryByID(id)
	if err != nil {
		cc.logger.Zap.Error("Failed to get category:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to get category:")
		return
	}
	newCategory.ID = category.ID
	if err := cc.categoryService.UpdateCategory(newCategory); err != nil {
		cc.logger.Zap.Error("Failed to update category:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to update category:")
		return
	}
	responses.JSON(c, http.StatusCreated, "Category Updated successfully")
}

// DeleteCategory -> deletes the category
func (cc CategoryController) DeleteCategory(c *gin.Context) {
	var category models.Category
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		cc.logger.Zap.Error("Error retriving id param:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to retrieve id param")
		return
	}
	category.ID = uint(id)
	if err := cc.categoryService.DeleteCategory(category); err != nil {
		cc.logger.Zap.Error("Failed to delete category:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to delete category:")
		return
	}
	responses.JSON(c, http.StatusOK, "Category deleted successfully")
}
