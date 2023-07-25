package memory

import (
	"sync"

	"github.com/thenanor/jsonapi-go/api/models"
)

type inMemoryStore struct {
	mutex    sync.RWMutex
	posts    map[string]models.Post
	authors  map[string]models.Author
	comments map[string]models.Comment
}

func New() (*inMemoryStore, error) {
	return &inMemoryStore{
		posts:    map[string]models.Post{},
		authors:  map[string]models.Author{},
		comments: map[string]models.Comment{},
	}, nil
}
