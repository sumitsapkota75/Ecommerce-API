package infrastructure

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RequestHandler function
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates a new request handler
func NewRequestHandler(env Env) RequestHandler {

	// Router handler
	httpRouter := gin.Default()
	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	httpRouter.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Travel Agency API Up and Running"})
	})
	return RequestHandler{Gin: httpRouter}
}
