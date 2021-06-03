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
	middleware     middlewares.AuthMiddleware
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

// Setup user Routes
func (u UserRoutes) Setup() {
	u.logger.Zap.Info("Setting up user routes 👤 -------------")
	user := u.handler.Gin.Group("/user")
	{
		user.GET("", u.middleware.Handle(), u.userController.GetAllUsers)
		user.POST("/signup", u.userController.CreateUser)
		user.GET("/profile", u.middleware.Handle(), u.userController.GetUserProfile)
		user.POST("/update", u.middleware.Handle(), u.userController.UpdateUser)
	}
}
