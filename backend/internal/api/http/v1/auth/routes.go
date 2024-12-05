package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(engine *gin.Engine, handler *Handler) {
	router := engine.Group("/auth")
	{
		router.POST("/register", handler.Register)
		router.POST("/refresh", handler.Refresh)
	}
}
