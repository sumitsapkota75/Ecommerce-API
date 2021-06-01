package services

import (
	"context"
	"travel/infrastructure"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

// FirebaseService represent firebase service
type FirebaseService struct {
	Firebase *auth.Client
	Logger   infrastructure.Logger
	Env      infrastructure.Env
}

// NewFirebaseService creates a new firebase service
func NewFirebaseService(fbClient *auth.Client, logger infrastructure.Logger, env infrastructure.Env) FirebaseService {
	return FirebaseService{
		Firebase: fbClient,
		Logger:   logger,
		Env:      env,
	}
}

// VerifyToken -> verify passed firebase id token
func (fb *FirebaseService) VerifyToken(idToken string) (*auth.Token, error) {
	token, err := fb.Firebase.VerifyIDToken(context.Background(), idToken)
	return token, err
}

// CreateToken creates new token from id
func (fb *FirebaseService) CreateToken(id string) (string, error) {
	token, err := fb.Firebase.CustomToken(context.Background(), id)
	return token, err
}

// GetUser -> gets firebase user from uid
func (fb *FirebaseService) GetUser(uid string) (*auth.UserRecord, error) {
	user, err := fb.Firebase.GetUser(context.Background(), uid)
	return user, err
}

// CreateUser -> creates a new user with email and password
func (fb *FirebaseService) CreateUser(email, password, username, role string) (string, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password).
		DisplayName(username)

	u, err := fb.Firebase.CreateUser(context.Background(), params)
	if err != nil {
		return "", err
	}

	// Adding claims as staff
	claims := map[string]interface{}{"role": "user"}
	err = fb.Firebase.SetCustomUserClaims(context.Background(), u.UID, claims)
	if err != nil {
		return "Internal Server Error", err
	}

	return u.UID, err
}

// UpdateUserAuth -> updates a user email & password
func (fb *FirebaseService) UpdateUserAuth(uid, email, password, username string, enabled bool) error {
	authParams := (&auth.UserToUpdate{})
	if email != "" {
		authParams = authParams.Email(email)
	}

	if password != "" {
		authParams = authParams.Password(password)
	}

	if username != "" {
		authParams = authParams.DisplayName(username)
	}

	if enabled {
		authParams = authParams.Disabled(false)
	}

	_, err := fb.Firebase.UpdateUser(context.Background(), uid, authParams)
	if err != nil {
		fb.Logger.Zap.Error("Update User Firebaser (UpdateUserAuth) ::", err.Error())
		return err
	}
	return nil
}

// DeleteUser -> deletes firebase user
func (fb *FirebaseService) DeleteUser(uid string) error {
	err := fb.Firebase.DeleteUser(context.Background(), uid)
	return err
}

// SetClaim set's claim to firebase user
func (fb *FirebaseService) SetClaim(uid string, claims gin.H) error {
	err := fb.Firebase.SetCustomUserClaims(context.Background(), uid, claims)
	return err

}

// CreateOrGet user in firebase
func (fb *FirebaseService) GetUserByEmail(email string) string {
	user, _ := fb.Firebase.GetUserByEmail(context.Background(), email)
	if user != nil {
		return user.UID
	}
	return ""
}

// CreateUser -> creates a new user with email and password
func (fb *FirebaseService) CreateDeactivatedUser(email, role string) (string, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		Disabled(true)

	u, err := fb.Firebase.CreateUser(context.Background(), params)
	if err != nil {
		return "", err
	}

	//Adding claims as staff
	claims := map[string]interface{}{"role": role}
	err = fb.Firebase.SetCustomUserClaims(context.Background(), u.UID, claims)
	if err != nil {
		return "Internal Server Error", err
	}
	return u.UID, err
}

func (fb *FirebaseService) GenerateEmailVerificationLink(email string) (string, error) {
	params := (&auth.ActionCodeSettings{
		URL: "http://localhost:8000/auth",
	})
	return fb.Firebase.EmailVerificationLinkWithSettings(context.Background(), email, params)
}
