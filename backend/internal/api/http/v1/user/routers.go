package user

import (
	"hackathons-app/internal/api/http/v1/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(engine *gin.Engine, handler *Handler) {
	router := engine.Group("/user/")
	{
		router.GET(":id", auth.AuthMiddleware, handler.GetUserById)
		router.GET("", handler.GetAll)
		router.GET("hackathons/", handler.GetAllWithHackathons)
		router.PUT("", handler.CreateUser)
		router.DELETE(":id", handler.DeleteUser)
		router.PATCH(":userId/:hackathonId", handler.AddHackathonById)
	}
}
