package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/thenanor/jsonapi-go/api/models"
)

type inMemoryStore struct {
	mutex    sync.RWMutex
	posts    map[uuid.UUID]models.Post
	authors  map[uuid.UUID]models.Author
	comments map[uuid.UUID]models.Comment
}

func New() (*inMemoryStore, error) {
	return &inMemoryStore{
		posts:    map[uuid.UUID]models.Post{},
		authors:  map[uuid.UUID]models.Author{},
		comments: map[uuid.UUID]models.Comment{},
	}, nil
}
