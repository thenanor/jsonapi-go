package memory

import (
	"context"
	"fmt"

	"github.com/thenanor/jsonapi-go/api/models"
)

var (
	ErrNotFound = fmt.Errorf("Resource not found")
)

func (dl *inMemoryStore) CreatePost(ctx context.Context, post models.Post) error {
	dl.mutex.Lock()
	defer dl.mutex.Unlock()

	dl.posts[post.ID] = post

	fmt.Println("check if we have added the post in the DL:", dl.posts)
	return nil
}

func (dl *inMemoryStore) GetPost(ctx context.Context, id string) (models.Post, error) {
	dl.mutex.Lock()
	defer dl.mutex.Unlock()

	if post, ok := dl.posts[id]; ok {
		return post, nil
	}
	return models.Post{}, ErrNotFound
}

func (dl *inMemoryStore) GetPosts(ctx context.Context) ([]models.Post, error) {
	dl.mutex.Lock()
	defer dl.mutex.Unlock()

	posts := make([]models.Post, 0)
	for _, post := range dl.posts {
		posts = append(posts, post)
	}
	return posts, nil
}
