package hackathon

import (
	"hackathons-app/internal/api/http/v1/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(engine *gin.Engine, handler *Handler, a *auth.Auth) {
	router := engine.Group("/hackathons/")
	{
		router.GET(":id", a.AuthorizedUser, handler.GetById)
		router.GET("", a.AuthorizedUser, handler.GetAll)
		router.POST("", a.AuthorizedUser, handler.Create)
		router.DELETE(":id", a.OnlyAdmin, handler.DeleteById)
		router.PATCH(":id", a.OnlyAdmin, handler.Update)
	}
}
