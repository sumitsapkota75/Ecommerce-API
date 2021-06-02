package routes

import (
	"travel/api/controllers"
	"travel/api/middlewares"
	"travel/infrastructure"
)

type CategoryRoutes struct {
	logger             infrastructure.Logger
	handler            infrastructure.RequestHandler
	categoryController controllers.CategoryController

	middleware middlewares.AuthMiddleware
}

// Setup category Routes
func (u CategoryRoutes) Setup() {
	u.logger.Zap.Info(" Setting up category routes 👤 -------------")
	category := u.handler.Gin.Group("/category")
	{
		category.GET("", u.categoryController.GetAllCategories)
		category.POST("", u.categoryController.CreateCategory)
		category.GET("/:id", u.categoryController.GetCategoryByID)
		category.POST("/:id", u.categoryController.UpdateCategory)

	}
}

// NewCategoryRoutes creates new category controller
func NewCategoryRoutes(
	logger infrastructure.Logger,
	handler infrastructure.RequestHandler,
	categoryController controllers.CategoryController,
	middleware middlewares.AuthMiddleware,
) CategoryRoutes {
	return CategoryRoutes{
		handler:            handler,
		logger:             logger,
		categoryController: categoryController,
		middleware:         middleware,
	}
}
