package routes

import (
	"travel/api/controllers"
	"travel/infrastructure"
)

// UtilityRoutes -> utility routes struct
type UtilityRoutes struct {
	Logger            infrastructure.Logger
	Handler           infrastructure.RequestHandler
	UtilityController controllers.UtilityController
}

//NewUtilityRoute -> returns new utility route
func NewUtilityRoutes(logger infrastructure.Logger, handler infrastructure.RequestHandler, UtilityController controllers.UtilityController) UtilityRoutes {
	return UtilityRoutes{
		Logger:            logger,
		Handler:           handler,
		UtilityController: UtilityController,
	}
}

//Setup -> sets up route for util entities
func (s UtilityRoutes) Setup() {
	s.Logger.Zap.Info("Setting up Utility Routes 🌴 ----------")
	util := s.Handler.Gin.Group("/utils")
	{
		util.POST("/upload-file", s.UtilityController.FileUploadHandler)
	}
}
