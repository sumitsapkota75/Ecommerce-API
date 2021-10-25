package controllers

import (
	"net/http"
	"travel/api/responses"
	"travel/api/services"
	"travel/constants"
	"travel/infrastructure"
	"travel/models"
	"travel/utils"

	"github.com/gin-gonic/gin"
)

// UserController -> data type
type UserController struct {
	logger          infrastructure.Logger
	userService     services.UserService
	firebaseService services.FirebaseService
	vendorService   services.VendorService
}

// NewUserController -> creates new user controller
func NewUserController(logger infrastructure.Logger, userService services.UserService, firebaseService services.FirebaseService, vendorService services.VendorService) UserController {
	return UserController{
		logger:          logger,
		userService:     userService,
		firebaseService: firebaseService,
		vendorService:   vendorService,
	}
}

//GetAllUsers -> get the list of users
func (u UserController) GetAllUsers(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	users, count, err := u.userService.GetAllUsers(pagination)
	if err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to get users")
		return
	}

	responses.JSONCount(c, http.StatusOK, users, int(count))
}

// CreateUser -> creates the user
func (u UserController) CreateUser(c *gin.Context) {
	requestUser := struct {
		models.User
		Password1 string `json:"password1"`
		Password2 string `json:"password2"`
	}{}
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		u.logger.Zap.Error("Error (ShouldBindJSON) ::", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "Error parsing json request")
		return
	}
	// check if two passwords match
	if requestUser.Password1 != requestUser.Password2 {
		u.logger.Zap.Error("Error [two passwords do not match] ::")
		responses.ErrorJSON(c, http.StatusBadRequest, "two passwords do not match")
		return
	}
	// Checks for firebase existing user
	firebaseID := u.firebaseService.GetUserByEmail(requestUser.Email)
	if firebaseID != "" {
		u.logger.Zap.Error("Error [create firebase user] ::")
		responses.ErrorJSON(c, http.StatusBadRequest, "The provided email is already in use")
		return
	}

	firebaseID, firebaseErr := u.firebaseService.CreateUser(requestUser.Email, requestUser.Password1, requestUser.Name, constants.CustomerUserType)
	if firebaseErr != nil {
		u.logger.Zap.Error("Error [create firebase user] ::", firebaseErr)
		responses.ErrorJSON(c, http.StatusBadRequest, "Error creating user in firebase")
		return
	}

	var user models.User
	user.ID = firebaseID
	user.Address = requestUser.Address
	user.Phone = requestUser.Phone
	user.Email = requestUser.Email
	user.Name = requestUser.Name
	user.UserType = constants.CustomerUserType
	_, err := u.userService.CreateUser(user)
	if err != nil {
		u.logger.Zap.Error("Error [db CreateUser]: ", err.Error())
		responses.ErrorJSON(c, http.StatusInternalServerError, "failed to Save User in Database")
		return
	}
	responses.SuccessJSON(c, http.StatusCreated, "User created successfully")

}

// GetUserProfile -> gets a user by ID
func (u UserController) GetUserProfile(c *gin.Context) {
	uid := c.MustGet(constants.UID).(string)
	user, err := u.userService.GetUserByID(uid)
	if err != nil {
		u.logger.Zap.Error("Error getting user ::", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "failed to get user")
		return
	}
	responses.JSON(c, http.StatusOK, user)
}

// UpdateUser
func (u UserController) UpdateUser(c *gin.Context) {
	uid := c.MustGet(constants.UID).(string)
	requestUser := struct {
		models.User
		Password string `json:"password"`
	}{}
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		u.logger.Zap.Error("Error [ShouldBindJSON] ::", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "failed to parse json data")
		return
	}
	user, err := u.userService.GetUserByID(uid)
	if err != nil {
		u.logger.Zap.Error("Error getting user data ::", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "failed to get user data")
		return
	}
	// Checks for firebase existing user and update firebase authentication details
	if err := u.firebaseService.UpdateUserAuth(
		uid,
		requestUser.Email,
		requestUser.Password,
		requestUser.Name,
		true,
	); err != nil {
		u.logger.Zap.Error("Error [UpdateUserAuth] ::", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "The provided email is already in use")
		return
	}
	user.Name = requestUser.Name
	user.Address = requestUser.Address
	user.Phone = requestUser.Phone
	user.Email = requestUser.Email
	if err := u.userService.UpdateUser(user); err != nil {
		u.logger.Zap.Error("Failed to update the user ::", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "failed to update the user data")
		return
	}
	responses.JSON(c, http.StatusBadRequest, "User updated successfully")

}
