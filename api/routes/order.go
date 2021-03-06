package routes

import (
	"travel/api/controllers"
	"travel/api/middlewares"
	"travel/infrastructure"
)

type OrderRoutes struct {
	logger          infrastructure.Logger
	handler         infrastructure.RequestHandler
	orderController controllers.OrderController

	authMiddleware middlewares.AuthMiddleware
	trxMiddleware  middlewares.DBTransactionMiddleware
}

// Setup category Routes
func (u OrderRoutes) Setup() {
	u.logger.Zap.Info(" Setting up order routes 👤 -------------")
	orders := u.handler.Gin.Group("/order").Use(u.authMiddleware.Handle())
	{
		orders.GET("", u.orderController.GetAllOrders)
		orders.POST("", u.trxMiddleware.Handle(), u.orderController.CreateOrder)
		orders.GET("/:id", u.orderController.GetOrderByID)
		orders.GET("/my-order", u.orderController.GetAllOrderByCustomer)

	}
}

// NewOrderRoutes creates new category controller
func NewOrderRoutes(
	logger infrastructure.Logger,
	handler infrastructure.RequestHandler,
	orderController controllers.OrderController,
	authMiddleware middlewares.AuthMiddleware,
	trxMiddleware middlewares.DBTransactionMiddleware,
) OrderRoutes {
	return OrderRoutes{
		handler:         handler,
		logger:          logger,
		orderController: orderController,
		authMiddleware:  authMiddleware,
		trxMiddleware:   trxMiddleware,
	}
}
