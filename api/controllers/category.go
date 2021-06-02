package controllers

import (
	"net/http"
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
