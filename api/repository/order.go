package repository

import (
	"travel/infrastructure"
	"travel/models"
	"travel/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//OrderRepository -> struct
type OrderRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewOrderRepository -> creates a new order repository
func NewOrderRepository(db infrastructure.Database, logger infrastructure.Logger) OrderRepository {
	return OrderRepository{
		db:     db,
		logger: logger,
	}
}

// WithTrx -> enables repository with transaction
func (c OrderRepository) WithTrx(trxHandle *gorm.DB) OrderRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}

// GetAllOrders -> lists all orders
func (c OrderRepository) GetAllOrders(searchParams models.OrderSearchParams, pagination utils.Pagination) ([]models.Order, int64, error) {
	var orders []models.Order
	var count int64

	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset)
	if pagination.All {
		queryBuilder = c.db.DB
	}
	if searchParams.Keyword != "" {
		query := "%" + searchParams.Keyword + "%"
		queryBuilder = queryBuilder.Where(
			c.db.DB.Where("name LIKE ? ", query))
	}
	err := queryBuilder.Model(&models.Order{}).
		Preload("OrderItem.Product").
		Preload(clause.Associations).
		Order("updated_at desc").
		Where(&orders).
		Find(&orders).
		Offset(-1).
		Limit(-1).
		Count(&count).Error
	return orders, count, err
}

// CreateOrder -> creates a new order
func (o OrderRepository) CreateOrder(order models.Order) (models.Order, error) {
	return order, o.db.DB.Omit("OrderItem").Create(&order).Error
}

// CreateOrderItem -> create order items
func (o OrderRepository) CreateOrderItem(orderItem models.OrderItem) error {
	return o.db.DB.Create(&orderItem).Error
}

// GetOrderByID -> returns a single order
func (o OrderRepository) GetOrderByID(order models.Order) (models.Order, error) {
	return order, o.db.DB.Model(&models.Order{}).Preload("OrderItem").Where("id = ?", order.ID).First(&order).Error
}
