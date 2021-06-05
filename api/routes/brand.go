package routes

import (
	"travel/api/controllers"
	"travel/api/middlewares"
	"travel/infrastructure"
)

type BrandRoutes struct {
	logger          infrastructure.Logger
	handler         infrastructure.RequestHandler
	brandController controllers.BrandController

	middleware middlewares.AuthMiddleware
}

// Setup brand Routes
func (u BrandRoutes) Setup() {
	u.logger.Zap.Info(" Setting up brand routes 👤 -------------")
	brand := u.handler.Gin.Group("/brand")
	{
		brand.GET("", u.brandController.GetAllBrands)
		brand.POST("", u.brandController.CreateBrand)
		brand.GET("/:id", u.brandController.GetBrandByID)
		brand.POST("/:id", u.brandController.UpdateBrand)
		brand.DELETE("/:id", u.brandController.DeleteBrand)

	}
}

// NewBrandRoute creates new brand controller
func NewBrandRoutes(
	logger infrastructure.Logger,
	handler infrastructure.RequestHandler,
	brandController controllers.BrandController,
	middleware middlewares.AuthMiddleware,
) BrandRoutes {
	return BrandRoutes{
		handler:         handler,
		logger:          logger,
		brandController: brandController,
		middleware:      middleware,
	}
}
