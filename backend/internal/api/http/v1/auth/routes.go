package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(engine *gin.Engine, handler *Handler) {
	engine.POST("/register", handler.RegisterHandler)
}
