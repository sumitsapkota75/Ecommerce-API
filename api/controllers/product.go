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

// ProductController -> data type
type ProductController struct {
	logger         infrastructure.Logger
	productService services.ProductService
}

// NewProductController -> creates new user controller
func NewProductController(logger infrastructure.Logger, productService services.ProductService, firebaseService services.FirebaseService) ProductController {
	return ProductController{
		logger:         logger,
		productService: productService,
	}
}

// GetAllProducts
func (cc ProductController) GetAllProducts(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	searchParams := models.ProductSearchParams{
		Keyword: c.Query("keyword"),
	}
	products, count, err := cc.productService.GetAllProducts(searchParams, pagination)
	if err != nil {
		cc.logger.Zap.Error("Failed to get products::", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to get products")
		return
	}
	responses.JSONCount(c, http.StatusOK, products, int(count))
}
