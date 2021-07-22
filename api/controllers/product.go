package controllers

import (
	"net/http"
	"strconv"
	"strings"
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

// AddProduct -> creates a new product
func (p ProductController) AddProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		p.logger.Zap.Error("Product params parse error in controller", err)
		responses.ErrorJSON(c, http.StatusBadGateway, "Failed to parse product params")
		return
	}
	// Create slug
	slug_string := strings.Trim(strings.ReplaceAll(strings.ToLower(product.Name), " ", "-"), " ")
	product.Slug = slug_string
	if err := p.productService.AddProduct(product); err != nil {
		p.logger.Zap.Error("Failed to save product", err)
		responses.ErrorJSON(c, http.StatusBadGateway, "Failed to save product")
		return
	}
	responses.JSON(c, http.StatusCreated, "Product Added Successfully")
}

// GetProductById -> gets a product by ID
func (p ProductController) GetProductByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		p.logger.Zap.Error("Error retriving id param:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to retrieve id param")
		return
	}
	product, err := p.productService.GetProductByID(id)
	if err != nil {
		p.logger.Zap.Error("Can not find product:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Can not find product")
		return
	}
	responses.JSON(c, http.StatusOK, product)
}

// UpdateProduct -> updates the existing product by ID
func (p ProductController) UpdateProduct(c *gin.Context) {
	var newProduct models.Product
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		p.logger.Zap.Error("Error retriving id param:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to retrieve id param")
		return
	}
	product, err := p.productService.GetProductByID(id)
	if err != nil {
		p.logger.Zap.Error("Error [UpdateProduct] [GetOneProduct] :: ", err.Error())
		responses.ErrorJSON(c, http.StatusInternalServerError, "failed to find Product")
		return
	}
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		p.logger.Zap.Error("Product params parse error in controller", err)
		responses.ErrorJSON(c, http.StatusBadGateway, "Failed to parse product params")
		return
	}
	newProduct.ID = product.ID
	if err := p.productService.UpdateProduct(newProduct); err != nil {
		p.logger.Zap.Error("Fail to update product", err)
		responses.ErrorJSON(c, http.StatusBadGateway, "Failed to update product")
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Product Updated successfully...")
}

// DeleteProduct -> soft delets a product from database
func (p ProductController) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		p.logger.Zap.Error("Error retriving id param:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to retrieve id param")
		return
	}
	if err := p.productService.DeleteProduct(id); err != nil {
		p.logger.Zap.Error("Fail to delete the product", err)
		responses.ErrorJSON(c, http.StatusBadGateway, "Failed to delete the product")
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Product Deleted successfully...")
}