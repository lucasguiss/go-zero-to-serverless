package services

import (
	"github.com/lucasguiss/go-zero-to-serverless/models"
	"github.com/lucasguiss/go-zero-to-serverless/repository"
)

func CreatePost(post models.Post) {
	repository.CreatePost(post)
}

func FindAllPosts() []models.Post {
	posts := repository.FindAllPosts()
	return posts
}
