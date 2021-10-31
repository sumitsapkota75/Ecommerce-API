package controllers

import (
	"net/http"
	"travel/api/responses"
	"travel/api/services"
	"travel/constants"
	"travel/infrastructure"
	"travel/models"

	"github.com/gin-gonic/gin"
)

// ShippingAddressController -> data type
type ShippingAddressController struct {
	logger  infrastructure.Logger
	service services.ShippingAddressService
}

// NewShippingAddressController -> creates new shipping address controller
func NewShippingAddressController(logger infrastructure.Logger, service services.ShippingAddressService) ShippingAddressController {
	return ShippingAddressController{
		logger:  logger,
		service: service,
	}
}

//GetAllShippingAddress -> get the list of shipping address of a user
func (u ShippingAddressController) GetAllShippingAddress(c *gin.Context) {
	var shipping_address models.ShippingAddress
	uid := c.MustGet(constants.UID).(string)
	shipping_address.UserID = uid
	addresses, err := u.service.GetAllShippingAddresses(shipping_address)
	if err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to get Vendors")
		return
	}

	responses.SuccessJSON(c, http.StatusOK, addresses)
}

// CreateShippingAddress -> creates new shipping address for the user
func (u ShippingAddressController) CreateShippingAddress(c *gin.Context) {
	var shipping_adress models.ShippingAddress
	uid := c.MustGet(constants.UID).(string)
	if err := c.ShouldBindJSON(&shipping_adress); err != nil {
		u.logger.Zap.Error("Shipping address params parse error in controller", err)
		responses.ErrorJSON(c, http.StatusBadGateway, "Failed to parse shipping address params")
		return
	}
	shipping_adress.UserID = uid
	if err := u.service.AddShippingAddress(shipping_adress); err != nil {
		u.logger.Zap.Error("Failed to add shipping_address::", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to save address")
		return
	}
	responses.JSON(c, http.StatusCreated, "Shipping Address Added Successfully")

}

// UpdateShippingAddress updates shipping address information
func (u ShippingAddressController) UpdateShippingAddress(c *gin.Context) {
	var shipping_address models.ShippingAddress
	idParam := c.Param("id")
	id, err := models.StringToBinary16(idParam)
	if err != nil {
		u.logger.Zap.Error("Failed to convert id to binary 16 :: ", err.Error())
		responses.ErrorJSON(c, http.StatusInternalServerError, "Failed to convert id to binary 16 ")
		return
	}
	shipping_address.ID = id
	address, err := u.service.GetShippingAddressByID(shipping_address)
	if err != nil {
		u.logger.Zap.Error("Error [UpdateShippingAddress] [GetOneShippingAddress] :: ", err.Error())
		responses.ErrorJSON(c, http.StatusInternalServerError, "failed to find shipping address")
		return
	}
	if err := c.ShouldBindJSON(&shipping_address); err != nil {
		u.logger.Zap.Error("Brand params parse error in Controller:", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to parse brand params")
		return
	}
	shipping_address.ID = address.ID
	if err := u.service.UpdateShippingAddress(shipping_address); err != nil {
		u.logger.Zap.Error("Failed to UpdateShippingAddress:::", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "failed to get address data")
		return
	}
	responses.SuccessJSON(c, http.StatusOK, shipping_address)

}

// GetShippingAddressByID
func (u ShippingAddressController) GetShippingAddressByID(c *gin.Context) {
	idParam := c.Param("id")
	var shipping_address models.ShippingAddress
	id, err := models.StringToBinary16(idParam)
	if err != nil {
		u.logger.Zap.Error("Failed to convert id to binary 16 :: ", err.Error())
		responses.ErrorJSON(c, http.StatusInternalServerError, "Failed to convert id to binary 16 ")
		return
	}
	shipping_address.ID = id
	add, err := u.service.GetShippingAddressByID(shipping_address)
	if err != nil {
		u.logger.Zap.Error("Fail to delete the product", err)
		responses.ErrorJSON(c, http.StatusBadGateway, "Failed to delete the product")
		return
	}
	responses.SuccessJSON(c, http.StatusOK, add)

}

// DeleteShippingAddress -> deletes the shipping_address
func (u ShippingAddressController) DeleteShippingAddress(c *gin.Context) {
	idParam := c.Param("id")
	var shipping_address models.ShippingAddress
	id, err := models.StringToBinary16(idParam)
	if err != nil {
		u.logger.Zap.Error("Failed to convert id to binary 16 :: ", err.Error())
		responses.ErrorJSON(c, http.StatusInternalServerError, "Failed to convert id to binary 16 ")
		return
	}
	shipping_address.ID = id
	if err := u.service.DeleteShippingAddress(shipping_address); err != nil {
		u.logger.Zap.Error("Fail to delete the product", err)
		responses.ErrorJSON(c, http.StatusBadGateway, "Failed to delete the product")
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Product Deleted successfully...")
}
