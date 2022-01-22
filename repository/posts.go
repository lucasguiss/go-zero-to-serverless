package repository

import (
	"fmt"
	"log"

	"github.com/lucasguiss/go-zero-to-serverless/config/database"
	"github.com/lucasguiss/go-zero-to-serverless/models"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
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

func FindAllPosts() []models.Post {
	collection, ctx := database.GetCollection("posts")
	iter := collection.Documents(ctx)
	var posts []models.Post
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
		data := doc.Data()
		var result models.Post
		err = mapstructure.Decode(data, &result)
		if err != nil {
			panic(err)
		}
		posts = append(posts, result)
	}

	return posts
}
