package businesslogic

import (
	"context"
	"fmt"

	"github.com/thenanor/jsonapi-go/api/models"
	"github.com/thenanor/jsonapi-go/datalayer"
)

type Businesslogic interface {
	CreatePost(context.Context, models.Post) (models.Post, error)

	CreateAuthor(context.Context, models.Author) (models.Author, error)
}

type BL struct {
	datalayer datalayer.Datalayer
}

func New() (Businesslogic, error) {
	dl, err := datalayer.New()
	if err != nil {
		return nil, fmt.Errorf("unable to create datalayer: %w", err)
	}

	return &BL{
		datalayer: dl,
	}, nil
}
