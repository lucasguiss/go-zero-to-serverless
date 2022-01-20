package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasguiss/go-zero-to-serverless/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		posts := main.Group("posts")
		{
			posts.POST("/", controllers.CreatePost)
		}
	}
	return router
}
