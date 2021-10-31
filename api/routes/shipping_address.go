package routes

import (
	"travel/api/controllers"
	"travel/api/middlewares"
	"travel/infrastructure"
)

type ShippingAddressRoute struct {
	logger             infrastructure.Logger
	handler            infrastructure.RequestHandler
	shippingController controllers.ShippingAddressController

	middleware middlewares.AuthMiddleware
}

// Setup Product Routes
func (u ShippingAddressRoute) Setup() {
	u.logger.Zap.Info(" Setting up Shipping Address Route routes 👤 -------------")
	product := u.handler.Gin.Group("/shipping-address")
	{
		product.GET("", u.middleware.Handle(), u.shippingController.GetAllShippingAddress)
		product.GET("/:id", u.middleware.Handle(), u.shippingController.GetShippingAddressByID)
		product.POST("", u.middleware.Handle(), u.shippingController.CreateShippingAddress)
		product.POST("/:id", u.middleware.Handle(), u.shippingController.UpdateShippingAddress)
		product.DELETE("/:id", u.middleware.Handle(), u.shippingController.DeleteShippingAddress)
	}
}

func NewShippingAddressRoute(
	logger infrastructure.Logger,
	handler infrastructure.RequestHandler,
	shippingController controllers.ShippingAddressController,
	middleware middlewares.AuthMiddleware,
) ShippingAddressRoute {
	return ShippingAddressRoute{
		handler:            handler,
		logger:             logger,
		shippingController: shippingController,
		middleware:         middleware,
	}
}
