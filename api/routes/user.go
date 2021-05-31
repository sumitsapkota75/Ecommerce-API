package routes

import (
	"travel/api/responses"
	"travel/infrastructure"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	logger  infrastructure.Logger
	handler infrastructure.RequestHandler
	// clientController controllers.ClientController

	// middleware middlewares.AuthMiddleware
}

// Setup user Routes
func (u UserRoutes) Setup() {
	u.logger.Zap.Info(" Setting up user routes 👤 -------------")
	user := u.handler.Gin.Group("/users")
	{
		user.GET("", func(c *gin.Context) {
			responses.JSON(c, 200, "Users api get running")
		})
	}
}

// NewUserRoutes creates new user controller
func NewUserRoutes(
	logger infrastructure.Logger,
	handler infrastructure.RequestHandler,
	// clientController controllers.ClientController,
	// middleware middlewares.AuthMiddleware,
) UserRoutes {
	return UserRoutes{
		handler: handler,
		logger:  logger,
		// clientController: clientController,
		// middleware: middleware,
	}
}
