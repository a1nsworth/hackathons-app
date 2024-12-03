package auth

import (
	"net/http"
	"strconv"
	"time"

	"hackathons-app/internal/config"
	"hackathons-app/internal/jwt"
	"hackathons-app/internal/models"
	"hackathons-app/internal/services"
	"hackathons-app/pkg/crypto"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	config  config.JWT
	service services.UserService
}

func NewAuthHandler(service services.UserService, config config.JWT) *Handler {
	return &Handler{service: service, config: config}
}

type RegisterRequest struct {
	Email          string `json:"email"`
	FirstName      string `json:"first_name"`
	SecondName     string `json:"second_name"`
	Password       string `json:"password"`
	VerifyPassword string `json:"verify_password"`
}

// RegisterHandler Register -   .
//
//	@Summary		Register a new user
//	@Description	.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body	RegisterRequest	true	"User data"
//	@Success		201
//	@Router			/register [post]
func (h *Handler) RegisterHandler(c *gin.Context) {
	var request RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	if request.Password != request.VerifyPassword {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "passwords do not match"})
		return
	}
	hashedPassword, err := crypto.HashPassword(request.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:          request.Email,
		FirstName:      request.FirstName,
		SecondName:     request.SecondName,
		HashedPassword: hashedPassword,
	}
	if err := h.service.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	payload := jwt.NewPayload(strconv.Itoa(int(user.ID)))
	accessToken, err := payload.CreateJwt(h.config.Secret, time.Second*time.Duration(h.config.Exp))
	c.JSON(
		http.StatusOK, gin.H{
			"access": accessToken,
		},
	)
}
