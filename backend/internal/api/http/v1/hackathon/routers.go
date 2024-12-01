package hackathon

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(engine *gin.Engine, handler *HackathonHandler) {
	router := engine.Group("/hackathons")
	router.GET(":id", handler.GetById)
	router.GET("", handler.GetAll)
	router.POST("", handler.Create)
	router.DELETE(":id", handler.DeleteById)
	router.PATCH(":id", handler.Update)
}
