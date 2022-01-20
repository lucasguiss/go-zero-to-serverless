package repository

import (
	"log"

	"github.com/lucasguiss/go-zero-to-serverless/config/database"
	"github.com/lucasguiss/go-zero-to-serverless/models"
)

func CreatePost(post models.Post) {
	collection, ctx := database.GetCollection("posts")
	_, _, err := collection.Add(ctx, map[string]interface{}{
		"author":  post.Author,
		"content": post.Content,
	})
	if err != nil {
		log.Fatalf("Error creating new post: %v", err)
	}
}
