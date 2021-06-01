package controllers

import (
	"fmt"
	"net/http"
	"travel/api/responses"
	"travel/api/services"
	"travel/infrastructure"
	"travel/models"

	"github.com/gin-gonic/gin"
)

// UserController -> data type
type UserController struct {
	logger          infrastructure.Logger
	userService     services.UserService
	firebaseService services.FirebaseService
}

// NewUserController -> creates new user controller
func NewUserController(logger infrastructure.Logger, userService services.UserService, firebaseService services.FirebaseService) UserController {
	return UserController{
		logger:          logger,
		userService:     userService,
		firebaseService: firebaseService,
	}
}

// CreateUser -> creates the user
func (u UserController) CreateUser(c *gin.Context) {
	// requestUser := struct {
	// 	models.User
	// }{}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		// u.logger.Zap.Error("Error (ShouldBindJSON) ::", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "Error parsing json request")
		return
	}
	fmt.Println("request user", user.Address)
	// _, firebaseErr := u.firebaseService.CreateUser(user.Email, "passowrd", user.Name, constants.CustomerUserType)
	// if firebaseErr != nil {
	// 	u.logger.Zap.Error("Error [create firebase user] ::", firebaseErr)
	// 	responses.ErrorJSON(c, http.StatusBadRequest, "The provided email is already in use")
	// 	return
	// }
	// var user models.User
	// user.Address = user.Address
	// user.Email = user.Email
	// user.Name = requestUser.Name
	// user.Name = "sumit"
	// u.logger.Zap.Info("Creating user ")
	_, err := u.userService.CreateUser(user)
	if err != nil {
		// u.logger.Zap.Error("Error [db CreateUser]: ", err.Error())
		responses.ErrorJSON(c, http.StatusInternalServerError, "failed to create User")
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "User created successfully")

}
