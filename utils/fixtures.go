package utils

import (
	"time"

	"github.com/google/uuid"
	"github.com/thenanor/jsonapi-go/api/models"
)

func FixturePost() *models.Post {
	return &models.Post{
		ID:    uuid.NewString(),
		Title: "My Post",
		Author: &models.Author{
			ID:   uuid.NewString(),
			Name: "Some Author",
		},
		CreatedAt: time.Now(),
		// Comments:  []*models.Comment{},
	}
}

func FixturePosts(count int) []*models.Post {
	posts := make([]*models.Post, count)
	for i := 0; i < count; i++ {
		posts = append(posts, FixturePost())
	}
	return posts
}
