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

// Setup user Routes
func (u CategoryRoutes) Setup() {
	u.logger.Zap.Info(" Setting up category routes 👤 -------------")
	user := u.handler.Gin.Group("/category")
	{
		user.GET("", u.categoryController.GetAllCategories)
		user.POST("", u.categoryController.CreateCategory)

	}
}

// NewCategoryRoutes creates new user controller
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
