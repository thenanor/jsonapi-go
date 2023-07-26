package businesslogic

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thenanor/jsonapi-go/api/models"
)

var (
	ErrInvalidID = fmt.Errorf("Invalid ID")
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

func (bl *BL) GetPost(ctx context.Context, id string) (models.Post, error) {
	if id == "" {
		return models.Post{}, ErrInvalidID
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return models.Post{}, ErrInvalidID
	}

	post, err := bl.datalayer.GetPost(ctx, id)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func (bl *BL) GetPosts(ctx context.Context) ([]models.Post, error) {
	posts, err := bl.datalayer.GetPosts(ctx)
	if err != nil {
		return []models.Post{}, err
	}
	return posts, nil
}
