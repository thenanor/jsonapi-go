package memory

import (
	"context"

	"github.com/thenanor/jsonapi-go/api/models"
)

func (dl *inMemoryStore) CreateAuthor(ctx context.Context, author *models.Author) error {
	dl.mutex.Lock()
	defer dl.mutex.Unlock()

	dl.authors[author.ID] = author

	return nil
}
