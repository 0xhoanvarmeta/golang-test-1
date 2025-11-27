package routers

import (
	"test-1/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterPublicEndpoint(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Public API Endpoint!",
		})
	})
}

func RegisterUserPublicEndpoint(r *gin.Engine, userHandler *handlers.UserHandler) {
	r.POST("/users", userHandler.CreateUser)
}
