package user

import (
	"hackathons-app/internal/api/http/v1/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(engine *gin.Engine, handler *Handler, a *auth.Auth) {
	router := engine.Group("/user/")
	{
		router.GET(":id", a.OnlyAdmin, handler.GetUserById)
		router.GET("", a.OnlyAdmin, handler.GetAll)
		router.GET("hackathons/", a.OnlyAdmin, handler.GetAllWithHackathons)
		router.PUT("", a.OnlyAdmin, handler.CreateUser)
		router.DELETE(":id", a.OnlyAdmin, handler.DeleteUser)
		router.PATCH(":userId/:hackathonId", a.AuthorizedUser, handler.AddHackathonById)
	}
}
