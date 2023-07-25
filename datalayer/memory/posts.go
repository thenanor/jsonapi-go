package memory

import (
	"context"

	"github.com/thenanor/jsonapi-go/api/models"
)

func (dl *inMemoryStore) CreatePost(ctx context.Context, post models.Post) error {
	dl.mutex.Lock()
	defer dl.mutex.Unlock()

	dl.posts[post.ID] = post

	return nil
}
