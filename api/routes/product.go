package routes

import (
	"travel/api/controllers"
	"travel/api/middlewares"
	"travel/infrastructure"
)

type ProductRoutes struct {
	logger            infrastructure.Logger
	handler           infrastructure.RequestHandler
	productController controllers.ProductController

	middleware middlewares.AuthMiddleware
}

// Setup Product Routes
func (u ProductRoutes) Setup() {
	u.logger.Zap.Info(" Setting up Product routes 👤 -------------")
	product := u.handler.Gin.Group("/product")
	{
		product.GET("", u.productController.GetAllProducts)
		product.GET("/:id", u.productController.GetProductByID)
		product.POST("", u.productController.AddProduct)
		product.POST("/:id", u.productController.UpdateProduct)
		product.DELETE("/:id", u.productController.DeleteProduct)
	}
}

// NewProductRoute creates new product controller
func NewProductRoute(
	logger infrastructure.Logger,
	handler infrastructure.RequestHandler,
	productController controllers.ProductController,
	middleware middlewares.AuthMiddleware,
) ProductRoutes {
	return ProductRoutes{
		handler:           handler,
		logger:            logger,
		productController: productController,
		middleware:        middleware,
	}
}
