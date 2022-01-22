package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasguiss/go-zero-to-serverless/models"
	"github.com/lucasguiss/go-zero-to-serverless/services"
)

func CreatePost(c *gin.Context) {

	var post models.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})

		return
	}

	services.CreatePost(post)

	c.Status(204)
}

func FindAllPosts(c *gin.Context) {
	posts := services.FindAllPosts()

	c.JSON(200, posts)
}
