package routes

import (
	"travel/api/controllers"
	"travel/api/middlewares"
	"travel/infrastructure"
)

type VendorRoutes struct {
	logger           infrastructure.Logger
	handler          infrastructure.RequestHandler
	vendorController controllers.VendorController

	middleware middlewares.AuthMiddleware
}

// Setup user Routes
func (u VendorRoutes) Setup() {
	u.logger.Zap.Info(" Setting up vendor routes 👤 -------------")
	user := u.handler.Gin.Group("/vendor")
	{
		user.GET("", u.vendorController.GetAllVendors)
		user.POST("/signup", u.vendorController.CreateVendor)
		user.GET("/profile", u.middleware.Handle(), u.vendorController.GetVendorProfile)
		user.POST("/update", u.middleware.Handle(), u.vendorController.UpdateVendorProfile)
	}
}

// NewVendorRoutes creates new user controller
func NewVendorRoutes(
	logger infrastructure.Logger,
	handler infrastructure.RequestHandler,
	vendorController controllers.VendorController,
	middleware middlewares.AuthMiddleware,
) VendorRoutes {
	return VendorRoutes{
		handler:          handler,
		logger:           logger,
		vendorController: vendorController,
		middleware:       middleware,
	}
}
