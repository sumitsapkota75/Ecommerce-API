package controllers

import (
	"net/http"
	"travel/api/responses"
	"travel/api/services"
	"travel/constants"
	"travel/infrastructure"
	"travel/models"
	"travel/utils"

	"github.com/gin-gonic/gin"
)

// VendorController -> data type
type VendorController struct {
	logger          infrastructure.Logger
	vendorService   services.VendorService
	firebaseService services.FirebaseService
}

// NewVendorController -> creates new Vendor controller
func NewVendorController(logger infrastructure.Logger, vendorService services.VendorService, firebaseService services.FirebaseService) VendorController {
	return VendorController{
		logger:          logger,
		vendorService:   vendorService,
		firebaseService: firebaseService,
	}
}

//GetAllVendors -> get the list of Vendors
func (u VendorController) GetAllVendors(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	vendors, count, err := u.vendorService.GetAllVendors(pagination)
	if err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to get Vendors")
		return
	}

	responses.JSONCount(c, http.StatusOK, vendors, int(count))
}

// CreateVendor -> create new vendor in database
func (v VendorController) CreateVendor(c *gin.Context) {
	requestVendor := struct {
		models.Vendor
		Password string `json:"password"`
	}{}
	if err := c.ShouldBindJSON(&requestVendor); err != nil {
		v.logger.Zap.Error("Error [create vendor ] ::")
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to bind vendors json data")
		return
	}
	// Checks for firebase existing user
	firebaseID := v.firebaseService.GetUserByEmail(requestVendor.Email)
	if firebaseID != "" {
		v.logger.Zap.Error("Error [create firebase user] ::")
		responses.ErrorJSON(c, http.StatusBadRequest, "The provided email is already in use")
		return
	}
	firebaseID, firebaseErr := v.firebaseService.CreateUser(requestVendor.Email, requestVendor.Password, requestVendor.Name, constants.VendorUserType)
	if firebaseErr != nil {
		v.logger.Zap.Error("Error [create firebase user] ::", firebaseErr)
		responses.ErrorJSON(c, http.StatusBadRequest, "Error creating user in firebase")
		return
	}

	var vendor models.Vendor
	vendor.ID = firebaseID
	vendor.Address = requestVendor.Address
	vendor.Phone = requestVendor.Phone
	vendor.Email = requestVendor.Email
	vendor.Name = requestVendor.Name
	vendor.StoreName = requestVendor.StoreName
	vendor.DocumentID = requestVendor.DocumentID
	vendor.DocumentType = requestVendor.DocumentType
	vendorData, err := v.vendorService.CreateVendor(vendor)
	if err != nil {
		v.logger.Zap.Error("Error [create vendor ] ::")
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to save vendor")
		return
	}
	responses.SuccessJSON(c, http.StatusCreated, vendorData)
}

// GetVendorProfile -> gets a Vendor by ID
func (u VendorController) GetVendorProfile(c *gin.Context) {
	uid := c.MustGet(constants.UID).(string)
	vendor, err := u.vendorService.GetVendorByID(uid)
	if err != nil {
		u.logger.Zap.Error("Error getting Vendor ::", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "failed to get Vendor")
		return
	}
	responses.JSON(c, http.StatusOK, vendor)
}

// UpdateVendorProfile -> updates vendor data in database
func (u VendorController) UpdateVendorProfile(c *gin.Context) {
	idParam := c.Param("id")
	requestVendor := struct {
		models.Vendor
		Password string `json:"password"`
	}{}
	vendor, err := u.vendorService.GetVendorByID(idParam)
	if err != nil {
		u.logger.Zap.Error("Error getting vendor data ::", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "failed to get vendor data")
		return
	}
	vendor.Name = requestVendor.Name
	vendor.StoreName = requestVendor.StoreName
	vendor.DocumentType = requestVendor.DocumentType
	vendor.DocumentID = requestVendor.DocumentID
	vendor.Email = requestVendor.Email
	vendor.Phone = requestVendor.Phone
	vendor.Thumbnail = requestVendor.Thumbnail
	vendor.Address = requestVendor.Address

	if err := u.vendorService.UpdateVendor(vendor); err != nil {
		u.logger.Zap.Error("Failed to update the vendor ::", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "failed to update the vendor data")
		return
	}
	responses.JSON(c, http.StatusBadRequest, "Vendor updated successfully")

}
