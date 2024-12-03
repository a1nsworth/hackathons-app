package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"hackathons-app/internal/jwt"
	"hackathons-app/internal/models"

	"github.com/gin-gonic/gin"
	goJwt "github.com/golang-jwt/jwt/v5"
)

func (a Auth) GetAuthorisedUser(c *gin.Context) (payload *jwt.Payload, err error) {
	a.log.Info("getting authorization header")
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is missing")
	}

	a.log.Info("checking authorization format Bearer")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, fmt.Errorf("invalid authorization format")
	}

	a.log.Info("getting token from header")
	token, found := strings.CutPrefix(authHeader, "Bearer ")
	if !found {
		return nil, fmt.Errorf("missing Bearer")
	}

	a.log.Info("validating token")
	payload, err = a.JWTAuthentication(token)
	if err != nil {
		switch {
		case errors.Is(err, goJwt.ErrTokenExpired):
			return nil, fmt.Errorf("token expired")
		default:
			return nil, fmt.Errorf("invalid token")
		}
	}
	a.log.Info("successfully validated token")
	return payload, nil
}

func (a Auth) ActiveUser(c *gin.Context) (user models.User, err error) {
	payload, err := a.GetAuthorisedUser(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(payload.Subject)
	if err != nil {
		return
	}

	a.log.Info("getting user by id")
	user, err = a.user.GetById(int64(id))
	if err != nil {
		return
	}
	a.log.Info("successfully got user by id")
	return
}

func (a Auth) AuthorizedUser(c *gin.Context) {
	user, err := a.ActiveUser(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set("user", user)
	c.Next()

}

func (a Auth) OnlyAdmin(c *gin.Context) {
	user, err := a.ActiveUser(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if user.Role != models.Admin {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	c.Set("user", user)
	c.Next()
}
