package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"travel/api/responses"
	"travel/api/services"
	"travel/constants"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware -> struct
type AuthMiddleware struct {
	firebaseService services.FirebaseService
}

//NewAuthMiddleware -> createa new auth middleware
func NewAuthMiddleware(firebaseService services.FirebaseService) AuthMiddleware {
	return AuthMiddleware{
		firebaseService: firebaseService,
	}
}

// Hnadle -> handles auth requests

func (m AuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := m.getTokenFromHeader(c)

		if err != nil {
			responses.ErrorJSON(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		fmt.Println("TOKEN:::::", token.Claims)
		c.Set(constants.UID, token.UID)
		c.Next()
	}
}

// getTokenFromHeader -> gets token from header
func (m AuthMiddleware) getTokenFromHeader(c *gin.Context) (*auth.Token, error) {
	header := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))
	token, err := m.firebaseService.VerifyToken(idToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}
