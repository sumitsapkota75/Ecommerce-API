package controllers

import (
	"travel/infrastructure"

	"github.com/gin-gonic/gin"
)

//UtilityController -> struct representing controller for utility
type UtilityController struct {
	logger infrastructure.Logger
	env    infrastructure.Env
}

//NewUtilityController -> returns new utility controller
func NewUtilityController(logger infrastructure.Logger,
	env infrastructure.Env,
) UtilityController {
	return UtilityController{
		logger: logger,
		env:    env,
	}
}

type Response struct {
	Success bool   `json:"success"`
	URL     string `json:"url"`
}

func (u UtilityController) FileUpload(c *gin.Context) {

}
