package routes

import (
	"travel/api/controllers"
	"travel/api/middlewares"
	"travel/infrastructure"
)

type UserRoutes struct {
	logger         infrastructure.Logger
	handler        infrastructure.RequestHandler
	userController controllers.UserController

	middleware middlewares.AuthMiddleware
}

// Setup user Routes
func (u UserRoutes) Setup() {
	u.logger.Zap.Info(" Setting up user routes 👤 -------------")
	user := u.handler.Gin.Group("/user")
	{
		user.POST("/signup", u.userController.CreateUser)
		user.GET("/profile", u.middleware.Handle(), u.userController.GetUserProfile)
	}
}

// NewUserRoutes creates new user controller
func NewUserRoutes(
	logger infrastructure.Logger,
	handler infrastructure.RequestHandler,
	userController controllers.UserController,
	middleware middlewares.AuthMiddleware,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
		middleware:     middleware,
	}
}
