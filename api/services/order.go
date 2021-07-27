package services

import (
	"travel/api/repository"
	"travel/models"
	"travel/utils"

	"gorm.io/gorm"
)

//OrderService -> struct
type OrderService struct {
	repository repository.OrderRepository
}

//NewOrderService -> constructor
func NewOrderService(repository repository.OrderRepository) OrderService {
	return OrderService{
		repository: repository,
	}
}

// WithTrx -> enables repository with transaction
func (c OrderService) WithTrx(trxHandle *gorm.DB) OrderService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

// GetAllOrder -> returns all orders
func (c OrderService) GetAllOrders(searchParams models.OrderSearchParams, pagination utils.Pagination) ([]models.Order, int64, error) {
	return c.repository.GetAllOrders(searchParams, pagination)
}

// CreateOrder -> creates a new order
func (o OrderService) CreateOrder(order models.Order) (models.Order, error) {
	return o.repository.CreateOrder(order)
}

// CreateOrderItem -> creates each order item
func (o OrderService) CreateOrderItem(orderItem models.OrderItem) error {
	return o.repository.CreateOrderItem(orderItem)
}
