package routes

import "github.com/gin-gonic/gin"

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		posts := main.Group("posts")
		{
			posts.GET("/")
		}
	}
	return router
}
