package auth

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"hackathons-app/internal/config"
	"hackathons-app/internal/jwt"
	"hackathons-app/internal/models"
	"hackathons-app/internal/services"
	"hackathons-app/pkg/crypto"

	"github.com/gin-gonic/gin"
	goJwt "github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	config  config.JWT
	service services.UserService
}

func NewAuthHandler(service services.UserService, config config.JWT) *Handler {
	return &Handler{service: service, config: config}
}

type Token = string

type RefreshRequest struct {
	Access  Token `json:"access_token" binding:"required"`
	Refresh Token `json:"refresh_token" binding:"required"`
}

type ResponseRefresh struct {
	Access Token `json:"access_token"`
}

// Refresh godoc
// @Summary Refresh Access Token
// @Description This endpoint refreshes the access token using a valid refresh token.
// @Tags auth
// @Accept json
// @Produce json
// @Param refreshRequest body RefreshRequest true "Request body containing access and refresh tokens"
// @Success 200 {object} ResponseRefresh
// @Security BearerAuth
// @Failure 400 "Invalid input or token issues"
// @Failure 401 "Access token is not expired"
// @Failure 500 "Internal server error"
// @Router /auth/refresh [post]
func (h *Handler) Refresh(c *gin.Context) {
	var request RefreshRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	_, err := jwt.DecodeJwt(request.Access, h.config.AccessSecret)
	if err == nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "access token is not expired"},
		)
		return
	}

	if !errors.Is(err, goJwt.ErrTokenExpired) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid access token"})
		return
	}

	refresh, err := jwt.DecodeJwt(request.Refresh, h.config.RefreshSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newAccess, err := refresh.CreateJwt(
		h.config.AccessSecret,
		time.Second*time.Duration(h.config.AccessExp),
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResponseRefresh{Access: newAccess})
}

type RegisterRequest struct {
	Email          string `json:"email"`
	FirstName      string `json:"first_name"`
	SecondName     string `json:"second_name"`
	Password       string `json:"password"`
	VerifyPassword string `json:"verify_password"`
}
type RegisterResponse struct {
	Access  Token `json:"access_token"`
	Refresh Token `json:"refresh_token"`
}

// Register Register -   .
//
//	@Summary		Register a new user
//	@Description	.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body	RegisterRequest	true	"User data"
//	@Success		201
//	@Router			/auth/register [post]
func (h *Handler) Register(c *gin.Context) {
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
	accessToken, err := payload.CreateJwt(
		h.config.AccessSecret,
		time.Second*time.Duration(h.config.AccessExp),
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	refresh, err := payload.CreateJwt(
		h.config.RefreshSecret,
		time.Second*time.Duration(h.config.RefreshExp),
	)
	response := RegisterResponse{
		Access:  accessToken,
		Refresh: refresh,
	}
	c.JSON(http.StatusOK, response)
}
