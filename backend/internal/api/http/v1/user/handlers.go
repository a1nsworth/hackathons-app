package user

import (
	"errors"
	"net/http"

	"hackathons-app/internal/models"
	"hackathons-app/internal/services"

	_ "hackathons-app/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *Handler {
	return &Handler{service: service}
}

type GetUserRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type Response struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
}
type ResponseHackathon struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ResponseWithHackathons struct {
	FirstName  string              `json:"first_name"`
	SecondName string              `json:"second_name"`
	Email      string              `json:"email"`
	Hackathons []ResponseHackathon `json:"hackathons"`
}

// GetUserById - ID
//
//	@Summary		Get user by ID
//	@Description	ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"User ID"
//	@Security		BearerAuth
//	@Success		200	{object}	Response
//	@Router			/user/{id} [get]
func (u *Handler) GetUserById(c *gin.Context) {
	var request GetUserRequest

	if err := c.BindUri(&request); err != nil {
		return
	}

	user, err := u.service.GetById(request.ID)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	case err == nil:
		c.JSON(
			http.StatusOK, Response{
				FirstName:  user.FirstName,
				SecondName: user.SecondName,
				Email:      user.Email,
			},
		)
	}
}

type CreateRequest struct {
	FirstName  string `json:"first_name" binding:"required"`
	SecondName string `json:"second_name" binding:"required"`
	Email      string `json:"email" binding:"required"`
}

// CreateUser -
//
//	@Summary		Create a new user
//	@Description	-
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body	CreateRequest	true	"User data"
//	@Success		201
//	@Security		BearerAuth
//	@Router			/user/ [put]
//	@title			Gin Swagger Example API
func (u *Handler) CreateUser(c *gin.Context) {
	var request CreateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user := models.User{
		Email:      request.Email,
		FirstName:  request.FirstName,
		SecondName: request.SecondName,
	}
	if err := u.service.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// GetAll -
//
//	@Summary		Get all users
//	@Description	-
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{array}	Response
//	@Router			/user/ [get]
func (u *Handler) GetAll(c *gin.Context) {
	users, err := u.service.GetAll()
	response := make([]Response, len(users))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for i, user := range users {
		response[i] = Response{
			FirstName:  user.FirstName,
			SecondName: user.SecondName,
			Email:      user.Email,
		}
	}
	c.JSON(http.StatusOK, response)
}

// GetAllWithHackathons -
//
//	@Summary		Get all users
//	@Description	-
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{array}	ResponseWithHackathons
//	@Router			/user/hackathons/ [get]
func (u *Handler) GetAllWithHackathons(c *gin.Context) {
	users, err := u.service.GetAllWithHackathons()
	response := make([]ResponseWithHackathons, len(users))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for i, user := range users {
		hackathons := make([]ResponseHackathon, len(user.Hackathons))
		for i, h := range user.Hackathons {
			hackathons[i] = ResponseHackathon{
				Name:        h.Name,
				Description: h.Description,
			}
		}
		response[i] = ResponseWithHackathons{
			FirstName:  user.FirstName,
			SecondName: user.SecondName,
			Email:      user.Email,
			Hackathons: hackathons,
		}
	}

	c.JSON(http.StatusOK, response)
}

// DeleteUser -
//
//	@Summary		Delete user by ID
//	@Description	ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"User ID"
//	@Security		BearerAuth
//	@Success		204
//	@Router			/user/{id} [delete]
func (u *Handler) DeleteUser(c *gin.Context) {
	var request GetUserRequest

	if err := c.BindUri(&request); err != nil {
		return
	}

	if err := u.service.DeleteById(request.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

type AddHackathonRequest struct {
	HackathonId int64 `uri:"hackathonId" binding:"required"`
	UserId      int64 `uri:"userId" binding:"required"`
}

// AddHackathonById -
//
//	@Summary		Add hackathon to user
//	@Description	,    ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			userId		path	int	true	"User ID"
//	@Param			hackathonId	path	int	true	"Hackathon ID"
//	@Security		BearerAuth
//	@Success		204
//	@Router			/user/{userId}/{hackathonId} [patch]
func (u *Handler) AddHackathonById(c *gin.Context) {
	var request AddHackathonRequest
	if err := c.BindUri(&request); err != nil {
		return
	}
	if err := u.service.AddHackathonById(request.UserId, request.HackathonId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
