package services

import (
	"travel/api/repository"
	"travel/models"
)

//ShippingAddressService -> struct
type ShippingAddressService struct {
	repository repository.ShippingAddressRepository
}

//NewShippingAddressService -> constructor
func NewShippingAddressService(repository repository.ShippingAddressRepository) ShippingAddressService {
	return ShippingAddressService{
		repository: repository,
	}
}

// GetAllShippingAddresses -> returns all categories
func (c ShippingAddressService) GetAllShippingAddresses(shippingAddress models.ShippingAddress) ([]models.ShippingAddress, error) {
	return c.repository.GetShippingAddresses(shippingAddress)
}

// GetShippingAddressByID ->
func (c ShippingAddressService) GetShippingAddressByID(shippingAddress models.ShippingAddress) (models.ShippingAddress, error) {
	return c.repository.GetShippingAddressByID(shippingAddress)
}

//AddShippingAddress -> creates new shippingAddress
func (c ShippingAddressService) AddShippingAddress(shippingAddress models.ShippingAddress) error {
	return c.repository.AddShippingAddress(shippingAddress)
}

//UpdateShippingAddress -> updates the ShippingAddress data
func (c ShippingAddressService) UpdateShippingAddress(shippingAddress models.ShippingAddress) error {
	return c.repository.UpdateShippingAddress(shippingAddress)
}

//DeleteShippingAddress -> deletes the ShippingAddress
func (c ShippingAddressService) DeleteShippingAddress(shippingAddress models.ShippingAddress) error {
	return c.repository.DeleteShippingAddress(shippingAddress)
}
