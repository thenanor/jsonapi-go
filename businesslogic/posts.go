package businesslogic

import (
	"context"

	"github.com/thenanor/jsonapi-go/api/models"
)

func (bl *BL) CreatePost(ctx context.Context, post models.Post) (models.Post, error) {
	return models.Post{}, nil
}
