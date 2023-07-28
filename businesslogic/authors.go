package businesslogic

import (
	"context"

	"github.com/google/uuid"
	"github.com/thenanor/jsonapi-go/api/models"
)

func (bl *BL) CreateAuthor(ctx context.Context, author *models.Author) (*models.Author, error) {
	// do some validations on author
	if author != nil {
		if author.ID == "" {
			author.ID = uuid.NewString()
		}

		// Send it to the DL
		err := bl.datalayer.CreateAuthor(ctx, author)
		if err != nil {
			return nil, err
		}
	}

	return author, nil
}
