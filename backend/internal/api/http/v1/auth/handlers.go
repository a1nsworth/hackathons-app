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

type TokenPair struct {
	Access  string
	Refresh string
}

func (h *Handler) generateTokenPair(userId int) (pair TokenPair, err error) {
	payload := jwt.NewPayload(strconv.Itoa(userId))
	accessToken, err := payload.CreateJwt(
		h.config.AccessSecret,
		time.Second*time.Duration(h.config.AccessExp),
	)
	if err != nil {
		return
	}

	refresh, err := payload.CreateJwt(
		h.config.RefreshSecret,
		time.Second*time.Duration(h.config.RefreshExp),
	)
	if err != nil {
		return
	}

	return TokenPair{accessToken, refresh}, nil
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Access  string `json:"accessToken"`
	Refresh string `json:"refreshToken"`
}

// Login handles user login and returns access and refresh tokens
// @Summary Login a user and return tokens
// @Description Login a user using email and password and return access and refresh tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param loginRequest body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse "Successfully logged in"
// @Failure 400 "Invalid input"
// @Failure 401 "User not found"
// @Failure 401 "Invalid password"
// @Failure 500 "Internal server error"
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.GetByEmail(request.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no such user"})
		return
	}

	if !crypto.CheckPasswordHash(request.Password, user.HashedPassword) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	userId := int(user.ID)
	pair, err := h.generateTokenPair(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := LoginResponse{
		Access:  pair.Access,
		Refresh: pair.Refresh,
	}
	c.JSON(http.StatusOK, response)
}

type Token = string

type RefreshRequest struct {
	Access  Token `json:"accessToken" binding:"required"`
	Refresh Token `json:"refreshToken" binding:"required"`
}

type ResponseRefresh struct {
	Access Token `json:"accessToken"`
}

// Refresh godoc
//
//	@Summary		Refresh Access Token
//	@Description	This endpoint refreshes the access token using a valid refresh token.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			refreshRequest	body		RefreshRequest	true	"Request body containing access and refresh tokens"
//	@Success		200				{object}	ResponseRefresh
//	@Security		BearerAuth
//	@Failure		400	"Invalid input or token issues"
//	@Failure		401	"Access token is not expired"
//	@Failure		500	"Internal server error"
//	@Router			/auth/refresh [post]
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
	Email           string `json:"email" binding:"required,email"`
	FirstName       string `json:"first_name"`
	SecondName      string `json:"second_name"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}
type RegisterResponse struct {
	Access  string `json:"accessToken"`
	Refresh string `json:"refreshToken"`
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
	if request.Password != request.ConfirmPassword {
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
	pair, err := h.generateTokenPair(int(user.ID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := RegisterResponse{
		Access:  pair.Access,
		Refresh: pair.Refresh,
	}
	c.JSON(http.StatusOK, response)
}
