package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(engine *gin.Engine, handler *Handler) {
	router := engine.Group("/user/")
	{
		router.GET(":id", handler.GetUserById)
		router.GET("", handler.GetAll)
		router.GET("hackathons/", handler.GetAllWithHackathons)
		router.PUT("", handler.CreateUser)
		router.DELETE(":id", handler.DeleteUser)
		router.PATCH(":userId/:hackathonId", handler.AddHackathonById)
	}
}
