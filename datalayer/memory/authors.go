package memory

import (
	"context"
	"fmt"

	"github.com/thenanor/jsonapi-go/api/models"
)

func (dl *inMemoryStore) CreateAuthor(ctx context.Context, author models.Author) error {
	dl.mutex.Lock()
	defer dl.mutex.Unlock()

	dl.authors[author.ID] = author

	fmt.Println("check if we have added the author in the DL:", dl.authors)
	return nil
}
