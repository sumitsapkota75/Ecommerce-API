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

// BrandController -> data type
type BrandController struct {
	logger       infrastructure.Logger
	brandService services.BrandService
}

// NewBrandController -> creates new user controller
func NewBrandController(logger infrastructure.Logger, brandService services.BrandService) BrandController {
	return BrandController{
		logger:       logger,
		brandService: brandService,
	}
}

// GetAllBrands
func (cc BrandController) GetAllBrands(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	searchParams := models.BrandSearchParams{
		Keyword: c.Query("keyword"),
	}
	brands, count, err := cc.brandService.GetAllBrands(searchParams, pagination)
	if err != nil {
		cc.logger.Zap.Error("Failed to get brands::", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to get brands")
		return
	}
	responses.JSONCount(c, http.StatusOK, brands, int(count))
}

// CreateBrand -> creates a new category
func (cc BrandController) CreateBrand(c *gin.Context) {
	// uid := c.MustGet("uid")
	var brand models.Brand

	if err := c.ShouldBindJSON(&brand); err != nil {
		cc.logger.Zap.Error("Brand params parse error in Controller:", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to parse brand params")
		return
	}
	if err := cc.brandService.CreateBrand(brand); err != nil {
		cc.logger.Zap.Error("Failed to save brand:", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to save brand")
		return
	}
	responses.JSON(c, http.StatusCreated, "Brand successfully created")
}

// GetBrandByID -> gets a new category by ID
func (cc BrandController) GetBrandByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		cc.logger.Zap.Error("Error retriving id param:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to retrieve id param")
		return
	}
	category, err := cc.brandService.GetBrandByID(id)
	if err != nil {
		cc.logger.Zap.Error("Can not find brand:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Can not find brand")
		return
	}
	responses.JSON(c, http.StatusOK, category)
}

// UpdateBrand -> updates the new category
func (cc BrandController) UpdateBrand(c *gin.Context) {
	// uid := c.MustGet("uid")
	var newBrand models.Brand
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		cc.logger.Zap.Error("Error retriving id param:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to retrieve id param")
		return
	}
	if err := c.ShouldBindJSON(&newBrand); err != nil {
		cc.logger.Zap.Error("Brand params parse error in Controller:", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to parse brand params")
		return
	}
	brand, err := cc.brandService.GetBrandByID(id)
	if err != nil {
		cc.logger.Zap.Error("Failed to get brand:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to get brand:")
		return
	}
	newBrand.ID = brand.ID
	if err := cc.brandService.UpdateBrand(newBrand); err != nil {
		cc.logger.Zap.Error("Failed to update brand:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to update brand:")
		return
	}
	responses.JSON(c, http.StatusCreated, "Brand Updated successfully")
}

// DeleteBrand -> deletes the category
func (cc BrandController) DeleteBrand(c *gin.Context) {
	var brand models.Brand
	idParam := c.Param("id")
	id, err := models.StringToBinary16(idParam)
	if err != nil {
		cc.logger.Zap.Error("Error retriving id param:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to retrieve id param")
		return
	}
	brand.ID = id
	if err := cc.brandService.DeleteBrand(brand); err != nil {
		cc.logger.Zap.Error("Failed to delete brand:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to delete brand:")
		return
	}
	responses.JSON(c, http.StatusOK, "Brand deleted successfully")
}
