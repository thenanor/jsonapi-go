package memory

import (
	"context"
	"fmt"

	"github.com/thenanor/jsonapi-go/api/models"
)

func (dl *inMemoryStore) CreatePost(ctx context.Context, post models.Post) error {
	dl.mutex.Lock()
	defer dl.mutex.Unlock()

	dl.posts[post.ID] = post

	fmt.Println("check if we have added the post in the DL:", dl.posts)
	return nil
}
