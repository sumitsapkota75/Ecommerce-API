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
	smtp            services.SMTPService
}

// NewUserController -> creates new user controller
func NewUserController(logger infrastructure.Logger, userService services.UserService, firebaseService services.FirebaseService, vendorService services.VendorService, smtp services.SMTPService) UserController {
	return UserController{
		logger:          logger,
		userService:     userService,
		firebaseService: firebaseService,
		vendorService:   vendorService,
		smtp:            smtp,
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
		Password string `json:"password"`
	}{}
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		u.logger.Zap.Error("Error (ShouldBindJSON) ::", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "Error parsing json request")
		return
	}

	// Checks for firebase existing user
	firebaseID := u.firebaseService.GetUserByEmail(requestUser.Email)
	if firebaseID != "" {
		u.logger.Zap.Error("Error [create firebase user] ::")
		responses.ErrorJSON(c, http.StatusBadRequest, "The provided email is already in use")
		return
	}

	firebaseID, firebaseErr := u.firebaseService.CreateUser(requestUser.Email, requestUser.Password, requestUser.Name, constants.CustomerUserType)
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
	link, _ := u.firebaseService.GenerateEmailVerificationLink(user.Email)
	u.logger.Zap.Info(" [LINK]: ", link)
	// Send Email for the user verification
	var emailSubject models.EmailSubject
	emailSubject.Title = "User Verification Link"

	var emailBody models.EmailBody
	emailBody.ToName = user.Name
	emailBody.Name = user.Name
	emailBody.Title = "Verify Email"
	emailBody.URL = link
	// email param
	var emailParam models.EmailParams
	emailParam.SubjectTemplate = "email_verification_subject.txt"
	emailParam.SubjectData = emailSubject
	emailParam.From = "sumitsapkota0@gmail.com"
	emailParam.To = user.Email
	emailParam.BodyTemplate = "email_verification_body.txt"
	emailParam.BodyData = emailBody
	// Send the message
	_, err = u.smtp.SendMail(emailParam)
	if err != nil {
		u.logger.Zap.Error("Error [sending email]: ", err.Error())
		responses.ErrorJSON(c, http.StatusInternalServerError, "failed to send email")
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
