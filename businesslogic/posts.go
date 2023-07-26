package businesslogic

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/thenanor/jsonapi-go/api/models"
)

func (bl *BL) CreatePost(ctx context.Context, post models.Post) (models.Post, error) {
	// do some validations on post
	if post.ID == "" {
		post.ID = uuid.NewString()
	}

	if post.CreatedAt.IsZero() {
		post.CreatedAt = time.Now()
	}

	if post.Author != nil {
		// do some validations on Author - maybe call the BL Author?
		_, err := bl.CreateAuthor(ctx, *post.Author)
		if err != nil {
			// do we want to continue with adding the post without author (anonymous)? or we want to return an error?
			return models.Post{}, err
		}
	}

	// Send it to the DL
	err := bl.datalayer.CreatePost(ctx, post)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil

}
