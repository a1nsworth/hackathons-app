package auth

import (
	"errors"
	"net/http"
	"strings"

	"hackathons-app/internal/jwt"

	goJwt "github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "Authorization header is missing"},
		)
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "Invalid authorization format"},
		)
	}

	token, found := strings.CutPrefix(authHeader, "Bearer ")
	if !found {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing Bearer"})
	}

	decodeJwt, err := jwt.DecodeJwt(
		token,
		"dc195722db37479538cd493bc876d912dc1bbc7da504420276f8c16e67ef511cace90546d26f27c19bfbc07d489c194791ebc9696bcf6dd5dd9d108e5a54b9497537420e27145b3a1f9a2acca5d2c319c534cbb63c06a8c2992c9d0aa2eeb9ecfc896f0c09df4f900e06c4e8d42001612dd1b41f91d3bd44cb2cbdabbdefe133f1eff7ad94cbbacde65bbb949ab8bfbabb7f77a71309a0acd934a1c589c998ceb75f0408d941fec95db3d747350df603b4c9232689a78c24112ceb8436c06100d1bf958c451c84966387f54bd839c0477ce704543134bfedebd3aa4d1078df8a412016fd2b98cb72a0fd5aa061a62a535c0a86122c779e7866d9f0760ad804ee",
	)
	if err != nil {
		switch {
		case errors.Is(err, goJwt.ErrTokenExpired):
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token expired"})
		default:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		}
	}
	c.Set("payload", decodeJwt)
	c.Next()
}
